package account

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hatlonely/account/internal/mysqldb"
	"github.com/hatlonely/account/internal/rule"
	"github.com/sirupsen/logrus"
)

type RegisterReqBody struct {
	Username   string `json:"username,omitempty"`
	FirstName  string `json:"firstName,omitempty"`
	SecondName string `json:"secondName,omitempty"`
	Phone      string `json:"phone,omitempty"`
	Email      string `json:"email,omitempty"`
	Password   string `json:"password,omitempty"`
	Birthday   string `json:"birthday,omitempty"`
	Gender     string `json:"gender,omitempty"`
}

type RegisterResBody struct {
	Success bool `json:"success,omitempty"`
}

func (s *Service) Register(c *gin.Context) {
	rid := c.DefaultQuery("rid", NewToken())
	req := &RegisterReqBody{}
	var res *RegisterResBody
	var err error
	var buf []byte
	status := http.StatusOK

	defer func() {
		AccessLog.WithFields(logrus.Fields{
			"host":   c.Request.Host,
			"body":   string(buf),
			"url":    c.Request.URL.String(),
			"req":    req,
			"res":    res,
			"rid":    rid,
			"err":    err,
			"status": status,
		}).Info()
	}()

	buf, err = c.GetRawData()
	if err != nil {
		err = fmt.Errorf("get raw data failed, err: [%v]", err)
		WarnLog.WithField("@rid", rid).WithField("err", err).Warn()
		status = http.StatusBadRequest
		c.String(status, err.Error())
		return
	}

	if err = json.Unmarshal(buf, req); err != nil {
		err = fmt.Errorf("json unmarshal body failed. body: [%v], err: [%v]", string(buf), err)
		WarnLog.WithField("@rid", rid).WithField("err", err).Warn()
		status = http.StatusBadRequest
		c.String(status, err.Error())
		return
	}

	if err = s.checkRegisterReqBody(req); err != nil {
		err = fmt.Errorf("check request body failed. body: [%v], err: [%v]", string(buf), err)
		WarnLog.WithField("@rid", rid).WithField("err", err).Warn()
		status = http.StatusBadRequest
		c.String(status, err.Error())
		return
	}

	res, err = s.register(req)
	if err != nil {
		err = fmt.Errorf("register failed. err: [%v]", err)
		WarnLog.WithField("@rid", rid).WithField("err", err).Warn()
		status = http.StatusInternalServerError
		c.String(status, err.Error())
		return
	}

	status = http.StatusOK
	c.JSON(status, res)
}

func (s *Service) checkRegisterReqBody(req *RegisterReqBody) error {
	if err := rule.Check(req.Username, []rule.Rule{rule.Required, rule.AtMost64Characters}); err != nil {
		return fmt.Errorf("username[%v] %v", req.Username, err)
	}

	if req.Phone == "" && req.Email == "" {
		return fmt.Errorf("email and phone should not be empty together")
	}
	if req.Phone != "" && !ValidatePhone(req.Phone) {
		return fmt.Errorf("invalid phone [%v]", req.Phone)
	}
	if req.Email != "" && !ValidateEmail(req.Email) {
		return fmt.Errorf("invalid email [%v]", req.Email)
	}
	return nil
}

func (s *Service) register(req *RegisterReqBody) (*RegisterResBody, error) {
	ok, err := s.db.InsertAccount(&mysqldb.Account{
		Username: req.Username,
		Phone:    req.Phone,
		Email:    req.Email,
		Password: req.Password,
	})

	return &RegisterResBody{Success: ok}, err
}

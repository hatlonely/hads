package account

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hatlonely/account/internal/mysqldb"
	"github.com/hatlonely/account/internal/rule"
	"github.com/sirupsen/logrus"
)

type SignUpReqBody struct {
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
	Phone     string `json:"phone,omitempty"`
	Email     string `json:"email,omitempty"`
	Password  string `json:"password,omitempty"`
	Birthday  string `json:"birthday,omitempty"`
	Gender    int    `json:"gender,omitempty"`
}

type SignUpResBody struct {
	Success bool `json:"success,omitempty"`
}

func (s *Service) SignUp(c *gin.Context) {
	rid := c.DefaultQuery("rid", NewToken())
	req := &SignUpReqBody{}
	var res *SignUpResBody
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

	if err = s.checkSignUpReqBody(req); err != nil {
		err = fmt.Errorf("check request body failed. body: [%v], err: [%v]", string(buf), err)
		WarnLog.WithField("@rid", rid).WithField("err", err).Warn()
		status = http.StatusBadRequest
		c.String(status, err.Error())
		return
	}

	res, err = s.signUp(req)
	if err != nil {
		err = fmt.Errorf("signUp failed. err: [%v]", err)
		WarnLog.WithField("@rid", rid).WithField("err", err).Warn()
		status = http.StatusInternalServerError
		c.String(status, err.Error())
		return
	}

	status = http.StatusOK
	c.JSON(status, res)
}

func (s *Service) checkSignUpReqBody(req *SignUpReqBody) error {
	if err := rule.Check(map[string][]rule.Rule{
		req.Phone: {rule.Required, rule.ValidPhone},
		req.Email: {rule.Required, rule.ValidEmail, rule.AtMost64Characters},
	}); err != nil {
		return err
	}

	return nil
}

func (s *Service) signUp(req *SignUpReqBody) (*SignUpResBody, error) {
	birthday, err := time.Parse("2006-01-02", req.Birthday)
	ok, err := s.db.InsertAccount(&mysqldb.Account{
		Phone:     req.Phone,
		Email:     req.Email,
		Password:  req.Password,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Birthday:  birthday,
		Gender:    req.Gender,
	})

	return &SignUpResBody{Success: ok}, err
}

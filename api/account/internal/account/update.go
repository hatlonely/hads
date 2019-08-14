package account

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hatlonely/account/internal/c"
	"github.com/hatlonely/account/internal/rule"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type UpdateReqBody struct {
	Token     string   `json:"token"`
	Field     string   `json:"field,omitempty"`
	Email     string   `json:"email,omitempty"`
	Phone     string   `json:"phone,omitempty"`
	FirstName string   `json:"firstName,omitempty"`
	LastName  string   `json:"lastName,omitempty"`
	Birthday  string   `json:"birthday,omitempty"`
	Password  string   `json:"password,omitempty"`
	Gender    c.Gender `json:"gender,omitempty"`
}

type UpdateResBody struct {
	OK  bool   `json:"ok"`
	Err string `json:"err"`
}

func (s *Service) Update(c *gin.Context) {
	rid := c.DefaultQuery("rid", NewToken())
	req := &UpdateReqBody{}
	var res *UpdateResBody
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

	if err = s.checkUpdateReqBody(req); err != nil {
		err = fmt.Errorf("check request body failed. body: [%v], err: [%v]", string(buf), err)
		WarnLog.WithField("@rid", rid).WithField("err", err).Warn()
		status = http.StatusBadRequest
		c.String(status, err.Error())
		return
	}

	res, err = s.update(req)
	if err != nil {
		WarnLog.WithField("@rid", rid).WithField("err", err).Warn("update failed")
		status = http.StatusInternalServerError
		c.String(status, err.Error())
		return
	}

	status = http.StatusOK
	c.JSON(status, res)
}

func (s *Service) checkUpdateReqBody(req *UpdateReqBody) error {
	if err := rule.Check(map[interface{}][]rule.Rule{
		req.Token: {rule.Required},
		req.Field: {rule.Required, rule.In(map[interface{}]struct{}{
			"phone": {}, "email": {}, "name": {}, "birthday": {}, "gender": {}, "password": {},
		})},
	}); err != nil {
		return err
	}

	switch req.Field {
	case "phone":
		return rule.Check(map[interface{}][]rule.Rule{
			req.Phone: {rule.Required, rule.ValidPhone},
		})
	case "email":
		return rule.Check(map[interface{}][]rule.Rule{
			req.Email: {rule.Required, rule.ValidEmail, rule.AtMost64Characters},
		})
	case "name":
		return rule.Check(map[interface{}][]rule.Rule{
			req.FirstName: {rule.Required},
			req.LastName:  {rule.Required},
		})
	case "birthday":
		return rule.Check(map[interface{}][]rule.Rule{
			req.Birthday: {rule.Required, rule.ValidBirthday},
		})
	case "gender":
		return rule.Check(map[interface{}][]rule.Rule{
			req.Gender: {rule.Required, rule.In(map[interface{}]struct{}{
				c.GenderUnknown: {}, c.Male: {}, c.Famale: {},
			})},
		})
	case "password":
		return rule.Check(map[interface{}][]rule.Rule{
			req.Birthday: {rule.Required, rule.AtLeast8Characters},
		})
	}

	return nil
}

func (s *Service) update(req *UpdateReqBody) (*UpdateResBody, error) {
	account, err := s.cache.GetAccount(req.Token)
	if err != nil {
		return nil, err
	}

	if account == nil {
		return &UpdateResBody{OK: false, Err: "会话已过期，请重新登录"}, nil
	}

	var ok bool
	switch req.Field {
	case "phone":
		ok, err = s.db.UpdateAccountPhone(account.ID, req.Phone)
	case "email":
		ok, err = s.db.UpdateAccountEmail(account.ID, req.Email)
	case "password":
		ok, err = s.db.UpdateAccountPassword(account.ID, req.Password)
	case "gender":
		ok, err = s.db.UpdateAccountGender(account.ID, req.Gender)
	case "name":
		ok, err = s.db.UpdateAccountName(account.ID, req.FirstName, req.LastName)
	case "birthday":
		birthday, _ := time.Parse("2006-01-02", req.Birthday)
		ok, err = s.db.UpdateAccountBirthday(account.ID, birthday)
	default:
		return &UpdateResBody{OK: false, Err: fmt.Sprintf("未知字段 [%v]", req.Field)}, nil
	}

	if err != nil {
		return nil, err
	}
	return &UpdateResBody{OK: ok}, nil
}

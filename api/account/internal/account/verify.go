package account

import (
	"encoding/json"
	"fmt"
	"github.com/hatlonely/account/internal/rule"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type VertifyReqBody struct {
	Field string `json:"field,omitempty"`
	Value string `json:"value,omitempty"`
}

type VertifyResBody struct {
	OK  bool   `json:"ok"`
	Tip string `json:"tip"`
}

func (s *Service) Vertify(c *gin.Context) {
	rid := c.DefaultQuery("rid", NewToken())
	req := &VertifyReqBody{}
	var res *VertifyResBody
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

	if err = s.checkVertifyReqBody(req); err != nil {
		err = fmt.Errorf("check request body failed. body: [%v], err: [%v]", string(buf), err)
		WarnLog.WithField("@rid", rid).WithField("err", err).Warn()
		status = http.StatusBadRequest
		c.String(status, err.Error())
		return
	}

	res, err = s.vertify(req)
	if err != nil {
		WarnLog.WithField("@rid", rid).WithField("err", err).Warn("vertify failed")
		status = http.StatusInternalServerError
		c.String(status, err.Error())
		return
	}

	status = http.StatusOK
	c.JSON(status, res)
}

func (s *Service) checkVertifyReqBody(req *VertifyReqBody) error {
	if err := rule.Check(map[interface{}][]rule.Rule{
		req.Field: {rule.Required, rule.In(map[interface{}]struct{}{"phone": {}, "email": {}, "username": {}})},
		req.Value: {rule.Required},
	}); err != nil {
		return err
	}

	return nil
}

func (s *Service) vertify(req *VertifyReqBody) (*VertifyResBody, error) {
	if req.Field == "phone" {
		account, err := s.db.SelectAccountByPhone(req.Value)
		if err != nil {
			return nil, err
		}
		if account == nil {
			return &VertifyResBody{OK: true}, nil
		}
		return &VertifyResBody{OK: false, Tip: "电话号码已存在"}, nil
	}

	if req.Field == "email" {
		account, err := s.db.SelectAccountByEmail(req.Value)
		if err != nil {
			return nil, err
		}
		if account == nil {
			return &VertifyResBody{OK: true}, nil
		}
		return &VertifyResBody{OK: false, Tip: "邮箱已存在"}, nil
	}

	if req.Field == "username" {
		account, err := s.db.SelectAccountByPhoneOrEmail(req.Value)
		if err != nil {
			return nil, err
		}
		if account == nil {
			return &VertifyResBody{OK: false, Tip: "账号不存在"}, nil
		}
		return &VertifyResBody{OK: true}, nil
	}

	return &VertifyResBody{OK: false, Tip: fmt.Sprintf("未知字段 [%v]", req.Field)}, nil
}

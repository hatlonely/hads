package account

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type SignInReqBody struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

type SignInResBody struct {
	Valid bool   `json:"valid"`
	Token string `json:"token"`
}

func (s *Service) SignIn(c *gin.Context) {
	rid := c.DefaultQuery("rid", NewToken())
	req := &SignInReqBody{}
	var res *SignInResBody
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

	if err = s.checkSignInReqBody(req); err != nil {
		err = fmt.Errorf("check request body failed. body: [%v], err: [%v]", string(buf), err)
		WarnLog.WithField("@rid", rid).WithField("err", err).Warn()
		status = http.StatusBadRequest
		c.String(status, err.Error())
		return
	}

	res, err = s.signIn(req)
	if err != nil {
		WarnLog.WithField("@rid", rid).WithField("err", err).Warn("signIn failed")
		status = http.StatusInternalServerError
		c.String(status, err.Error())
		return
	}

	if res.Valid {
		c.SetCookie("token", res.Token, 7*24*3600, "/", c.Request.Host, false, true)
	}

	status = http.StatusOK
	c.JSON(status, res)
}

func (s *Service) checkSignInReqBody(req *SignInReqBody) error {
	if req.Username == "" || req.Password == "" {
		return fmt.Errorf("username or password is null")
	}
	return nil
}

func (s *Service) signIn(req *SignInReqBody) (*SignInResBody, error) {
	account, err := s.db.SelectAccountByPhoneOrEmail(req.Username)
	if err != nil {
		return nil, err
	}

	if account == nil {
		return &SignInResBody{Valid: false}, nil
	}

	if account.Password != req.Password {
		return &SignInResBody{Valid: false}, nil
	}

	token := NewToken()
	if err := s.cache.SetAccount(token, account); err != nil {
		return nil, err
	}

	return &SignInResBody{Valid: true, Token: token}, nil
}

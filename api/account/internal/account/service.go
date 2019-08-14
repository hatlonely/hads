package account

import (
	"encoding/hex"
	"fmt"
	"github.com/hatlonely/account/internal/mysqldb"
	"github.com/hatlonely/account/internal/rediscache"
	"github.com/sirupsen/logrus"
	"math/rand"
	"time"
)

var InfoLog *logrus.Logger
var WarnLog *logrus.Logger
var AccessLog *logrus.Logger

func init() {
	InfoLog = logrus.New()
	WarnLog = logrus.New()
	AccessLog = logrus.New()
}

type Service struct {
	db    *mysqldb.MysqlDB
	cache *rediscache.RedisCache
}

func NewService(db *mysqldb.MysqlDB, cache *rediscache.RedisCache) *Service {
	return &Service{
		db:    db,
		cache: cache,
	}
}

func NewToken() string {
	buf := make([]byte, 32)
	token := make([]byte, 16)
	rand.New(rand.NewSource(time.Now().UnixNano())).Read(token)
	hex.Encode(buf, token)
	return string(buf)
}

func NewCode() string {
	return fmt.Sprintf("%06d", int(rand.NewSource(time.Now().UnixNano()).(rand.Source64).Uint64()%1000000))
}

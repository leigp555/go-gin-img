package global

import (
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"img/server/config"
)

var (
	Config *config.Config
	Mdb    *gorm.DB
	Rdb    *redis.Client
	Slog   *zap.SugaredLogger
	Glog   *zap.SugaredLogger
	Mlog   *zap.SugaredLogger
)

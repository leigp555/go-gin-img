package global

import (
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"img/server/config"
)

var (
	Config      config.Config
	Mydb        *gorm.DB
	Redb        *redis.Client
	SugarLog    *zap.SugaredLogger
	GinSugarLog *zap.SugaredLogger
	Logger      *zap.Logger
)

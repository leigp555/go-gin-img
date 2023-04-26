package global

import (
	es "github.com/elastic/go-elasticsearch/v7"
	"github.com/panjf2000/ants/v2"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"img/server/config"
)

var (
	Config  *config.Config
	Mdb     *gorm.DB
	Rdb     *redis.Client
	Elastic *es.Client
	Slog    *zap.SugaredLogger
	Glog    *zap.SugaredLogger
	Mlog    *zap.SugaredLogger
	Pool    *ants.Pool
)

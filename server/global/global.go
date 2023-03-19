package global

import (
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"server/config"
)

var (
	Config config.Config
	Log    *zap.Logger
	DB     *gorm.DB
	RDB    *redis.Client
	Viper  *viper.Viper
	Gpt    *config.Gpt
)

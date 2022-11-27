package global

import (
	"chains-gotest-backend/config"
	"chains-gotest-backend/utils/timer"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)
/**
Global Parameters:
1. REDIS
2. DATABASE
3. Viper Config
4. Timer
 */
var (
	GVA_DB     *gorm.DB
	GVA_REDIS  *redis.Client
	GVA_CONFIG config.Server
	GVA_VP     *viper.Viper
	//GVA_LOG    *oplogging.Logger
	GVA_LOG   *zap.Logger
	GVA_Timer timer.Timer = timer.NewTimerTask()
)

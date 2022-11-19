package g

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

var (
	Logger        *logrus.Logger
	stdFormatter  *prefixed.TextFormatter // 命令行输出格式
	fileFormatter *prefixed.TextFormatter // 文件输出格式
	//GV_DB     *gorm.DB
	//GV_REDIS  *redis.Client
	//GV_SERVER types.Server
	//GV_VP     *viper.Viper
	//GV_LOG    *ZapLogger
	//GV_MODEL  *types.BaseModel

	//GVA_Timer               timer.Timer = timer.NewTimerTask()
	//GVA_Concurrency_Control             = &singleflight.Group{}
	//
	//BlackCache local_cache.Cache
)

package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type SystemConfig struct {
	App   App
	Mysql Mysql
	Redis Redis
}

type App struct {
	RunEnv   string `yaml:"run_env"`
	HttpPort int    `yaml:"http_port"`
	LogFile  string `yaml:"log_file"`
}

type Mysql struct {
	Db         string `yaml:"db" default:"mysql"`
	DbHost     string `yaml:"db_host" default:"10.211.55.6"`
	DbPort     int    `yaml:"db_port" default:3306`
	DbUser     string `yaml:"db_user" defautl:"root"`
	DbPassword string `yaml:"db_password" default:"hnh666666"`
	DbName     string `yaml:"db_name" default:"ips"`
}

type Redis struct {
	Db       string `yaml:"db" default:"redis"`
	Host     string `yaml:"host" default:"10.211.55.6:6379"`
	Password string `yaml:"password"`
	Port     string `yaml:"port" default:"3"`
}

func InitConfig(cfgPath string) {

	info := fmt.Sprintf("加载配置文件:%s", cfgPath)
	vp := viper.New()
	vp.SetConfigFile(cfgPath)
	err := vp.ReadInConfig()
	if err != nil {
		//panic(fmt.Errorf("读取配置文件失败: %s \n", err))
		global.Logger.Errorf("%s,失败: %s \n", info, err)
		return
	}
	vp.WatchConfig()
	vp.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := vp.Unmarshal(&GConfig); err != nil {
			//panic(fmt.Sprintf("加载配置文件失败:%s", err.Error()))
			global.Logger.Errorf("%s,失败:%s", info, err.Error())
		}
	})
	if err := vp.Unmarshal(&GConfig); err != nil {
		global.Logger.Errorf("%s,失败:%s", info, err.Error())
	}
	global.Logger.Infof("%s,成功", info)

}

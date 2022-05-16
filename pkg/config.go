package pkg

import (
	"github.com/LianJianTech/lj-go-common/log"
	"github.com/fsnotify/fsnotify"
	"github.com/go-redis/redis"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

var (
	Cfg      *Config
	DB       *sqlx.DB
	RedisCli *redis.ClusterClient
)

const (
	YYYY_MM_DD          = "2006-01-02"
	YYYY_MM_DD_HH_MM_SS = "2006-01-02 15:04:05"
	HANDLE_ERROR        = "handle error"
)

type Config struct {
	Logger *LoggerConfig
}

func Init(cfgName string) {
	setConfig(cfgName)
	Cfg = loadConfig()
	initConfig(Cfg)
	watchConfig()
}

func setConfig(cfgName string) {
	if cfgName != "" {
		viper.SetConfigFile(cfgName)
	} else {
		viper.AddConfigPath("conf")
		viper.SetConfigName("config-local")
	}
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		panic("initConfig error")
	}
}

func loadConfig() *Config {
	cfg := &Config{
		Logger: LoadLoggerConfig(viper.Sub("logger")),
	}
	return cfg
}

func initConfig(cfg *Config) {
	cfg.Logger.InitLogger()
}

func watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Infof("Config file changed: %s", e.Name)
	})
}

package configs

import (
	"github.com/spf13/viper"
	"os"
)

type dBConfig struct {
	DBUrl      string `mapstructure:"DB_URL"`
	DBUsername string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBName     string `mapstructure:"DB_NAME"`
}

type rabbitMqConfig struct {
	Host         string `mapstructure:"RABBITMQ_HOST"`
	UserName     string `mapstructure:"RABBITMQ_USERNAME"`
	Password     string `mapstructure:"RABBITMQ_PASSWORD"`
	Port         int32  `mapstructure:"RABBITMQ_POR"`
	UseSSL       bool   `mapstructure:"RABBITMQ_USE_SSL"`
	SSLAlgorithm string `mapstructure:"RABBITMQ_SSL_ALGO"`
	CustPart     string `mapstructure:"RABBITMQ_CUST_PART_1"`
	VHost        string `mapstructure:"RABBITMQ_VHOST"`
}

type redisCacheConfig struct {
	ADDRESS  string `mapstructure:"CLUSTER_ADDRESS"`
	PASSWORD string `mapstructure:"CLUSTER_PASSWORD"`
	UseSSL   bool   `mapstructure:"CLUSTER_USE_SSL"`
}

type applicationConfiguration struct {
	EnableDB    bool `mapstructure:"APP_ENABLE_DB"`
	EnableCache bool `mapstructure:"APP_ENABLE_CACHE"`
}

var dbConfig dBConfig
var rabbitConfig rabbitMqConfig
var redisConfig redisCacheConfig

var appConfig applicationConfiguration

func init() {
	viper.SetConfigFile(os.Getenv("APP_HOME") + "/conf/application.properties")

	err := viper.ReadInConfig()
	if err != nil {
		println("cannot read cofiguration", err)
	}

	viper.SetDefault("TIMEZONE", "UTC")

	err = viper.Unmarshal(&dbConfig)
	if err != nil {
		println("db setting cant be loaded: ", err)
	}

	err = viper.Unmarshal(&rabbitConfig)
	if err != nil {
		println("rabbit setting cant be loaded: ", err)
	}

	err = viper.Unmarshal(&redisConfig)
	if err != nil {
		println("rabbit setting cant be loaded: ", err)
	}

	err = viper.Unmarshal(&appConfig)
	if err != nil {
		println("application setting cant be loaded: ", err)
	}
}

func GetDBConfigs() dBConfig {
	return dbConfig
}

func GetRabbitMqConfigs() rabbitMqConfig {
	return rabbitConfig
}

func GetRedisCacheConfigs() redisCacheConfig {
	return redisConfig
}

func GetAppConfig() applicationConfiguration {
	return appConfig
}

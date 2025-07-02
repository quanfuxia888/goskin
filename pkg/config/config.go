// pkg/config/config.go
package config

import (
	"fmt"
	v "github.com/spf13/viper"
	"os"
	"strings"
)

var Cfg *Config

type Config struct {
	App      AppConfig
	Log      LogConfig
	MySQL    MySQLConfig
	Redis    RedisConfig
	RabbitMQ RabbitMQConfig
	JWT      JWTConfig
}

type AppConfig struct {
	Name string `mapstructure:"name"`
	Mode string `mapstructure:"mode"`
	Port string `mapstructure:"port"`
}

type LogConfig struct {
	Level string `mapstructure:"level"`
	File  string `mapstructure:"file"`
}

type MySQLConfig struct {
	DSN string `mapstructure:"dsn"`
}

type RedisConfig struct {
	Addr     string `mapstructure:"addr"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

type RabbitMQConfig struct {
	URL string `mapstructure:"url"`
}

type JWTConfig struct {
	Secret        string `mapstructure:"secret"`
	AccessExpire  int    `mapstructure:"access_expire"`  // 分钟
	RefreshExpire int    `mapstructure:"refresh_expire"` // 分钟
}

func Init(path ...string) {

	// 默认路径为 ./configs/config.yaml
	configPath := "configs/config.yaml"
	if len(path) > 0 && path[0] != "" {
		configPath = path[0]
	}

	// 设置文件名与路径
	v.SetConfigFile(configPath)
	v.SetConfigType("yaml")
	v.AutomaticEnv()                                   // 支持环境变量覆盖
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_")) // 支持 APP_PORT => app.port

	if err := v.ReadInConfig(); err != nil {
		fmt.Printf("读取配置文件失败: %v\n", err)
		os.Exit(1)
	}

	// 映射到结构体
	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		fmt.Printf("解析配置失败: %v\n", err)
		os.Exit(1)
	}

	Cfg = &cfg
	fmt.Println("✅ 配置加载成功")
}

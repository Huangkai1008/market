package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

// Config 应用配置
type Config struct {
	RunMode string
	Log
	Jwt
	HTTP
	Gorm
	Database
}

// New 返回一个配置实例
func New() (*Config, error) {
	var (
		err    error
		config *Config
	)

	v := viper.New()
	v.AddConfigPath("configs")
	v.SetConfigType("toml")
	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	v.SetConfigFile(".env")
	if err := v.MergeInConfig(); err != nil {
		return nil, err
	}

	if err = v.Unmarshal(&config); err != nil {
		return nil, err
	}
	return config, err
}

// Log 日志配置参数
type Log struct {
	Level    int
	FileName string
}

// Jwt JWT配置参数
type Jwt struct {
	SecretKey         string
	JwtExpireDuration time.Duration
	JwtIssuer         string
}

// HTTP http配置参数
type HTTP struct {
	HttpHost     string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

// Gorm Gorm配置参数
type Gorm struct {
	MaxIdleConnections int
	MaxOpenConnections int
	EnableAutoMigrate  bool
}

// Database 数据库配置参数
type Database struct {
	DBType     string
	User       string
	Password   string
	Host       string
	Port       int
	DBName     string
	Parameters string
}

// DSN 数据库连接串
func (d Database) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s",
		d.User, d.Password, d.Host, d.Port, d.DBName, d.Parameters)
}

package config

import (
	"bufio"
	"encoding/json"
	"os"
)

var (
	GlobalConfig AppConfig
)

type AppConfig struct {
	Config         Config
	DatabaseConfig DatabaseConfig
	RedisConfig    RedisConfig
	JwtConfig      JwtConfig
}
type Config struct {
	AppName     string         `json:"app_name"`
	AppModel    string         `json:"app_model"`
	AppHost     string         `json:"app_host"`
	AppPort     string         `json:"app_port"`
	Database    DatabaseConfig `json:"database"`
	RedisConfig RedisConfig    `json:"redis_config"`
	LogLevel    string         `json:"log_level"`
	JwtConfig   JwtConfig      `json:"Jwt_config"`
}

// 数据库配置
type DatabaseConfig struct {
	Driver   string `json:"driver"`
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	DbName   string `json:"db_name"`
	Chartset string `json:"charset"`
	ShowSql  bool   `json:"show_sql"`
}

// Redis配置
type RedisConfig struct {
	Addr     string `json:"addr"`
	Port     int    `json:"port"`
	Password string `json:"password"`
	Db       int    `json:"db"`
}
type JwtConfig struct {
	JwtSecretKey *string
	PassSalt     *string
}

// 获取配置，外部使用"config.GetConfig()"调用
func GetConfig() *Config {
	return cfg
}

// 存储配置的全局对象
var cfg *Config = nil

func ParseConfig(path string) (*Config, error) {
	file, err := os.Open(path) //读取文件
	defer file.Close()
	if err != nil {
		return nil, err
	}
	reader := bufio.NewReader(file)
	decoder := json.NewDecoder(reader) //解析json
	if err = decoder.Decode(&cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}

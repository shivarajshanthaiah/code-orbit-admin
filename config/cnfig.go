package config

import "github.com/spf13/viper"

type Config struct {
	SECRETKEY string `mapstructure:"JWTSECRET"`
	Host      string `mapstructure:"HOST"`
	User      string `mapstructure:"USER"`
	Password  string `mapstructure:"PASSWORD"`
	Database  string `mapstructure:"DBNAME"`
	Port      string `mapstructure:"PORT"`
	Sslmode   string `mapstructure:"SSL"`
	GrpcPort  string `mapstructure:"GRPCPORT"`
	ProblemPort string `mapstructure:"GRPCPROBLEMPORT"`
	UserPort string `mapstructure:"GRPCUSERPORT"`
}

func LoadConfig() *Config {
	var config Config
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	viper.Unmarshal(&config)
	return &config
}

package config

import "github.com/spf13/viper"

type Config struct {
	DBName    string `mapstructure:"DB_NAME"`
	AppName   string `mapstructure:"APP_NAME"`
	IsResetDB bool   `mapstructure:"IS_RESET_DB"`
	MaxProcs  int    `mapstructure:"MAX_PROCS"`
}

func InitConfig(configPath string) (config Config, err error) {
	viper.AddConfigPath(configPath)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	if err = viper.ReadInConfig(); err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}

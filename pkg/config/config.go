package config

import "github.com/spf13/viper"

type Config struct {
	Port string `mapstructure:"PORT"`
	Dburl string `mapstructure:"DB_URL"`
	ProductSvc string `mapstructure:"PRODUCT_SVC_URL"`
}

func Loadconfig()(config Config, err error){
	viper.AddConfigPath("./pkg/config/envs")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)

	return
}
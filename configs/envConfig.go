package configs

import (
	"log"

	"github.com/spf13/viper"
)

type EnvConfig struct{
Port string `mapstructure:"PORT"`

}

func Envload()(config *EnvConfig){
	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	if err:=viper.ReadInConfig();err!=nil{
		log.Fatal("error loading config file",err)
	}
	if err:=viper.Unmarshal(&config);err!=nil{
		log.Fatal("error unmarshaling config file",err)
	}

	return config
}
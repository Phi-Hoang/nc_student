package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

var Config ConfigSchema

func init() {
	viper.SetConfigName("config") // name of config file (without extension)
	//viper.AddConfigPath("/etc/appname/")  // path to look for the config file in
	//viper.AddConfigPath("$HOME/.appname") // call multiple times to add many search paths
	//viper.AddConfigPath(".")              // optionally look for config in the working directory
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	err = viper.Unmarshal(&Config)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
}

type ConfigSchema struct {
	Mongo struct {
		URI      string `mapstructure:"Uri"`
		Host     string `mapstructure:"HostName"`
		Username string `mapstructure:"UserName"`
		Password string `mapstructure:"Password"`
	} `mapstructure:"MongoDB"`
	JWTSecret struct {
		JWTKey string
	}
}

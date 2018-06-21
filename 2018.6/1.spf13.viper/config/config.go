package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var TestData string

func Setconfig() {
	// 1. config file
	viper.SetConfigName("config/config")  // name of config file (without extension)
	viper.AddConfigPath("/etc/appname/")  // path to look for the config file in
	viper.AddConfigPath("$HOME/.appname") // call multiple times to add many search paths
	viper.AddConfigPath(".")              // optionally look for config in the working directory
	err := viper.ReadInConfig()           // Find and read the config file
	if err != nil {                       // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	result := viper.Get("websites")
	fmt.Println(result)
	// 2. env
	viper.SetEnvPrefix("spf") // will be uppercased automatically
	viper.BindEnv("id")

	id := viper.Get("id") // 13
	if sid, ok := id.(string); ok {
		fmt.Println(sid)
	}
}
func init() {
	TestData = "dd"
	fmt.Println("config package init")
}

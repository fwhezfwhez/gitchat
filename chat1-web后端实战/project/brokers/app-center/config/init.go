package config

import (
	"flag"
	"fmt"
	"github.com/spf13/viper"
)

var Cfg *viper.Viper
var configPath string
var Mode string

func init() {
	Cfg = viper.New()
	flag.StringVar(&Mode, "mode", "local", "go run main.go -mode dev")
	flag.Parse()
	switch Mode {
	case "dev","":
		configPath = "/home/tom/projects/gitchat/config/dev.yaml"
	case "pro":
		configPath= "/home/tom/projects/gitchat/config/pro.yaml"
	case "local":
		configPath = "G:\\go_workspace\\GOPATH\\src\\gitchat\\chat1-web后端实战\\project\\brokers\app-center\\config\\local.yaml"
	}

	Cfg.SetConfigFile(configPath)
	if e := Cfg.ReadInConfig(); e != nil {
		panic(e)
	}
	fmt.Println(Cfg.GetString("title"))
}

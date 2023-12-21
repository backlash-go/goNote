package main

import (
	"flag"
	"fmt"
	"github.com/spf13/viper"
	"log"
)

func main() {

	configFile := flag.String("conf", "./app-dev.yaml", "path of config file")
	flag.Parse()
    fmt.Println(*configFile)
	viper.SetConfigFile(*configFile)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("viper read config is failed, err is")
	}

	fmt.Println(viper.GetStringMapString("database")["host"])

}

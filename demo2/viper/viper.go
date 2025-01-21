package main

import (
	"fmt"
	"github.com/spf13/viper"
)

func main() {
	viper.AutomaticEnv()
	fmt.Println(viper.Get("NODE_OPTIONS"))

}

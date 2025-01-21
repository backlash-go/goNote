package main

import (
	"fmt"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func main() {
	p := new(string)
	p2 := new(string)
	pflag.StringVar(p, "name", "xxb", "input your name buddy")
	pflag.StringVar(p2, "email", "", "input your email buddy")
	pflag.Parse()
	fmt.Printf("name: %v\n", *p)
	fmt.Printf("email: %v\n", *p2)

	viper.AddConfigPath("./")

	viper.SetConfigType("yaml")
	viper.SetConfigName("app-dev")
	if err := viper.ReadInConfig() ; err  != nil {
		fmt.Println(err)
	}
	fmt.Println(viper.GetString("address"))









}

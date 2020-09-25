package main

import (
	"flag"
	"fmt"
)

func main() {

	//period := flag.Duration("period", 1 * time.Second, "period sleep")
	////flag.Parse()
	//fmt.Printf("Sleeping for %v...\n", *period)
	//time.Sleep(*period)
	//fmt.Println("hello")

	//读取yaml文件
	//go run flag -conf app-dev.yaml
	//var configFile = flag.String("conf", "./app-dev.yaml", "this is a species")
	//flag.Parse()
	//fmt.Println(*configFile)
	//
	//viper.SetConfigFile(*configFile)
	//
	//if err := viper.ReadInConfig(); err != nil {
	//	log.Fatalf("viper read config is failed, err is %v configFile is %s ", err, configFile)
	//}
	//
	//db := viper.GetStringMapString("database")
	//fmt.Println(db["name"])
	//
	//hostList := viper.GetStringMapStringSlice("hostSlice")
	//fmt.Println(hostList["hostlist"][0])
	//fmt.Println(hostList["datalist"][1])

	var gopherType string
	const (
		defaultGopher = "pocket"
		usage         = "the variety of gopher"
	)
	flag.StringVar(&gopherType, "gopher_type", defaultGopher, usage)
	flag.StringVar(&gopherType, "g", defaultGopher, usage+" (shorthand)")
	flag.Parse()

//go run flag.go -g aaa
//go run flag.go -g
//go run flag.go -gopher_type
	fmt.Println(gopherType)

}

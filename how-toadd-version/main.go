package main

import (
	"flag"
	"fmt"
)

var (
	GitVersion = "v0.0.0-master+$Format:%h$"
	BuildDate = "1970-01-01T00:00:00Z"
)

func main() {
	version := flag.Bool("version", false, "Print version info.")
	flag.Parse()

	if *version {
		fmt.Println("GitVersion", GitVersion)
		fmt.Println("BuildDate", BuildDate)
	}
	fmt.Println("ok")
}


//go build -ldflags "-X main.GitVersion=v1.0.0 -X main.BuildDate=$(date +%F)" -o version main.go
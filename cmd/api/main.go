package main

import (
	"flag"
	"fmt"

	"sb_social_network/internal/server"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/test.json", "path to config file")
}

func main() {
	flag.Parse()

	config, err := server.NewConfig(configPath)
	if err != nil {
		fmt.Println(err)
	}

	if err := server.Start(config); err != nil {
		fmt.Println(err)
	}
}

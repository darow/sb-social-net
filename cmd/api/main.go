package main

import (
	"fmt"
	"sb_social_network/internal/server"
)

func main() {
	if err := server.Start(); err != nil {
		fmt.Println(err)
	}
}

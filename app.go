package main

import (
	"fmt"
	"github.com/meifif/taboola-home-test/config"
	"os"
)

func main() {
	fmt.Println("starting app")
	conf, err := config.InitConfig()
	if err != nil {
		fmt.Println("missing conf file... exit!")
		os.Exit(1)
	}
	fmt.Println(conf)

}

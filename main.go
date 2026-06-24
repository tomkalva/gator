package main

import (
	"fmt"
	"gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println(err)
		return
	}
	err = cfg.SetUser("Tom")
	if err != nil {
		fmt.Println(err)
		return
	}

	cfg, err = config.Read()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", cfg)
}

package main

import (
	"fmt"

	"github.com/mskydream/aport/config"
	"github.com/mskydream/aport/db"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Println("config error")
	}

	db := new(db.DB).InitDatabase(&cfg)
	fmt.Println(db)
}

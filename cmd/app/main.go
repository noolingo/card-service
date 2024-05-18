package main

import (
	"fmt"
	"log"
	"os"

	"github.com/noolingo/card-service/internal/app"
)

const configPath = "configs/config.yml"

func main() {
	err := app.Run(configPath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("card service work!!!")
	os.Exit(0)
}

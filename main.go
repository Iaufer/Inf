package main

import (
	"Inf/internal/app/apiserver"
	"flag"
	"fmt"
	"log"

	"github.com/BurntSushi/toml"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "../../configs/apiserver.toml", "path to config file")
}

func main() {
	flag.Parse()

	config := apiserver.NewConfig()

	_, err := toml.DecodeFile(configPath, config)

	fmt.Println(config)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Start...")

	if err := apiserver.Start(config); err != nil {
		fmt.Println("Тут ошибка")

		log.Fatal(err)
	}

}

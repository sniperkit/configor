package main

import (
	"fmt"

	// internal
	"github.com/sniperkit/snk.golang.configor/pkg/configor"
)

var Config = struct {
	APPName string `default:"app name"`

	DB struct {
		Name     string
		User     string `default:"root"`
		Password string `required:"true" env:"DBPassword"`
		Port     uint   `default:"3306"`
	}

	Contacts []struct {
		Name  string
		Email string `required:"true"`
	}
}{}

func main() {
	configor.Load(&Config, "../../shared/config/config-example.yml")
	fmt.Printf("config: %#v", Config)
}

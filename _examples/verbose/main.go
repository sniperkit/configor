package main

import (
	"fmt"

	// external
	"github.com/k0kubun/pp"

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

	// Enable debug mode or set env `CONFIGOR_DEBUG_MODE` to true when running your application
	configor.New(&configor.Config{Verbose: true}).Load(&Config, "../../shared/config/config-example.yml")

	pp.Println("config:", Config)
}

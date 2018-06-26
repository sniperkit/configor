package main

import (
	// external
	"github.com/k0kubun/pp"

	// internal
	configor "github.com/sniperkit/snk.golang.configor/pkg/configor/v1"
)

var configFiles = []string{
	"../../shared/config/config.example.yml"
}

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
	configor.New(&configor.Config{Debug: true}).Load(&Config, configFiles...)

	pp.Println("config:", Config)
}

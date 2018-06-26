package main

import (
	// external
	"github.com/k0kubun/pp"

	// internal
	configor "github.com/sniperkit/snk.golang.configor/pkg/configor/v1"
)

// Earlier configurations have higher priority !!!
// Available formats: YAML, JSON and TOML
// load mixed config files (formats: .yaml/.yml, .json and .toml)
// you can change other extensions as all config filenames have been converted to all 3 formats
var configFiles = []string{
	"../../shared/config/yaml/advanced.example.yml", // yaml format
	// "../../shared/config/json/server.example.json",  // json format
}

var Config = struct {
	App      AppConfig     `json:"application" yaml:"application" toml:"application"`
	Server   ServerConfig  `json:"server" yaml:"server" toml:"server"`
	Database DBConfig      `json:"database" yaml:"database" toml:"database"`
	Service  ServiceConfig `json:"service" yaml:"service" toml:"service"`
}{}

func main() {

	configor.New(&configor.Config{
		Environment:          "../../shared/config/.env-example",
		ENVPrefix:            "",
		Verbose:              true,
		Debug:                true,
		ErrorOnUnmatchedKeys: false,
	}).Load(&Config, configFiles...)

	pp.Println("Config:", Config)
}

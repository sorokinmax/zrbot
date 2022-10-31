package main

import (
	"log"
	"os"

	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v2"
)

// Config struct
type Config struct {
	Zrbot struct {
		Login         string `yaml:"login"`
		Password      string `yaml:"password"`
		ZabbixRootURL string `yaml:"zabbix_root_url"`
	} `yaml:"zrbot"`
	Reportlinks struct {
		Daily  string `yaml:"daily"`
		Weekly string `yaml:"weekly"`
	} `yaml:"reportlinks"`
	SMTP struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
		User string `yaml:"user"`
		Pass string `yaml:"pass"`
		From string `yaml:"from"`
	} `yaml:"smtp"`
}

func ReadConfigFile(cfg *Config, file string) {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
	if err != nil {
		log.Fatal(err)
	}
}

func ReadConfigEnv(cfg *Config) {
	err := envconfig.Process("", cfg)
	if err != nil {
		log.Fatal(err)
	}
}

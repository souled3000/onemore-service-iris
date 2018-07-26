package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	AppName  string `yaml:"appname,omitempty"`
	HttpPort string `yaml:"http_port,omitempty"`
}

var Conf Config

func InitConfig() {
	yamlFile, err := ioutil.ReadFile("config/config.yaml")
	if err != nil {
		fmt.Println(err.Error())
	}
	err = yaml.Unmarshal(yamlFile, &Conf)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(Conf)
}

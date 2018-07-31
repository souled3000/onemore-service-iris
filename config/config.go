package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	AppName  string `yaml:"appname,omitempty"`
	HttpPort string `yaml:"http_port,omitempty"`
	Mysql    Mysql  `yaml:"mysql,omitempty"`
	Redis    Redis  `yaml:"redis,omitempty"`
}

type Redis struct {
	Addr      string `yaml:"addr,omitempty"`
	Pwd       string `yaml:"pwd,omitempty"`
	MaxActive int    `yaml:"maxActive,omitempty"`
	Idle      int    `yaml:"idle,omitempty"`
}

type Mysql struct {
	Host     string `yaml:"host,omitempty"`
	Port     string `yaml:"port,omitempty"`
	UserName string `yaml:"user_name,omitempty"`
	Password string `yaml:"password,omitempty"`
	DBName   string `yaml:"db_name,omitempty"`
}

var Conf Config

func init() {
	InitConfig()
}
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

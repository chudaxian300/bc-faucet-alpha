package conf

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type conf struct {
	Server server `yaml:"server"`
	Db     db     `yaml:"db"`
	MyLog  myLog  `yaml:"myLog"`
	Cache  cache  `yaml:"cache"`
	Geth   geth   `yaml:"geth"`
}

type server struct {
	Address string `yaml:"address"`
	Model   string `yaml:model`
}

type db struct {
	Dialects string `yaml:"dialects"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Db       string `yaml:"db"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Charset  string `yaml:"charset"`
	MaxIdle  int    `yaml:maxIdle`
	MaxOpen  int    `yaml:"mxOpen"`
}

type myLog struct {
	Path  string `yaml:"path"`
	Name  string `yaml:"name"`
	Model string `yaml:"model"`
}

type cache struct {
	Expire int `yaml:"expire"`
	Clear  int `yaml:"clearUp"`
}

type geth struct {
	Url string `yaml:"url"`
}

var Conf *conf

func init() {
	yamlFile, err := ioutil.ReadFile("./conf.yaml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(yamlFile, &Conf)
	if err != nil {
		panic(err)
	}
}

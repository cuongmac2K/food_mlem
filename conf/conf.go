package conf

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var MYSQL_CONFIG map[string]string

const Path = "conf/config.yaml"

var str = "READ FILE CONFIG SUCCESS.\n"

type ConfigFile struct {
	MONGODB_CONFIG struct {
		ENDPOINT string `yaml:"ENDPOINT"`
		DATABASE string `yaml:"DATABASE"`
	} `yaml:"MONGODB_CONFIG"`
}

var Config *ConfigFile

func init() {
	buf, err := ioutil.ReadFile(Path)
	if err != nil {
		fmt.Println("READ FILE ERROR: " + err.Error())
		panic(err)
	}

	c := &ConfigFile{}
	err = yaml.Unmarshal(buf, c)
	if err != nil {
		fmt.Println("READ FILE ERROR: " + err.Error())
		panic(err)
	}
	Config = c
	fmt.Print(str)
	str = ""
}

package internal

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"

	"gopkg.in/yaml.v2"
)

type Config struct {
	ListId string `yaml:"listid"`
}

func LoadConfig() *Config {
	var configDir string
	home := os.Getenv("HOME")
	if home == "" && runtime.GOOS == "windows" {
		configDir = os.Getenv("APPDATA")
	} else {
		configDir = filepath.Join(home, ".config")
	}
	fname := filepath.Join(configDir, "gtodo", "config.yaml")
	f, err := ioutil.ReadFile(fname)
	if err != nil {
		fmt.Println(err)
	}
	config := &Config{}
	err = yaml.Unmarshal(f, &config)
	if err != nil {
		fmt.Println(err)
	}
	return config
}

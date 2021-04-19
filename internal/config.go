package internal

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"

	"gopkg.in/yaml.v2"
)

// type Config struct {
// 	Gtodo Gtodo `yaml:gtodo`
// }

// type Gtodo struct {
// 	ListId string `yaml:"listid"`
// }

// func Load() (*Config, error) {
// 	viper.SetConfigName("config")
// 	viper.SetConfigType("yaml")
// 	viper.AddConfigPath(filepath.Join("$HOME", ".config", "gtodo"))

// 	err := viper.ReadInConfig
// 	if err != nil {
// 		return nil, fmt.Errorf("can't read config file: %s \n", err)
// 	}

// 	var cfg Config

// }

// func ReadDefaultListID() string {
// 	cfg := Configure()
// 	return cfg.gtodo.listid
// }
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

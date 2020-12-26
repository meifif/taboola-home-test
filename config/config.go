package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Config struct {
	DBUser     string `json:"DB_USER"`
	DBPassword string `json:"DB_PASSWORD"`
	DBHost     string `json:"DB_HOST"`
	DBName     string `json:"DB_NAME"`
}

const configFile = "conf.json"

func InitConfig() (Config, error) {
	fmt.Println("loading conf file")
	fmt.Println("opening conf file")
	jsonFile, err := os.Open("config/" + configFile)
	if err != nil {
		fmt.Println(err)
		return Config{}, err
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var conf Config
	if err := json.Unmarshal(byteValue, &conf); err != nil {
		fmt.Println(err)
		return Config{}, err
	}
	fmt.Println("conf file loaded")
	return conf, nil
}

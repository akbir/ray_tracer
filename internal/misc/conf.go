package misc

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Dimensions DimensionsConfig `json:"dimensions"`
	World      []ObjectConfig   `json:"world"`
}

type DimensionsConfig struct {
	Height   int `json:"height"`
	Width    int `json:"width"`
	Aliasing int `json:"aliasing"`
}

type ObjectConfig struct {
	Center   []float64      `json:"center"`
	Radius   float64        `json:"radius"`
	Material materialConfig `json:"material"`
}

type materialConfig struct {
	Material string    `json:"type"`
	Color    []float64 `json:"color"`
	Property float64   `json:"property"`
}

func GetConfig() Config {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("error:", err)
	}
	file, err := os.Open(dir + "/configs/conf.json")
	defer file.Close()

	decoder := json.NewDecoder(file)
	configuration := Config{}
	err = decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}
	return configuration

}

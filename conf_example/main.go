package main

import (
	"bytes"
	"fmt"
	"io/ioutil"

	"github.com/spf13/viper"
)

func main() {

	viper.SetConfigType("yaml") // or viper.SetConfigType("YAML")
	yamlExample, _ := ioutil.ReadFile("./test.yaml")

	// any approach to require this configuration into your program.

	err := viper.ReadConfig(bytes.NewBuffer(yamlExample))
	if err != nil {
		fmt.Println(err)
	}

	// fmt.Println(viper.Get("eyes"))

	yey := viper.Get("eyes")
	for _, key := range yey.([]interface{}) {
		fmt.Println(key.(string))
	}
}

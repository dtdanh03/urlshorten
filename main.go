package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

//RedirectionMap type
type RedirectionMap map[string][]string

func main() {
	// m := make(map[string]string)
	// m["something"] = "value"
	// fmt.Println(m)
	filename, _ := filepath.Abs("./map.yml")
	yamlFile, error := ioutil.ReadFile(filename)

	if error != nil {
		fmt.Println("Error: --")
		fmt.Println(error)
	} else {
		var redirectionMap RedirectionMap
		// m := make(map[interface{}]interface{})
		err := yaml.Unmarshal(yamlFile, &redirectionMap)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(redirectionMap["apple"])
		}
	}

}

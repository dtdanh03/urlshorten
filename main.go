package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

//RedirectionMap type
type RedirectionMap map[string][]string

func main() {

	pattern := flag.String("a", "", "Pattern")
	destination := flag.String("u", "", "Full destination URL")
	list := flag.Bool("l", false, "List all redirection map")
	help := flag.Bool("h", false, "Help")

	flag.Usage = func() {
		fmt.Printf("Usage: %s [options] [<dir>]\nOptions are:\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()
	if *help {
		flag.Usage()
	}

	fmt.Print(*pattern, *destination, *list)

	filename, _ := filepath.Abs("./map.yml")
	yamlFile, error := ioutil.ReadFile(filename)

	if error != nil {
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

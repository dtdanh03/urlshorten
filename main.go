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

	list := flag.Bool("l", false, "List all redirection map")
	help := flag.Bool("h", false, "Help")
	remove := flag.String("d", "", "Remove from the list")

	configure := flag.NewFlagSet("configure", flag.ContinueOnError)
	pattern := configure.String("a", "", "Pattern")
	destination := configure.String("u", "", "Full destination URL")
	configureHelp := configure.Bool("h", false, "Help")

	flag.Parse()

	if *list {
		listMap()
		os.Exit(0)
	}

	if *remove != "" {
		removeFromMap(*remove)
		os.Exit(0)
	}

	if *help {
		fmt.Println("Options:")
		flag.PrintDefaults()
		fmt.Println("Use `configure` subcommand to add to map")
	}

	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "configure":
			handleConfigureCommand(configure, pattern, destination, configureHelp)
		case "run":
			//start http server
		}
	}
}

func getRedirectionMap() RedirectionMap {
	filename, _ := filepath.Abs("./map.yml")
	yamlFile, readFileError := ioutil.ReadFile(filename)
	if readFileError != nil {
		return makeEmptyRedirectionMap()
	}

	var redirectionMap RedirectionMap
	parsingError := yaml.Unmarshal(yamlFile, &redirectionMap)
	if parsingError != nil {
		return makeEmptyRedirectionMap()
	}
	return redirectionMap
}

func saveRedirectionMap(redirectionMap RedirectionMap) {
	data, error := yaml.Marshal(&redirectionMap)
	if error != nil {
		fmt.Println(error)
		return
	}
	ioutil.WriteFile("./map.yml", data, 0644)
}

func makeEmptyRedirectionMap() RedirectionMap {
	return make(map[string][]string)
}

func listMap() {
	redirectionMap := getRedirectionMap()
	for key, values := range redirectionMap {
		fmt.Printf("%s:\n", key)
		for _, value := range values {
			fmt.Println("--", value)
		}
	}

}

func removeFromMap(pattern string) {
	redirectionMap := getRedirectionMap()
	for key, values := range redirectionMap {
		for index, value := range values {
			if pattern == value {
				slice := redirectionMap[key]
				redirectionMap[key] = append(slice[:index], slice[index+1:]...)
				if len(redirectionMap[key]) == 0 {
					delete(redirectionMap, key)
				}
				break
			}
		}
	}
	saveRedirectionMap(redirectionMap)

}

func handleConfigureCommand(flagSet *flag.FlagSet, pattern *string, destination *string, help *bool) {
	err := flagSet.Parse(os.Args[2:])
	if err != nil {
		return
	}
	if *help {
		flagSet.PrintDefaults()
		os.Exit(0)
	}

	if *pattern == "" && *destination == "" {
		return
	}

	redirectionMap := getRedirectionMap()
	slice := redirectionMap[*destination]
	if slice == nil {
		redirectionMap[*destination] = []string{*pattern}
	} else {
		redirectionMap[*destination] = append(slice, *pattern)
	}
	saveRedirectionMap(redirectionMap)
}

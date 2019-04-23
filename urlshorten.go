package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

//RedirectionMap type
type RedirectionMap map[string]string

func main() {

	list := flag.Bool("l", false, "List all redirection map")
	help := flag.Bool("h", false, "Help")
	remove := flag.String("d", "", "Remove from the list")

	configure := flag.NewFlagSet("configure", flag.ContinueOnError)
	pattern := configure.String("a", "", "Pattern")
	destination := configure.String("u", "", "Full destination URL")
	configureHelp := configure.Bool("h", false, "Help")

	run := flag.NewFlagSet("run", flag.ContinueOnError)
	port := run.String("p", "8080", "Port")

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
			handleRunCommand(run, port)
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
	return make(map[string]string)
}

func listMap() {
	redirectionMap := getRedirectionMap()
	for key, value := range redirectionMap {
		fmt.Printf("%s:\n", key)
		fmt.Println("--", value)
	}

}

func removeFromMap(pattern string) {
	redirectionMap := getRedirectionMap()
	delete(redirectionMap, pattern)
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
	redirectionMap[*pattern] = *destination
	saveRedirectionMap(redirectionMap)
}

func handleRunCommand(runFlagSet *flag.FlagSet, port *string) {

	err := runFlagSet.Parse(os.Args[2:])
	if err != nil {
		return
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Welcome to my website!")
	})

	formattedPortString := fmt.Sprintf(":%s", *port)
	http.ListenAndServe(formattedPortString, nil)
}

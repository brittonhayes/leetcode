package main

import (
	"github.com/BurntSushi/toml"
	"io/ioutil"
	"log"
	"os"
	"text/template"
)

type Readme struct {
	Title       string
	Subtitle    string
	Description string
	ImageURL    string
	ImageSize   int
	Sections    map[string]Section
	Packages    map[string]Package
}

type Section struct {
	Title string
	Body  string
	Emoji string
}

type Package struct {
	Title       string
	Description string
	Path        string
}

func main() {

	var conf Readme

	// Create the readme file
	f, err := os.Create("README.md")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Read in README template
	readme, err := fileToString("README.tpl")
	if err != nil {
		log.Fatal("Error reading README template: ", err)
	}

	// Read in config
	confData, err := fileToString("config.toml")
	if err != nil {
		log.Fatal("Error reading config file: ", err)
	}

	// Put the data from the config into the conf struct
	if _, err := toml.Decode(confData, &conf); err != nil {
		log.Fatal("Error decoding toml info struct: ", err)
	}

	// Execute the template
	t, err := template.New("README").Parse(readme)
	if err != nil {
		panic(err)
	}

	err = t.Execute(f, conf)
	if err != nil {
		panic(err)
	}

}

func fileToString(filename string) (string, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

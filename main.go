package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Configurations represents the main content block of c_cpp_properties.json
type Configurations struct {
	Configurations []Configuration `json:"configurations"`
}

// Configuration represents a single configuration inside Configurations
type Configuration struct {
	CompilerPath     string   `json:"compilerPath"`
	CStandard        string   `json:"cStandard"`
	IntelliSenseMode string   `json:"intelliSenseMode"`
	Name             string   `json:"name"`
	Defines          []string `json:"defines"`
	IncludePath      []string `json:"includePath"`
}

func main() {
	args := os.Args[1:]

	jsonFile, err := os.Open(args[0])

	if err != nil {
		fmt.Println(err)
		return
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result Configurations
	var newIncludes []string
	json.Unmarshal([]byte(byteValue), &result)

	for _, path := range result.Configurations[0].IncludePath {
		if _, err := os.Stat(path); os.IsNotExist(err) {
		} else {
			newIncludes = append(newIncludes, path)
		}
	}

	result.Configurations[0].IncludePath = newIncludes
	newJSON, err := json.Marshal(result)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(newJSON))
}
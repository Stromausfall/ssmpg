package input

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type ConfigData struct {
	Topbar    string
	Bottombar string
}

func CreateConfigData(sourcefile string) ConfigData {
	value := ConfigData{}

	data, readFileError := ioutil.ReadFile(sourcefile)

	if readFileError != nil {
		fmt.Printf(fmt.Sprintf("%v", readFileError))
		panic("error while opening configFile")
	}

	unmarshallError := yaml.Unmarshal(data, &value)

	if unmarshallError != nil {
		fmt.Printf(fmt.Sprintf("%v", unmarshallError))
		panic("error while unmarshaling configFile")
	}

	return value
}

package input

import (
	"io/ioutil"
	"encoding/json"
)

type ConfigFileData struct {
	TopBar string
	BottomBar string
}

func CreateConfigFile(sourcefile string) *ConfigFileData {
	value := &ConfigFileData{}
	
	data, _ := ioutil.ReadFile(sourcefile)
	stringData := string(data)
	
	json.Unmarshal([]byte(stringData), value)
	
	return value
}

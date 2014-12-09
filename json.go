package goconfigure

import (
	"encoding/json"
	"io/ioutil"
)

func ReadJsonConfigFromFile(fileName string, o *interface{}) error {
	fileRead, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}

	return ReadJsonConfig(fileRead, o)
}

func ReadJsonConfig(fileRead []byte, o *interface{}) error {
	err := json.Unmarshal(fileRead, o)
	if err != nil {
		return err
	}

	return nil
}

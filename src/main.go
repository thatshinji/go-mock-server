package src

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func ParsePath() (path string, err error) {
	pwd, _ := os.Getwd()
	fileList, err := ioutil.ReadDir(pwd)
	if err != nil {
		log.Fatal(err)
	}
	var curJsonFile string
	for _, v := range fileList {
		name := v.Name()
		if isJson := strings.Contains(name, "json"); isJson {
			curJsonFile = name
		}
	}
	if curJsonFile == "" {
		return curJsonFile, errors.New("dont have json file")
	}
	absPath := pwd + "/" + curJsonFile
	return absPath, nil
}

func ReadJSON(path string) ([]byte, error) {
	jsonFile, err := os.Open(path)
	if err != nil {
		log.Fatal("open file err")
	}
	defer jsonFile.Close()
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}
	return byteValue, nil
}

func DecodeJSONString(jsonBytes []byte) (map[string]interface{}, error) {
	var m map[string]interface{}
	err := json.Unmarshal(jsonBytes, &m)
	if err != nil {
		log.Fatal("json unmarshal fail", err)
		return nil, err
	}
	return m, nil
}

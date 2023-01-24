package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"reflect"
	"strings"
)

const path = "files"

func main() {

	info := Run(path)

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Enter Key: ")
		scanner.Scan()

		input := scanner.Text()
		if len(input) != 0 {
			val, err := info.Get(input)
			if err != nil {
				fmt.Println(err)

			} else {
				printVal(val)
			}
		} else {
			break
		}
	}

	if scanner.Err() != nil {
		fmt.Println("Error: ", scanner.Err())
	}

}

type Config interface {
	Get(k string) (interface{}, error)
}

// Info describes result of merge operation
type Info struct {
	Errors   []error
	Mappings map[string]interface{}
}

func Run(path string) *Info {
	info := &Info{
		Mappings: make(map[string]interface{}),
	}

	f, err := os.Open(path)
	if err != nil {
		fmt.Println(err)

	}
	files, err := f.Readdir(0)
	if err != nil {
		fmt.Println(err)
	}

	for _, v := range files {

		fileContent, err := os.Open(path + "/" + v.Name())

		if err != nil {
			info.Errors = append(info.Errors, err)
		}

		defer fileContent.Close()

		byteResult, _ := io.ReadAll(fileContent)

		var data interface{}

		err = unmarshalJSON([]byte(byteResult), &data)

		if err != nil {
			info.Errors = append(info.Errors, err)
		}

		info.mergeConfig(data, nil)

	}

	return info

}

func (info *Info) mergeConfig(config interface{}, path []string) {
	if configObject, ok := config.(map[string]interface{}); ok {

		for k, v := range configObject {

			switch vv := v.(type) {
			case string:
				info.insertValue(k, vv, path)
			case int:
				info.insertValue(k, vv, path)
			case bool:
				info.insertValue(k, vv, path)
			case float64:
				info.insertValue(k, vv, path)
			case []interface{}:
				info.insertValue(k, vv, path)
			case map[string]interface{}:
				info.insertValue(k, vv, path)
				info.mergeConfig(vv, append(path, k))
			default:
				fmt.Println(k, "unknown type")
			}
		}

	}

}

func (info *Info) insertValue(k string, value interface{}, path []string) {
	path = append(path, k)
	pathStr := strings.Join(path, ".")

	config, valuePresent := info.Mappings[pathStr]

	if valuePresent && reflect.DeepEqual(value, config) {
		return
	}

	info.Mappings[pathStr] = value
}

func unmarshalJSON(buff []byte, data interface{}) error {
	decoder := json.NewDecoder(bytes.NewReader(buff))

	return decoder.Decode(data)
}

func printVal(val interface{}) {
	jsonStr, err := json.Marshal(val)

	if err != nil {
		fmt.Println("parsing error")
	}

	fmt.Println(string(jsonStr))
}

func (info *Info) Get(k string) (interface{}, error) {

	if len(strings.TrimSpace(k)) == 0 {
		return nil, errors.New("empty key")
	}

	configValue, configAvaialble := info.Mappings[k]

	if !configAvaialble {
		return nil, errors.New("config not present")
	}

	return configValue, nil

}

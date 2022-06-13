package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	if err := run(os.Args, os.Stdin, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run(args []string, stdin *os.File, stdout io.Writer) error {
	// We remove the first "useless" argument
	args = args[1:]

	// if args is empty return error
	inputBytes, err := ioutil.ReadAll(stdin)
	if err != nil {
		return err
	}

	items := []map[string]interface{}{}
	if err := json.Unmarshal(inputBytes, &items); err != nil {
		return err
	}

	outputMap, err := createLeavesFromLoop(args, items)
	if err != nil {
		return err
	}
	outBytes, err := json.MarshalIndent(outputMap, "", "  ")
	if err != nil {
		return err
	}

	file, err := os.Create("output.json")
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err := file.Write(outBytes); err != nil {
		return err
	}

	return nil
}

func createLeavesFromLoop(args []string, items []map[string]interface{}) (map[string]interface{}, error) {
	if len(args) == 0 {
		return nil, errors.New("arguments cannot be blank")
	}

	// if arg doesn't exist on the object, we use the zero value as a key
	outputMap := map[string]interface{}{}
	for _, item := range items {

		leafMap := outputMap
		for i, arg := range args {
			value := item[arg]
			// In case that the key is not of type string we will return an error, so we can return a valid json
			// If the value is nil, we treat it as an empty string
			valueStr, ok := value.(string)
			if !ok && value != nil {
				return nil, fmt.Errorf("invalid type for arg '%s': only strings are supported", arg)
			}

			if i < len(args)-1 {
				// First arguments: we create the leaves
				leaf, ok := leafMap[valueStr]
				if !ok {
					leaf = map[string]interface{}{}
					leafMap[valueStr] = leaf
				}
				newleafMap, ok := leaf.(map[string]interface{})
				if !ok {
					return nil, errors.New("unexpected object")
				}
				leafMap = newleafMap
			} else {
				// Last argument: we create and append to the array
				leaf, ok := leafMap[valueStr]
				if !ok {
					leaf = []interface{}{}
					leafMap[valueStr] = leaf
				}
				newleafSlc, ok := leaf.([]interface{})
				if !ok {
					return nil, errors.New("unexpected object")
				}
				newleafSlc = append(newleafSlc, item)
				leafMap[valueStr] = newleafSlc
			}

			delete(item, arg)
		}
	}
	return outputMap, nil
}

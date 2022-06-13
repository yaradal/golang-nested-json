package nesting

import (
	"errors"
	"fmt"
)

type Service interface {
	CreateLeavesFromLoop(args []string, items []map[string]interface{}) (map[string]interface{}, error)
}

func NewService() *service {
	return &service{}
}

type service struct{}

func (s *service) CreateLeavesFromLoop(args []string, items []map[string]interface{}) (map[string]interface{}, error) {
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

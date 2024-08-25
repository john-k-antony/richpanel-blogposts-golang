package utils

import (
	"encoding/json"
	"fmt"
	"strconv"
)

var idCounter int

// ToStringP creates a pointer to string representation of the passed interface{} type (any type)
func ToStringP(anything interface{}) *string {
	str := fmt.Sprintf("%v", anything)
	return &str
}

// StringifyJSON returns json in map[string]interface{} to string
func StringifyJSON(j map[string]interface{}, isPretty bool) string {
	if j == nil {
		return ""
	}
	var jsonStr []byte
	var err error
	if isPretty {
		jsonStr, err = json.MarshalIndent(j, "", "  ")
	} else {
		jsonStr, err = json.Marshal(j)
	}
	if err != nil {
		return err.Error()
	}
	return string(jsonStr)
}

func GetNextId() string {
	idCounter++
	return strconv.Itoa(idCounter)
}

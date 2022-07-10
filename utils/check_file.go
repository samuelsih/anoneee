package utils

import "strings"

func IsYAMLFile(filename string) bool {
	s := strings.Split(filename, ".")
	
	return s[len(s) - 1] == "yaml"
}

func IsJSONFile(filename string) bool {
	s := strings.Split(filename, ".")
	
	return s[len(s) - 1] == "json"
}
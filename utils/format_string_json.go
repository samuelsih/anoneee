package utils

import "strings"

func FormatString(value string) string {
	splittedStr := Split(value)

	var sb strings.Builder

	if len(splittedStr) == 1 {
		sb.WriteString(splittedStr[0])
		return sb.String()
	}

	for i := 0; i < len(splittedStr); i++ {
		str := strings.ToLower(splittedStr[i])

		if i == len(splittedStr)-1 {
			sb.WriteString(str)
		} else {
			sb.WriteString(str + "_")
		}
	}

	return sb.String()
}

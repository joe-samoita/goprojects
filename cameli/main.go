package main

import (
	"fmt"
)

func main() {
	fmt.Println(CamelToSnakeCase("HelloWorld"))
	fmt.Println(CamelToSnakeCase("helloWorld"))
	fmt.Println(CamelToSnakeCase("camelCase"))
	fmt.Println(CamelToSnakeCase("CAMELtoSnackCASE"))
	fmt.Println(CamelToSnakeCase("camelToSnakeCase"))
	fmt.Println(CamelToSnakeCase("hey2"))
}

func CamelToSnakeCase(s string) string {

	if s == "" {
		return s
	}
	runes := []rune(s)

	for i := 0; i < len(runes); i++ {
		if !(runes[i] >= 'a' && runes[i] <= 'z') && !(runes[i] >= 'A' && runes[i] <= 'Z') {
			return s

		}

		if i == len(runes)-1 && runes[i] >= 'A' && runes[i] <= 'Z' {
			return s
		}
		if i > 0 && runes[i] >= 'A' && runes[i] <= 'Z' && runes[i-1] >= 'A' && runes[i-1] <= 'Z' {

			return s
		}
	}

	result := make([]rune, 0, len(runes))

	for i, r := range runes {

		if r >= 'A' && r <= 'Z' {

			if i != 0 {
				result = append(result, '_')
			}
		}
		result = append(result, r)

	}
	return string(result)

}

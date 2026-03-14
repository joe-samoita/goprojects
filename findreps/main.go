// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func CodeRepeats(s string) string {

	s = s + "#"
	result := ""
	count := 1

	if len(s) == 0 {
		return ""
	}
	for i := 1; i < len(s); i++ {

		if s[i] == s[i-1] {
			count++
		} else {
			result += string(s[i-1]) + string(rune('0'+count))
			count = 1

		}
	}

	return result

}

func main() {
	fmt.Println(CodeRepeats("AAABBBCCC"))
	fmt.Println(CodeRepeats("!!@@##"))
	fmt.Println(CodeRepeats("     "))
	fmt.Println(CodeRepeats("JjjJohhnnnNn"))

}

package main

import "fmt"

func HashCode(dec string) string {

	hash := []byte{}
	size := len(dec)

	for _, ch := range dec {
		hashed := (int(ch) + size)

		if hashed > 127 {
			hashed += 33
		}
		hash = append(hash, byte(hashed))
	}

	return string(hash)

}

func main() {
	fmt.Println(HashCode("A"))
	fmt.Println(HashCode("AB"))
	fmt.Println(HashCode("BAC"))
	fmt.Println(HashCode("Hello World"))
}

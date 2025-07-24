package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go SHA256|SHA384|SHA512>")
		return
	}
	if os.Args[1] == "SHA384" {
		c1 := sha512.Sum384([]byte("x"))
		c2 := sha512.Sum384([]byte("X"))
		fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)

	} else if os.Args[1] == "SHA512" {
		c1 := sha512.Sum512([]byte("x"))
		c2 := sha512.Sum512([]byte("X"))
		fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
	} else {
		c1 := sha256.Sum256([]byte("x"))
		c2 := sha256.Sum256([]byte("X"))
		fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
	}

}

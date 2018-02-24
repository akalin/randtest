package main

import (
	"alice/eightball"
	"crypto/rand"
	"fmt"
)

func main() {
	fmt.Println(eightball.Ask())

	var b [32]byte
	rand.Read(b[:])
	fmt.Printf("%x\n", b)
}

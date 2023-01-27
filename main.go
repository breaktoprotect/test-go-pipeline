package main

import (
	"fmt"
	"test-go-pipeline/cmd/shapes"
)

func main() {
	fmt.Println("[*] Welcome to simple Rectangle calculator")
	r := shapes.Rectangle{Width: 10, Height: 20}
	fmt.Printf("[+] Rectangle of %f by %f has been instantiated!\n", r.Width, r.Height)
	fmt.Print("[+] Rectangle of")
}

package main

import (
	"fmt"
	"test-go-pipeline/cmd/shapes"
	"test-go-pipeline/internal/simpleHttpClient"
)

func main() {
	fmt.Println("[*] Welcome to simple Rectangle calculator")
	r := shapes.Rectangle{Width: 10, Height: 20}
	fmt.Printf("[+] Rectangle of %f by %f has been instantiated!\n", r.Width, r.Height)
	fmt.Print("[+] Rectangle of")

	fmt.Println()

	fmt.Println("[*] Welcome to net/http requests to httpbin.org")
	simpleHttpClient.GetHttpBin()

	fmt.Println()

	fmt.Println("[*] Welcome to go-openapi requests to jsonplaceholder.typicode.org")
	simpleHttpClient.GetTodo()
}

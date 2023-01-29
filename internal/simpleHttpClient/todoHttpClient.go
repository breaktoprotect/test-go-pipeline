package simplehttpclient

import (
	"fmt"
	"os"

	"test-go-pipeline/internal/simplehttpclient/generated/client"
	"test-go-pipeline/internal/simplehttpclient/generated/client/todos"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/joho/godotenv"
)

func RunTodoClientMain() {
	// Load env
	err := godotenv.Load(".env")

	// Create a transport
	transport := httptransport.New(os.Getenv("TODO_HOSTNAME"), "", nil)

	// Auth
	transport.DefaultAuthentication = httptransport.BearerToken(os.Getenv("X_API_KEY"))

	//debug
	fmt.Printf("TODO_HOSTNAME: %s\n", os.Getenv("TODO_HOSTNAME"))

	// Initialize the client with the httpClient
	c := client.New(transport, strfmt.Default)

	// Auth
	//bearerTokenAuth := httptransport.BearerToken(os.Getenv("X-API-KEY"))

	// Make the GET request to the /todos endpoint
	params := todos.NewGetTodosListParams()
	resp, err := c.Todos.GetTodosList(params)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to make GET request: %s\n", err)
		os.Exit(1)
	}

	// Print the response
	//fmt.Printf("%#v\n", resp.GetPayload()[0])

	// Print a human-readable content
	for _, item := range resp.GetPayload() {
		fmt.Printf("%v - %v\n", *item.ID, *item.Title)
		fmt.Printf("Created by user (userID): %v || ", *item.UserID)
		fmt.Printf("Completed: %v\n", *item.Completed)
		fmt.Println()
	}

}

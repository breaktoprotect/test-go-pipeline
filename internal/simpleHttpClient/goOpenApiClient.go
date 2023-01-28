package simplehttpclient

import (
	"fmt"
	"os"

	"test-go-pipeline/internal/simplehttpclient/generated/client"
	"test-go-pipeline/internal/simplehttpclient/generated/client/todos"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

func GoOpenApiClient() {
	// Create a transport
	transport := httptransport.New(os.Getenv("TODO_HOSTNAME"), "", nil)

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
	fmt.Println(resp.Payload)
}

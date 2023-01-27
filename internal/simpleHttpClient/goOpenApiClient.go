package simpleHttpClient

import (
	"context"
	"fmt"
	"log"
	"net/http"
	jphclient "test-go-pipeline/internal/generated/jsonplaceholder/client"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

var apiHostname = "https://jsonplaceholder.typicode.com"
var endpointTodo = "/todos"

type Todo struct {
	userId    string `json:"userId"`
	id        string `json:"id"`
	title     string `json:"title"`
	completed bool   `json:"completed"`
}

func GetTodo(id int) *Todo {
	transport := runtime.New(http.DefaultClient, "http://jsonplaceholder.typicode.com", []string{"http"})

	client := jphclient.New(transport, strfmt.Default)

	ctx := context.Background()

	params := client.NewGetTodosParams()

	resp, err := client.Operations.GetTodos(ctx, params)
	if err != nil {
		log.Fatalf("GetTodo() encountered an error when performing get request: %v", err)
	}

	fmt.Println(resp)

	// Convert JSON response to struct
	/* var todo Todo
	err = json.Unmarshal(resp.Payload, &todo)
	if err != nil {
		log.Fatalf("GetTodo() failed to unmarshal JSON: %v", err)
	}
	return &todo*/

	return resp.Payload
}

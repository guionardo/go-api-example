package main

//NOCOVER

import (
	"fmt"
	"os"
	"strings"

	"github.com/guionardo/go-api-example/api"
	"github.com/guionardo/go-api-example/infra"
	"github.com/guionardo/go-api-example/repository"
)

func usage() {
	fmt.Printf("Usage: %s <command>\n", os.Args[0])
	fmt.Printf(`Available commands:
	api : Run the API
	reset : Reset the database

`)
	os.Exit(1)
}
func main() {
	fmt.Println("API")
	infra.SetupLog()
	var command string
	if len(os.Args) > 1 {
		command = strings.ToLower(os.Args[1])
	}
	switch command {
	case "api":
		api.RunAPI()
	case "reset":
		repository.RunReset()
	default:
		usage()
	}

}

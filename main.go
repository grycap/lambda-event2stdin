package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/aws/aws-lambda-go/lambda"
)

// Handler function to start a executable passing the event to stdin
func Handler(ctx context.Context, event json.RawMessage) error {
	// Read the application's path from the first argument
	if len(os.Args) < 2 {
		log.Println("Please, set a executable file as first argument")
		return fmt.Errorf("Please, set a executable file as first argument")
	}
	app := os.Args[1]

	// Execute the program passing the event to stdin
	cmd := exec.Command(app)
	cmd.Stdin = strings.NewReader(string(event))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	log.Printf("Executing \"%s\"", app)
	err := cmd.Run()
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

// Run the Lambda function
func main() {
	lambda.Start(Handler)
}

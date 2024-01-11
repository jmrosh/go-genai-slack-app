package main

import (
	"context"
	"github.com/GoogleCloudPlatform/functions-framework-go/funcframework"
	run "github.com/jmrosh/go-genai-slack-app"
	"log"
	"os"
	"time"
)

func runAssistant(ctx context.Context) {
	funcframework.RegisterEventFunctionContext(ctx, "assistant", run.Assistant)

	port := "3000"
	log.Printf("/assistant running on port %s", port)
	if err := funcframework.Start(port); err != nil {
		log.Fatalf("funcframework.Start: %v\n", err)
	}

}

func runSubscriber(ctx context.Context) {
	funcframework.RegisterHTTPFunctionContext(ctx, "subscriber", run.Subscriber)

	port := "3001"
	log.Printf("/subscriber running on port %s", port)
	if err := funcframework.Start(port); err != nil {
		log.Fatalf("funcframework.Start: %v\n", err)
	}
}

func main() {
	os.Setenv("FUNCTION_TARGET", "-")

	ctx := context.Background()
	go runSubscriber(ctx)
	go runAssistant(ctx)

	// Timout after 1 hour
	time.Sleep(time.Hour)
}

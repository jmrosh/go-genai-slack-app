build:
	go build -o go-genai-slack-app
build-local:
	go build -o go-genai-slack-app-local cmd/local.go
test:
	go test -v ./...
firestore:
	gcloud emulators firestore start
run:
	go run cmd/local.go

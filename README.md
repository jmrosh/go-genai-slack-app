# go-genai-slack-app

## Overview
This repository contains Google Cloud Functions that are used for a GenAI Slack App.

## Functions
### Subscriber
Triggered by slack events via HTTP POST requests.
* [Google Cloud Functions with HTTP Triggers (1st gen)](https://cloud.google.com/functions/docs/calling/http)
* [Slack Events API Documentation](https://api.slack.com/apis/connections/events-api)

### Assistant
Triggered by document changes to a Firestore collection.
* [Google Cloud Functions with Cloud Firestore Triggers (1st gen)](https://cloud.google.com/functions/docs/calling/cloud-firestore-1st-gen)

## Development
### Prerequisites
* [Go](https://golang.org/doc/install)
* [Google Cloud CLI](https://cloud.google.com/sdk/gcloud/reference) (for firestore emulator)

### Run tests
```sh
make test
```
### Run functions locally
```sh
make firestore
```

```sh
make run
```

## Deployment
```sh

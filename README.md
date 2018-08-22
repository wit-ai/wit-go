# wit.ai 🤖 [![Build Status](https://travis-ci.org/plutov/wit.ai.svg?branch=master)](https://travis-ci.org/plutov/wit.ai) [![GoDoc](https://godoc.org/github.com/plutov/wit.ai?status.svg)](https://godoc.org/github.com/plutov/wit.ai) [![Go Report Card](https://goreportcard.com/badge/github.com/plutov/wit.ai)](https://goreportcard.com/report/github.com/plutov/wit.ai)

Go client for [wit.ai](https://wit.ai/) HTTP API.

## Install

```
go get -u github.com/plutov/wit.ai
```

## Usage

### Parse text

```go
package main

import (
	"os"

	witai "github.com/plutov/wit.ai"
)

func main() {
	client := witai.NewClient(os.Getenv("WIT_AI_TOKEN"))

	msg, _ := client.Parse(&witai.MessageRequest{
		Query: "hello",
	})
	fmt.Printf("%v", msg)
}
```

### Send an audio file

```go
msg, _ := client.Speech(&witai.MessageRequest{
	Speech: &witai.Speech{
		File:        file,
		ContentType: "audio/raw;encoding=unsigned-integer;bits=16;rate=16k;endian=little",
	},
})
```

### Create entity

```go
client.CreateEntity(witai.NewEntity{
	ID:  "favorite_city",
	Doc: "A city that I like",
})
```

### Get entity

```go
client.GetEntity("favorite_city")
```

## Testing

### Unit tests

```
go test -race -v ./...
```

### Integration tests

Integration tests are connecting to real Wit.ai API, so you need to provide a valid token:

```
WITAI_INTEGRATION_TOKEN=your_secret_token_here go test -v -tags=integration
```
# wit.ai 🤖 

Go client for [wit.ai](https://wit.ai/) HTTP API. [![Build Status](https://travis-ci.org/plutov/wit.ai.svg?branch=master)](https://travis-ci.org/plutov/wit.ai) [![GoDoc](https://godoc.org/github.com/plutov/wit.ai?status.svg)](https://godoc.org/github.com/plutov/wit.ai) [![Go Report Card](https://goreportcard.com/badge/github.com/plutov/wit.ai)](https://goreportcard.com/report/github.com/plutov/wit.ai)

## Install

```
go get -u github.com/plutov/wit.ai
```

## Usage

```go
package main

import (
	"os"

	witai "github.com/plutov/wit.ai"
)

func main() {
	// WIT_AI_TOKEN is a Server Access Token
	client := witai.NewClient(os.Getenv("WIT_AI_TOKEN"))

	// /messages?q=hello
	msg, _ := client.Parse(&witai.MessageRequest{
		Query: "hello",
	})
	fmt.Printf("%v", msg)
}
```
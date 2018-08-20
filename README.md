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

	// https://wit.ai/docs/http/20170307#get__message_link
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
# wit.ai 🤖 

Go client for wit.ai HTTP API. [![Build Status](https://travis-ci.org/plutov/wit.ai.svg?branch=master)](https://travis-ci.org/plutov/logrus) [![GoDoc](https://godoc.org/github.com/plutov/wit.ai?status.svg)](https://godoc.org/github.com/plutov/wit.ai) [![Go Report Card](https://goreportcard.com/badge/github.com/plutov/wit.ai)](https://goreportcard.com/report/github.com/plutov/wit.ai)

## Install

```
go get -u github.com/plutov/wit.ai
```

## Usage

```
package main

import (
	"os"

	witai "github.com/plutov/wit.ai"
)

func main() {
	client := witai.NewClient(os.Getenv("WIT_AI_TOKEN"))
}
```
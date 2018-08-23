![wit.ai](https://s3.amazonaws.com/pliutau.com/wit.png)

[![Build Status](https://travis-ci.org/plutov/wit.ai.svg?branch=master)](https://travis-ci.org/plutov/wit.ai) [![GoDoc](https://godoc.org/github.com/plutov/wit.ai?status.svg)](https://godoc.org/github.com/plutov/wit.ai) [![Go Report Card](https://goreportcard.com/badge/github.com/plutov/wit.ai)](https://goreportcard.com/report/github.com/plutov/wit.ai)

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

### Update entity

```go
client.UpdateEntity("favorite_city", witai.UpdateEntityFields{
	Doc: "My favorite city",
})
```

### Delete entity

```go
client.DeleteEntity("favorite_city")
```

### Add entity value

```go
client.AddEntityValue("favorite_city", witai.EntityValue{
	Value: "HCMC",
	Expressions: ["HoChiMinh", "HCMC"],
})
```

### Delete entity value

```go
client.DeleteEntityValue("favorite_city", "HCMC")
```

### Add value expression

```go
client.AddEntityValueExpression("favorite_city", "HCMC", "HoChiMinh")
```

### Delete value expression

```go
client.DeleteEntityValueExpression("favorite_city", "HCMC", "HoChiMinh")
```

## Testing

### Unit tests

```
go test -race -v
```

### Integration tests

Integration tests are connecting to real Wit.ai API, so you need to provide a valid token:

```
export WITAI_INTEGRATION_TOKEN=your_secret_token_here
go test -v -tags=integration
```
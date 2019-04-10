![wit.ai](https://s3.amazonaws.com/pliutau.com/wit.png)

[![Build Status](https://travis-ci.org/wit-ai/wit-go.svg?branch=master)](https://travis-ci.org/wit-ai/wit-go) [![GoDoc](https://godoc.org/github.com/wit-ai/wit-go?status.svg)](https://godoc.org/github.com/wit-ai/wit-go) [![Go Report Card](https://goreportcard.com/badge/github.com/wit-ai/wit-go)](https://goreportcard.com/report/github.com/wit-ai/wit-go)

*This repository is community-maintained. We gladly accept pull requests. Please see the [Wit HTTP Reference](https://wit.ai/docs/http/latest) for all supported endpoints.*

Go client for [wit.ai](https://wit.ai/) HTTP API.

## Install

```
go get -u github.com/wit-ai/wit-go
```

## Usage

### Parse text

```go
package main

import (
	"os"
	"fmt"

	witai "github.com/wit-ai/wit-go"
)

func main() {
	client := witai.NewClient(os.Getenv("WIT_AI_TOKEN"))
	// Use client.SetHTTPClient() to set custom http.Client

	msg, _ := client.Parse(&witai.MessageRequest{
		Query: "hello",
	})
	fmt.Printf("%v", msg)
}
```

### Send an audio file

```go
file, _ := os.Open("speech.wav")

msg, _ := client.Speech(&witai.MessageRequest{
	Speech: &witai.Speech{
		File:        file,
		ContentType: "audio/raw;encoding=unsigned-integer;bits=16;rate=16k;endian=little",
	},
})
```

### Entities

Create:
```go
client.CreateEntity(witai.Entity{
	ID:  "favorite_city",
	Doc: "A city that I like",
})
```

Get:
```go
client.GetEntity("favorite_city")
```

Update:
```go
client.UpdateEntity("favorite_city", witai.Entity{
	Doc: "My favorite city",
})
```

Delete:
```go
client.DeleteEntity("favorite_city")
```

### Entity values

Add:
```go
client.AddEntityValue("favorite_city", witai.EntityValue{
	Value: "HCMC",
	Expressions: ["HoChiMinh", "HCMC"],
})
```

Delete:
```go
client.DeleteEntityValue("favorite_city", "HCMC")
```

### Value expressions

Add:
```go
client.AddEntityValueExpression("favorite_city", "HCMC", "HoChiMinh")
```

Delete:
```go
client.DeleteEntityValueExpression("favorite_city", "HCMC", "HoChiMinh")
```

### Training

Validate samples (sentence + entities annotations) to train your app programmatically:
```go
client.ValidateSamples([]witai.Sample{
	Sample{
		Text: "I live in HCMC",
	},
})
```

Get validate samples:
```go
limit := 10
offset := 0
client.GetSamples(limit, offset)
```

### Export

```go
downloadURL := client.Export()
```

## Testing

Both Unit / Integration tests are executed by TravisCI.

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


## License

The license for wit-go can be found in LICENSE file in the root directory of this source tree.

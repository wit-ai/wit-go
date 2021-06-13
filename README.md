![wit.ai](https://s3.amazonaws.com/pliutau.com/wit.png)

[![GoDoc](https://godoc.org/github.com/wit-ai/wit-go?status.svg)](https://godoc.org/github.com/wit-ai/wit-go) [![Go Report Card](https://goreportcard.com/badge/github.com/wit-ai/wit-go)](https://goreportcard.com/report/github.com/wit-ai/wit-go)

*This repository is community-maintained. We gladly accept pull requests. Please see the [Wit HTTP Reference](https://wit.ai/docs/http/latest) for all supported endpoints.*

Go client for [wit.ai](https://wit.ai/) HTTP API.

## Install

```
go get -u github.com/wit-ai/wit-go/v2
```

## Usage

```go
package main

import (
	"os"
	"fmt"

	witai "github.com/wit-ai/wit-go/v2"
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

## Testing

Both Unit / Integration tests are executed by Github Actions.

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

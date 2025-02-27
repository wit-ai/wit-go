![wit.ai](https://s3.amazonaws.com/pliutau.com/wit.png)

[![Go Reference](https://pkg.go.dev/badge/github.com/wit-ai/wit-go)](https://pkg.go.dev/github.com/wit-ai/wit-go)

_This repository is community-maintained. We gladly accept pull requests. Please
see the [Wit HTTP Reference](https://wit.ai/docs/http/latest) for all supported
endpoints._

Go client for [wit.ai](https://wit.ai/) HTTP API.

API version: 20240304

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

Unit tests are executed by Github Actions.

### Unit tests

```
go test -race -v
```

### Integration tests

Integration tests have to be executed manually by providing a valid token via
`WITAI_INTEGRATION_TOKEN` env var.

Integration tests are connecting to real Wit.ai API, so you need to provide a
valid token:

```
WITAI_INTEGRATION_TOKEN={SERVER_ACCESS_TOKEN} go test -v -tags=integration
```

## License

The license for wit-go can be found in LICENSE file in the root directory of
this source tree.

## Terms of Use

Our terms of use can be found at https://opensource.facebook.com/legal/terms.

## Privacy Policy

Our privacy policy can be found at
https://opensource.facebook.com/legal/privacy.

## IDE
- Visutal Studio Code
- jetbrains goland


## Getting start

```go
// main package !!!
package main

import "fmt"

const (
	PublicConst  = "public const"
	privateConst = "package private const"
)

var (
	PublicVar  = "public var"
	privateVar = "package private var"
)

func init() {
	fmt.Println("before main")
}

func main() {
	fmt.Println("Hello world")
}

func sayHello() string {
	return "Hello world"
}
```

### Build

`go build -o hello`

### Run

`go run start.go`

### Unit Test

```go
package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldSayHello(t *testing.T) {
	assert := assert.New(t)

	assert.Equal("Hello world", sayHello())
}
```

- All: `go test ./...`
- Single: `go test -run ^(TestShouldSayHello)$`
- With timeout: `go test -timeout 30s -run ^(TestShouldSayHello)$`

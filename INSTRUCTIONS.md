# Installation and First Program

## Download Installer

<https://go.dev/dl>

## Create a module

```bash
mkdir hello
cd hello
go mod init github.com/rjha1-godaddy/go-fundamentals
```

### Generated go.mod

```bash
module github.com/rjha1-godaddy/go-fundamentals

go 1.24.0
```

## Hello World

Create `main.go` with the following content:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

## Build and Run

```bash
go build main.go
./main
```

```bash
go run main.go
```

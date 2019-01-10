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

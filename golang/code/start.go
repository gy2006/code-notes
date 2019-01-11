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

func sayHello() string {
	return "Hello world"
}

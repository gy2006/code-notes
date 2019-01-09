package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldSayHello(t *testing.T) {
	assert := assert.New(t)

	assert.Equal("Hello world", sayHello())
}

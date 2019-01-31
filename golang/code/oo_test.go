package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldNilForPointerType(t *testing.T) {
	assert := assert.New(t)

	pointer := new(Person)
	assert.Nil(pointer.Address)
}

func TestShouldPrintPersonName(t *testing.T) {
	assert := assert.New(t)

	pointer := &Person{
		Name: "yang",
		Age:  18,
	}

	assert.NotNil(pointer)
	assert.Equal("name:yang", pointer.PrintName())
}

func TestShouldNewClassWithInherit(t *testing.T) {
	assert := assert.New(t)

	inherit := &Student{
		Person: Person{
			Name: "yang",
			Age:  18,
		},
		School: "xxxx",
	}

	assert.NotNil(inherit)
	assert.Equal("yang", inherit.Name)
	assert.NotNil(18, inherit.Age)
}

func TestShouldImplementJsonable(t *testing.T) {
	assert := assert.New(t)

	var jsonable Jsonable
	jsonable = &Person{
		Name: "yang",
		Age:  18,
	}

	otherJsonable := &Person{
		Name: "yang",
		Age:  18,
	}

	assert.NotNil(jsonable)
	assert.Equal(`{"name":"yang","age":18,"address":null}`, jsonable.ToJson())
	assert.Equal(`{"name":"yang","age":18,"address":null}`, otherJsonable.ToJson())
}

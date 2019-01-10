package main

import "encoding/json"

// Define a interface
type Jsonable interface {
	ToJson() string
}

type Address struct {
	Zipcode  string
	Country  string
	Province string
	Addr     string
}

type Person struct {
	Name    string   `json:"name"`
	Age     int      `json:"age"`
	Address *Address `json:"address"`
}

func (p Person) PrintName() string {
	return "name:" + p.Name
}

func (p Person) ToJson() string {
	bytes, err := json.Marshal(p)
	if err != nil {
		return ""
	}
	return string(bytes)
}

// Inherit
type Student struct {
	Person
	// Other struct..
	// Other struct..

	School string
}

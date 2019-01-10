# OO

```go
// Every type implements at least zero methods
// Hold values of any type
interface{}
```


## Class

```go
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
```

```go
func TestShouldPrintPersonName(t *testing.T) {
	assert := assert.New(t)

	pointer := &Person{
		Name: "yang",
		Age:  18,
	}

	assert.NotNil(pointer)
	assert.Equal("name:yang", pointer.PrintName())
}
```

## Inherit

No keyword such as `extends` in java

```go
type Student struct {
    Person
    // Other struct..
    // Other struct..

    School string
}
```

```go
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
```

## Interface

```go
// Define a interface
type Jsonable interface {
	ToJson() string
}
```

```go
// Implement Jsonable 
type Person struct {
	Name    string   `json:"name"`
	Age     int      `json:"age"`
	Address *Address `json:"address"`
}

func (p Person) ToJson() string {
	bytes, err := json.Marshal(p)
	if err != nil {
		return ""
	}
	return string(bytes)
}
```

```go
func TestShouldImplementJsonable(t *testing.T) {
	assert := assert.New(t)

	var jsonable Jsonable
	jsonable = &Person{
		Name: "yang",
		Age:  18,
	}

	assert.NotNil(jsonable)
	assert.Equal(`{"name":"yang","age":18,"address":null}`, jsonable.ToJson())
}
```
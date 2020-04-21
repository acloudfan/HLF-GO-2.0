package main

import "fmt"
import "github.com/golang/protobuf/proto"


func main() {
	// Create an instance of the structure Person defined in proto.buf.go
	john := &Person{
		Name: "john",
		Age: 22,
	}

	// Generate the proto buffer representation
	data, error := proto.Marshal(john)
	if error != nil {
		fmt.Println("Marshalling error: ", error)
	}

	fmt.Println(data)

	// Convert the proto buffer to struct
	newPerson   := &Person{}
	// Populate the struct from the data in the buffer
	error = proto.Unmarshal(data, newPerson)
	fmt.Println(newPerson.GetAge(), newPerson.GetName())
}
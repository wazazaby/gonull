package main

import (
	"encoding/json"
	"fmt"

	"github.com/LukaGiorgadze/gonull/v2"
)

type MyCustomInt int
type MyCustomFloat32 float32

type Person struct {
	Name    string                           `json:"name"`
	Age     gonull.Nullable[MyCustomInt]     `json:"age"`
	Address gonull.Nullable[string]          `json:"address"`
	Height  gonull.Nullable[MyCustomFloat32] `json:"height"`
	HasPet  gonull.Nullable[bool]            `json:"has_pet"`
	IsZero  gonull.Nullable[bool]            `json:"is_zero,omitzero"` // This property will be omitted from the output if it is false.
}

func main() {
	jsonData := []byte(`{"name":"Alice","age":15,"address":null,"height":null}`)

	var person Person
	if err := json.Unmarshal(jsonData, &person); err != nil {
		panic(err)
	}

	// Age is present and valid.
	fmt.Printf("Person.Age is valid: %t, present: %t\n", person.Age.Valid, person.Age.Present)

	// Address is present but invalid (explicit null).
	fmt.Printf("Person.Address is valid: %t, present: %t\n", person.Address.Valid, person.Address.Present)

	// Same for the height.
	fmt.Printf("Person.Height is valid: %t, present: %t\n", person.Height.Valid, person.Height.Present)

	// HasPet is not present.
	fmt.Printf("Person.HasPet is valid: %t, present: %t\n", person.HasPet.Valid, person.HasPet.Present)

	// IsZero is not present.
	fmt.Printf("Person.IsZero is valid: %t, present: %t\n", person.IsZero.Valid, person.IsZero.Present)

	marshalledData, err := json.Marshal(person)
	if err != nil {
		panic(err)
	}
	// Null values will be kept when marshalling to JSON.
	fmt.Println(string(marshalledData))

	// Output:
	// Person.Age is valid: true, present: true
	// Person.Address is valid: false, present: true
	// Person.Height is valid: false, present: true
	// Person.HasPet is valid: false, present: false
	// Person.IsZero is valid: false, present: false
	// {"name":"Alice","age":15,"address":null,"height":null,"has_pet":null}

	// As you see, IsZero is not present in the output, because we used the omitzero tag introduced in go v1.24.

}

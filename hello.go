package main

import (
	"fmt"
	"go-service/internal/data_structure/array"
	"go-service/internal/data_structure/basic"
	"go-service/internal/data_structure/slice"
	"go-service/internal/data_structure/structs"
	"go-service/internal/function"
	"go-service/internal/operators"
)

func main() {
	fmt.Println("Basic types")
	basic.Integer()
	basic.String()
	basic.Float()
	basic.Boolean()

	fmt.Println("Arrays")
	array.Dynamic()
	array.DefinedLength()

	fmt.Println("Slices")
	slice.Basic()
	slice.ArrayToSlice("!")
	slice.MakeSlice()
	slice.AccessSlice()
	slice.ChangeElSlice()
	fmt.Println("Append sem passar valores")
	slice.AppendSlice()
	fmt.Println("Append passando valores")
	slice.AppendSlice("Valores", "Inseridos", "Posteriormente")
	slice.Copy()

	fmt.Println("Operators")
	operators.Comparison()

	fmt.Println("Functions")
	fmt.Println(function.WithReturn(12, 32))
	fmt.Println(function.WithNamedReturn(33, 44))
	res, msg := function.WithMultipleNamedReturn(26, 212)
	fmt.Printf("MSG: %v RES: %d\n", msg, res)
	res, msg = function.WithMultipleReturn(66, 23)
	fmt.Printf("MSG: %v RES: %d\n", msg, res)
	fmt.Println("test assign empty without declaring")
	_, msg = function.WithMultipleNamedReturn(11, 2)
	fmt.Printf("MSG: %v\n", msg)

	_, msg2 := function.WithMultipleNamedReturn(43, 33)
	fmt.Printf("MSG: %v\n", msg2)
	res2, _ := function.WithMultipleReturn(87, 347)
	fmt.Printf("RES: %d\n", res2)

	structs.GenerateUser("Jamal")
	structs.GenerateUser("Josue")

	users := structs.GetUsers()

	for _, user := range users {
		fmt.Printf("- User data -\nID: %d | Name: %s\n", user.ID, user.Name)
	}
}

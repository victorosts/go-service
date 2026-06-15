package basic

import (
	"fmt"
)

func Integer() {
	var i int = 1259

	fmt.Println(i)
}

func String() {
	var s string = "Hello World"

	fmt.Println(s)
}

func Float() {
	var f32limit float32 = 3.4e+38
	var f64limit float64 = 1.7e+308

	fmt.Printf("float32 limit: %v, float64 limit: %v\n", f32limit, f64limit)
}

func Boolean() {
	var btrue bool = true
	var bfalse bool = false

	fmt.Println(btrue)
	fmt.Println(bfalse)
}

package array

import "fmt"

func Dynamic() {
	intArr := [...]int{1, 2, 5, 8, 8}

	fmt.Println(intArr)
}

func DefinedLength() {
	strArr := [5]string{}

	fmt.Println(strArr)

	strArr[3] = "something"
	fmt.Println(strArr)
}

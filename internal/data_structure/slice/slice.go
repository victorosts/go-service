package slice

import "fmt"

func Basic() {
	slc := []int{}
	fmt.Println(slc)

	slc = append(slc, 2, 3)
	fmt.Println(slc)

	slcWithLengthAndCap := []bool{true, false, true, true}
	fmt.Println(slcWithLengthAndCap)
}

func ArrayToSlice(elementToAdd string) {
	arr := [3]string{"test", "hello", "world"}
	fmt.Println(arr)

	slice := arr[:]
	fmt.Println(slice)
	slice = append(slice, elementToAdd)
	fmt.Println(slice)
}

func MakeSlice() {
	slcwithLength := make([]float32, 5)

	fmt.Printf("Slice with length: %v\n", slcwithLength)
	fmt.Printf("Slice length: %d\n", len(slcwithLength))
	fmt.Printf("Slice capacity: %d\n", cap(slcwithLength))

	slcWithLengthAndCap := make([]int, 3, 5)

	fmt.Printf("Slice with length and cap: %v\n", slcWithLengthAndCap)
	fmt.Printf("Slice length: %v\n", len(slcWithLengthAndCap))
	fmt.Printf("Slice cap: %v\n", cap(slcWithLengthAndCap))
}

func AccessSlice() {
	prices := []int{10, 12, 30}

	fmt.Println(prices[2])
	fmt.Println(prices[0])
}

func ChangeElSlice() {
	prices := []int{10, 14, 22}
	fmt.Printf("Prices before: %v\n", prices)

	prices[2] = 33
	fmt.Printf("Prices after: %v\n", prices)
}

func AppendSlice(elements ...string) {
	items := []string{"Default", "Values"}
	fmt.Printf("Items are: %v\n", items)

	items = append(items, elements...) // el expansion

	fmt.Printf("Items now are: %v\n", items)
}

func Copy() {
	items := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("Items: %v\n", items)
	fmt.Printf("Items length: %v\n", len(items))
	fmt.Printf("Items capacity: %v\n", cap(items))

	newItems := items[1:5]
	fmt.Printf("Part of Items: %v\n", newItems)
	fmt.Printf("Part of Items length: %v\n", len(newItems))
	fmt.Printf("Part of Items capacity: %v\n", cap(newItems))

	sliceCopy := make([]int, 5)
	copy(sliceCopy, items[2:])

	fmt.Printf("Copy of Items: %v\n", sliceCopy)
	fmt.Printf("Copy of Items length: %v\n", len(sliceCopy))
	fmt.Printf("Copy of Items capacity: %v\n", cap(sliceCopy))
}

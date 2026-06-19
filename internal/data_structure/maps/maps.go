package maps

import "fmt"

func CreateMap() {
	zeldaGameMap := map[string]string{"game": "The Legend of Zelda", "category": "ação-aventura", "launched_at": "1986-02-21"}

	for key, value := range zeldaGameMap {
		fmt.Printf("Key: %s - Value: %s\n", key, value)
	}

	makeMap := make(map[string]int)

	makeMap["testKey"] = 44
	makeMap["79"] = 79

	fmt.Printf("makeMap - %v\n", makeMap)

	delete(makeMap, "79")
	fmt.Printf("makeMap - %v\n", makeMap)

	newMap := make(map[string]string)
	newMap["brand"] = "ford"
	newMap["model"] = "fiesta"

	newMapBrand, ok1 := newMap["brand"]
	newMapYear, ok2 := newMap["year"]
	newMapModel, ok3 := newMap["model"]

	fmt.Println(newMapBrand, ok1)
	fmt.Println(newMapYear, ok2)
	fmt.Println(newMapModel, ok3)

	refA := map[string]bool{"success": true, "error": false}
	refB := refA

	fmt.Printf("refA: %v\n", refA)
	fmt.Printf("refB: %v\n", refB)

	refA["logged"] = true

	fmt.Printf("refA: %v\n", refA)
	fmt.Printf("refB: %v\n", refB)
}

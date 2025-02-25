package iterate

import "fmt"

func MainMapFunctions() {
	myMap := make(map[int]int)
	myMap[1] = 3
	myMap[2] = 4
	myMap[4] = 6

	for key, value := range myMap {
		fmt.Printf("key : %d and value : %d", key, value)
	}
}

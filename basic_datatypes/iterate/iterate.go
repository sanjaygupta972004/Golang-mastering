package iterate

import "fmt"

func main() {

	// this is most common way to implement like all language
	for i := 0; i < 100; i++ {

		if i%5 == 0 {
			continue
		}

		if i == 98 {
			break
		}

		fmt.Printf("the value of number :%d\n", i)

	}

	// how to implement while loop in golang if there is no while key-word
	n := 20
	for {
		if n <= 0 {
			break
		}
		fmt.Print(n * 2)
		n--
	}

	// how to use range key-word to aceess index and respectivew value from the array

	for i, val := range []int{4, 6, 7, 4, 8, 9} {
		if val == 9 {
			fmt.Println("how to check a condition while iterating on a loop")
		}

		if i < 0 {
			fmt.Print("index can not be negative")
		}

		fmt.Print("\n", val)
	}

	// how to do iterate on map data-type

	studentScores := map[string]int{"Alice": 90, "Bob": 85, "Charlie": 88}

	for key, val := range studentScores {
		fmt.Printf("%s of score : %d", key, val)
	}
}

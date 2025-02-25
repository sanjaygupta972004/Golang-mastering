package reference

import "fmt"

// go pointers

func getPointers(n *int) {
	*n = *n * *n
}

func returnPointers(n int) *int {
	n = n * n

	return &n
}

func MainPointers() {
	n := 4
	nj := &n
	getPointers(nj)
	returnPointers(n)
	getAge()

}

// most common use of pointer in golang
type Person struct {
	Name string
	Age  int
}

func updateAge(p *Person) {
	p.Age = 40
}

func getAge() {
	person := Person{Name: "Chai", Age: 20}
	updateAge(&person)

	fmt.Printf("age of the person %d", person.Age)
}

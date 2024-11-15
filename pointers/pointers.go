package pointers

import "fmt"

func App() {
	age := 32
	var userAge *int
	userAge = &age
	fmt.Println("Age:", *userAge)

	adultYears := getAdultYears(age)
	fmt.Println("Adult years:", adultYears)
}

func getAdultYears(age int) int {
	return age - 18
}

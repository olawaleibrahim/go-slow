package pointers

import "fmt"

func App() {
	age := 32
	// var userAge *int
	userAge := &age
	fmt.Println("Age:", *userAge)

	adultYears := getAdultYears(userAge)
	fmt.Println("Adult years:", adultYears)
	fmt.Println("age years:", age)
}

func getAdultYears(age *int) int {
	*age = *age - 18
	return *age
}

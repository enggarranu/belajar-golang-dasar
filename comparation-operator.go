package main

import (
	"fmt"
)

func main() {
	var name1 = "Enggar"
	var name2 = "Novan"

	var result bool = name1 > name2
	fmt.Println(result) // The result will be false because E less than N (ascending)

	result = name1 == name2
	fmt.Println(result) // The result will be false because Enggar not equals Novan

	var value1 = 100
	var value2 = 200

	fmt.Println(value1 > value2)
	fmt.Println(value1 < value2)
	fmt.Println(value1 == value2)
	fmt.Println(value1 != value2)

}

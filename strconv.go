package main

import (
	"fmt"
	"strconv"
)

func main() {
	boolean, err := strconv.ParseBool("true")
	if err == nil {
		fmt.Println(boolean)

	} else {
		fmt.Println(err.Error())
	}

	number, err := strconv.ParseInt("100000", 10, 64)
	if err == nil {
		fmt.Println(number)

	} else {
		fmt.Println(err.Error())
	}

	value := strconv.FormatInt(1000000, 10)
	fmt.Println(value)

	valueInt, _ := strconv.Atoi("2000000")
	fmt.Println(valueInt)

	valueString := strconv.Itoa(3000000)
	fmt.Println(valueString)
}

//go to https://golang.org/pkg/strconv/ for more details

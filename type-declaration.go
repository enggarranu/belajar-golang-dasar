package main

import (
	"fmt"
)

func main() {
	type NoKTP string
	type Married bool

	var noKtpEnggar NoKTP = "1238904567890"
	var marriedStatus Married = true

	fmt.Println(noKtpEnggar)
	fmt.Println(marriedStatus)
}

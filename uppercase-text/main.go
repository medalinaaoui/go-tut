package main

import (
	"fmt"
	"strings"

	"github.com/atotto/clipboard"
)

func main() {

	str := "Elevate Your Business with Expert Consultation"
	upperStr := strings.ToUpper(str)
	fmt.Println(upperStr)

	err := clipboard.WriteAll(upperStr)
	if err != nil {
		fmt.Println("Error copying to clipboard:", err)
		return
	}

	fmt.Println("Uppercase text copied to clipboard!")

}
package main

import (
	"errors"
	"fmt"
)

func main() {

	n1 := 100000
	fmt.Println(n1)

	var n2 int = -1000000000000000000
	fmt.Println(n2)

	var n3 int8 = -100
	fmt.Println(n3)

	var n4 int16 = -10000
	fmt.Println(n4)

	var n5 int32 = -1000000000
	fmt.Println(n5)

	var n6 int64 = -1000000000000000000
	fmt.Println(n6)

	var n7 uint = 1000000000000000000
	fmt.Println(n7)

	// alias = byte
	var n8 uint8 = 100
	fmt.Println(n8)

	var n9 uint16 = 10000
	fmt.Println(n9)

	// alias = rune
	var n10 uint32 = 1000000000
	fmt.Println(n10)

	var n11 uint64 = 1000000000000000000
	fmt.Println(n11)

	var n12 float32 = 1415.34
	fmt.Println(n12)

	var n13 float64 = 564061415.34
	fmt.Println(n13)

	n14 := 45089.11
	fmt.Println(n14)

	var str1 string = "Text"
	fmt.Println(str1)

	str2 := "Text 2"
	fmt.Println(str2)

	char := 'B'
	fmt.Println(char)

	var zero int16
	fmt.Println(zero)

	var boolean bool
	fmt.Println(boolean)

	var problem error = errors.New("Database Error")
	fmt.Println(problem)
}

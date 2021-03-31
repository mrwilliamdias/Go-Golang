package main

import "fmt"

func main() {
	var v1 string = "First variable"
	v2 := "Second variable"

	fmt.Println(v1)
	fmt.Println(v2)

	var (
		v3 string = "Thrid variable"
		v4 string = "Fourth variable"
	)

	fmt.Println(v3, v4)

	v5, v6 := "Fifth Variable", "Sixth Variable"

	v5, v6 = v6, v5

	fmt.Println(v5, v6)

	const c1 string = "First Constant"
	fmt.Println(c1)

}

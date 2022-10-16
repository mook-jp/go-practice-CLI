package main

import "fmt"

const (
	first1 = 200 + iota
	second1
)
const (
	first2 = iota
	second2
)

func main() {
	fmt.Println("vim-go")

	const HTTPStatusOK = 200

	const (
		StatusOk              = 0
		StatusConnectionReset = 1
		StatusOtherError      = 2
	)

	println("go begin")
	fmt.Printf("first1  = %d\n", first1)
	fmt.Printf("second1 = %d\n", second1)
	println("go end")
}

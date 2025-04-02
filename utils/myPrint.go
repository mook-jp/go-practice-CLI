package utils

import "fmt"

func redln(args ...any) {
	args = append([]any{"\033[31m"}, args...)
	fmt.Println(args...)
}

func greln(args ...any) {
	args = append([]any{"\033[32m"}, args...)
	fmt.Println(args...)
}

func redf(format string, args ...any) {
	fmt.Printf("\033[31m"+format, args...)
}

func gref(format string, args ...any) {
	fmt.Printf("\033[32m"+format, args...)
}

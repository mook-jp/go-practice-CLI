package utils

import "log"

func redln(args ...any) {
	args = append([]any{"\033[31m"}, args...)
	log.Println(args...)
}

func greln(args ...any) {
	args = append([]any{"\033[32m"}, args...)
	log.Println(args...)
}

func redf(format string, args ...any) {
	log.Printf("\033[31m"+format, args...)
}

func gref(format string, args ...any) {
	log.Printf("\033[32m"+format, args...)
}

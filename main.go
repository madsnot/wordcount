package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var count int
	if len(os.Args) > 1 {
		str := os.Args[1]
		strArray := strings.Split(str, " ")
		for _, word := range strArray {
			if word != "" {
				count++
			}
		}
	}
	fmt.Println(count)
}

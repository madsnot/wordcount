package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	str := os.Args[1]
	strArray := strings.Split(str, " ")
	fmt.Println(len(strArray))
}

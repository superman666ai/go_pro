package main

import (
	"fmt"
	"strings"
)

func main() {

	var s = []string{"phthon", "and", "go"}
	c := strings.Join(s, "--")

	fmt.Println(c)
}


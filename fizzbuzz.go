package main

import (
	"fmt"
	"strings"
)

func main() {
	for x := 1; x <= 100; x++ {
		a := []string{}
		if (x % 3) == 0 {
			a = append(a, "Fizz")
		}
		if (x % 5) == 0 {
			a = append(a, "Buzz")
		}
		if len(a) > 0 {
			fmt.Printf("%s\n", strings.Join(a, ""))
		}
	}
}

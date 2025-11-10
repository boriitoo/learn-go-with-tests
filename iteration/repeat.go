package iteration

import (
	"fmt"
	"strings"
)

func Repeat(c string, n int) string {
	var repeated strings.Builder
	for i := 0; i < n; i++ {
		repeated.WriteString(c)
	}
	return repeated.String()
}

func ExampleRepeat() {
	fmt.Println("We want to print a character n times.")
	outcome := Repeat("a", 10)
	fmt.Println(outcome)
}

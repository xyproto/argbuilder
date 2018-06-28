package argbuilder

import (
	"fmt"
)

func Example() {
	ab := New("hello")
	fmt.Println(ab)
	// Output:
	// 0: hello
}

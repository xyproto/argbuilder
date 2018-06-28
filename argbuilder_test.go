package argbuilder

import (
	"fmt"
)

func ExampleArgs() {
	ab := New("hello", "there")
	fmt.Println(ab)
	// Output:
	// 00: hello
	// 01: there
}

func ExampleRun() {
	ab := New("ls", "argbuilder_test.go")
	fmt.Println(ab.TrimOutput())
	// Output: argbuilder_test.go
}

func ExampleSplit() {
	ab := New("ls argbuilder_test.go")
	fmt.Println(ab.TrimOutput())
	// Output: argbuilder_test.go
}

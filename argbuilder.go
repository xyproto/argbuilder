// Package argbuilder can be used for collecting arguments and then execute commands
package argbuilder

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type ArgBuilder struct {
	args []string
}

// New creates a new ArgBuilder struct
func New(args ...string) *ArgBuilder {
	if len(args) == 1 && strings.Contains(args[0], " ") {
		return &ArgBuilder{strings.Split(args[0], " ")}
	}
	return &ArgBuilder{args}
}

// String returns the argument list as a string with one numbered argument per line
func (ab *ArgBuilder) String() string {
	var buf bytes.Buffer
	for i, v := range ab.args {
		buf.WriteString(fmt.Sprintf("%02d: %s\n", i, v))
	}
	return buf.String()
}

// Add adds the argument to the argument builder, without any checks or modifications
func (ab *ArgBuilder) Add(arg string) {
	ab.args = append(ab.args, arg)
}

// TrimAdd adds the argument to the argument builder.
// The argument is trimmed first. If it's empty, it's not added.
func (ab *ArgBuilder) TrimAdd(arg string) {
	if trimmed := strings.TrimSpace(arg); trimmed != "" {
		ab.args = append(ab.args, trimmed)
	}
}

func (ab *ArgBuilder) Run() error {
	var cmd *exec.Cmd
	cmd = exec.Command(ab.args[0], ab.args[1:]...)
	cmd.Stdin, cmd.Stdout, cmd.Stderr = os.Stdin, os.Stdout, os.Stderr
	return cmd.Run()
}

// Output2 executes a command and returns the combined output, or an error
func (ab *ArgBuilder) Output2() (string, error) {
	var cmdo *exec.Cmd
	if len(ab.args) == 1 {
		cmdo = exec.Command(ab.args[0])
	} else {
		cmdo = exec.Command(ab.args[0], ab.args[1:]...)
	}
	if b, err := cmdo.CombinedOutput(); err != nil {
		return "", err
	} else {
		return string(b), nil
	}
}

// Output executes a command and returns the combined output
func (ab *ArgBuilder) Output() string {
	output, err := ab.Output2()
	if err != nil {
		return ""
	}
	return output
}

// TrimOutput executs a command and returns the combined output, trimmed
func (ab *ArgBuilder) TrimOutput() string {
	output, err := ab.Output2()
	if err != nil {
		return ""
	}
	return strings.TrimSpace(output)
}

// AddStrings adds a slice of strings as arguments
func (ab *ArgBuilder) AddStrings(args []string) {
	for _, val := range args {
		ab.args = append(ab.args, val)
	}
}

// AddValues adds a map of strings as arguments
func (ab *ArgBuilder) AddValues(args map[string]string) {
	for _, val := range args {
		ab.args = append(ab.args, val)
	}
}

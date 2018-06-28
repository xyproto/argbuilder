package argbuilder

import (
	"bytes"
	"fmt"
)

type ArgBuilder struct {
	args []string
}

func New(args ...string) *ArgBuilder {
	return &ArgBuilder{args}
}

func (ab *ArgBuilder) String() string {
	var buf bytes.Buffer
	for i, v := range ab.args {
		buf.WriteString(fmt.Sprintf("%d: %s", i, v))
	}
	return buf.String()
}

package dependency_inject

import (
	"fmt"
	"io"
)

func greet(writer io.Writer, name string) {
	_, err := fmt.Fprintf(writer, "hello %s", name)
	if err != nil {
		return
	}
}

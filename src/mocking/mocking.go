package main

import (
	"fmt"
	"io"
)

func main() {

}

func countdown(writer io.Writer) {
	for i := 3;i > 0;i-- {
		fmt.Fprintf(writer, "%d\n", i)
	}
	fmt.Fprintf(writer, "GO")
}



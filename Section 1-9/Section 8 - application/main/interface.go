package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("hello")
	fmt.Fprintln(os.Stdout, "hello this is from fprintln")
	io.WriteString(os.Stdout, "hello this is from write string")
}

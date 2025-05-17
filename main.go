package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	greet(os.Stdout)
}

func greet(w io.Writer) {
	fmt.Fprintf(w, "Hello, GitHub Actions!")
}

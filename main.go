package main

import (
	"fmt"
	"io"
	"os"

	"github.com/rs/zerolog/log"
)

func main() {
	log.Info().Msg("Starting the application")
	greet(os.Stdout)
}

func greet(w io.Writer) {
	fmt.Fprintf(w, "Hello, GitHub Actions!")
}

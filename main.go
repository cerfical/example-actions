package main

import (
	"fmt"
	"io"
	"os"

	"github.com/rs/zerolog/log"
)

func main() {
	log.Info().Msg("Starting the application")
	log.Info().Msg("Oh... Wow! New feature just dropped!")
	greet(os.Stdout)
}

func greet(w io.Writer) {
	fmt.Fprintf(w, "Hello, GitHub Actions!")
}

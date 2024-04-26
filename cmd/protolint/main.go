package main

import (
	"os"

	"github.com/Jordandevpg/protolint-pg/internal/cmd"
)

func main() {
	os.Exit(int(
		cmd.Do(
			os.Args[1:],
			os.Stdout,
			os.Stderr,
		),
	))
}

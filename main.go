package main

import (
	"os"

	"github.com/Mstuart712/rm-bank/internal/cmd"
)

var version = "dev"

func main() {
	cmd.Execute(
		version,
		os.Exit,
		os.Args[1:],
	)
}

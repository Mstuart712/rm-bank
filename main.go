package main

import (
	"os"

	"github.com/mstuart712/rm-bank/internal/cmd"
)

var version = "dev"

func main() {
	cmd.Execute(
		version,
		os.Exit,
		os.Args[1:],
	)
}

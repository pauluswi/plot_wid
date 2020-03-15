package main

import (
	"flag"

	"github.com/pauluswi/plot_wid/app"
)

func main() {
	// Initiate main
	flag.Parse()
	// Check condition flag arguments
	if len(flag.Args()) > 0 {
		// Fulfill File Condition
		app.ExecuteFile(flag.Args()[0])
	} else {
		// Fulfill Command Condition
		app.ExecuteInlineCommands()
	}
}

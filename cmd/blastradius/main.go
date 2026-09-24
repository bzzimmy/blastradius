// Command blastradius verifies a leaked credential and rates its impact.
package main

import (
	"fmt"
	"os"
)

var version = "dev"

func main() {
	if len(os.Args) > 1 && os.Args[1] == "version" {
		fmt.Println("blastradius", version)
		return
	}
	fmt.Fprintln(os.Stderr, "blastradius: not yet implemented")
	os.Exit(1)
}

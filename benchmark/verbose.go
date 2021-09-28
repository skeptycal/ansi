package benchmark

import "fmt"

func verbosef(format string, args ...interface{}) {
	if config.verbose {
		fmt.Fprintf(config.output, format, args...)
	}
}

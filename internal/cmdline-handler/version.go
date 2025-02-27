package cmdline_handler

import "fmt"

func Version(version string) func() {
	return func() {
		fmt.Println("gas version", version)
	}
}

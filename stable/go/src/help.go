// help.go
package main

import "fmt"

// DisplayHelp prints the usage instructions and available commands.
func DisplayHelp() {
	fmt.Println("Usage: codersh <command>")
	fmt.Println("Commands:")
	fmt.Println("  run         - Runs a c program and removes output binary")
	fmt.Println("  run -r      - Compiles an c program into binary only.")
	fmt.Println("  --v         - Gives version")
	fmt.Println("  --h         - Gives some help")
}

package main

import (
	"fmt"
	"os"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Please provide a command (run) and a C file name.")
        return
    }

    command := os.Args[1]
    if command == "run" {
        if len(os.Args) < 3 {
            fmt.Println("Please provide the C file name to run.")
            return
        }
        runCProgram(os.Args[2]) // Call the function from run.go
    } else {
        fmt.Println("Unknown command. Use 'run' followed by the C file name.")
    }
}

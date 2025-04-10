package main

import (
	"fmt"
	"os"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Please provide a command (run or run -r) and a C file name.")
        return
    }

    command := os.Args[1]
    if command == "run" {
        if len(os.Args) < 3 {
            fmt.Println("Please provide the C file name to run.")
            return
        }
        if len(os.Args) == 4 && os.Args[2] == "--r" {
            compileOnly(os.Args[3]) // Call the function to compile only
        } else {
            runCProgram(os.Args[2]) // Call the function to compile, link, and run
        }
    } else {
        fmt.Println("Unknown command. Use 'run' followed by the C file name or 'run -r' to compile only.")
    }
}

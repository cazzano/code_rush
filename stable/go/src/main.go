package main

import (
	"fmt"
	"os"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Please provide a command (run, --h, or --v).")
        return
    }

    command := os.Args[1]
    
    switch command {
    case "run":
        if len(os.Args) < 3 {
            fmt.Println("Please provide the C file name to run.")
            return
        }
        if len(os.Args) >= 4 && os.Args[2] == "--r" {
            compileOnly(os.Args[3]) // Call the function to compile only
        } else {
            runCProgram(os.Args[2]) // Call the function to compile, link, and run
        }
    case "--h":
        DisplayHelp() // Call the help function from help.go
    case "--v":
        DisplayVersion() // Call the version function from version.go
    default:
        fmt.Println("Unknown command. Use 'run' followed by the C file name, '--h' for help, or '--v' for version information.")
    }
}

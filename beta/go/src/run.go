package main

import (
	"fmt"
	"os"
	"os/exec"
)

// runCProgram compiles, links, and runs the specified C program.
func runCProgram(fileName string) {
    // Compile the C program
    compileCmd := exec.Command("gcc", fileName, "-o", "output")
    err := compileCmd.Run()
    if err != nil {
        fmt.Printf("Error compiling %s: %v\n", fileName, err)
        return
    }

    // Run the compiled program
    runCmd := exec.Command("./output")
    runCmd.Stdout = os.Stdout
    runCmd.Stderr = os.Stderr
    err = runCmd.Run()
    if err != nil {
        fmt.Printf("Error running the program: %v\n", err)
        return
    }

    // Remove the compiled binary
    err = os.Remove("output")
    if err != nil {
        fmt.Printf("Error removing the output file: %v\n", err)
    }
}

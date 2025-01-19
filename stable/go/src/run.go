package main

import (
	"fmt"
	"os"
	"os/exec"
)

// runCProgram compiles, links, and runs the specified C program.
func runCProgram(fileName string) {
    fmt.Printf("Attempting to compile and run: %s\n", fileName)

    // Compile the C program
    compileCmd := exec.Command("gcc", fileName, "-o", "output")
    fmt.Printf("Running command: %s\n", compileCmd.String())
    err := compileCmd.Run()
    if err != nil {
        fmt.Printf("Error compiling %s: %v\n", fileName, err)
        return
    }
    fmt.Printf("Successfully compiled %s to 'output'.\n", fileName)

    // Run the compiled program
    runCmd := exec.Command("./output")
    runCmd.Stdout = os.Stdout
    runCmd.Stderr = os.Stderr
    fmt.Printf("Running the compiled program...\n")
    err = runCmd.Run()
    if err != nil {
        fmt.Printf("Error running the program: %v\n", err)
        return
    }
    fmt.Println("Program executed successfully.")

    // Remove the compiled binary
    err = os.Remove("output")
    if err != nil {
        fmt.Printf("Error removing the output file: %v\n", err)
    } else {
        fmt.Println("Removed the output file successfully.")
    }
}

// compileOnly compiles the specified C program without running it.
func compileOnly(fileName string) {
    fmt.Printf("Attempting to compile only: %s\n", fileName)

    compileCmd := exec.Command("gcc", fileName, "-o", "output")
    fmt.Printf("Running command: %s\n", compileCmd.String())
    err := compileCmd.Run()
    if err != nil {
        fmt.Printf("Error compiling %s: %v\n", fileName, err)
        return
    }
    fmt.Printf("Compiled %s successfully. Output binary is 'output'.\n", fileName)
}

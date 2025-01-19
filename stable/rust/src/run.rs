use std::env;
use std::fs;
use std::process::Command;

pub fn run_c_program(file_name: &str) {
    println!("Attempting to compile and run: {}", file_name);

    // Compile the C program
    let compile_status = Command::new("gcc")
        .arg(file_name)
        .arg("-o")
        .arg("output")
        .status()
        .expect("Failed to execute compile command");

    if !compile_status.success() {
        println!("Error compiling {}: {}", file_name, compile_status);
        return;
    }
    println!("Successfully compiled {} to 'output'.", file_name);

    // Run the compiled program
    let run_status = Command::new("./output")
        .status()
        .expect("Failed to execute run command");

    if !run_status.success() {
        println!("Error running the program: {}", run_status);
        return;
    }
    println!("Program executed successfully.");

    // Remove the compiled binary
    if let Err(err) = fs::remove_file("output") {
        println!("Error removing the output file: {}", err);
    } else {
        println!("Removed the output file successfully.");
    }
}

pub fn compile_only(file_name: &str) {
    println!("Attempting to compile only: {}", file_name);

    let compile_status = Command::new("gcc")
        .arg(file_name)
        .arg("-o")
        .arg("output")
        .status()
        .expect("Failed to execute compile command");

    if !compile_status.success() {
        println!("Error compiling {}: {}", file_name, compile_status);
        return;
    }
    println!(
        "Compiled {} successfully. Output binary is 'output'.",
        file_name
    );
}

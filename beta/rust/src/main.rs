mod run; // Import the run module

use std::env;

fn main() {
    let args: Vec<String> = env::args().collect();

    if args.len() < 2 {
        println!("Please provide a command (run or run -r) and a C file name.");
        return;
    }

    let command = &args[1];
    if command == "run" {
        if args.len() < 3 {
            println!("Please provide the C file name to run.");
            return;
        }
        if args.len() == 4 && args[2] == "-r" {
            run::compile_only(&args[3]); // Call the function to compile only
        } else {
            run::run_c_program(&args[2]); // Call the function to compile, link, and run
        }
    } else {
        println!(
            "Unknown command. Use 'run' followed by the C file name or 'run -r' to compile only."
        );
    }
}

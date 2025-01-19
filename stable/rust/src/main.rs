mod help; // Import the help module
mod run; // Import the run module
mod version; // Import the version module

use std::env;

fn main() {
    let args: Vec<String> = env::args().collect();

    if args.len() < 2 {
        println!("Please provide a command (run, run -r, -h, or -v) and a C file name.");
        return;
    }

    let command = &args[1];
    match command.as_str() {
        "run" => {
            if args.len() < 3 {
                println!("Please provide the C file name to run.");
                return;
            }
            if args.len() == 4 && args[2] == "-r" {
                run::compile_only(&args[3]); // Call the function to compile only
            } else {
                run::run_c_program(&args[2]); // Call the function to compile, link, and run
            }
        }
        "-h" => {
            help::display_help(); // Call the function to display help
        }
        "-v" => {
            version::display_version(); // Call the function to display version
        }
        _ => {
            println!("Unknown command. Use 'run' followed by the C file name, '-h' for help, or '-v' for version.");
        }
    }
}

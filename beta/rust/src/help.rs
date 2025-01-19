// src/help.rs

pub fn display_help() {
    println!("Help Information:");
    println!("Usage:");
    println!("  main <command> [options]");
    println!();
    println!("Commands:");
    println!("  run <C file>       Compile and run the specified C file.");
    println!("  run -r <C file>    Compile the specified C file without running it.");
    println!("  -h                 Display this help message.");
    println!("  -v                 Display the version information.");
}

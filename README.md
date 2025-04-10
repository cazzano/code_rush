# 🛠️ CoderSH

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
**CoderSH** is a lightweight command-line tool designed to simplify C program compilation and execution.

## ✨ Features

- **Simple C Program Execution** - Compile, run, and clean up in one command
- **Compilation Only Mode** - Generate binaries without execution
- **Streamlined Workflow** - Optimize your C development experience

## 📦 Build && Installation From Source

```bash
# Clone the repository
git clone https://github.com/cazzano/code_rush.git

# Navigate to the directory
cd code_rush

# Install the tool
go build && mv main codersh && sudo mv codersh /usr/bin && echo "Installation Complete!"
```

## Installation From Release

```bash
curl -L -o codersh_v1.0_x86_64.zip https://github.com/cazzano/code_rush/releases/download/v1.0/codersh_v1.0_x86_64.zip

unzip codersh_v1.0_x86_64.zip && sudo mv codersh /usr/bin/ && rm codersh_v1.0_x86_64.zip && echo "Installation Successful!"
```

## 🚀 Usage

CoderSH simplifies your C development workflow with these commands:

```
Usage: codersh <command>
Commands:
  run         - Runs a C program and removes output binary
  run -r      - Compiles a C program into binary only
  --v         - Gives version
  --h         - Gives some help
```

## 📝 Examples

### Example 1: Compile and Run

```bash
# Compile, run, and clean up hello.c
codersh run hello.c
```

### Example 2: Compile Only

```bash
# Generate binary from hello.c without running
codersh run -r hello.c
```

### Example 3: Version Information

```bash
# Display version information
codersh --v
```

### Example 4: Help

```bash
# Show help message
codersh --h
```

## 🛠️ Command Reference

| Command | Description |
|---------|-------------|
| `run <file.c>` | Compiles and runs the C program, then removes the binary |
| `run -r <file.c>` | Compiles the C program into binary only |
| `--v` | Displays version information |
| `--h` | Shows help message |

## 📄 License

This project is licensed under the MIT License - see the LICENSE file for details.

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📞 Support

For support, please open an issue in the GitHub repository or contact the maintainers.

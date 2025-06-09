# Go README Generator: Installation Guide

This guide provides comprehensive instructions for installing the **Go README Generator**, a command-line tool designed to create `README.md` files for your projects.

## 🚀 Installation

There are two primary methods to install the tool on a Linux or macOS system. The `make` method is recommended for its simplicity.

### Prerequisites

Before you begin, ensure you have the following installed on your system:
- **Go**: Version 1.18 or newer.
- **make**: A standard build automation tool, typically pre-installed on Linux and macOS.

---

### Method 1: Using `make` (Recommended)

This is the easiest and quickest way to get the tool up and running.

1.  **Clone the Repository**
    First, download the source code from its repository.
    ```sh
    git clone https://github.com/BobOfTheHawk/Go-README-Generator.git
    cd go-readme-generator
    ```

2.  **Build and Install**
    Run the `install` command from the Makefile. This will compile the source code and move the executable to `/usr/local/bin`, making it accessible from anywhere in your terminal.
    ```sh
    sudo make install
    ```
    *You will be prompted for your password as `sudo` is required to write to system directories.*

---

### Method 2: Manual Installation

If you prefer a step-by-step approach or do not have `make`, follow these instructions.

1.  **Clone the Repository** (if not already done)
    ```sh
    git clone https://github.com/BobOfTheHawk/Go-README-Generator.git
    cd go-readme-generator
    ```

2.  **Build the Go Binary**
    Compile the project into a single executable file named `readme`.
    ```sh
    go build -o readme .
    ```

3.  **Move the Binary to Your PATH**
    To make the `readme` command globally available, move the file you just created to `/usr/local/bin`.
    ```sh
    sudo mv readme /usr/local/bin/
    ```

## ✅ Verify Your Installation

Whichever method you used, you can now verify that the tool is installed correctly. Open a new terminal window and run:

```sh
readme
```

This command should launch the interactive README generator.

## 🗑️ Uninstallation

To remove the program from your system, navigate back to the cloned repository directory and use the `uninstall` command from the Makefile.

```sh
# This must be run from the project directory
sudo make uninstall
```

This will safely remove the `readme` executable from your system.

# Makefile for the Go README Generator

# The name of the final executable binary
BINARY_NAME=readme

# The path where the binary will be installed. /usr/local/bin is the standard
# place for user-installed executables on Linux/macOS.
INSTALL_PATH=/usr/local/bin

# Default command to run when user just types `make`.
# This will build the Go project into an executable file named 'readme'.
all: build

# The 'build' command compiles the source code.
# The -o flag sets the output file name.
build:
	@echo "Building executable..."
	go build -o $(BINARY_NAME) .
	@echo "Build complete! Binary '$(BINARY_NAME)' created."

# The 'install' command moves the built binary to the system-wide command path.
# This usually requires administrator privileges, so it should be run with 'sudo'.
install: build
	@echo "Installing $(BINARY_NAME) to $(INSTALL_PATH)..."
	sudo mv $(BINARY_NAME) $(INSTALL_PATH)/$(BINARY_NAME)
	@echo "Installation complete! You can now run the 'readme' command from anywhere."

# The 'uninstall' command removes the binary from the system.
# This also requires 'sudo'.
uninstall:
	@echo "Uninstalling $(BINARY_NAME) from $(INSTALL_PATH)..."
	sudo rm -f $(INSTALL_PATH)/$(BINARY_NAME)
	@echo "Uninstall complete."

# The 'clean' command removes the compiled binary from the project folder.
clean:
	@echo "Cleaning up build artifacts..."
	rm -f $(BINARY_NAME)
	@echo "Cleanup complete."

.PHONY: all build install uninstall clean


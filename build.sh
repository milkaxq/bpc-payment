#!/bin/bash

# Set the program name (replace with your actual program name)
PROGRAM_NAME="bpcpayment"

# Define build targets for each OS
TARGET_OS=(
  "darwin/arm64"  # macOS (amd64 architecture)
  "windows/amd64" # Windows (amd64 architecture)
  "linux/amd64"   # Linux (amd64 architecture)
)

# Loop through each target OS
for os_arch in "${TARGET_OS[@]}"; do
  # Extract OS and architecture from target
  os="${os_arch%%/*}"
  arch="${os_arch##*/}"

  # Set environment variables for GOOS and GOARCH
  export GOOS="$os"
  export GOARCH="$arch"

  # Build the binary
  echo "Building for ${os}/${arch}..."
  go build -o "bin/${PROGRAM_NAME}-${os}-${arch}" .

  # Reset environment variables (optional, good practice)
  unset GOOS
  unset GOARCH
done

echo "Build completed! Find binaries in dist folder."
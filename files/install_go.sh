#!/bin/bash

# --- Colors ---
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[0;33m'
BLUE='\033[0;34m'
NC='\033[0m'    # No Color

# --- ASCII Art ---
LOGO="
   ______   __    __        ______  __    __  
  /      \ /  |  /  |      /      |/  |  /  | 
 /$$$$$$  |$$ |  $$ |     /$$$$$$/ $$ |  $$ | 
 $$ |  $$/ $$ |  $$ |     $$ |    $$ |  $$ |  
 $$ |      $$ |  $$ |     $$ |    $$ |  $$ |  
 $$ |   __ $$ |  $$ |     $$ |    $$ |  $$ |  
 $$ \__/  |$$ \__$$ |     $$ \____$$ \__$$ | 
 $$    $$/ $$    $$/      $$    $$/ $$    $$/ 
 $$$$$$/   $$$$$$/        $$$$$$/   $$$$$$/  
"
# Welcome Banner
echo ""
echo ""
echo "################################################"
echo "#     Welcome to Space Invaders on Demand!     #"
echo "#                Made by Dave 🍺               #"
echo "################################################"
echo ""

# --- Function to check if a command exists ---
command_exists() {
  command -v "$1" >/dev/null 2>&1
}

# --- Function to install Go ---
install_go() {
  echo -e "${YELLOW}------------------------------------------------${NC}"
  echo -e "${BLUE} Go is not installed. Attempting to install...${NC}"
  echo -e "${YELLOW}------------------------------------------------${NC}"
  

  # Determine OS
  if [[ "$OSTYPE" == "darwin"* ]]; then
    OS="darwin"
    echo -e "${BLUE}Detected macOS${NC}"
  elif [[ "$OSTYPE" == "linux"* ]]; then
    OS="linux"
    echo -e "${BLUE}Detected Linux${NC}"
  else
      echo -e "${RED}Unsupported OS detected: $OSTYPE${NC}"
      echo -e "${YELLOW}Please download and install go from https://golang.org/dl/${NC}"
      exit 1
  fi

  # Download Go
  GO_VERSION="1.21.5" # Change this to a different version if needed
  GO_ARCH=$(uname -m)
  GO_FILE="go$GO_VERSION.$OS-$GO_ARCH.tar.gz"
  GO_URL="https://go.dev/dl/$GO_FILE"
   echo -e "${BLUE}Downloading Go from $GO_URL${NC}"
  curl -s -O "$GO_URL"
  if [ $? -ne 0 ]; then
      echo -e "${RED}Error downloading Go.${NC}"
    exit 1
  fi

  # Verify SHA256
   GO_SHA256_DARWIN="7995040f7f5705b182407122982a9c5733e269b733e551a188f7293cb4c5128e" # Change this if version changes
    GO_SHA256_LINUX="4f61b76d7b58c2f78b28ebf08971886e34f723597996c453d7d877701492d103"  # Change this if version changes

  if [ "$OS" = "darwin" ]; then
      EXPECTED_SHA256="$GO_SHA256_DARWIN"
  elif [ "$OS" = "linux" ]; then
      EXPECTED_SHA256="$GO_SHA256_LINUX"
  fi

  ACTUAL_SHA256=$(shasum -a 256 "$GO_FILE" | awk '{print $1}')

  if [ "$ACTUAL_SHA256" != "$EXPECTED_SHA256" ]; then
      echo -e "${RED}Error: SHA256 checksum does not match.${NC}"
      echo -e "${RED}Expected: $EXPECTED_SHA256${NC}"
      echo -e "${RED}Actual:   $ACTUAL_SHA256${NC}"
      echo -e "${YELLOW}Download has been tampered with or a wrong version, please try a different version from https://go.dev/dl/${NC}"
      rm "$GO_FILE"
      exit 1
  fi

  # Extract Go
   echo -e "${BLUE}Extracting Go...${NC}"
  sudo tar -C /usr/local -xzf "$GO_FILE"
  if [ $? -ne 0 ]; then
      echo -e "${RED}Error extracting Go to /usr/local.${NC}"
      echo -e "${YELLOW}Please try to install manually from https://golang.org/dl/${NC}"
      rm "$GO_FILE"
    exit 1
  fi
  # Remove the go file after installing
  rm "$GO_FILE"

  # Set PATH (only for the current terminal session, use user's .bashrc or .zshrc for persistent PATH change)
  echo -e "${BLUE}Adding Go to PATH${NC}"
  export PATH=$PATH:/usr/local/go/bin

    echo -e "${GREEN}Go installed successfully. Please close this terminal window and open a new terminal window, for the changes to apply.${NC}"
    echo -e "${YELLOW}To install it for all users, check: https://go.dev/doc/install${NC}"

  # Test Go installation
  if command_exists go; then
    echo -e "${BLUE}Go version: $(go version)${NC}"
  else
       echo -e "${RED}Error: Go was not installed, check the output messages for errors${NC}"
       exit 1
  fi

}

echo -e "${GREEN}$LOGO${NC}"

# Check if go is installed
if ! command_exists go; then
    install_go
fi

# -------------------------- CRUCIAL CHANGE --------------------------
# Change to the project root directory
cd "$(dirname "$0")/.."  # Goes up one directory from the script's location

echo -e "${GREEN}------------------------------------------------${NC}"
echo -e "${BLUE}Installing required packages...${NC}"
echo -e "${GREEN}------------------------------------------------${NC}"

# Initialize the module (CRUCIAL - use the correct module path)
go mod init github.com/davidtkeane/invaders # Or your actual repo URL

# Tidy up the modules (This handles dependency installation)
go mod tidy

# Get all packages in current project (Installs dependencies and updates go.mod)
go get ./... # Now works correctly because we're in the project root

echo -e "${GREEN}------------------------------------------------${NC}"
echo -e "${BLUE}All required packages and dependencies have been installed.${NC}"
echo -e "${BLUE}You can now run the game by typing:${NC}"
echo -e "${YELLOW}go run .${NC}" # Correct command to run the main package
echo -e "${GREEN}------------------------------------------------${NC}"


# Run the Game
read -p "Do you want to run the game now? 🍺 [y/N] " -n 1 -r
echo    # (optional) move to a new line
if [[ $REPLY =~ ^[Yy]$ ]]
then
    echo "Running the game....🍺 "
    go run ./main.go
    echo "Thank you for playing!" # This will always print after the game finishes.
else
    echo ""
    echo "THere's a little message from our solar system  🍺 :"
    echo ""
    cat font/girlfriend.txt  # Assuming font/girlfriend.txt exists
    echo ""
    echo "Thank you for playing. 🍺"
    echo ""
fi

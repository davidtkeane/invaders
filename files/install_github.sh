#!/bin/bash

# Configuration Variables
ENV_NAME="ranger"  # Change this to the desired Conda environment name
REPO_URL="https://github.com/attreyabhatt/Space-Invaders-Pygame.git"  # GitHub repository URL
REPO_NAME="Space-Invaders-Pygame"  # Folder name after cloning the repository
SUBFOLDER=""  # Change this if there's a subfolder to go into (e.g., "subfolder_name/")
PYTHON_VERSION="3.11"  # Python version to use
TORCH_INDEX_URL="https://download.pytorch.org/whl/nightly/cpu"  # PyTorch index URL
COMMAND="main.py"  # Command to run the app
cd '/Volumes/KaliPro/Applications/Documents/Ranger_Python/Github' || exit

# Initial Message
echo -e "\033[1;32mSetting up $REPO_NAME - Let's Get Started!\033[0m"

# Function to check if a command exists
command_exists() {
  command -v "$1" > /dev/null 2>&1
}

# Function to check Git credentials
check_git_credentials() {
  echo -e "\033[1;33mChecking Git credentials...\033[0m"
  git config --global user.name
  if [ $? -ne 0 ]; then
    echo -e "\033[1;31mGit user.name is not set! Please configure it with:\033[0m"
    echo -e "\033[1;33mgit config --global user.name 'Your Name'\033[0m"
  else
    echo -e "\033[1;32mGit user.name is set.\033[0m"
  fi

  git config --global user.email
  if [ $? -ne 0 ]; then
    echo -e "\033[1;31mGit user.email is not set! Please configure it with:\033[0m"
    echo -e "\033[1;33mgit config --global user.email 'youremail@example.com'\033[0m"
  else
    echo -e "\033[1;32mGit user.email is set.\033[0m"
  fi
}

# Function to check GitHub SSH connection
check_github_connection() {
  echo -e "\033[1;33mChecking GitHub SSH connection...\033[0m"
  ssh -T git@github.com 2>&1 | grep -q "You’ve successfully authenticated"
  if [ $? -eq 0 ]; then
    echo -e "\033[1;32mGitHub SSH connection successful.\033[0m"
  else
    echo -e "\033[1;31mGitHub SSH connection failed. You may need to configure SSH keys or use HTTPS.\033[0m"
  fi
}

# Function to check and install necessary packages
check_install_package() {
  local package=$1
  if ! command_exists "$package"; then
    echo -e "\033[1;33m$package is missing. Installing...\033[0m"
    brew install "$package" || { echo "Failed to install $package. Please install it manually."; exit 1; }
  else
    echo -e "\033[1;32m$package is already installed.\033[0m"
  fi
}

# Install Git if not present
check_install_package "git"

# Check if Git credentials are configured
check_git_credentials

# Check for GitHub connection
check_github_connection

# Install Miniconda if not present
check_install_package "conda"

# Ensure Conda is initialized in the current shell
eval "$(conda shell.bash hook)"

# Activate the 'comfyui' environment
# Check if the environment exists
if ! conda env list | grep -q "$ENV_NAME"; then
  echo -e "\033[1;33mCreating the $ENV_NAME environment...\033[0m"
  conda create -n "$ENV_NAME" python="$PYTHON_VERSION" -y
  echo -e "\033[1;33mActivating the $ENV_NAME environment...\033[0m"
  conda activate "$ENV_NAME"
  echo -e "\033[1;33mInstalling PyTorch (CPU version)...\033[0m"
  pip install --pre torch torchvision torchaudio --index-url "$TORCH_INDEX_URL"
else
  echo -e "\033[1;32m$ENV_NAME environment already exists. Activating it...\033[0m"
  conda activate "$ENV_NAME"
fi

# Clone the repository if it doesn't exist
if [ ! -d "$REPO_NAME" ]; then
  echo -e "\033[1;33mCloning the repository $REPO_URL...\033[0m"
  git clone "$REPO_URL"
else
  echo -e "\033[1;32mRepository $REPO_NAME already exists. Skipping clone.\033[0m"
fi

# Navigate to the repository directory
cd "$REPO_NAME" || exit 1

# If there is a subfolder, change to that directory
if [ -n "$SUBFOLDER" ]; then
  echo -e "\033[1;33mChanging to subfolder $SUBFOLDER...\033[0m"
  cd "$SUBFOLDER" || { echo "Subfolder $SUBFOLDER does not exist."; exit 1; }
fi

# Install Python dependencies
echo -e "\033[1;33mInstalling Python dependencies from requirements.txt...\033[0m"
pip install -r requirements.txt

# Run the application
echo -e "\033[1;35mStarting $REPO_NAME... This may take a while to load.\033[0m"
python "$COMMAND"

echo -e "\033[1;32mSuccessfully activated the 'comfyui' environment.\033[0m"
echo -e "\033[1;32m$REPO_NAME is now running. Enjoy!\033[0m"

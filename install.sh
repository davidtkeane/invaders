#!/bin/bash

# Check if go is installed
if ! command -v go &> /dev/null
then
    echo "Go is not installed. Please download and install go from https://golang.org/dl/"
    exit 1
fi

# download dependencies
go get github.com/golang/freetype/truetype
go mod download golang.org/x/image
go get github.com/hajimehoshi/ebiten/v2
go mod tidy
go mod init space_invaders
go get github.com/disintegration/gift github.com/nsf/termbox-go


echo "All required packages and dependencies have been installed. You can now run the game by typing:"
echo "go run ./main.go"
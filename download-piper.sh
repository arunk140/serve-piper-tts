#!/bin/bash

# Function to get the machine architecture
get_architecture() {
    case $(uname -m) in
        x86_64) echo "amd64" ;;
        aarch64) echo "arm64" ;;
        armv7l) echo "armv7" ;;
    esac
}

# Get the machine architecture
architecture=$(get_architecture)

# GitHub URL
url="https://github.com/rhasspy/piper/releases/latest"

# Get the redirect URL of the latest release
redirect_url=$(curl -sL -w %{url_effective} -o /dev/null $url)

# Extract the release tag from the redirect URL
release_tag=$(basename $redirect_url)

# Construct the release file URL
release_file_url="https://github.com/rhasspy/piper/releases/download/$release_tag/piper_$architecture.tar.gz"

# Download the release file
curl -L -o piper.tar.gz $release_file_url

mkdir bin

# Extract the tar gz archive
tar -xzf piper.tar.gz -C bin

# Clean up the downloaded archive
rm piper.tar.gz

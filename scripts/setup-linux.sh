#!/bin/bash


# Check if the correct number of arguments is provided
# if [ "$#" -ne 2 ]; then
#     echo "Usage: $0 <remotes-service-file-url> <executable-path>"
#     exit 1
# fi

# Variables
#SERVICE_FILE_REMOTE="$1"
SERVICE_NAME=goml.service
SERVICE_FILE_LOCAL="./scripts/service/$SERVICE_NAME"
WORKING_DIRECTORY="/etc/mlayer/"
EXECUTABLE_PATH=$WORKING_DIRECTORY/goml
USERNAME=$(whoami)

# Check if the working directory exists, if not create it
if [ ! -d "$WORKING_DIRECTORY" ]; then
    echo "Creating working directory: $WORKING_DIRECTORY"
    mkdir -p "$WORKING_DIRECTORY"
    if [[ $? -ne 0 ]]; then
        echo "Failed to create working directory."
        exit 1
    fi
fi

env GOOS=linux GOARCH=amd64 CGO_ENABLED=1  CXX="x86_64-linux-musl-g++" go build -v -o $EXECUTABLE_PATH  .

# Function to replace placeholders in the service file
replace_placeholders() {
    sed -i "s|/path/to/executable|$EXECUTABLE_PATH|g" $SERVICE_FILE_LOCAL
    sed -i "s|/path/to/working_dir|$WORKING_DIRECTORY|g" $SERVICE_FILE_LOCAL
    sed -i "s|username|$USERNAME|g" $SERVICE_FILE_LOCAL
}

# # Download the service file
# curl -o $SERVICE_FILE_LOCAL $SERVICE_FILE_REMOTE

# # Check if the download was successful
# if [[ $? -ne 0 ]]; then
#     echo "Failed to download the service file."
#     exit 1
# fi

# Replace placeholders in the service file
replace_placeholders

# Reload systemd to recognize the new service
sudo systemctl daemon-reload

# Enable the service to start on boot
sudo systemctl enable $SERVICE_NAME

# Start the service
sudo systemctl start $SERVICE_NAME

# Check the status of the service
sudo systemctl status $SERVICE_NAME
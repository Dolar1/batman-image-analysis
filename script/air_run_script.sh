#!/bin/bash

# Get the PID of the process using port 8080
PID=$(lsof -ti :8080)

# Check if a process was found and kill it
if [ -n "$PID" ]; then
    echo "Killing process $PID using port 8080..."
    kill -9 $PID
else
    echo "No process is using port 8080."
fi

# Wait for a moment (optional)
sleep 2

# Start the Go application
echo "Starting the Go application..."
go run cmd/main.go start
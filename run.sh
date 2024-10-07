#!/bin/bash

# Function to kill the servers if already running
cleanup_existing_processes() {
  echo "Cleaning up any existing compute and auth servers..."

  # Find and kill the process running on Render's default PORT or specific port for services
  compute_pid=$(lsof -t -i :6942)
  if [ -n "$compute_pid" ]; then
    kill $compute_pid
    echo "Terminated existing compute server (PID: $compute_pid)"
  else
    echo "No existing process found on port 6942"
  fi

  auth_pid=$(lsof -t -i :6969)
  if [ -n "$auth_pid" ]; then
    kill $auth_pid
    echo "Terminated existing auth server (PID: $auth_pid)"
  else
    echo "No existing process found on port 6969"
  fi
}

# Ensure cleanup of old processes before starting new ones
cleanup_existing_processes

# Start Nginx
./nginx_control.sh start

# Function to kill the servers on exit
cleanup() {
  echo "Terminating compute and auth servers..."

  compute_pid=$(lsof -t -i :6942)
  if [ -n "$compute_pid" ]; then
    kill $compute_pid
    echo "Terminated compute server (PID: $compute_pid)"
  else
    echo "No process found on port 6942"
  fi

  auth_pid=$(lsof -t -i :6969)
  if [ -n "$auth_pid" ]; then
    kill $auth_pid
    echo "Terminated auth server (PID: $auth_pid)"
  else
    echo "No process found on port 6969"
  fi

  # Stop Nginx
  ./nginx_control.sh stop
}

# Trap SIGINT (Ctrl+C) and SIGTERM signals to run cleanup
trap cleanup SIGINT SIGTERM

# Start the compute server on Render-provided port or fallback
PORT=${PORT:-6942} cd ./compute_server && go run . &
COMPUTE_PID=$!

# Start the auth server on Render-provided port or fallback
PORT=${PORT:-6969} cd ./auth_server && go run . &
AUTH_PID=$!

# Wait for the servers to finish
wait $COMPUTE_PID
wait $AUTH_PID

# Cleanup before exiting
cleanup

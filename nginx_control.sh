#!/bin/bash

# Function to kill any existing compute and auth servers
cleanup_existing_processes() {
  echo "Cleaning up any existing compute and auth servers..."

  # Loop through the compute server ports and kill the processes if running
  for port in 6942 6943 6944; do
    pid=$(lsof -t -i :$port)
    if [ -n "$pid" ]; then
      kill $pid
      echo "Terminated existing compute server on port $port (PID: $pid)"
    else
      echo "No existing process found on port $port"
    fi
  done

  # Find and kill the auth server process on port 6969
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

# Start the compute servers on ports 6942, 6943, and 6944
echo "Starting compute servers..."
PORT=6942 cd ./compute_server && go run . &
PORT=6943 cd ./compute_server && go run . &
PORT=6944 cd ./compute_server && go run . &

# Start the auth server on port 6969
echo "Starting auth server..."
PORT=6969 cd ./auth_server && go run . &

# Start Nginx using nginx_control.sh
echo "Starting Nginx..."
./nginx_control.sh start

# Function to kill all servers and stop Nginx on script exit
cleanup() {
  echo "Terminating all servers..."

  # Terminate compute servers
  for port in 6942 6943 6944; do
    pid=$(lsof -t -i :$port)
    if [ -n "$pid" ]; then
      kill $pid
      echo "Terminated compute server on port $port (PID: $pid)"
    else
      echo "No process found on port $port"
    fi
  done

  # Terminate auth server
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

# Wait for all background processes (servers) to finish
wait

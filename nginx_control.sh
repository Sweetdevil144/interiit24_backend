#!/bin/bash

# Check if nginx is installed, and if not, install it (only for Linux)
if [[ "$(uname)" == "Linux" ]]; then
  if ! command -v nginx &> /dev/null; then
    echo "Nginx not found. Installing Nginx..."
    sudo apt-get update && sudo apt-get install -y nginx
  fi
fi
command="$1"
if [[ $command == "start" ]]; then
  nginx -c $(pwd)/nginx/nginx.conf && echo "Started Service"
elif [[ $command == "stop" ]]; then
  nginx -s stop && echo "Stopping Service"
elif [[ $command == "reload" ]]; then
  nginx -s reload -c $(pwd)/nginx/nginx.conf
else
  echo "Usage: $0 <start|stop|reload>"
  exit 1
fi

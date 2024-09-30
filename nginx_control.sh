#!/bin/bash

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

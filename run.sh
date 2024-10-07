#!/bin/bash

# Start Nginx
./nginx_control.sh start

# Start the compute server on port 6942
PORT=6942 ./compute_server &

# Start the auth server on port 6969
PORT=6969 ./auth_server &

wait

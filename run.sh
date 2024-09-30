./nginx_control.sh start

# For the compute server on port 6942
PORT=6942 cd ./compute_server && go run . && cd ../ &

# For the auth server on port 6969
PORT=6969 cd ./auth_server && go run . && cd ../ &

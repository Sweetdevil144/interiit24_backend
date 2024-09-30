# Readme

Run the server using nginx as follow :

```bash
# Provide necessary permissions
chmod +x nginx_control.sh && chmod +x run.sh
./run.sh
```

to stop/reload the nginx service, kill the servers and kill/reload nginx using the below command :

```bash
./nginx_control.sh stop
```

Alternatively you can run the below commands to start, stop or reload the nginx configuration locally.

```bash
nginx -s stop -c $(pwd)/nginx/nginx.conf # Stop the service.
nginx -s start -c $(pwd)/nginx/nginx.conf # Start the service.
nginx -s reload -c $(pwd)/nginx/nginx.conf # Reload the service.
```

Frontend link: [https://github.com/the-dg04/interiit24_frontend](https://github.com/the-dg04/interiit24_frontend)

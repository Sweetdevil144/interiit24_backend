# Use the official Nginx image
FROM nginx:1.21-alpine

# Copy your custom Nginx configuration file
COPY nginx.conf /etc/nginx/nginx.conf

# Expose the Nginx default port (80)
EXPOSE 80

# Start Nginx with the configuration
CMD ["nginx", "-g", "daemon off;"]

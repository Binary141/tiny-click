# tiny-click

This is a very basic url shortener. Currently has no optimizations for speed or customization of the shortened link. 


# Deployment

For the most part you just need to run `npm run build` to get the static assets built, then run `docker compose up -d --build` to start all the containers. A db, backend, and two different frontend containers will be spun up.

# Containers

The backend exposes two ports, `5000` and `5001`, where the former is only used to accept shortened links and redirect to the chosen destination, and the latter is used for management (creating, deleting, and updating redirects).

There exist two containers for the frontend, one is running Apache and serves the built static site running listening on port `80`. The other container runs the vite dev instance and provides hot reloading to the client and listens on port `8080`. 

# Env
```
VITE_SERVER_URL="<url of the backend running on port 5000>"
VITE_ADMIN_SERVER_URL="<url of the management backend running on port 5001>"
```

# Nginx config
```
server {
    listen 443;
    server_name server.tiny-click.com;
    location / { 
        proxy_read_timeout 1;
        proxy_connect_timeout 1;
        proxy_send_timeout 1;
        send_timeout 1;

        include /etc/nginx/mime.types;
        proxy_pass http://192.168.24.105:5001/;
    }   
}

server {
    listen 80; 
    server_name server.tiny-click.com;
    location / { 
        proxy_read_timeout 1;
        proxy_connect_timeout 1;
        proxy_send_timeout 1;
        send_timeout 1;

        include /etc/nginx/mime.types;
        proxy_pass http://192.168.24.105:5001/;
    }   
}

server {
    listen 443;
    server_name admin.tiny-click.com;
    location / {
        proxy_read_timeout 1;
        proxy_connect_timeout 1;
        proxy_send_timeout 1;
        send_timeout 1;

        include /etc/nginx/mime.types;
        proxy_pass http://192.168.24.105:8080/;
    }
}

server {
    listen 80;
    server_name admin.tiny-click.com;
    location / {
        include /etc/nginx/mime.types;
        proxy_pass http://192.168.24.105:8080/;
    }
}

server {
    listen 80;
    server_name tiny-click.com;

    location / {
        expires -1;
        add_header Cache-Control 'no-store, no-cache';
        return 301 http://192.168.24.105:5000$request_uri;
    }
}

server {
    listen 443 ssl;
    server_name tiny-click.com;

    ssl_certificate /etc/nginx/conf.d/fullchain-wildcard.key;
    ssl_certificate_key /etc/nginx/conf.d/privkey-wildcard.pem;

    location / {
        expires -1;
        add_header Cache-Control 'no-store, no-cache';
        return 301 http://192.168.24.105:5000$request_uri;
    }
}
```

This is a pretty basic Nginx config so that the redirect to the destination works, and can navigate to the frontend and management ports


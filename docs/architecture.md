# Architecture

## Frontend

The frontend is a Vue project. You can refer to [frontend](../frontend/) directory for its code.

### Frontend proxy

The frontend uses an nginx proxy to serve frontend and backend under same URL:

```nginx
events {
}
http {
    include mime.types;
    default_type application/octet-stream;

    server {
        listen      80;
        listen      [::]:80;
        server_name _;

        include /etc/nginx/conf.d/*.variable;

        # logging
        access_log  /var/log/nginx/access.log combined buffer=512k flush=1m;
        error_log   /var/log/nginx/error.log warn;

        location / {
            root /var/www/html/public;
            index  index.html;
            try_files $uri $uri/ /index.html;
        }

        # for ingress-auth
        location  /ingress-auth {
            rewrite ^/ingress-auth/(.*)$ /v1/ingress-auth/$1 redirect;
        }

        # reverse proxy
        location  ~ /v1/(.*) {
             resolver 127.0.0.1 10.96.0.10 127.0.0.11 1.1.1.1;
             proxy_connect_timeout 2s;
             proxy_read_timeout 600s;
             proxy_send_timeout 600s;
             proxy_pass ${AUTHPROXY_BACKEND_URL};
             proxy_set_header        Host    $host:80;
             proxy_set_header        X-Real-IP       $remote_addr;
             proxy_set_header        X-Forwarded-For $proxy_add_x_forwarded_for;
             client_max_body_size    1000m;
         }
    }
}
```

- The `/v1/*` is proxied to backend `/v1/*`
- The backend is specified in `AUTHPROXY_BACKEND_URL` variable in `/etc/nginx/conf.d/env.variable`

In container deployment, this variable is set by an `entrypoint.sh` script

```shell
#!/bin/bash

# shellcheck disable=SC2016
echo "set \$AUTHPROXY_BACKEND_URL $AUTHPROXY_BACKEND_URL ;" > /etc/nginx/conf.d/env.variable

nginx -g 'daemon off;'
```

### Frontend APIs

| Route            | Description                                  |
| ---------------- | -------------------------------------------- |
| `/`              | Login Page                                   |
| `/passwordreset` | Self-serviced password reset page            |
| `/admin`         | Admin page, only accessible to system admins |
| `/v1`            | Proxied to backend server                    |

## Backend

Backend server API 功能

| Route                                  | Description                                                                        | Authentication            | Access Policy |
| -------------------------------------- | ---------------------------------------------------------------------------------- | ------------------------- | ------------- |
| `/v1/jwt/<login,refresh>`              | JWt creation                                                                       | username, password (POST) | admin only    |
| `/admin/users/*`                       | User CRUD                                                                          | JWT                       | admin only    |
| `/admin/policies/*`                    | Policy CRUD                                                                        | JWT                       | admin only    |
| `/admin/server/<option,shutdown,sync>` | Server Ctl                                                                         | JWt                       | admin only    |
| `/v1/ingress-auth/:resource`           | BasicAuth interface for each resource                                              | BasicAuth                 | every user    |
| `/v1/user/:name`                       | User CRUD for normal user, only accept GET/UPDATE operations of authenticated user | username, password (POST) | every user    |
| `/v1/healthz`                          | Health detection                                                                   | none                      | everyone      |
| `/v1/version`                          | Get server version                                                                 | none                      | everyone      |
| `/healthz`                             | Health detection \[duplicated\]                                                    | none                      | everyone      |
| `/version`                             | Get server version \[duplicated\]                                                  | none                      | everyone      |

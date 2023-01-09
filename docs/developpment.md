# Development

Here are some guide for debugging this project

## Backend

The backend is a standard Go project, its entrance is at `cmd/authproxy/app.go`. To prepare the database, use `manifests/mysql/docker-compose.yml`:

```shell
cd manifests/mysql
docker-compose up -d
```

Once the mysql server has started, you can use `manifests/mysql/init.sql` to populate it with default tables

```shell
cat manifests/mysql/init.sql | mysql -h 127.0.0.1 -P 3306 -uauthproxy -pauthproxy
```

Before launching the backend app, export mysql options in the shell

```shell
export AUTHPROXY_DEBUG=1;AUTHPROXY_MYSQL_DATABASE=authproxy;AUTHPROXY_MYSQL_HOSTNAME=localhost;AUTHPROXY_MYSQL_PORT=3306;AUTHPROXY_NETWORK_PORT=50032;
```

Finally, launch the application with `go run cmd/authproxy/app.go`

> The default admin user will not be created the database is manually created with mysql command.

The app will listen at port specified in `$AUTHPROXY_NETWORK_PORT`, which is port `50032`

## Frontend

To build the frontend, install node and yarn, then run

```shell
cd frontend
yarn install && yarn build
```

Run the develope server:

```shell
yarn serve
```

The frontend will show up at [http://localhost:8080](http://localhost:8080)

### Developement server

The Cross-Origin Resource Sharing(CORS) protection will stop frontend to send request to a different origin, e.g. `http://localhost:50032`. Therefore, in development environment we have to configure Vue development servers.

Vue supports development server configuration to proxy backend service. All you need is to edit the `frontend/vue.config.js`:

```javascript
module.exports = {
  transpileDependencies: [
    'vuetify'
  ],
  publicPath: '/',
  devServer: {
    proxy: {
      "/v1": {
        target: 'http://localhost:50032',
        changeOrigin: true,
      }
    }
  }
}
```

You can now test frontend at [http://localhost:8080](http://localhost:8080)

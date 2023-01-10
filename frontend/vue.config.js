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
      },
      "/oidc/": {
        target: 'http://localhost:9998',
        changeOrigin: true,
      },
      "/.well-know/openid-configuration": {
        target: 'http://localhost:9998',
        changeOrigin: true,
      }
    }
  }
}

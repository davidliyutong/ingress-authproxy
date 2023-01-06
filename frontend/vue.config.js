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

module.exports = {
  devServer: {
    disableHostCheck: true,
    proxy: {
      '/api': {
        target: process.env.SERVER_HOST,
      },
      '/images/': {
        target: process.env.SERVER_HOST,
      },
      '/info': {
        target: process.env.SERVER_HOST,
      },
    },
  }
}

const webpack = require('webpack');

module.exports = {
  devtool: 'inline-source-map',

  devServer: {
    host,
    port,
    publicPath,
    inline: true,
    lazy: false,
    hot: true,
    headers: { 'Access-Control-Allow-Origin': '*' },
    watchOptions: {
      aggregateTimeout: 300,
      poll: 100,
    },
    setup(app) {
      app.get('/ping', (_req, res) => {
        res.send('pong');
      });
    },
  },

  plugins: [
    new webpack.NoEmitOnErrorsPlugin(),
    new webpack.LoaderOptionsPlugin({ minimize: false, debug: true }),
  ],
};

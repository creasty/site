var webpack = require('webpack');
var config = require('./config');
var WebpackDevServer = require('webpack-dev-server');

const port = process.env.PORT || '9100'

new WebpackDevServer(webpack(config), {
  publicPath: config.output.publicPath,
  hot: true,
  historyApiFallback: true,
  stats: {
    colors: true
  }
}).listen(port, '0.0.0.0', function (err) {
  if (err) {
    console.log(err);
  }

  console.log(`Listening at 0.0.0.0:${port}`);
});

const webpack = require('webpack');
const config = require('./config');
const WebpackDevServer = require('webpack-dev-server');

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
    return;
  }

  console.log(`Listening at 0.0.0.0:${port}`);
});

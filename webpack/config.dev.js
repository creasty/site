var path = require('path');
var webpack = require('webpack');
var autoprefixer = require('autoprefixer');

const DEV_SERVER = 'http://localhost:9100';

module.exports = {
  devtool: 'eval',
  context: path.resolve(__dirname, '../'),
  entry: [
    'webpack-dev-server/client?' + DEV_SERVER,
    'src/main.js',
  ],
  resolve: {
    extensions: ['', '.js', 'index.js'],
  },
  output: {
    path: path.resolve(__dirname, '../public/assets'),
    filename: '[name].js',
    publicPath: DEV_SERVER + '/assets/',
  },
  plugins: [
    new webpack.DefinePlugin({
      'process.env.NODE_ENV': JSON.stringify(process.env.NODE_ENV)
    }),
  ],
  module: {
    loaders: [
      {
        test: /\.js$/,
        loaders: ['babel'],
        exclude: /node_modules/,
      },
      {
        test: /\.(css|scss)$/,
        loaders: [
          'style',
          'css',
          'sass',
          'import-glob-loader',
          'postcss',
        ]
      },
      {
        test: /.(png|jpg|gif|mp4)$/,
        loader: 'file?name=[path][name].[ext]',
        exclude: /node_modules/,
      },
    ],
  },
  postcss: function() {
    return [autoprefixer];
  },
};

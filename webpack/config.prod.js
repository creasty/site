var path = require('path');
var webpack = require('webpack');
var autoprefixer = require('autoprefixer');
var ExtractTextPlugin = require("extract-text-webpack-plugin");

module.exports = {
  devtool: 'eval',
  context: path.resolve(__dirname, '../'),
  entry: [
    'src/main.js',
  ],
  resolve: {
    extensions: ['', '.js', 'index.js'],
  },
  output: {
    path: path.resolve(__dirname, '../public/assets'),
    filename: '[name]-[hash].js',
    publicPath: '/assets/',
  },
  plugins: [
    new webpack.DefinePlugin({
      'process.env.NODE_ENV': JSON.stringify(process.env.NODE_ENV)
    }),
    new ExtractTextPlugin("[name]-[hash].css"),
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
        loader: ExtractTextPlugin.extract('style', [
          'css',
          'sass',
          'import-glob-loader',
          'postcss',
        ]),
      },
      {
        test: /.(png|jpg|gif|mp4)$/,
        loader: 'file?name=[path][name]-[hash].[ext]',
        exclude: /node_modules/,
      }
    ],
  },
  postcss: function() {
    return [autoprefixer];
  },
  'uglify-loader': {
    compress: true,
    minimize: true,
    sourceMap: true,
    compressor: {
      warnings: false
    }
  }
};

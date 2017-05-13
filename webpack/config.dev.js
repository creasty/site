const path = require('path');
const webpack = require('webpack');
const autoprefixer = require('autoprefixer');
const HtmlWebpackPlugin = require('html-webpack-plugin');

const DEV_SERVER = 'http://localhost:9100';

module.exports = {
  devtool: 'eval',
  context: path.resolve(__dirname, '../src'),
  entry: [
    'webpack-dev-server/client?' + DEV_SERVER,
    './main.js',
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
    new HtmlWebpackPlugin({
      filename: 'index.html',
      template: 'index.html',
      inject: true,
      nodeEnv: process.env.NODE_ENV,
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

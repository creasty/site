const webpack = require('webpack');
const merge = require('webpack-merge');
const ExtractTextPlugin = require('extract-text-webpack-plugin');
const HtmlWebpackPlugin = require('html-webpack-plugin');
const path = require('path');

const srcPath = path.resolve(__dirname, '../src');

const styleExtractor = new ExtractTextPlugin({
  filename: '[name]-[hash].css',
  allChunks: true,
  disable: false,
});

const base = {
  context: srcPath,
  entry: [
    './main.js',
  ],
  resolve: {
    modules: [srcPath, 'node_modules'],
    extensions: ['.js', 'index.js'],
  },
  module: {
    rules: [
      {
        test: /\.js$/,
        use: [
          {
            loader: 'babel-loader',
            options: { cacheDirectory: true },
          },
        ],
        exclude: /node_modules/,
      },
      {
        test: /\.css$/,
        use: styleExtractor.extract({
          fallback: 'style-loader',
          use: [
            {
              loader: 'css-loader',
              options: { sourceMap: true },
            },
          ],
        }),
      },
      {
        test: /\.(scss|sass|css)$/,
        use: styleExtractor.extract({
          fallback: 'style-loader',
          use: [
            {
              loader: 'css-loader',
              options: {
                modules: true,
                sourceMap: true,
                localIdentName: '[local]',
                importLoaders: 1,
              },
            },
            {
              loader: 'sass-loader',
              options: { sourceMap: false },
            },
          ],
        }),
      },
      {
        test: /\.(png|jpg|gif|mp4|eot|woff2?|ttf|svg)$/,
        exclude: /node_modules/,
        use: [
          {
            loader: 'file-loader',
            options: { name: '[name]-[hash].[ext]' },
          },
        ],
      },
    ],
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

    styleExtractor,
    new HtmlWebpackPlugin({
      filename: 'index.html',
      template: 'index.html',
      inject: true,
      nodeEnv: process.env.NODE_ENV,
    }),
  ],
};

const config = (process.env.NODE_ENV === 'production')
  ? require('./config.prod')
  : require('./config.dev');

module.exports = merge(base, config);

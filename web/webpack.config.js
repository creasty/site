'use strict';

var BowerWebpackPlugin = require('bower-webpack-plugin');
var path               = require('path');

module.exports = {
    entry: [
        './src/js/index.js',
    ],
    devtool: process.env.WEBPACK_DEVTOOL || 'source-map',
    output: {
        path:     path.join(__dirname, 'static'),
        filename: 'bundle.js',
    },
    resolve: {
        extensions: ['', '.js', '.jsx'],
    },
    module: {
        noParse: /\.min\.js$/,
        loaders: [
        {
            test:    /\.css$/,
            loaders: ['style-loader', 'css-loader'],
        },
        {
            test:    /\.scss$/,
            loaders: ['style-loader', 'css-loader', 'sass-loader'],
        },
        {
            test:    /\.html$/,
            loaders: ['html-loader'],
        },
        {
            test:    /\.jsx?$/,
            exclude: /(node_modules|bower_components)/,
            loaders: ['jsx-loader?harmony=1', 'babel-loader'],
        },
        {
            test:    /\.js$/,
            exclude: /(node_modules|bower_components)/,
            loaders: ['babel-loader?stage=1'],
        },
        ]
    },
    plugins: [
        new BowerWebpackPlugin({
            excludes: /\.map$/
        }),
    ],
};

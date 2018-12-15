const webpack = require('webpack');
const merge = require('webpack-merge');

module.exports = {
  devtool: '#source-map',

  plugins: [
    new webpack.LoaderOptionsPlugin({ minimize: true, debug: false }),

    new webpack.optimize.OccurrenceOrderPlugin(),

    new webpack.optimize.UglifyJsPlugin({
      compress: {
        screw_ie8: true,
        warnings: false,
      },
      sourceMap: true,
    }),
  ],
};

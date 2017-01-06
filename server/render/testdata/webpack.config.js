var webpack = require('webpack')
var NodeSourcePlugin = require('webpack/lib/node/NodeSourcePlugin')

var config = {
  entry: './main.js',
  target: 'node',
  devtool: false,
  output: {
    filename: 'index.js',
    libraryTarget: 'commonjs2'
  },
  plugins: [
    new NodeSourcePlugin(
      {
        process: true,
        global: true,
        Buffer: true,
        setImmediate: true,
        module: 'empty',
        __filename: 'mock',
        __dirname: 'mock'
      }),
    new webpack.optimize.UglifyJsPlugin({
      compress: {
        warnings: false
      }
    }),
    new webpack.LoaderOptionsPlugin({
      minimize: true
    })
  ],
  module: {
    rules: [
      {
        test: /\.vue$/,
        loader: 'vue-loader',
        options: {}
      },
      {
        test: /\.js$/,
        loader: 'babel-loader'
      },
      {
        test: /\.json$/,
        loader: 'json-loader'
      }
    ]
  },
}

var binding = process.binding
process.binding = function (name) {
  res = name
  if (name === 'natives') return {}
  return binding.apply(process, arguments)
}

module.exports = config
const { join } = require("path");

module.exports = {
  entry: "./index.js",
  mode: "production",
  output: {
    library: "PDIGrammar",
    path: join(__dirname, "./dist"),
    filename: "grammar-pdi-umd.js",
    libraryTarget: "umd"
  },
  resolve: {
    extensions: [".js"]
  },
  node: {
    module: "empty",
    net: "empty",
    fs: "empty"
  },
  devtool: "source-map"
};
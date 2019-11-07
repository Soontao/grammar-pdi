const antlr4 = require("antlr4");

var { ABSLLexer } = require("./js/ABSLLexer");
var { ABSLParser } = require("./js/ABSLParser");
var { ProgramContext } = ABSLParser;

/**
 * Parse ABSL source ast
 *
 * @param {String} source code string
 * @returns {ProgramContext} tree
 */
const parseSource = (source = "") => {
  var is = new antlr4.InputStream(source);
  var lexer = new ABSLLexer(is);
  var tokens = new antlr4.CommonTokenStream(lexer);
  var parser = new ABSLParser(tokens);
  parser.buildParseTrees = true;
  return parser.program();
};


module.exports = { parseSource };

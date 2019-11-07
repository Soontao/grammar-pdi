const antlr4 = require("antlr4");

var { BODLParser } = require("./js/BODLParser");
var { BODLLexer } = require("./js/BODLLexer");
var { ProgramContext } = BODLParser;

/**
 * Parse BODL source ast
 *
 * @param {String} source code string
 * @returns {ProgramContext} tree
 */
const parseSource = (source = "") => {
  var is = new antlr4.InputStream(source);
  var lexer = new BODLLexer(is);
  var tokens = new antlr4.CommonTokenStream(lexer);
  var parser = new BODLParser(tokens);
  parser.buildParseTrees = true;
  return parser.program();
};


module.exports = { parseSource };

var { parseSource: parseBODLSource } = require("./bodl/BODLParser");
var { parseSource: parseABSLSource } = require("./absl/ABSLParser");

module.exports = { parseBODLSource, parseABSLSource };

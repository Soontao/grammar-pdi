var fs = require("fs");
var path = require("path");
var { parseSource } = require("./ABSLParser");

var readFile = (relativePath) => fs.readFileSync(
  path.join(__dirname, relativePath), { encoding: "UTF-8" }
);

test("should parse basic ABSL code source", () => {

  const source = readFile("./examples/basic.absl");
  const program = parseSource(source);
  const elements = program.sourceElements().sourceElement();
  expect(elements.length).toBe(2);

  const importStatement = elements[0].statement().importStatement();

  expect(importStatement.singleExpression().children[0].Identifier().getText()).toBe("ABSL");

});

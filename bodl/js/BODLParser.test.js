var fs = require("fs");
var path = require("path");
var { parseSource } = require(".");

var readFile = (relativePath) => fs.readFileSync(
  path.join(__dirname, relativePath), { encoding: "UTF-8" }
);


test("should parse basic BODL source", () => {

  var source = readFile("../examples/basic.bo");
  var program = parseSource(source);

  var importStatements = program.statements();
  var imports = importStatements.children;
  expect(imports.length).toBe(1); // single import

  var m = imports[0].memberExpression();
  expect(m.getText()).toBe("ABSL"); // import 'ABSL'

  var definitions = program.definitions().children;
  expect(definitions.length).toBe(1); // single BO definition

  var firstDefinition = definitions[0];
  var id = firstDefinition.identifier();
  expect(id.getText()).toBe("Hello"); // business object name is 'Hello'

});

test("should parse complex BODL source", () => {

  var source = readFile("../examples/complex.bo");
  var program = parseSource(source);

  var importStatements = program.statements();
  var imports = importStatements.children;
  expect(imports.length).toBe(2); // import 2 lines

  var firstImport = imports[0];
  expect(firstImport.identifier().getText()).toBe("Hello");

  var firstDefinition = program.definitions().children[0];
  var annotations = firstDefinition.annotations().children;
  expect(annotations.length).toBe(2);
  expect(firstDefinition.raiseMessage().identifierList().identifier()[0].getText()).toBe("ERROR");

  var secondAnnotation = annotations[1];
  var ids = secondAnnotation.identifier();
  var firstId = ids[0];
  var secondId = ids[1];
  expect(firstId.getText()).toBe("DeploymentUnit");
  expect(secondId.getText()).toBe("CustomerRelationshipManagement");

  var firstNode = firstDefinition.block().itemList().node()[0];
  expect(firstNode.identifier().getText()).toBe("FNDCommon");

  var firstNodeFirstSubMessage = firstNode.block().itemList().message()[0];
  var template = firstNodeFirstSubMessage.StringLiteral();
  expect(template.getText()).toBe("\"&1\"");

  var firstNodeFirstSubNode = firstNode.block().itemList().node()[0];
  expect(firstNodeFirstSubNode.identifier().getText()).toBe("Address");
  expect(firstNodeFirstSubNode.multiplicity().literal()[1].getText()).toBe("1");


});


test("should parse custom annotation BODL source", () => {

  var program = parseSource(readFile("../examples/custom_annotation.bo"));

  var firstDefinition = program.definitions().children[0];
  var annotations = firstDefinition.annotations().children;
  expect(annotations.length).toBe(3);
  expect(annotations[1].literal().children[0].getText())
    .toBe("\"https://stackoverflow.com\"");

  var elements = firstDefinition.block().itemList().element();
  expect(elements.length).toBe(2);

  var firstElementsAnnotation = elements[0].annotations().annotation();
  expect(firstElementsAnnotation.length).toBe(1);
  expect(firstElementsAnnotation[0].identifier()[0].getText()).toBe("Text");

});

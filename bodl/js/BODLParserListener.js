// Generated from ./BODLParser.g4 by ANTLR 4.7.2
// jshint ignore: start
var antlr4 = require("antlr4/index");

// This class defines a complete listener for a parse tree produced by BODLParser.
function BODLParserListener() {
  antlr4.tree.ParseTreeListener.call(this);
  return this;
}

BODLParserListener.prototype = Object.create(antlr4.tree.ParseTreeListener.prototype);
BODLParserListener.prototype.constructor = BODLParserListener;

// Enter a parse tree produced by BODLParser#program.
BODLParserListener.prototype.enterProgram = function(ctx) {
};

// Exit a parse tree produced by BODLParser#program.
BODLParserListener.prototype.exitProgram = function(ctx) {
};


// Enter a parse tree produced by BODLParser#statements.
BODLParserListener.prototype.enterStatements = function(ctx) {
};

// Exit a parse tree produced by BODLParser#statements.
BODLParserListener.prototype.exitStatements = function(ctx) {
};


// Enter a parse tree produced by BODLParser#importStatement.
BODLParserListener.prototype.enterImportStatement = function(ctx) {
};

// Exit a parse tree produced by BODLParser#importStatement.
BODLParserListener.prototype.exitImportStatement = function(ctx) {
};


// Enter a parse tree produced by BODLParser#definitions.
BODLParserListener.prototype.enterDefinitions = function(ctx) {
};

// Exit a parse tree produced by BODLParser#definitions.
BODLParserListener.prototype.exitDefinitions = function(ctx) {
};


// Enter a parse tree produced by BODLParser#definition.
BODLParserListener.prototype.enterDefinition = function(ctx) {
};

// Exit a parse tree produced by BODLParser#definition.
BODLParserListener.prototype.exitDefinition = function(ctx) {
};


// Enter a parse tree produced by BODLParser#block.
BODLParserListener.prototype.enterBlock = function(ctx) {
};

// Exit a parse tree produced by BODLParser#block.
BODLParserListener.prototype.exitBlock = function(ctx) {
};


// Enter a parse tree produced by BODLParser#itemList.
BODLParserListener.prototype.enterItemList = function(ctx) {
};

// Exit a parse tree produced by BODLParser#itemList.
BODLParserListener.prototype.exitItemList = function(ctx) {
};


// Enter a parse tree produced by BODLParser#element.
BODLParserListener.prototype.enterElement = function(ctx) {
};

// Exit a parse tree produced by BODLParser#element.
BODLParserListener.prototype.exitElement = function(ctx) {
};


// Enter a parse tree produced by BODLParser#boAction.
BODLParserListener.prototype.enterBoAction = function(ctx) {
};

// Exit a parse tree produced by BODLParser#boAction.
BODLParserListener.prototype.exitBoAction = function(ctx) {
};


// Enter a parse tree produced by BODLParser#message.
BODLParserListener.prototype.enterMessage = function(ctx) {
};

// Exit a parse tree produced by BODLParser#message.
BODLParserListener.prototype.exitMessage = function(ctx) {
};


// Enter a parse tree produced by BODLParser#node.
BODLParserListener.prototype.enterNode = function(ctx) {
};

// Exit a parse tree produced by BODLParser#node.
BODLParserListener.prototype.exitNode = function(ctx) {
};


// Enter a parse tree produced by BODLParser#association.
BODLParserListener.prototype.enterAssociation = function(ctx) {
};

// Exit a parse tree produced by BODLParser#association.
BODLParserListener.prototype.exitAssociation = function(ctx) {
};


// Enter a parse tree produced by BODLParser#associationUsingDefinition.
BODLParserListener.prototype.enterAssociationUsingDefinition = function(ctx) {
};

// Exit a parse tree produced by BODLParser#associationUsingDefinition.
BODLParserListener.prototype.exitAssociationUsingDefinition = function(ctx) {
};


// Enter a parse tree produced by BODLParser#valuationDefinition.
BODLParserListener.prototype.enterValuationDefinition = function(ctx) {
};

// Exit a parse tree produced by BODLParser#valuationDefinition.
BODLParserListener.prototype.exitValuationDefinition = function(ctx) {
};


// Enter a parse tree produced by BODLParser#valutaionExpressionList.
BODLParserListener.prototype.enterValutaionExpressionList = function(ctx) {
};

// Exit a parse tree produced by BODLParser#valutaionExpressionList.
BODLParserListener.prototype.exitValutaionExpressionList = function(ctx) {
};


// Enter a parse tree produced by BODLParser#valutaionExpression.
BODLParserListener.prototype.enterValutaionExpression = function(ctx) {
};

// Exit a parse tree produced by BODLParser#valutaionExpression.
BODLParserListener.prototype.exitValutaionExpression = function(ctx) {
};


// Enter a parse tree produced by BODLParser#raiseMessage.
BODLParserListener.prototype.enterRaiseMessage = function(ctx) {
};

// Exit a parse tree produced by BODLParser#raiseMessage.
BODLParserListener.prototype.exitRaiseMessage = function(ctx) {
};


// Enter a parse tree produced by BODLParser#annotations.
BODLParserListener.prototype.enterAnnotations = function(ctx) {
};

// Exit a parse tree produced by BODLParser#annotations.
BODLParserListener.prototype.exitAnnotations = function(ctx) {
};


// Enter a parse tree produced by BODLParser#annotation.
BODLParserListener.prototype.enterAnnotation = function(ctx) {
};

// Exit a parse tree produced by BODLParser#annotation.
BODLParserListener.prototype.exitAnnotation = function(ctx) {
};


// Enter a parse tree produced by BODLParser#typeList.
BODLParserListener.prototype.enterTypeList = function(ctx) {
};

// Exit a parse tree produced by BODLParser#typeList.
BODLParserListener.prototype.exitTypeList = function(ctx) {
};


// Enter a parse tree produced by BODLParser#typeDeclaration.
BODLParserListener.prototype.enterTypeDeclaration = function(ctx) {
};

// Exit a parse tree produced by BODLParser#typeDeclaration.
BODLParserListener.prototype.exitTypeDeclaration = function(ctx) {
};


// Enter a parse tree produced by BODLParser#typeDefaultValue.
BODLParserListener.prototype.enterTypeDefaultValue = function(ctx) {
};

// Exit a parse tree produced by BODLParser#typeDefaultValue.
BODLParserListener.prototype.exitTypeDefaultValue = function(ctx) {
};


// Enter a parse tree produced by BODLParser#valueAssignList.
BODLParserListener.prototype.enterValueAssignList = function(ctx) {
};

// Exit a parse tree produced by BODLParser#valueAssignList.
BODLParserListener.prototype.exitValueAssignList = function(ctx) {
};


// Enter a parse tree produced by BODLParser#valueAssign.
BODLParserListener.prototype.enterValueAssign = function(ctx) {
};

// Exit a parse tree produced by BODLParser#valueAssign.
BODLParserListener.prototype.exitValueAssign = function(ctx) {
};


// Enter a parse tree produced by BODLParser#multiplicity.
BODLParserListener.prototype.enterMultiplicity = function(ctx) {
};

// Exit a parse tree produced by BODLParser#multiplicity.
BODLParserListener.prototype.exitMultiplicity = function(ctx) {
};


// Enter a parse tree produced by BODLParser#memberExpression.
BODLParserListener.prototype.enterMemberExpression = function(ctx) {
};

// Exit a parse tree produced by BODLParser#memberExpression.
BODLParserListener.prototype.exitMemberExpression = function(ctx) {
};


// Enter a parse tree produced by BODLParser#identifierList.
BODLParserListener.prototype.enterIdentifierList = function(ctx) {
};

// Exit a parse tree produced by BODLParser#identifierList.
BODLParserListener.prototype.exitIdentifierList = function(ctx) {
};


// Enter a parse tree produced by BODLParser#keyword.
BODLParserListener.prototype.enterKeyword = function(ctx) {
};

// Exit a parse tree produced by BODLParser#keyword.
BODLParserListener.prototype.exitKeyword = function(ctx) {
};


// Enter a parse tree produced by BODLParser#literal.
BODLParserListener.prototype.enterLiteral = function(ctx) {
};

// Exit a parse tree produced by BODLParser#literal.
BODLParserListener.prototype.exitLiteral = function(ctx) {
};


// Enter a parse tree produced by BODLParser#comments.
BODLParserListener.prototype.enterComments = function(ctx) {
};

// Exit a parse tree produced by BODLParser#comments.
BODLParserListener.prototype.exitComments = function(ctx) {
};


// Enter a parse tree produced by BODLParser#compareOperator.
BODLParserListener.prototype.enterCompareOperator = function(ctx) {
};

// Exit a parse tree produced by BODLParser#compareOperator.
BODLParserListener.prototype.exitCompareOperator = function(ctx) {
};


// Enter a parse tree produced by BODLParser#logicOperator.
BODLParserListener.prototype.enterLogicOperator = function(ctx) {
};

// Exit a parse tree produced by BODLParser#logicOperator.
BODLParserListener.prototype.exitLogicOperator = function(ctx) {
};


// Enter a parse tree produced by BODLParser#identifier.
BODLParserListener.prototype.enterIdentifier = function(ctx) {
};

// Exit a parse tree produced by BODLParser#identifier.
BODLParserListener.prototype.exitIdentifier = function(ctx) {
};



exports.BODLParserListener = BODLParserListener;
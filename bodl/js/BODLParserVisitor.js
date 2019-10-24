// Generated from ./BODLParser.g4 by ANTLR 4.7.2
// jshint ignore: start
var antlr4 = require("antlr4/index");

// This class defines a complete generic visitor for a parse tree produced by BODLParser.

function BODLParserVisitor() {
  antlr4.tree.ParseTreeVisitor.call(this);
  return this;
}

BODLParserVisitor.prototype = Object.create(antlr4.tree.ParseTreeVisitor.prototype);
BODLParserVisitor.prototype.constructor = BODLParserVisitor;

// Visit a parse tree produced by BODLParser#program.
BODLParserVisitor.prototype.visitProgram = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by BODLParser#statements.
BODLParserVisitor.prototype.visitStatements = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by BODLParser#importStatement.
BODLParserVisitor.prototype.visitImportStatement = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by BODLParser#definitions.
BODLParserVisitor.prototype.visitDefinitions = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by BODLParser#definition.
BODLParserVisitor.prototype.visitDefinition = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by BODLParser#block.
BODLParserVisitor.prototype.visitBlock = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by BODLParser#itemList.
BODLParserVisitor.prototype.visitItemList = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by BODLParser#element.
BODLParserVisitor.prototype.visitElement = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by BODLParser#boAction.
BODLParserVisitor.prototype.visitBoAction = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by BODLParser#message.
BODLParserVisitor.prototype.visitMessage = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by BODLParser#node.
BODLParserVisitor.prototype.visitNode = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by BODLParser#association.
BODLParserVisitor.prototype.visitAssociation = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by BODLParser#associationUsingDefinition.
BODLParserVisitor.prototype.visitAssociationUsingDefinition = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by BODLParser#valuationDefinition.
BODLParserVisitor.prototype.visitValuationDefinition = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by BODLParser#valutaionExpressionList.
BODLParserVisitor.prototype.visitValutaionExpressionList = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by BODLParser#valutaionExpression.
BODLParserVisitor.prototype.visitValutaionExpression = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by BODLParser#raiseMessage.
BODLParserVisitor.prototype.visitRaiseMessage = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by BODLParser#annotations.
BODLParserVisitor.prototype.visitAnnotations = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by BODLParser#annotation.
BODLParserVisitor.prototype.visitAnnotation = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by BODLParser#typeList.
BODLParserVisitor.prototype.visitTypeList = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by BODLParser#typeDeclaration.
BODLParserVisitor.prototype.visitTypeDeclaration = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by BODLParser#typeDefaultValue.
BODLParserVisitor.prototype.visitTypeDefaultValue = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by BODLParser#valueAssignList.
BODLParserVisitor.prototype.visitValueAssignList = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by BODLParser#valueAssign.
BODLParserVisitor.prototype.visitValueAssign = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by BODLParser#multiplicity.
BODLParserVisitor.prototype.visitMultiplicity = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by BODLParser#memberExpression.
BODLParserVisitor.prototype.visitMemberExpression = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by BODLParser#identifierList.
BODLParserVisitor.prototype.visitIdentifierList = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by BODLParser#keyword.
BODLParserVisitor.prototype.visitKeyword = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by BODLParser#literal.
BODLParserVisitor.prototype.visitLiteral = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by BODLParser#comments.
BODLParserVisitor.prototype.visitComments = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by BODLParser#compareOperator.
BODLParserVisitor.prototype.visitCompareOperator = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by BODLParser#logicOperator.
BODLParserVisitor.prototype.visitLogicOperator = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by BODLParser#identifier.
BODLParserVisitor.prototype.visitIdentifier = function(ctx) {
  return this.visitChildren(ctx);
};



exports.BODLParserVisitor = BODLParserVisitor;
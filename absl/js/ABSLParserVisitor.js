// Generated from ./ABSLParser.g4 by ANTLR 4.7.2
// jshint ignore: start
var antlr4 = require("antlr4/index");

// This class defines a complete generic visitor for a parse tree produced by ABSLParser.

function ABSLParserVisitor() {
  antlr4.tree.ParseTreeVisitor.call(this);
  return this;
}

ABSLParserVisitor.prototype = Object.create(antlr4.tree.ParseTreeVisitor.prototype);
ABSLParserVisitor.prototype.constructor = ABSLParserVisitor;

// Visit a parse tree produced by ABSLParser#program.
ABSLParserVisitor.prototype.visitProgram = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by ABSLParser#sourceElement.
ABSLParserVisitor.prototype.visitSourceElement = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by ABSLParser#statement.
ABSLParserVisitor.prototype.visitStatement = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by ABSLParser#block.
ABSLParserVisitor.prototype.visitBlock = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by ABSLParser#statementList.
ABSLParserVisitor.prototype.visitStatementList = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by ABSLParser#importStatement.
ABSLParserVisitor.prototype.visitImportStatement = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by ABSLParser#importFromBlock.
ABSLParserVisitor.prototype.visitImportFromBlock = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by ABSLParser#moduleItems.
ABSLParserVisitor.prototype.visitModuleItems = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by ABSLParser#importDefault.
ABSLParserVisitor.prototype.visitImportDefault = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by ABSLParser#importNamespace.
ABSLParserVisitor.prototype.visitImportNamespace = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by ABSLParser#importFrom.
ABSLParserVisitor.prototype.visitImportFrom = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by ABSLParser#aliasName.
ABSLParserVisitor.prototype.visitAliasName = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by ABSLParser#ExportDeclaration.
ABSLParserVisitor.prototype.visitExportDeclaration = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by ABSLParser#ExportDefaultDeclaration.
ABSLParserVisitor.prototype.visitExportDefaultDeclaration = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by ABSLParser#exportFromBlock.
ABSLParserVisitor.prototype.visitExportFromBlock = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by ABSLParser#declaration.
ABSLParserVisitor.prototype.visitDeclaration = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by ABSLParser#variableStatement.
ABSLParserVisitor.prototype.visitVariableStatement = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by ABSLParser#variableDeclarationList.
ABSLParserVisitor.prototype.visitVariableDeclarationList = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by ABSLParser#variableDeclaration.
ABSLParserVisitor.prototype.visitVariableDeclaration = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by ABSLParser#variableType.
ABSLParserVisitor.prototype.visitVariableType = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by ABSLParser#emptyStatement.
ABSLParserVisitor.prototype.visitEmptyStatement = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by ABSLParser#expressionStatement.
ABSLParserVisitor.prototype.visitExpressionStatement = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by ABSLParser#ifStatement.
ABSLParserVisitor.prototype.visitIfStatement = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by ABSLParser#DoStatement.
ABSLParserVisitor.prototype.visitDoStatement = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by ABSLParser#WhileStatement.
ABSLParserVisitor.prototype.visitWhileStatement = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by ABSLParser#ForStatement.
ABSLParserVisitor.prototype.visitForStatement = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by ABSLParser#ForInStatement.
ABSLParserVisitor.prototype.visitForInStatement = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by ABSLParser#ForOfStatement.
ABSLParserVisitor.prototype.visitForOfStatement = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by ABSLParser#ForeachStatement.
ABSLParserVisitor.prototype.visitForeachStatement = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by ABSLParser#varModifier.
ABSLParserVisitor.prototype.visitVarModifier = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by ABSLParser#continueStatement.
ABSLParserVisitor.prototype.visitContinueStatement = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by ABSLParser#breakStatement.
ABSLParserVisitor.prototype.visitBreakStatement = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by ABSLParser#returnStatement.
ABSLParserVisitor.prototype.visitReturnStatement = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by ABSLParser#yieldStatement.
ABSLParserVisitor.prototype.visitYieldStatement = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by ABSLParser#withStatement.
ABSLParserVisitor.prototype.visitWithStatement = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by ABSLParser#switchStatement.
ABSLParserVisitor.prototype.visitSwitchStatement = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by ABSLParser#caseBlock.
ABSLParserVisitor.prototype.visitCaseBlock = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by ABSLParser#caseClauses.
ABSLParserVisitor.prototype.visitCaseClauses = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by ABSLParser#caseClause.
ABSLParserVisitor.prototype.visitCaseClause = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by ABSLParser#defaultClause.
ABSLParserVisitor.prototype.visitDefaultClause = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by ABSLParser#labelledStatement.
ABSLParserVisitor.prototype.visitLabelledStatement = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by ABSLParser#throwStatement.
ABSLParserVisitor.prototype.visitThrowStatement = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by ABSLParser#tryStatement.
ABSLParserVisitor.prototype.visitTryStatement = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by ABSLParser#catchProduction.
ABSLParserVisitor.prototype.visitCatchProduction = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by ABSLParser#finallyProduction.
ABSLParserVisitor.prototype.visitFinallyProduction = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by ABSLParser#debuggerStatement.
ABSLParserVisitor.prototype.visitDebuggerStatement = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by ABSLParser#functionDeclaration.
ABSLParserVisitor.prototype.visitFunctionDeclaration = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by ABSLParser#classDeclaration.
ABSLParserVisitor.prototype.visitClassDeclaration = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by ABSLParser#classTail.
ABSLParserVisitor.prototype.visitClassTail = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by ABSLParser#classElement.
ABSLParserVisitor.prototype.visitClassElement = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by ABSLParser#methodDefinition.
ABSLParserVisitor.prototype.visitMethodDefinition = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by ABSLParser#formalParameterList.
ABSLParserVisitor.prototype.visitFormalParameterList = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by ABSLParser#formalParameterArg.
ABSLParserVisitor.prototype.visitFormalParameterArg = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by ABSLParser#lastFormalParameterArg.
ABSLParserVisitor.prototype.visitLastFormalParameterArg = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by ABSLParser#functionBody.
ABSLParserVisitor.prototype.visitFunctionBody = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by ABSLParser#sourceElements.
ABSLParserVisitor.prototype.visitSourceElements = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by ABSLParser#arrayLiteral.
ABSLParserVisitor.prototype.visitArrayLiteral = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by ABSLParser#elementList.
ABSLParserVisitor.prototype.visitElementList = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by ABSLParser#arrayElement.
ABSLParserVisitor.prototype.visitArrayElement = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by ABSLParser#objectLiteral.
ABSLParserVisitor.prototype.visitObjectLiteral = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by ABSLParser#PropertyExpressionAssignment.
ABSLParserVisitor.prototype.visitPropertyExpressionAssignment = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by ABSLParser#ComputedPropertyExpressionAssignment.
ABSLParserVisitor.prototype.visitComputedPropertyExpressionAssignment = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by ABSLParser#FunctionProperty.
ABSLParserVisitor.prototype.visitFunctionProperty = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by ABSLParser#PropertyGetter.
ABSLParserVisitor.prototype.visitPropertyGetter = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by ABSLParser#PropertySetter.
ABSLParserVisitor.prototype.visitPropertySetter = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by ABSLParser#PropertyShorthand.
ABSLParserVisitor.prototype.visitPropertyShorthand = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by ABSLParser#propertyName.
ABSLParserVisitor.prototype.visitPropertyName = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by ABSLParser#arguments.
ABSLParserVisitor.prototype.visitArguments = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by ABSLParser#argument.
ABSLParserVisitor.prototype.visitArgument = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by ABSLParser#expressionSequence.
ABSLParserVisitor.prototype.visitExpressionSequence = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by ABSLParser#TemplateStringExpression.
ABSLParserVisitor.prototype.visitTemplateStringExpression = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by ABSLParser#TernaryExpression.
ABSLParserVisitor.prototype.visitTernaryExpression = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by ABSLParser#LogicalAndExpression.
ABSLParserVisitor.prototype.visitLogicalAndExpression = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by ABSLParser#PowerExpression.
ABSLParserVisitor.prototype.visitPowerExpression = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by ABSLParser#PreIncrementExpression.
ABSLParserVisitor.prototype.visitPreIncrementExpression = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by ABSLParser#ObjectLiteralExpression.
ABSLParserVisitor.prototype.visitObjectLiteralExpression = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by ABSLParser#MetaExpression.
ABSLParserVisitor.prototype.visitMetaExpression = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by ABSLParser#InExpression.
ABSLParserVisitor.prototype.visitInExpression = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by ABSLParser#LogicalOrExpression.
ABSLParserVisitor.prototype.visitLogicalOrExpression = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by ABSLParser#NotExpression.
ABSLParserVisitor.prototype.visitNotExpression = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by ABSLParser#PreDecreaseExpression.
ABSLParserVisitor.prototype.visitPreDecreaseExpression = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by ABSLParser#ArgumentsExpression.
ABSLParserVisitor.prototype.visitArgumentsExpression = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by ABSLParser#AwaitExpression.
ABSLParserVisitor.prototype.visitAwaitExpression = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by ABSLParser#ThisExpression.
ABSLParserVisitor.prototype.visitThisExpression = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by ABSLParser#FunctionExpression.
ABSLParserVisitor.prototype.visitFunctionExpression = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by ABSLParser#UnaryMinusExpression.
ABSLParserVisitor.prototype.visitUnaryMinusExpression = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by ABSLParser#AssignmentExpression.
ABSLParserVisitor.prototype.visitAssignmentExpression = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by ABSLParser#PostDecreaseExpression.
ABSLParserVisitor.prototype.visitPostDecreaseExpression = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by ABSLParser#TypeofExpression.
ABSLParserVisitor.prototype.visitTypeofExpression = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by ABSLParser#InstanceofExpression.
ABSLParserVisitor.prototype.visitInstanceofExpression = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by ABSLParser#UnaryPlusExpression.
ABSLParserVisitor.prototype.visitUnaryPlusExpression = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by ABSLParser#DeleteExpression.
ABSLParserVisitor.prototype.visitDeleteExpression = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by ABSLParser#ImportExpression.
ABSLParserVisitor.prototype.visitImportExpression = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by ABSLParser#EqualityExpression.
ABSLParserVisitor.prototype.visitEqualityExpression = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by ABSLParser#BitXOrExpression.
ABSLParserVisitor.prototype.visitBitXOrExpression = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by ABSLParser#SuperExpression.
ABSLParserVisitor.prototype.visitSuperExpression = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by ABSLParser#MultiplicativeExpression.
ABSLParserVisitor.prototype.visitMultiplicativeExpression = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by ABSLParser#BitShiftExpression.
ABSLParserVisitor.prototype.visitBitShiftExpression = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by ABSLParser#ParenthesizedExpression.
ABSLParserVisitor.prototype.visitParenthesizedExpression = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by ABSLParser#AdditiveExpression.
ABSLParserVisitor.prototype.visitAdditiveExpression = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by ABSLParser#RelationalExpression.
ABSLParserVisitor.prototype.visitRelationalExpression = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by ABSLParser#PostIncrementExpression.
ABSLParserVisitor.prototype.visitPostIncrementExpression = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by ABSLParser#YieldExpression.
ABSLParserVisitor.prototype.visitYieldExpression = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by ABSLParser#BitNotExpression.
ABSLParserVisitor.prototype.visitBitNotExpression = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by ABSLParser#NewExpression.
ABSLParserVisitor.prototype.visitNewExpression = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by ABSLParser#LiteralExpression.
ABSLParserVisitor.prototype.visitLiteralExpression = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by ABSLParser#ArrayLiteralExpression.
ABSLParserVisitor.prototype.visitArrayLiteralExpression = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by ABSLParser#MemberDotExpression.
ABSLParserVisitor.prototype.visitMemberDotExpression = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by ABSLParser#ClassExpression.
ABSLParserVisitor.prototype.visitClassExpression = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by ABSLParser#MemberIndexExpression.
ABSLParserVisitor.prototype.visitMemberIndexExpression = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by ABSLParser#IdentifierExpression.
ABSLParserVisitor.prototype.visitIdentifierExpression = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by ABSLParser#BitAndExpression.
ABSLParserVisitor.prototype.visitBitAndExpression = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by ABSLParser#BitOrExpression.
ABSLParserVisitor.prototype.visitBitOrExpression = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by ABSLParser#AssignmentOperatorExpression.
ABSLParserVisitor.prototype.visitAssignmentOperatorExpression = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by ABSLParser#VoidExpression.
ABSLParserVisitor.prototype.visitVoidExpression = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by ABSLParser#CoalesceExpression.
ABSLParserVisitor.prototype.visitCoalesceExpression = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by ABSLParser#assignable.
ABSLParserVisitor.prototype.visitAssignable = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by ABSLParser#FunctionDecl.
ABSLParserVisitor.prototype.visitFunctionDecl = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by ABSLParser#AnoymousFunctionDecl.
ABSLParserVisitor.prototype.visitAnoymousFunctionDecl = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by ABSLParser#ArrowFunction.
ABSLParserVisitor.prototype.visitArrowFunction = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by ABSLParser#arrowFunctionParameters.
ABSLParserVisitor.prototype.visitArrowFunctionParameters = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by ABSLParser#arrowFunctionBody.
ABSLParserVisitor.prototype.visitArrowFunctionBody = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by ABSLParser#assignmentOperator.
ABSLParserVisitor.prototype.visitAssignmentOperator = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by ABSLParser#literal.
ABSLParserVisitor.prototype.visitLiteral = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by ABSLParser#numericLiteral.
ABSLParserVisitor.prototype.visitNumericLiteral = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by ABSLParser#bigintLiteral.
ABSLParserVisitor.prototype.visitBigintLiteral = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by ABSLParser#identifierName.
ABSLParserVisitor.prototype.visitIdentifierName = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by ABSLParser#reservedWord.
ABSLParserVisitor.prototype.visitReservedWord = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by ABSLParser#keyword.
ABSLParserVisitor.prototype.visitKeyword = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by ABSLParser#getter.
ABSLParserVisitor.prototype.visitGetter = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by ABSLParser#setter.
ABSLParserVisitor.prototype.visitSetter = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by ABSLParser#eos.
ABSLParserVisitor.prototype.visitEos = function(ctx) {
  return this.visitChildren(ctx);
};



exports.ABSLParserVisitor = ABSLParserVisitor;
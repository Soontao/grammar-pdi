// Generated from ./ABSLParser.g4 by ANTLR 4.7.2
// jshint ignore: start
var antlr4 = require("antlr4/index");

// This class defines a complete listener for a parse tree produced by ABSLParser.
function ABSLParserListener() {
  antlr4.tree.ParseTreeListener.call(this);
  return this;
}

ABSLParserListener.prototype = Object.create(antlr4.tree.ParseTreeListener.prototype);
ABSLParserListener.prototype.constructor = ABSLParserListener;

// Enter a parse tree produced by ABSLParser#program.
ABSLParserListener.prototype.enterProgram = function(ctx) {
};

// Exit a parse tree produced by ABSLParser#program.
ABSLParserListener.prototype.exitProgram = function(ctx) {
};


// Enter a parse tree produced by ABSLParser#sourceElement.
ABSLParserListener.prototype.enterSourceElement = function(ctx) {
};

// Exit a parse tree produced by ABSLParser#sourceElement.
ABSLParserListener.prototype.exitSourceElement = function(ctx) {
};


// Enter a parse tree produced by ABSLParser#statement.
ABSLParserListener.prototype.enterStatement = function(ctx) {
};

// Exit a parse tree produced by ABSLParser#statement.
ABSLParserListener.prototype.exitStatement = function(ctx) {
};


// Enter a parse tree produced by ABSLParser#block.
ABSLParserListener.prototype.enterBlock = function(ctx) {
};

// Exit a parse tree produced by ABSLParser#block.
ABSLParserListener.prototype.exitBlock = function(ctx) {
};


// Enter a parse tree produced by ABSLParser#statementList.
ABSLParserListener.prototype.enterStatementList = function(ctx) {
};

// Exit a parse tree produced by ABSLParser#statementList.
ABSLParserListener.prototype.exitStatementList = function(ctx) {
};


// Enter a parse tree produced by ABSLParser#importStatement.
ABSLParserListener.prototype.enterImportStatement = function(ctx) {
};

// Exit a parse tree produced by ABSLParser#importStatement.
ABSLParserListener.prototype.exitImportStatement = function(ctx) {
};


// Enter a parse tree produced by ABSLParser#importFromBlock.
ABSLParserListener.prototype.enterImportFromBlock = function(ctx) {
};

// Exit a parse tree produced by ABSLParser#importFromBlock.
ABSLParserListener.prototype.exitImportFromBlock = function(ctx) {
};


// Enter a parse tree produced by ABSLParser#moduleItems.
ABSLParserListener.prototype.enterModuleItems = function(ctx) {
};

// Exit a parse tree produced by ABSLParser#moduleItems.
ABSLParserListener.prototype.exitModuleItems = function(ctx) {
};


// Enter a parse tree produced by ABSLParser#importDefault.
ABSLParserListener.prototype.enterImportDefault = function(ctx) {
};

// Exit a parse tree produced by ABSLParser#importDefault.
ABSLParserListener.prototype.exitImportDefault = function(ctx) {
};


// Enter a parse tree produced by ABSLParser#importNamespace.
ABSLParserListener.prototype.enterImportNamespace = function(ctx) {
};

// Exit a parse tree produced by ABSLParser#importNamespace.
ABSLParserListener.prototype.exitImportNamespace = function(ctx) {
};


// Enter a parse tree produced by ABSLParser#importFrom.
ABSLParserListener.prototype.enterImportFrom = function(ctx) {
};

// Exit a parse tree produced by ABSLParser#importFrom.
ABSLParserListener.prototype.exitImportFrom = function(ctx) {
};


// Enter a parse tree produced by ABSLParser#aliasName.
ABSLParserListener.prototype.enterAliasName = function(ctx) {
};

// Exit a parse tree produced by ABSLParser#aliasName.
ABSLParserListener.prototype.exitAliasName = function(ctx) {
};


// Enter a parse tree produced by ABSLParser#ExportDeclaration.
ABSLParserListener.prototype.enterExportDeclaration = function(ctx) {
};

// Exit a parse tree produced by ABSLParser#ExportDeclaration.
ABSLParserListener.prototype.exitExportDeclaration = function(ctx) {
};


// Enter a parse tree produced by ABSLParser#ExportDefaultDeclaration.
ABSLParserListener.prototype.enterExportDefaultDeclaration = function(ctx) {
};

// Exit a parse tree produced by ABSLParser#ExportDefaultDeclaration.
ABSLParserListener.prototype.exitExportDefaultDeclaration = function(ctx) {
};


// Enter a parse tree produced by ABSLParser#exportFromBlock.
ABSLParserListener.prototype.enterExportFromBlock = function(ctx) {
};

// Exit a parse tree produced by ABSLParser#exportFromBlock.
ABSLParserListener.prototype.exitExportFromBlock = function(ctx) {
};


// Enter a parse tree produced by ABSLParser#declaration.
ABSLParserListener.prototype.enterDeclaration = function(ctx) {
};

// Exit a parse tree produced by ABSLParser#declaration.
ABSLParserListener.prototype.exitDeclaration = function(ctx) {
};


// Enter a parse tree produced by ABSLParser#variableStatement.
ABSLParserListener.prototype.enterVariableStatement = function(ctx) {
};

// Exit a parse tree produced by ABSLParser#variableStatement.
ABSLParserListener.prototype.exitVariableStatement = function(ctx) {
};


// Enter a parse tree produced by ABSLParser#variableDeclarationList.
ABSLParserListener.prototype.enterVariableDeclarationList = function(ctx) {
};

// Exit a parse tree produced by ABSLParser#variableDeclarationList.
ABSLParserListener.prototype.exitVariableDeclarationList = function(ctx) {
};


// Enter a parse tree produced by ABSLParser#variableDeclaration.
ABSLParserListener.prototype.enterVariableDeclaration = function(ctx) {
};

// Exit a parse tree produced by ABSLParser#variableDeclaration.
ABSLParserListener.prototype.exitVariableDeclaration = function(ctx) {
};


// Enter a parse tree produced by ABSLParser#variableType.
ABSLParserListener.prototype.enterVariableType = function(ctx) {
};

// Exit a parse tree produced by ABSLParser#variableType.
ABSLParserListener.prototype.exitVariableType = function(ctx) {
};


// Enter a parse tree produced by ABSLParser#programEmptyStatement.
ABSLParserListener.prototype.enterProgramEmptyStatement = function(ctx) {
};

// Exit a parse tree produced by ABSLParser#programEmptyStatement.
ABSLParserListener.prototype.exitProgramEmptyStatement = function(ctx) {
};


// Enter a parse tree produced by ABSLParser#expressionStatement.
ABSLParserListener.prototype.enterExpressionStatement = function(ctx) {
};

// Exit a parse tree produced by ABSLParser#expressionStatement.
ABSLParserListener.prototype.exitExpressionStatement = function(ctx) {
};


// Enter a parse tree produced by ABSLParser#ifStatement.
ABSLParserListener.prototype.enterIfStatement = function(ctx) {
};

// Exit a parse tree produced by ABSLParser#ifStatement.
ABSLParserListener.prototype.exitIfStatement = function(ctx) {
};


// Enter a parse tree produced by ABSLParser#DoStatement.
ABSLParserListener.prototype.enterDoStatement = function(ctx) {
};

// Exit a parse tree produced by ABSLParser#DoStatement.
ABSLParserListener.prototype.exitDoStatement = function(ctx) {
};


// Enter a parse tree produced by ABSLParser#WhileStatement.
ABSLParserListener.prototype.enterWhileStatement = function(ctx) {
};

// Exit a parse tree produced by ABSLParser#WhileStatement.
ABSLParserListener.prototype.exitWhileStatement = function(ctx) {
};


// Enter a parse tree produced by ABSLParser#ForStatement.
ABSLParserListener.prototype.enterForStatement = function(ctx) {
};

// Exit a parse tree produced by ABSLParser#ForStatement.
ABSLParserListener.prototype.exitForStatement = function(ctx) {
};


// Enter a parse tree produced by ABSLParser#ForInStatement.
ABSLParserListener.prototype.enterForInStatement = function(ctx) {
};

// Exit a parse tree produced by ABSLParser#ForInStatement.
ABSLParserListener.prototype.exitForInStatement = function(ctx) {
};


// Enter a parse tree produced by ABSLParser#ForOfStatement.
ABSLParserListener.prototype.enterForOfStatement = function(ctx) {
};

// Exit a parse tree produced by ABSLParser#ForOfStatement.
ABSLParserListener.prototype.exitForOfStatement = function(ctx) {
};


// Enter a parse tree produced by ABSLParser#ForeachStatement.
ABSLParserListener.prototype.enterForeachStatement = function(ctx) {
};

// Exit a parse tree produced by ABSLParser#ForeachStatement.
ABSLParserListener.prototype.exitForeachStatement = function(ctx) {
};


// Enter a parse tree produced by ABSLParser#varModifier.
ABSLParserListener.prototype.enterVarModifier = function(ctx) {
};

// Exit a parse tree produced by ABSLParser#varModifier.
ABSLParserListener.prototype.exitVarModifier = function(ctx) {
};


// Enter a parse tree produced by ABSLParser#continueStatement.
ABSLParserListener.prototype.enterContinueStatement = function(ctx) {
};

// Exit a parse tree produced by ABSLParser#continueStatement.
ABSLParserListener.prototype.exitContinueStatement = function(ctx) {
};


// Enter a parse tree produced by ABSLParser#breakStatement.
ABSLParserListener.prototype.enterBreakStatement = function(ctx) {
};

// Exit a parse tree produced by ABSLParser#breakStatement.
ABSLParserListener.prototype.exitBreakStatement = function(ctx) {
};


// Enter a parse tree produced by ABSLParser#returnStatement.
ABSLParserListener.prototype.enterReturnStatement = function(ctx) {
};

// Exit a parse tree produced by ABSLParser#returnStatement.
ABSLParserListener.prototype.exitReturnStatement = function(ctx) {
};


// Enter a parse tree produced by ABSLParser#yieldStatement.
ABSLParserListener.prototype.enterYieldStatement = function(ctx) {
};

// Exit a parse tree produced by ABSLParser#yieldStatement.
ABSLParserListener.prototype.exitYieldStatement = function(ctx) {
};


// Enter a parse tree produced by ABSLParser#withStatement.
ABSLParserListener.prototype.enterWithStatement = function(ctx) {
};

// Exit a parse tree produced by ABSLParser#withStatement.
ABSLParserListener.prototype.exitWithStatement = function(ctx) {
};


// Enter a parse tree produced by ABSLParser#switchStatement.
ABSLParserListener.prototype.enterSwitchStatement = function(ctx) {
};

// Exit a parse tree produced by ABSLParser#switchStatement.
ABSLParserListener.prototype.exitSwitchStatement = function(ctx) {
};


// Enter a parse tree produced by ABSLParser#caseBlock.
ABSLParserListener.prototype.enterCaseBlock = function(ctx) {
};

// Exit a parse tree produced by ABSLParser#caseBlock.
ABSLParserListener.prototype.exitCaseBlock = function(ctx) {
};


// Enter a parse tree produced by ABSLParser#caseClauses.
ABSLParserListener.prototype.enterCaseClauses = function(ctx) {
};

// Exit a parse tree produced by ABSLParser#caseClauses.
ABSLParserListener.prototype.exitCaseClauses = function(ctx) {
};


// Enter a parse tree produced by ABSLParser#caseClause.
ABSLParserListener.prototype.enterCaseClause = function(ctx) {
};

// Exit a parse tree produced by ABSLParser#caseClause.
ABSLParserListener.prototype.exitCaseClause = function(ctx) {
};


// Enter a parse tree produced by ABSLParser#defaultClause.
ABSLParserListener.prototype.enterDefaultClause = function(ctx) {
};

// Exit a parse tree produced by ABSLParser#defaultClause.
ABSLParserListener.prototype.exitDefaultClause = function(ctx) {
};


// Enter a parse tree produced by ABSLParser#labelledStatement.
ABSLParserListener.prototype.enterLabelledStatement = function(ctx) {
};

// Exit a parse tree produced by ABSLParser#labelledStatement.
ABSLParserListener.prototype.exitLabelledStatement = function(ctx) {
};


// Enter a parse tree produced by ABSLParser#throwStatement.
ABSLParserListener.prototype.enterThrowStatement = function(ctx) {
};

// Exit a parse tree produced by ABSLParser#throwStatement.
ABSLParserListener.prototype.exitThrowStatement = function(ctx) {
};


// Enter a parse tree produced by ABSLParser#tryStatement.
ABSLParserListener.prototype.enterTryStatement = function(ctx) {
};

// Exit a parse tree produced by ABSLParser#tryStatement.
ABSLParserListener.prototype.exitTryStatement = function(ctx) {
};


// Enter a parse tree produced by ABSLParser#catchProduction.
ABSLParserListener.prototype.enterCatchProduction = function(ctx) {
};

// Exit a parse tree produced by ABSLParser#catchProduction.
ABSLParserListener.prototype.exitCatchProduction = function(ctx) {
};


// Enter a parse tree produced by ABSLParser#finallyProduction.
ABSLParserListener.prototype.enterFinallyProduction = function(ctx) {
};

// Exit a parse tree produced by ABSLParser#finallyProduction.
ABSLParserListener.prototype.exitFinallyProduction = function(ctx) {
};


// Enter a parse tree produced by ABSLParser#debuggerStatement.
ABSLParserListener.prototype.enterDebuggerStatement = function(ctx) {
};

// Exit a parse tree produced by ABSLParser#debuggerStatement.
ABSLParserListener.prototype.exitDebuggerStatement = function(ctx) {
};


// Enter a parse tree produced by ABSLParser#functionDeclaration.
ABSLParserListener.prototype.enterFunctionDeclaration = function(ctx) {
};

// Exit a parse tree produced by ABSLParser#functionDeclaration.
ABSLParserListener.prototype.exitFunctionDeclaration = function(ctx) {
};


// Enter a parse tree produced by ABSLParser#classDeclaration.
ABSLParserListener.prototype.enterClassDeclaration = function(ctx) {
};

// Exit a parse tree produced by ABSLParser#classDeclaration.
ABSLParserListener.prototype.exitClassDeclaration = function(ctx) {
};


// Enter a parse tree produced by ABSLParser#classTail.
ABSLParserListener.prototype.enterClassTail = function(ctx) {
};

// Exit a parse tree produced by ABSLParser#classTail.
ABSLParserListener.prototype.exitClassTail = function(ctx) {
};


// Enter a parse tree produced by ABSLParser#classElement.
ABSLParserListener.prototype.enterClassElement = function(ctx) {
};

// Exit a parse tree produced by ABSLParser#classElement.
ABSLParserListener.prototype.exitClassElement = function(ctx) {
};


// Enter a parse tree produced by ABSLParser#methodDefinition.
ABSLParserListener.prototype.enterMethodDefinition = function(ctx) {
};

// Exit a parse tree produced by ABSLParser#methodDefinition.
ABSLParserListener.prototype.exitMethodDefinition = function(ctx) {
};


// Enter a parse tree produced by ABSLParser#formalParameterList.
ABSLParserListener.prototype.enterFormalParameterList = function(ctx) {
};

// Exit a parse tree produced by ABSLParser#formalParameterList.
ABSLParserListener.prototype.exitFormalParameterList = function(ctx) {
};


// Enter a parse tree produced by ABSLParser#formalParameterArg.
ABSLParserListener.prototype.enterFormalParameterArg = function(ctx) {
};

// Exit a parse tree produced by ABSLParser#formalParameterArg.
ABSLParserListener.prototype.exitFormalParameterArg = function(ctx) {
};


// Enter a parse tree produced by ABSLParser#lastFormalParameterArg.
ABSLParserListener.prototype.enterLastFormalParameterArg = function(ctx) {
};

// Exit a parse tree produced by ABSLParser#lastFormalParameterArg.
ABSLParserListener.prototype.exitLastFormalParameterArg = function(ctx) {
};


// Enter a parse tree produced by ABSLParser#functionBody.
ABSLParserListener.prototype.enterFunctionBody = function(ctx) {
};

// Exit a parse tree produced by ABSLParser#functionBody.
ABSLParserListener.prototype.exitFunctionBody = function(ctx) {
};


// Enter a parse tree produced by ABSLParser#sourceElements.
ABSLParserListener.prototype.enterSourceElements = function(ctx) {
};

// Exit a parse tree produced by ABSLParser#sourceElements.
ABSLParserListener.prototype.exitSourceElements = function(ctx) {
};


// Enter a parse tree produced by ABSLParser#arrayLiteral.
ABSLParserListener.prototype.enterArrayLiteral = function(ctx) {
};

// Exit a parse tree produced by ABSLParser#arrayLiteral.
ABSLParserListener.prototype.exitArrayLiteral = function(ctx) {
};


// Enter a parse tree produced by ABSLParser#elementList.
ABSLParserListener.prototype.enterElementList = function(ctx) {
};

// Exit a parse tree produced by ABSLParser#elementList.
ABSLParserListener.prototype.exitElementList = function(ctx) {
};


// Enter a parse tree produced by ABSLParser#arrayElement.
ABSLParserListener.prototype.enterArrayElement = function(ctx) {
};

// Exit a parse tree produced by ABSLParser#arrayElement.
ABSLParserListener.prototype.exitArrayElement = function(ctx) {
};


// Enter a parse tree produced by ABSLParser#objectLiteral.
ABSLParserListener.prototype.enterObjectLiteral = function(ctx) {
};

// Exit a parse tree produced by ABSLParser#objectLiteral.
ABSLParserListener.prototype.exitObjectLiteral = function(ctx) {
};


// Enter a parse tree produced by ABSLParser#PropertyExpressionAssignment.
ABSLParserListener.prototype.enterPropertyExpressionAssignment = function(ctx) {
};

// Exit a parse tree produced by ABSLParser#PropertyExpressionAssignment.
ABSLParserListener.prototype.exitPropertyExpressionAssignment = function(ctx) {
};


// Enter a parse tree produced by ABSLParser#ComputedPropertyExpressionAssignment.
ABSLParserListener.prototype.enterComputedPropertyExpressionAssignment = function(ctx) {
};

// Exit a parse tree produced by ABSLParser#ComputedPropertyExpressionAssignment.
ABSLParserListener.prototype.exitComputedPropertyExpressionAssignment = function(ctx) {
};


// Enter a parse tree produced by ABSLParser#FunctionProperty.
ABSLParserListener.prototype.enterFunctionProperty = function(ctx) {
};

// Exit a parse tree produced by ABSLParser#FunctionProperty.
ABSLParserListener.prototype.exitFunctionProperty = function(ctx) {
};


// Enter a parse tree produced by ABSLParser#PropertyGetter.
ABSLParserListener.prototype.enterPropertyGetter = function(ctx) {
};

// Exit a parse tree produced by ABSLParser#PropertyGetter.
ABSLParserListener.prototype.exitPropertyGetter = function(ctx) {
};


// Enter a parse tree produced by ABSLParser#PropertySetter.
ABSLParserListener.prototype.enterPropertySetter = function(ctx) {
};

// Exit a parse tree produced by ABSLParser#PropertySetter.
ABSLParserListener.prototype.exitPropertySetter = function(ctx) {
};


// Enter a parse tree produced by ABSLParser#PropertyShorthand.
ABSLParserListener.prototype.enterPropertyShorthand = function(ctx) {
};

// Exit a parse tree produced by ABSLParser#PropertyShorthand.
ABSLParserListener.prototype.exitPropertyShorthand = function(ctx) {
};


// Enter a parse tree produced by ABSLParser#propertyName.
ABSLParserListener.prototype.enterPropertyName = function(ctx) {
};

// Exit a parse tree produced by ABSLParser#propertyName.
ABSLParserListener.prototype.exitPropertyName = function(ctx) {
};


// Enter a parse tree produced by ABSLParser#arguments.
ABSLParserListener.prototype.enterArguments = function(ctx) {
};

// Exit a parse tree produced by ABSLParser#arguments.
ABSLParserListener.prototype.exitArguments = function(ctx) {
};


// Enter a parse tree produced by ABSLParser#argument.
ABSLParserListener.prototype.enterArgument = function(ctx) {
};

// Exit a parse tree produced by ABSLParser#argument.
ABSLParserListener.prototype.exitArgument = function(ctx) {
};


// Enter a parse tree produced by ABSLParser#expressionSequence.
ABSLParserListener.prototype.enterExpressionSequence = function(ctx) {
};

// Exit a parse tree produced by ABSLParser#expressionSequence.
ABSLParserListener.prototype.exitExpressionSequence = function(ctx) {
};


// Enter a parse tree produced by ABSLParser#TemplateStringExpression.
ABSLParserListener.prototype.enterTemplateStringExpression = function(ctx) {
};

// Exit a parse tree produced by ABSLParser#TemplateStringExpression.
ABSLParserListener.prototype.exitTemplateStringExpression = function(ctx) {
};


// Enter a parse tree produced by ABSLParser#TernaryExpression.
ABSLParserListener.prototype.enterTernaryExpression = function(ctx) {
};

// Exit a parse tree produced by ABSLParser#TernaryExpression.
ABSLParserListener.prototype.exitTernaryExpression = function(ctx) {
};


// Enter a parse tree produced by ABSLParser#LogicalAndExpression.
ABSLParserListener.prototype.enterLogicalAndExpression = function(ctx) {
};

// Exit a parse tree produced by ABSLParser#LogicalAndExpression.
ABSLParserListener.prototype.exitLogicalAndExpression = function(ctx) {
};


// Enter a parse tree produced by ABSLParser#PowerExpression.
ABSLParserListener.prototype.enterPowerExpression = function(ctx) {
};

// Exit a parse tree produced by ABSLParser#PowerExpression.
ABSLParserListener.prototype.exitPowerExpression = function(ctx) {
};


// Enter a parse tree produced by ABSLParser#PreIncrementExpression.
ABSLParserListener.prototype.enterPreIncrementExpression = function(ctx) {
};

// Exit a parse tree produced by ABSLParser#PreIncrementExpression.
ABSLParserListener.prototype.exitPreIncrementExpression = function(ctx) {
};


// Enter a parse tree produced by ABSLParser#ObjectLiteralExpression.
ABSLParserListener.prototype.enterObjectLiteralExpression = function(ctx) {
};

// Exit a parse tree produced by ABSLParser#ObjectLiteralExpression.
ABSLParserListener.prototype.exitObjectLiteralExpression = function(ctx) {
};


// Enter a parse tree produced by ABSLParser#MetaExpression.
ABSLParserListener.prototype.enterMetaExpression = function(ctx) {
};

// Exit a parse tree produced by ABSLParser#MetaExpression.
ABSLParserListener.prototype.exitMetaExpression = function(ctx) {
};


// Enter a parse tree produced by ABSLParser#InExpression.
ABSLParserListener.prototype.enterInExpression = function(ctx) {
};

// Exit a parse tree produced by ABSLParser#InExpression.
ABSLParserListener.prototype.exitInExpression = function(ctx) {
};


// Enter a parse tree produced by ABSLParser#LogicalOrExpression.
ABSLParserListener.prototype.enterLogicalOrExpression = function(ctx) {
};

// Exit a parse tree produced by ABSLParser#LogicalOrExpression.
ABSLParserListener.prototype.exitLogicalOrExpression = function(ctx) {
};


// Enter a parse tree produced by ABSLParser#NotExpression.
ABSLParserListener.prototype.enterNotExpression = function(ctx) {
};

// Exit a parse tree produced by ABSLParser#NotExpression.
ABSLParserListener.prototype.exitNotExpression = function(ctx) {
};


// Enter a parse tree produced by ABSLParser#PreDecreaseExpression.
ABSLParserListener.prototype.enterPreDecreaseExpression = function(ctx) {
};

// Exit a parse tree produced by ABSLParser#PreDecreaseExpression.
ABSLParserListener.prototype.exitPreDecreaseExpression = function(ctx) {
};


// Enter a parse tree produced by ABSLParser#ArgumentsExpression.
ABSLParserListener.prototype.enterArgumentsExpression = function(ctx) {
};

// Exit a parse tree produced by ABSLParser#ArgumentsExpression.
ABSLParserListener.prototype.exitArgumentsExpression = function(ctx) {
};


// Enter a parse tree produced by ABSLParser#AwaitExpression.
ABSLParserListener.prototype.enterAwaitExpression = function(ctx) {
};

// Exit a parse tree produced by ABSLParser#AwaitExpression.
ABSLParserListener.prototype.exitAwaitExpression = function(ctx) {
};


// Enter a parse tree produced by ABSLParser#ThisExpression.
ABSLParserListener.prototype.enterThisExpression = function(ctx) {
};

// Exit a parse tree produced by ABSLParser#ThisExpression.
ABSLParserListener.prototype.exitThisExpression = function(ctx) {
};


// Enter a parse tree produced by ABSLParser#FunctionExpression.
ABSLParserListener.prototype.enterFunctionExpression = function(ctx) {
};

// Exit a parse tree produced by ABSLParser#FunctionExpression.
ABSLParserListener.prototype.exitFunctionExpression = function(ctx) {
};


// Enter a parse tree produced by ABSLParser#UnaryMinusExpression.
ABSLParserListener.prototype.enterUnaryMinusExpression = function(ctx) {
};

// Exit a parse tree produced by ABSLParser#UnaryMinusExpression.
ABSLParserListener.prototype.exitUnaryMinusExpression = function(ctx) {
};


// Enter a parse tree produced by ABSLParser#AssignmentExpression.
ABSLParserListener.prototype.enterAssignmentExpression = function(ctx) {
};

// Exit a parse tree produced by ABSLParser#AssignmentExpression.
ABSLParserListener.prototype.exitAssignmentExpression = function(ctx) {
};


// Enter a parse tree produced by ABSLParser#PostDecreaseExpression.
ABSLParserListener.prototype.enterPostDecreaseExpression = function(ctx) {
};

// Exit a parse tree produced by ABSLParser#PostDecreaseExpression.
ABSLParserListener.prototype.exitPostDecreaseExpression = function(ctx) {
};


// Enter a parse tree produced by ABSLParser#TypeofExpression.
ABSLParserListener.prototype.enterTypeofExpression = function(ctx) {
};

// Exit a parse tree produced by ABSLParser#TypeofExpression.
ABSLParserListener.prototype.exitTypeofExpression = function(ctx) {
};


// Enter a parse tree produced by ABSLParser#InstanceofExpression.
ABSLParserListener.prototype.enterInstanceofExpression = function(ctx) {
};

// Exit a parse tree produced by ABSLParser#InstanceofExpression.
ABSLParserListener.prototype.exitInstanceofExpression = function(ctx) {
};


// Enter a parse tree produced by ABSLParser#UnaryPlusExpression.
ABSLParserListener.prototype.enterUnaryPlusExpression = function(ctx) {
};

// Exit a parse tree produced by ABSLParser#UnaryPlusExpression.
ABSLParserListener.prototype.exitUnaryPlusExpression = function(ctx) {
};


// Enter a parse tree produced by ABSLParser#DeleteExpression.
ABSLParserListener.prototype.enterDeleteExpression = function(ctx) {
};

// Exit a parse tree produced by ABSLParser#DeleteExpression.
ABSLParserListener.prototype.exitDeleteExpression = function(ctx) {
};


// Enter a parse tree produced by ABSLParser#ImportExpression.
ABSLParserListener.prototype.enterImportExpression = function(ctx) {
};

// Exit a parse tree produced by ABSLParser#ImportExpression.
ABSLParserListener.prototype.exitImportExpression = function(ctx) {
};


// Enter a parse tree produced by ABSLParser#EqualityExpression.
ABSLParserListener.prototype.enterEqualityExpression = function(ctx) {
};

// Exit a parse tree produced by ABSLParser#EqualityExpression.
ABSLParserListener.prototype.exitEqualityExpression = function(ctx) {
};


// Enter a parse tree produced by ABSLParser#BitXOrExpression.
ABSLParserListener.prototype.enterBitXOrExpression = function(ctx) {
};

// Exit a parse tree produced by ABSLParser#BitXOrExpression.
ABSLParserListener.prototype.exitBitXOrExpression = function(ctx) {
};


// Enter a parse tree produced by ABSLParser#SuperExpression.
ABSLParserListener.prototype.enterSuperExpression = function(ctx) {
};

// Exit a parse tree produced by ABSLParser#SuperExpression.
ABSLParserListener.prototype.exitSuperExpression = function(ctx) {
};


// Enter a parse tree produced by ABSLParser#MultiplicativeExpression.
ABSLParserListener.prototype.enterMultiplicativeExpression = function(ctx) {
};

// Exit a parse tree produced by ABSLParser#MultiplicativeExpression.
ABSLParserListener.prototype.exitMultiplicativeExpression = function(ctx) {
};


// Enter a parse tree produced by ABSLParser#BitShiftExpression.
ABSLParserListener.prototype.enterBitShiftExpression = function(ctx) {
};

// Exit a parse tree produced by ABSLParser#BitShiftExpression.
ABSLParserListener.prototype.exitBitShiftExpression = function(ctx) {
};


// Enter a parse tree produced by ABSLParser#ParenthesizedExpression.
ABSLParserListener.prototype.enterParenthesizedExpression = function(ctx) {
};

// Exit a parse tree produced by ABSLParser#ParenthesizedExpression.
ABSLParserListener.prototype.exitParenthesizedExpression = function(ctx) {
};


// Enter a parse tree produced by ABSLParser#AdditiveExpression.
ABSLParserListener.prototype.enterAdditiveExpression = function(ctx) {
};

// Exit a parse tree produced by ABSLParser#AdditiveExpression.
ABSLParserListener.prototype.exitAdditiveExpression = function(ctx) {
};


// Enter a parse tree produced by ABSLParser#RelationalExpression.
ABSLParserListener.prototype.enterRelationalExpression = function(ctx) {
};

// Exit a parse tree produced by ABSLParser#RelationalExpression.
ABSLParserListener.prototype.exitRelationalExpression = function(ctx) {
};


// Enter a parse tree produced by ABSLParser#PostIncrementExpression.
ABSLParserListener.prototype.enterPostIncrementExpression = function(ctx) {
};

// Exit a parse tree produced by ABSLParser#PostIncrementExpression.
ABSLParserListener.prototype.exitPostIncrementExpression = function(ctx) {
};


// Enter a parse tree produced by ABSLParser#YieldExpression.
ABSLParserListener.prototype.enterYieldExpression = function(ctx) {
};

// Exit a parse tree produced by ABSLParser#YieldExpression.
ABSLParserListener.prototype.exitYieldExpression = function(ctx) {
};


// Enter a parse tree produced by ABSLParser#BitNotExpression.
ABSLParserListener.prototype.enterBitNotExpression = function(ctx) {
};

// Exit a parse tree produced by ABSLParser#BitNotExpression.
ABSLParserListener.prototype.exitBitNotExpression = function(ctx) {
};


// Enter a parse tree produced by ABSLParser#NewExpression.
ABSLParserListener.prototype.enterNewExpression = function(ctx) {
};

// Exit a parse tree produced by ABSLParser#NewExpression.
ABSLParserListener.prototype.exitNewExpression = function(ctx) {
};


// Enter a parse tree produced by ABSLParser#LiteralExpression.
ABSLParserListener.prototype.enterLiteralExpression = function(ctx) {
};

// Exit a parse tree produced by ABSLParser#LiteralExpression.
ABSLParserListener.prototype.exitLiteralExpression = function(ctx) {
};


// Enter a parse tree produced by ABSLParser#ArrayLiteralExpression.
ABSLParserListener.prototype.enterArrayLiteralExpression = function(ctx) {
};

// Exit a parse tree produced by ABSLParser#ArrayLiteralExpression.
ABSLParserListener.prototype.exitArrayLiteralExpression = function(ctx) {
};


// Enter a parse tree produced by ABSLParser#MemberDotExpression.
ABSLParserListener.prototype.enterMemberDotExpression = function(ctx) {
};

// Exit a parse tree produced by ABSLParser#MemberDotExpression.
ABSLParserListener.prototype.exitMemberDotExpression = function(ctx) {
};


// Enter a parse tree produced by ABSLParser#ClassExpression.
ABSLParserListener.prototype.enterClassExpression = function(ctx) {
};

// Exit a parse tree produced by ABSLParser#ClassExpression.
ABSLParserListener.prototype.exitClassExpression = function(ctx) {
};


// Enter a parse tree produced by ABSLParser#MemberIndexExpression.
ABSLParserListener.prototype.enterMemberIndexExpression = function(ctx) {
};

// Exit a parse tree produced by ABSLParser#MemberIndexExpression.
ABSLParserListener.prototype.exitMemberIndexExpression = function(ctx) {
};


// Enter a parse tree produced by ABSLParser#IdentifierExpression.
ABSLParserListener.prototype.enterIdentifierExpression = function(ctx) {
};

// Exit a parse tree produced by ABSLParser#IdentifierExpression.
ABSLParserListener.prototype.exitIdentifierExpression = function(ctx) {
};


// Enter a parse tree produced by ABSLParser#BitAndExpression.
ABSLParserListener.prototype.enterBitAndExpression = function(ctx) {
};

// Exit a parse tree produced by ABSLParser#BitAndExpression.
ABSLParserListener.prototype.exitBitAndExpression = function(ctx) {
};


// Enter a parse tree produced by ABSLParser#BitOrExpression.
ABSLParserListener.prototype.enterBitOrExpression = function(ctx) {
};

// Exit a parse tree produced by ABSLParser#BitOrExpression.
ABSLParserListener.prototype.exitBitOrExpression = function(ctx) {
};


// Enter a parse tree produced by ABSLParser#AssignmentOperatorExpression.
ABSLParserListener.prototype.enterAssignmentOperatorExpression = function(ctx) {
};

// Exit a parse tree produced by ABSLParser#AssignmentOperatorExpression.
ABSLParserListener.prototype.exitAssignmentOperatorExpression = function(ctx) {
};


// Enter a parse tree produced by ABSLParser#VoidExpression.
ABSLParserListener.prototype.enterVoidExpression = function(ctx) {
};

// Exit a parse tree produced by ABSLParser#VoidExpression.
ABSLParserListener.prototype.exitVoidExpression = function(ctx) {
};


// Enter a parse tree produced by ABSLParser#CoalesceExpression.
ABSLParserListener.prototype.enterCoalesceExpression = function(ctx) {
};

// Exit a parse tree produced by ABSLParser#CoalesceExpression.
ABSLParserListener.prototype.exitCoalesceExpression = function(ctx) {
};


// Enter a parse tree produced by ABSLParser#assignable.
ABSLParserListener.prototype.enterAssignable = function(ctx) {
};

// Exit a parse tree produced by ABSLParser#assignable.
ABSLParserListener.prototype.exitAssignable = function(ctx) {
};


// Enter a parse tree produced by ABSLParser#FunctionDecl.
ABSLParserListener.prototype.enterFunctionDecl = function(ctx) {
};

// Exit a parse tree produced by ABSLParser#FunctionDecl.
ABSLParserListener.prototype.exitFunctionDecl = function(ctx) {
};


// Enter a parse tree produced by ABSLParser#AnoymousFunctionDecl.
ABSLParserListener.prototype.enterAnoymousFunctionDecl = function(ctx) {
};

// Exit a parse tree produced by ABSLParser#AnoymousFunctionDecl.
ABSLParserListener.prototype.exitAnoymousFunctionDecl = function(ctx) {
};


// Enter a parse tree produced by ABSLParser#ArrowFunction.
ABSLParserListener.prototype.enterArrowFunction = function(ctx) {
};

// Exit a parse tree produced by ABSLParser#ArrowFunction.
ABSLParserListener.prototype.exitArrowFunction = function(ctx) {
};


// Enter a parse tree produced by ABSLParser#arrowFunctionParameters.
ABSLParserListener.prototype.enterArrowFunctionParameters = function(ctx) {
};

// Exit a parse tree produced by ABSLParser#arrowFunctionParameters.
ABSLParserListener.prototype.exitArrowFunctionParameters = function(ctx) {
};


// Enter a parse tree produced by ABSLParser#arrowFunctionBody.
ABSLParserListener.prototype.enterArrowFunctionBody = function(ctx) {
};

// Exit a parse tree produced by ABSLParser#arrowFunctionBody.
ABSLParserListener.prototype.exitArrowFunctionBody = function(ctx) {
};


// Enter a parse tree produced by ABSLParser#assignmentOperator.
ABSLParserListener.prototype.enterAssignmentOperator = function(ctx) {
};

// Exit a parse tree produced by ABSLParser#assignmentOperator.
ABSLParserListener.prototype.exitAssignmentOperator = function(ctx) {
};


// Enter a parse tree produced by ABSLParser#literal.
ABSLParserListener.prototype.enterLiteral = function(ctx) {
};

// Exit a parse tree produced by ABSLParser#literal.
ABSLParserListener.prototype.exitLiteral = function(ctx) {
};


// Enter a parse tree produced by ABSLParser#numericLiteral.
ABSLParserListener.prototype.enterNumericLiteral = function(ctx) {
};

// Exit a parse tree produced by ABSLParser#numericLiteral.
ABSLParserListener.prototype.exitNumericLiteral = function(ctx) {
};


// Enter a parse tree produced by ABSLParser#bigintLiteral.
ABSLParserListener.prototype.enterBigintLiteral = function(ctx) {
};

// Exit a parse tree produced by ABSLParser#bigintLiteral.
ABSLParserListener.prototype.exitBigintLiteral = function(ctx) {
};


// Enter a parse tree produced by ABSLParser#identifierName.
ABSLParserListener.prototype.enterIdentifierName = function(ctx) {
};

// Exit a parse tree produced by ABSLParser#identifierName.
ABSLParserListener.prototype.exitIdentifierName = function(ctx) {
};


// Enter a parse tree produced by ABSLParser#reservedWord.
ABSLParserListener.prototype.enterReservedWord = function(ctx) {
};

// Exit a parse tree produced by ABSLParser#reservedWord.
ABSLParserListener.prototype.exitReservedWord = function(ctx) {
};


// Enter a parse tree produced by ABSLParser#keyword.
ABSLParserListener.prototype.enterKeyword = function(ctx) {
};

// Exit a parse tree produced by ABSLParser#keyword.
ABSLParserListener.prototype.exitKeyword = function(ctx) {
};


// Enter a parse tree produced by ABSLParser#getter.
ABSLParserListener.prototype.enterGetter = function(ctx) {
};

// Exit a parse tree produced by ABSLParser#getter.
ABSLParserListener.prototype.exitGetter = function(ctx) {
};


// Enter a parse tree produced by ABSLParser#setter.
ABSLParserListener.prototype.enterSetter = function(ctx) {
};

// Exit a parse tree produced by ABSLParser#setter.
ABSLParserListener.prototype.exitSetter = function(ctx) {
};


// Enter a parse tree produced by ABSLParser#eos.
ABSLParserListener.prototype.enterEos = function(ctx) {
};

// Exit a parse tree produced by ABSLParser#eos.
ABSLParserListener.prototype.exitEos = function(ctx) {
};



exports.ABSLParserListener = ABSLParserListener;
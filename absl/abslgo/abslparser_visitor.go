// Code generated from ./ABSLParser.g4 by ANTLR 4.7.2. DO NOT EDIT.

package abslgo // ABSLParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by ABSLParser.
type ABSLParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by ABSLParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by ABSLParser#sourceElement.
	VisitSourceElement(ctx *SourceElementContext) interface{}

	// Visit a parse tree produced by ABSLParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by ABSLParser#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by ABSLParser#statementList.
	VisitStatementList(ctx *StatementListContext) interface{}

	// Visit a parse tree produced by ABSLParser#importStatement.
	VisitImportStatement(ctx *ImportStatementContext) interface{}

	// Visit a parse tree produced by ABSLParser#importFromBlock.
	VisitImportFromBlock(ctx *ImportFromBlockContext) interface{}

	// Visit a parse tree produced by ABSLParser#moduleItems.
	VisitModuleItems(ctx *ModuleItemsContext) interface{}

	// Visit a parse tree produced by ABSLParser#importDefault.
	VisitImportDefault(ctx *ImportDefaultContext) interface{}

	// Visit a parse tree produced by ABSLParser#importNamespace.
	VisitImportNamespace(ctx *ImportNamespaceContext) interface{}

	// Visit a parse tree produced by ABSLParser#importFrom.
	VisitImportFrom(ctx *ImportFromContext) interface{}

	// Visit a parse tree produced by ABSLParser#aliasName.
	VisitAliasName(ctx *AliasNameContext) interface{}

	// Visit a parse tree produced by ABSLParser#ExportDeclaration.
	VisitExportDeclaration(ctx *ExportDeclarationContext) interface{}

	// Visit a parse tree produced by ABSLParser#ExportDefaultDeclaration.
	VisitExportDefaultDeclaration(ctx *ExportDefaultDeclarationContext) interface{}

	// Visit a parse tree produced by ABSLParser#exportFromBlock.
	VisitExportFromBlock(ctx *ExportFromBlockContext) interface{}

	// Visit a parse tree produced by ABSLParser#declaration.
	VisitDeclaration(ctx *DeclarationContext) interface{}

	// Visit a parse tree produced by ABSLParser#variableStatement.
	VisitVariableStatement(ctx *VariableStatementContext) interface{}

	// Visit a parse tree produced by ABSLParser#variableDeclarationList.
	VisitVariableDeclarationList(ctx *VariableDeclarationListContext) interface{}

	// Visit a parse tree produced by ABSLParser#variableDeclaration.
	VisitVariableDeclaration(ctx *VariableDeclarationContext) interface{}

	// Visit a parse tree produced by ABSLParser#variableType.
	VisitVariableType(ctx *VariableTypeContext) interface{}

	// Visit a parse tree produced by ABSLParser#emptyStatement.
	VisitEmptyStatement(ctx *EmptyStatementContext) interface{}

	// Visit a parse tree produced by ABSLParser#expressionStatement.
	VisitExpressionStatement(ctx *ExpressionStatementContext) interface{}

	// Visit a parse tree produced by ABSLParser#ifStatement.
	VisitIfStatement(ctx *IfStatementContext) interface{}

	// Visit a parse tree produced by ABSLParser#DoStatement.
	VisitDoStatement(ctx *DoStatementContext) interface{}

	// Visit a parse tree produced by ABSLParser#WhileStatement.
	VisitWhileStatement(ctx *WhileStatementContext) interface{}

	// Visit a parse tree produced by ABSLParser#ForStatement.
	VisitForStatement(ctx *ForStatementContext) interface{}

	// Visit a parse tree produced by ABSLParser#ForInStatement.
	VisitForInStatement(ctx *ForInStatementContext) interface{}

	// Visit a parse tree produced by ABSLParser#ForOfStatement.
	VisitForOfStatement(ctx *ForOfStatementContext) interface{}

	// Visit a parse tree produced by ABSLParser#ForeachStatement.
	VisitForeachStatement(ctx *ForeachStatementContext) interface{}

	// Visit a parse tree produced by ABSLParser#varModifier.
	VisitVarModifier(ctx *VarModifierContext) interface{}

	// Visit a parse tree produced by ABSLParser#continueStatement.
	VisitContinueStatement(ctx *ContinueStatementContext) interface{}

	// Visit a parse tree produced by ABSLParser#breakStatement.
	VisitBreakStatement(ctx *BreakStatementContext) interface{}

	// Visit a parse tree produced by ABSLParser#returnStatement.
	VisitReturnStatement(ctx *ReturnStatementContext) interface{}

	// Visit a parse tree produced by ABSLParser#yieldStatement.
	VisitYieldStatement(ctx *YieldStatementContext) interface{}

	// Visit a parse tree produced by ABSLParser#withStatement.
	VisitWithStatement(ctx *WithStatementContext) interface{}

	// Visit a parse tree produced by ABSLParser#switchStatement.
	VisitSwitchStatement(ctx *SwitchStatementContext) interface{}

	// Visit a parse tree produced by ABSLParser#caseBlock.
	VisitCaseBlock(ctx *CaseBlockContext) interface{}

	// Visit a parse tree produced by ABSLParser#caseClauses.
	VisitCaseClauses(ctx *CaseClausesContext) interface{}

	// Visit a parse tree produced by ABSLParser#caseClause.
	VisitCaseClause(ctx *CaseClauseContext) interface{}

	// Visit a parse tree produced by ABSLParser#defaultClause.
	VisitDefaultClause(ctx *DefaultClauseContext) interface{}

	// Visit a parse tree produced by ABSLParser#labelledStatement.
	VisitLabelledStatement(ctx *LabelledStatementContext) interface{}

	// Visit a parse tree produced by ABSLParser#throwStatement.
	VisitThrowStatement(ctx *ThrowStatementContext) interface{}

	// Visit a parse tree produced by ABSLParser#tryStatement.
	VisitTryStatement(ctx *TryStatementContext) interface{}

	// Visit a parse tree produced by ABSLParser#catchProduction.
	VisitCatchProduction(ctx *CatchProductionContext) interface{}

	// Visit a parse tree produced by ABSLParser#finallyProduction.
	VisitFinallyProduction(ctx *FinallyProductionContext) interface{}

	// Visit a parse tree produced by ABSLParser#debuggerStatement.
	VisitDebuggerStatement(ctx *DebuggerStatementContext) interface{}

	// Visit a parse tree produced by ABSLParser#functionDeclaration.
	VisitFunctionDeclaration(ctx *FunctionDeclarationContext) interface{}

	// Visit a parse tree produced by ABSLParser#classDeclaration.
	VisitClassDeclaration(ctx *ClassDeclarationContext) interface{}

	// Visit a parse tree produced by ABSLParser#classTail.
	VisitClassTail(ctx *ClassTailContext) interface{}

	// Visit a parse tree produced by ABSLParser#classElement.
	VisitClassElement(ctx *ClassElementContext) interface{}

	// Visit a parse tree produced by ABSLParser#methodDefinition.
	VisitMethodDefinition(ctx *MethodDefinitionContext) interface{}

	// Visit a parse tree produced by ABSLParser#formalParameterList.
	VisitFormalParameterList(ctx *FormalParameterListContext) interface{}

	// Visit a parse tree produced by ABSLParser#formalParameterArg.
	VisitFormalParameterArg(ctx *FormalParameterArgContext) interface{}

	// Visit a parse tree produced by ABSLParser#lastFormalParameterArg.
	VisitLastFormalParameterArg(ctx *LastFormalParameterArgContext) interface{}

	// Visit a parse tree produced by ABSLParser#functionBody.
	VisitFunctionBody(ctx *FunctionBodyContext) interface{}

	// Visit a parse tree produced by ABSLParser#sourceElements.
	VisitSourceElements(ctx *SourceElementsContext) interface{}

	// Visit a parse tree produced by ABSLParser#arrayLiteral.
	VisitArrayLiteral(ctx *ArrayLiteralContext) interface{}

	// Visit a parse tree produced by ABSLParser#elementList.
	VisitElementList(ctx *ElementListContext) interface{}

	// Visit a parse tree produced by ABSLParser#arrayElement.
	VisitArrayElement(ctx *ArrayElementContext) interface{}

	// Visit a parse tree produced by ABSLParser#objectLiteral.
	VisitObjectLiteral(ctx *ObjectLiteralContext) interface{}

	// Visit a parse tree produced by ABSLParser#PropertyExpressionAssignment.
	VisitPropertyExpressionAssignment(ctx *PropertyExpressionAssignmentContext) interface{}

	// Visit a parse tree produced by ABSLParser#ComputedPropertyExpressionAssignment.
	VisitComputedPropertyExpressionAssignment(ctx *ComputedPropertyExpressionAssignmentContext) interface{}

	// Visit a parse tree produced by ABSLParser#FunctionProperty.
	VisitFunctionProperty(ctx *FunctionPropertyContext) interface{}

	// Visit a parse tree produced by ABSLParser#PropertyGetter.
	VisitPropertyGetter(ctx *PropertyGetterContext) interface{}

	// Visit a parse tree produced by ABSLParser#PropertySetter.
	VisitPropertySetter(ctx *PropertySetterContext) interface{}

	// Visit a parse tree produced by ABSLParser#PropertyShorthand.
	VisitPropertyShorthand(ctx *PropertyShorthandContext) interface{}

	// Visit a parse tree produced by ABSLParser#propertyName.
	VisitPropertyName(ctx *PropertyNameContext) interface{}

	// Visit a parse tree produced by ABSLParser#arguments.
	VisitArguments(ctx *ArgumentsContext) interface{}

	// Visit a parse tree produced by ABSLParser#argument.
	VisitArgument(ctx *ArgumentContext) interface{}

	// Visit a parse tree produced by ABSLParser#expressionSequence.
	VisitExpressionSequence(ctx *ExpressionSequenceContext) interface{}

	// Visit a parse tree produced by ABSLParser#TemplateStringExpression.
	VisitTemplateStringExpression(ctx *TemplateStringExpressionContext) interface{}

	// Visit a parse tree produced by ABSLParser#TernaryExpression.
	VisitTernaryExpression(ctx *TernaryExpressionContext) interface{}

	// Visit a parse tree produced by ABSLParser#LogicalAndExpression.
	VisitLogicalAndExpression(ctx *LogicalAndExpressionContext) interface{}

	// Visit a parse tree produced by ABSLParser#PowerExpression.
	VisitPowerExpression(ctx *PowerExpressionContext) interface{}

	// Visit a parse tree produced by ABSLParser#PreIncrementExpression.
	VisitPreIncrementExpression(ctx *PreIncrementExpressionContext) interface{}

	// Visit a parse tree produced by ABSLParser#ObjectLiteralExpression.
	VisitObjectLiteralExpression(ctx *ObjectLiteralExpressionContext) interface{}

	// Visit a parse tree produced by ABSLParser#MetaExpression.
	VisitMetaExpression(ctx *MetaExpressionContext) interface{}

	// Visit a parse tree produced by ABSLParser#InExpression.
	VisitInExpression(ctx *InExpressionContext) interface{}

	// Visit a parse tree produced by ABSLParser#LogicalOrExpression.
	VisitLogicalOrExpression(ctx *LogicalOrExpressionContext) interface{}

	// Visit a parse tree produced by ABSLParser#NotExpression.
	VisitNotExpression(ctx *NotExpressionContext) interface{}

	// Visit a parse tree produced by ABSLParser#PreDecreaseExpression.
	VisitPreDecreaseExpression(ctx *PreDecreaseExpressionContext) interface{}

	// Visit a parse tree produced by ABSLParser#ArgumentsExpression.
	VisitArgumentsExpression(ctx *ArgumentsExpressionContext) interface{}

	// Visit a parse tree produced by ABSLParser#AwaitExpression.
	VisitAwaitExpression(ctx *AwaitExpressionContext) interface{}

	// Visit a parse tree produced by ABSLParser#ThisExpression.
	VisitThisExpression(ctx *ThisExpressionContext) interface{}

	// Visit a parse tree produced by ABSLParser#FunctionExpression.
	VisitFunctionExpression(ctx *FunctionExpressionContext) interface{}

	// Visit a parse tree produced by ABSLParser#UnaryMinusExpression.
	VisitUnaryMinusExpression(ctx *UnaryMinusExpressionContext) interface{}

	// Visit a parse tree produced by ABSLParser#AssignmentExpression.
	VisitAssignmentExpression(ctx *AssignmentExpressionContext) interface{}

	// Visit a parse tree produced by ABSLParser#PostDecreaseExpression.
	VisitPostDecreaseExpression(ctx *PostDecreaseExpressionContext) interface{}

	// Visit a parse tree produced by ABSLParser#TypeofExpression.
	VisitTypeofExpression(ctx *TypeofExpressionContext) interface{}

	// Visit a parse tree produced by ABSLParser#InstanceofExpression.
	VisitInstanceofExpression(ctx *InstanceofExpressionContext) interface{}

	// Visit a parse tree produced by ABSLParser#UnaryPlusExpression.
	VisitUnaryPlusExpression(ctx *UnaryPlusExpressionContext) interface{}

	// Visit a parse tree produced by ABSLParser#DeleteExpression.
	VisitDeleteExpression(ctx *DeleteExpressionContext) interface{}

	// Visit a parse tree produced by ABSLParser#ImportExpression.
	VisitImportExpression(ctx *ImportExpressionContext) interface{}

	// Visit a parse tree produced by ABSLParser#EqualityExpression.
	VisitEqualityExpression(ctx *EqualityExpressionContext) interface{}

	// Visit a parse tree produced by ABSLParser#BitXOrExpression.
	VisitBitXOrExpression(ctx *BitXOrExpressionContext) interface{}

	// Visit a parse tree produced by ABSLParser#SuperExpression.
	VisitSuperExpression(ctx *SuperExpressionContext) interface{}

	// Visit a parse tree produced by ABSLParser#MultiplicativeExpression.
	VisitMultiplicativeExpression(ctx *MultiplicativeExpressionContext) interface{}

	// Visit a parse tree produced by ABSLParser#BitShiftExpression.
	VisitBitShiftExpression(ctx *BitShiftExpressionContext) interface{}

	// Visit a parse tree produced by ABSLParser#ParenthesizedExpression.
	VisitParenthesizedExpression(ctx *ParenthesizedExpressionContext) interface{}

	// Visit a parse tree produced by ABSLParser#AdditiveExpression.
	VisitAdditiveExpression(ctx *AdditiveExpressionContext) interface{}

	// Visit a parse tree produced by ABSLParser#RelationalExpression.
	VisitRelationalExpression(ctx *RelationalExpressionContext) interface{}

	// Visit a parse tree produced by ABSLParser#PostIncrementExpression.
	VisitPostIncrementExpression(ctx *PostIncrementExpressionContext) interface{}

	// Visit a parse tree produced by ABSLParser#YieldExpression.
	VisitYieldExpression(ctx *YieldExpressionContext) interface{}

	// Visit a parse tree produced by ABSLParser#BitNotExpression.
	VisitBitNotExpression(ctx *BitNotExpressionContext) interface{}

	// Visit a parse tree produced by ABSLParser#NewExpression.
	VisitNewExpression(ctx *NewExpressionContext) interface{}

	// Visit a parse tree produced by ABSLParser#LiteralExpression.
	VisitLiteralExpression(ctx *LiteralExpressionContext) interface{}

	// Visit a parse tree produced by ABSLParser#ArrayLiteralExpression.
	VisitArrayLiteralExpression(ctx *ArrayLiteralExpressionContext) interface{}

	// Visit a parse tree produced by ABSLParser#MemberDotExpression.
	VisitMemberDotExpression(ctx *MemberDotExpressionContext) interface{}

	// Visit a parse tree produced by ABSLParser#ClassExpression.
	VisitClassExpression(ctx *ClassExpressionContext) interface{}

	// Visit a parse tree produced by ABSLParser#MemberIndexExpression.
	VisitMemberIndexExpression(ctx *MemberIndexExpressionContext) interface{}

	// Visit a parse tree produced by ABSLParser#IdentifierExpression.
	VisitIdentifierExpression(ctx *IdentifierExpressionContext) interface{}

	// Visit a parse tree produced by ABSLParser#BitAndExpression.
	VisitBitAndExpression(ctx *BitAndExpressionContext) interface{}

	// Visit a parse tree produced by ABSLParser#BitOrExpression.
	VisitBitOrExpression(ctx *BitOrExpressionContext) interface{}

	// Visit a parse tree produced by ABSLParser#AssignmentOperatorExpression.
	VisitAssignmentOperatorExpression(ctx *AssignmentOperatorExpressionContext) interface{}

	// Visit a parse tree produced by ABSLParser#VoidExpression.
	VisitVoidExpression(ctx *VoidExpressionContext) interface{}

	// Visit a parse tree produced by ABSLParser#CoalesceExpression.
	VisitCoalesceExpression(ctx *CoalesceExpressionContext) interface{}

	// Visit a parse tree produced by ABSLParser#assignable.
	VisitAssignable(ctx *AssignableContext) interface{}

	// Visit a parse tree produced by ABSLParser#FunctionDecl.
	VisitFunctionDecl(ctx *FunctionDeclContext) interface{}

	// Visit a parse tree produced by ABSLParser#AnoymousFunctionDecl.
	VisitAnoymousFunctionDecl(ctx *AnoymousFunctionDeclContext) interface{}

	// Visit a parse tree produced by ABSLParser#ArrowFunction.
	VisitArrowFunction(ctx *ArrowFunctionContext) interface{}

	// Visit a parse tree produced by ABSLParser#arrowFunctionParameters.
	VisitArrowFunctionParameters(ctx *ArrowFunctionParametersContext) interface{}

	// Visit a parse tree produced by ABSLParser#arrowFunctionBody.
	VisitArrowFunctionBody(ctx *ArrowFunctionBodyContext) interface{}

	// Visit a parse tree produced by ABSLParser#assignmentOperator.
	VisitAssignmentOperator(ctx *AssignmentOperatorContext) interface{}

	// Visit a parse tree produced by ABSLParser#literal.
	VisitLiteral(ctx *LiteralContext) interface{}

	// Visit a parse tree produced by ABSLParser#numericLiteral.
	VisitNumericLiteral(ctx *NumericLiteralContext) interface{}

	// Visit a parse tree produced by ABSLParser#bigintLiteral.
	VisitBigintLiteral(ctx *BigintLiteralContext) interface{}

	// Visit a parse tree produced by ABSLParser#identifierName.
	VisitIdentifierName(ctx *IdentifierNameContext) interface{}

	// Visit a parse tree produced by ABSLParser#reservedWord.
	VisitReservedWord(ctx *ReservedWordContext) interface{}

	// Visit a parse tree produced by ABSLParser#keyword.
	VisitKeyword(ctx *KeywordContext) interface{}

	// Visit a parse tree produced by ABSLParser#getter.
	VisitGetter(ctx *GetterContext) interface{}

	// Visit a parse tree produced by ABSLParser#setter.
	VisitSetter(ctx *SetterContext) interface{}

	// Visit a parse tree produced by ABSLParser#eos.
	VisitEos(ctx *EosContext) interface{}
}

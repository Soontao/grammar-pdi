// Code generated from ./ABSLParser.g4 by ANTLR 4.7.2. DO NOT EDIT.

package abslgo // ABSLParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseABSLParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseABSLParserVisitor) VisitProgram(ctx *ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseABSLParserVisitor) VisitSourceElement(ctx *SourceElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseABSLParserVisitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseABSLParserVisitor) VisitBlock(ctx *BlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseABSLParserVisitor) VisitStatementList(ctx *StatementListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseABSLParserVisitor) VisitImportStatement(ctx *ImportStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseABSLParserVisitor) VisitImportFromBlock(ctx *ImportFromBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseABSLParserVisitor) VisitModuleItems(ctx *ModuleItemsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseABSLParserVisitor) VisitImportDefault(ctx *ImportDefaultContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseABSLParserVisitor) VisitImportNamespace(ctx *ImportNamespaceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseABSLParserVisitor) VisitImportFrom(ctx *ImportFromContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseABSLParserVisitor) VisitAliasName(ctx *AliasNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseABSLParserVisitor) VisitExportDeclaration(ctx *ExportDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseABSLParserVisitor) VisitExportDefaultDeclaration(ctx *ExportDefaultDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseABSLParserVisitor) VisitExportFromBlock(ctx *ExportFromBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseABSLParserVisitor) VisitDeclaration(ctx *DeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseABSLParserVisitor) VisitVariableStatement(ctx *VariableStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseABSLParserVisitor) VisitVariableDeclarationList(ctx *VariableDeclarationListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseABSLParserVisitor) VisitVariableDeclaration(ctx *VariableDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseABSLParserVisitor) VisitVariableType(ctx *VariableTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseABSLParserVisitor) VisitProgramEmptyStatement(ctx *ProgramEmptyStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseABSLParserVisitor) VisitExpressionStatement(ctx *ExpressionStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseABSLParserVisitor) VisitIfStatement(ctx *IfStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseABSLParserVisitor) VisitDoStatement(ctx *DoStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseABSLParserVisitor) VisitWhileStatement(ctx *WhileStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseABSLParserVisitor) VisitForStatement(ctx *ForStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseABSLParserVisitor) VisitForInStatement(ctx *ForInStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseABSLParserVisitor) VisitForOfStatement(ctx *ForOfStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseABSLParserVisitor) VisitForeachStatement(ctx *ForeachStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseABSLParserVisitor) VisitVarModifier(ctx *VarModifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseABSLParserVisitor) VisitContinueStatement(ctx *ContinueStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseABSLParserVisitor) VisitBreakStatement(ctx *BreakStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseABSLParserVisitor) VisitReturnStatement(ctx *ReturnStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseABSLParserVisitor) VisitYieldStatement(ctx *YieldStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseABSLParserVisitor) VisitWithStatement(ctx *WithStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseABSLParserVisitor) VisitSwitchStatement(ctx *SwitchStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseABSLParserVisitor) VisitCaseBlock(ctx *CaseBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseABSLParserVisitor) VisitCaseClauses(ctx *CaseClausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseABSLParserVisitor) VisitCaseClause(ctx *CaseClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseABSLParserVisitor) VisitDefaultClause(ctx *DefaultClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseABSLParserVisitor) VisitLabelledStatement(ctx *LabelledStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseABSLParserVisitor) VisitThrowStatement(ctx *ThrowStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseABSLParserVisitor) VisitTryStatement(ctx *TryStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseABSLParserVisitor) VisitCatchProduction(ctx *CatchProductionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseABSLParserVisitor) VisitFinallyProduction(ctx *FinallyProductionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseABSLParserVisitor) VisitDebuggerStatement(ctx *DebuggerStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseABSLParserVisitor) VisitFunctionDeclaration(ctx *FunctionDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseABSLParserVisitor) VisitClassDeclaration(ctx *ClassDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseABSLParserVisitor) VisitClassTail(ctx *ClassTailContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseABSLParserVisitor) VisitClassElement(ctx *ClassElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseABSLParserVisitor) VisitMethodDefinition(ctx *MethodDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseABSLParserVisitor) VisitFormalParameterList(ctx *FormalParameterListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseABSLParserVisitor) VisitFormalParameterArg(ctx *FormalParameterArgContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseABSLParserVisitor) VisitLastFormalParameterArg(ctx *LastFormalParameterArgContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseABSLParserVisitor) VisitFunctionBody(ctx *FunctionBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseABSLParserVisitor) VisitSourceElements(ctx *SourceElementsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseABSLParserVisitor) VisitArrayLiteral(ctx *ArrayLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseABSLParserVisitor) VisitElementList(ctx *ElementListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseABSLParserVisitor) VisitArrayElement(ctx *ArrayElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseABSLParserVisitor) VisitObjectLiteral(ctx *ObjectLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseABSLParserVisitor) VisitPropertyExpressionAssignment(ctx *PropertyExpressionAssignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseABSLParserVisitor) VisitComputedPropertyExpressionAssignment(ctx *ComputedPropertyExpressionAssignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseABSLParserVisitor) VisitFunctionProperty(ctx *FunctionPropertyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseABSLParserVisitor) VisitPropertyGetter(ctx *PropertyGetterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseABSLParserVisitor) VisitPropertySetter(ctx *PropertySetterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseABSLParserVisitor) VisitPropertyShorthand(ctx *PropertyShorthandContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseABSLParserVisitor) VisitPropertyName(ctx *PropertyNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseABSLParserVisitor) VisitArguments(ctx *ArgumentsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseABSLParserVisitor) VisitArgument(ctx *ArgumentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseABSLParserVisitor) VisitExpressionSequence(ctx *ExpressionSequenceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseABSLParserVisitor) VisitTemplateStringExpression(ctx *TemplateStringExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseABSLParserVisitor) VisitTernaryExpression(ctx *TernaryExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseABSLParserVisitor) VisitLogicalAndExpression(ctx *LogicalAndExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseABSLParserVisitor) VisitPowerExpression(ctx *PowerExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseABSLParserVisitor) VisitPreIncrementExpression(ctx *PreIncrementExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseABSLParserVisitor) VisitObjectLiteralExpression(ctx *ObjectLiteralExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseABSLParserVisitor) VisitMetaExpression(ctx *MetaExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseABSLParserVisitor) VisitInExpression(ctx *InExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseABSLParserVisitor) VisitLogicalOrExpression(ctx *LogicalOrExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseABSLParserVisitor) VisitNotExpression(ctx *NotExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseABSLParserVisitor) VisitPreDecreaseExpression(ctx *PreDecreaseExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseABSLParserVisitor) VisitArgumentsExpression(ctx *ArgumentsExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseABSLParserVisitor) VisitAwaitExpression(ctx *AwaitExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseABSLParserVisitor) VisitThisExpression(ctx *ThisExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseABSLParserVisitor) VisitFunctionExpression(ctx *FunctionExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseABSLParserVisitor) VisitUnaryMinusExpression(ctx *UnaryMinusExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseABSLParserVisitor) VisitAssignmentExpression(ctx *AssignmentExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseABSLParserVisitor) VisitPostDecreaseExpression(ctx *PostDecreaseExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseABSLParserVisitor) VisitTypeofExpression(ctx *TypeofExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseABSLParserVisitor) VisitInstanceofExpression(ctx *InstanceofExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseABSLParserVisitor) VisitUnaryPlusExpression(ctx *UnaryPlusExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseABSLParserVisitor) VisitDeleteExpression(ctx *DeleteExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseABSLParserVisitor) VisitImportExpression(ctx *ImportExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseABSLParserVisitor) VisitEqualityExpression(ctx *EqualityExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseABSLParserVisitor) VisitBitXOrExpression(ctx *BitXOrExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseABSLParserVisitor) VisitSuperExpression(ctx *SuperExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseABSLParserVisitor) VisitMultiplicativeExpression(ctx *MultiplicativeExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseABSLParserVisitor) VisitBitShiftExpression(ctx *BitShiftExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseABSLParserVisitor) VisitParenthesizedExpression(ctx *ParenthesizedExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseABSLParserVisitor) VisitAdditiveExpression(ctx *AdditiveExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseABSLParserVisitor) VisitRelationalExpression(ctx *RelationalExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseABSLParserVisitor) VisitPostIncrementExpression(ctx *PostIncrementExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseABSLParserVisitor) VisitYieldExpression(ctx *YieldExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseABSLParserVisitor) VisitBitNotExpression(ctx *BitNotExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseABSLParserVisitor) VisitNewExpression(ctx *NewExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseABSLParserVisitor) VisitLiteralExpression(ctx *LiteralExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseABSLParserVisitor) VisitArrayLiteralExpression(ctx *ArrayLiteralExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseABSLParserVisitor) VisitMemberDotExpression(ctx *MemberDotExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseABSLParserVisitor) VisitClassExpression(ctx *ClassExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseABSLParserVisitor) VisitMemberIndexExpression(ctx *MemberIndexExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseABSLParserVisitor) VisitIdentifierExpression(ctx *IdentifierExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseABSLParserVisitor) VisitBitAndExpression(ctx *BitAndExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseABSLParserVisitor) VisitBitOrExpression(ctx *BitOrExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseABSLParserVisitor) VisitAssignmentOperatorExpression(ctx *AssignmentOperatorExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseABSLParserVisitor) VisitVoidExpression(ctx *VoidExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseABSLParserVisitor) VisitCoalesceExpression(ctx *CoalesceExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseABSLParserVisitor) VisitAssignable(ctx *AssignableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseABSLParserVisitor) VisitFunctionDecl(ctx *FunctionDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseABSLParserVisitor) VisitAnoymousFunctionDecl(ctx *AnoymousFunctionDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseABSLParserVisitor) VisitArrowFunction(ctx *ArrowFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseABSLParserVisitor) VisitArrowFunctionParameters(ctx *ArrowFunctionParametersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseABSLParserVisitor) VisitArrowFunctionBody(ctx *ArrowFunctionBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseABSLParserVisitor) VisitAssignmentOperator(ctx *AssignmentOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseABSLParserVisitor) VisitLiteral(ctx *LiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseABSLParserVisitor) VisitNumericLiteral(ctx *NumericLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseABSLParserVisitor) VisitBigintLiteral(ctx *BigintLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseABSLParserVisitor) VisitIdentifierName(ctx *IdentifierNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseABSLParserVisitor) VisitReservedWord(ctx *ReservedWordContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseABSLParserVisitor) VisitKeyword(ctx *KeywordContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseABSLParserVisitor) VisitGetter(ctx *GetterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseABSLParserVisitor) VisitSetter(ctx *SetterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseABSLParserVisitor) VisitEos(ctx *EosContext) interface{} {
	return v.VisitChildren(ctx)
}

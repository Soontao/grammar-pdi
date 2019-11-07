// Code generated from ./ABSLParser.g4 by ANTLR 4.7.2. DO NOT EDIT.

package abslgo // ABSLParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseABSLParserListener is a complete listener for a parse tree produced by ABSLParser.
type BaseABSLParserListener struct{}

var _ ABSLParserListener = &BaseABSLParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseABSLParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseABSLParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseABSLParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseABSLParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseABSLParserListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseABSLParserListener) ExitProgram(ctx *ProgramContext) {}

// EnterSourceElement is called when production sourceElement is entered.
func (s *BaseABSLParserListener) EnterSourceElement(ctx *SourceElementContext) {}

// ExitSourceElement is called when production sourceElement is exited.
func (s *BaseABSLParserListener) ExitSourceElement(ctx *SourceElementContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseABSLParserListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseABSLParserListener) ExitStatement(ctx *StatementContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseABSLParserListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseABSLParserListener) ExitBlock(ctx *BlockContext) {}

// EnterStatementList is called when production statementList is entered.
func (s *BaseABSLParserListener) EnterStatementList(ctx *StatementListContext) {}

// ExitStatementList is called when production statementList is exited.
func (s *BaseABSLParserListener) ExitStatementList(ctx *StatementListContext) {}

// EnterImportStatement is called when production importStatement is entered.
func (s *BaseABSLParserListener) EnterImportStatement(ctx *ImportStatementContext) {}

// ExitImportStatement is called when production importStatement is exited.
func (s *BaseABSLParserListener) ExitImportStatement(ctx *ImportStatementContext) {}

// EnterImportFromBlock is called when production importFromBlock is entered.
func (s *BaseABSLParserListener) EnterImportFromBlock(ctx *ImportFromBlockContext) {}

// ExitImportFromBlock is called when production importFromBlock is exited.
func (s *BaseABSLParserListener) ExitImportFromBlock(ctx *ImportFromBlockContext) {}

// EnterModuleItems is called when production moduleItems is entered.
func (s *BaseABSLParserListener) EnterModuleItems(ctx *ModuleItemsContext) {}

// ExitModuleItems is called when production moduleItems is exited.
func (s *BaseABSLParserListener) ExitModuleItems(ctx *ModuleItemsContext) {}

// EnterImportDefault is called when production importDefault is entered.
func (s *BaseABSLParserListener) EnterImportDefault(ctx *ImportDefaultContext) {}

// ExitImportDefault is called when production importDefault is exited.
func (s *BaseABSLParserListener) ExitImportDefault(ctx *ImportDefaultContext) {}

// EnterImportNamespace is called when production importNamespace is entered.
func (s *BaseABSLParserListener) EnterImportNamespace(ctx *ImportNamespaceContext) {}

// ExitImportNamespace is called when production importNamespace is exited.
func (s *BaseABSLParserListener) ExitImportNamespace(ctx *ImportNamespaceContext) {}

// EnterImportFrom is called when production importFrom is entered.
func (s *BaseABSLParserListener) EnterImportFrom(ctx *ImportFromContext) {}

// ExitImportFrom is called when production importFrom is exited.
func (s *BaseABSLParserListener) ExitImportFrom(ctx *ImportFromContext) {}

// EnterAliasName is called when production aliasName is entered.
func (s *BaseABSLParserListener) EnterAliasName(ctx *AliasNameContext) {}

// ExitAliasName is called when production aliasName is exited.
func (s *BaseABSLParserListener) ExitAliasName(ctx *AliasNameContext) {}

// EnterExportDeclaration is called when production ExportDeclaration is entered.
func (s *BaseABSLParserListener) EnterExportDeclaration(ctx *ExportDeclarationContext) {}

// ExitExportDeclaration is called when production ExportDeclaration is exited.
func (s *BaseABSLParserListener) ExitExportDeclaration(ctx *ExportDeclarationContext) {}

// EnterExportDefaultDeclaration is called when production ExportDefaultDeclaration is entered.
func (s *BaseABSLParserListener) EnterExportDefaultDeclaration(ctx *ExportDefaultDeclarationContext) {}

// ExitExportDefaultDeclaration is called when production ExportDefaultDeclaration is exited.
func (s *BaseABSLParserListener) ExitExportDefaultDeclaration(ctx *ExportDefaultDeclarationContext) {}

// EnterExportFromBlock is called when production exportFromBlock is entered.
func (s *BaseABSLParserListener) EnterExportFromBlock(ctx *ExportFromBlockContext) {}

// ExitExportFromBlock is called when production exportFromBlock is exited.
func (s *BaseABSLParserListener) ExitExportFromBlock(ctx *ExportFromBlockContext) {}

// EnterDeclaration is called when production declaration is entered.
func (s *BaseABSLParserListener) EnterDeclaration(ctx *DeclarationContext) {}

// ExitDeclaration is called when production declaration is exited.
func (s *BaseABSLParserListener) ExitDeclaration(ctx *DeclarationContext) {}

// EnterVariableStatement is called when production variableStatement is entered.
func (s *BaseABSLParserListener) EnterVariableStatement(ctx *VariableStatementContext) {}

// ExitVariableStatement is called when production variableStatement is exited.
func (s *BaseABSLParserListener) ExitVariableStatement(ctx *VariableStatementContext) {}

// EnterVariableDeclarationList is called when production variableDeclarationList is entered.
func (s *BaseABSLParserListener) EnterVariableDeclarationList(ctx *VariableDeclarationListContext) {}

// ExitVariableDeclarationList is called when production variableDeclarationList is exited.
func (s *BaseABSLParserListener) ExitVariableDeclarationList(ctx *VariableDeclarationListContext) {}

// EnterVariableDeclaration is called when production variableDeclaration is entered.
func (s *BaseABSLParserListener) EnterVariableDeclaration(ctx *VariableDeclarationContext) {}

// ExitVariableDeclaration is called when production variableDeclaration is exited.
func (s *BaseABSLParserListener) ExitVariableDeclaration(ctx *VariableDeclarationContext) {}

// EnterVariableType is called when production variableType is entered.
func (s *BaseABSLParserListener) EnterVariableType(ctx *VariableTypeContext) {}

// ExitVariableType is called when production variableType is exited.
func (s *BaseABSLParserListener) ExitVariableType(ctx *VariableTypeContext) {}

// EnterProgramEmptyStatement is called when production programEmptyStatement is entered.
func (s *BaseABSLParserListener) EnterProgramEmptyStatement(ctx *ProgramEmptyStatementContext) {}

// ExitProgramEmptyStatement is called when production programEmptyStatement is exited.
func (s *BaseABSLParserListener) ExitProgramEmptyStatement(ctx *ProgramEmptyStatementContext) {}

// EnterExpressionStatement is called when production expressionStatement is entered.
func (s *BaseABSLParserListener) EnterExpressionStatement(ctx *ExpressionStatementContext) {}

// ExitExpressionStatement is called when production expressionStatement is exited.
func (s *BaseABSLParserListener) ExitExpressionStatement(ctx *ExpressionStatementContext) {}

// EnterIfStatement is called when production ifStatement is entered.
func (s *BaseABSLParserListener) EnterIfStatement(ctx *IfStatementContext) {}

// ExitIfStatement is called when production ifStatement is exited.
func (s *BaseABSLParserListener) ExitIfStatement(ctx *IfStatementContext) {}

// EnterDoStatement is called when production DoStatement is entered.
func (s *BaseABSLParserListener) EnterDoStatement(ctx *DoStatementContext) {}

// ExitDoStatement is called when production DoStatement is exited.
func (s *BaseABSLParserListener) ExitDoStatement(ctx *DoStatementContext) {}

// EnterWhileStatement is called when production WhileStatement is entered.
func (s *BaseABSLParserListener) EnterWhileStatement(ctx *WhileStatementContext) {}

// ExitWhileStatement is called when production WhileStatement is exited.
func (s *BaseABSLParserListener) ExitWhileStatement(ctx *WhileStatementContext) {}

// EnterForStatement is called when production ForStatement is entered.
func (s *BaseABSLParserListener) EnterForStatement(ctx *ForStatementContext) {}

// ExitForStatement is called when production ForStatement is exited.
func (s *BaseABSLParserListener) ExitForStatement(ctx *ForStatementContext) {}

// EnterForInStatement is called when production ForInStatement is entered.
func (s *BaseABSLParserListener) EnterForInStatement(ctx *ForInStatementContext) {}

// ExitForInStatement is called when production ForInStatement is exited.
func (s *BaseABSLParserListener) ExitForInStatement(ctx *ForInStatementContext) {}

// EnterForOfStatement is called when production ForOfStatement is entered.
func (s *BaseABSLParserListener) EnterForOfStatement(ctx *ForOfStatementContext) {}

// ExitForOfStatement is called when production ForOfStatement is exited.
func (s *BaseABSLParserListener) ExitForOfStatement(ctx *ForOfStatementContext) {}

// EnterForeachStatement is called when production ForeachStatement is entered.
func (s *BaseABSLParserListener) EnterForeachStatement(ctx *ForeachStatementContext) {}

// ExitForeachStatement is called when production ForeachStatement is exited.
func (s *BaseABSLParserListener) ExitForeachStatement(ctx *ForeachStatementContext) {}

// EnterVarModifier is called when production varModifier is entered.
func (s *BaseABSLParserListener) EnterVarModifier(ctx *VarModifierContext) {}

// ExitVarModifier is called when production varModifier is exited.
func (s *BaseABSLParserListener) ExitVarModifier(ctx *VarModifierContext) {}

// EnterContinueStatement is called when production continueStatement is entered.
func (s *BaseABSLParserListener) EnterContinueStatement(ctx *ContinueStatementContext) {}

// ExitContinueStatement is called when production continueStatement is exited.
func (s *BaseABSLParserListener) ExitContinueStatement(ctx *ContinueStatementContext) {}

// EnterBreakStatement is called when production breakStatement is entered.
func (s *BaseABSLParserListener) EnterBreakStatement(ctx *BreakStatementContext) {}

// ExitBreakStatement is called when production breakStatement is exited.
func (s *BaseABSLParserListener) ExitBreakStatement(ctx *BreakStatementContext) {}

// EnterReturnStatement is called when production returnStatement is entered.
func (s *BaseABSLParserListener) EnterReturnStatement(ctx *ReturnStatementContext) {}

// ExitReturnStatement is called when production returnStatement is exited.
func (s *BaseABSLParserListener) ExitReturnStatement(ctx *ReturnStatementContext) {}

// EnterYieldStatement is called when production yieldStatement is entered.
func (s *BaseABSLParserListener) EnterYieldStatement(ctx *YieldStatementContext) {}

// ExitYieldStatement is called when production yieldStatement is exited.
func (s *BaseABSLParserListener) ExitYieldStatement(ctx *YieldStatementContext) {}

// EnterWithStatement is called when production withStatement is entered.
func (s *BaseABSLParserListener) EnterWithStatement(ctx *WithStatementContext) {}

// ExitWithStatement is called when production withStatement is exited.
func (s *BaseABSLParserListener) ExitWithStatement(ctx *WithStatementContext) {}

// EnterSwitchStatement is called when production switchStatement is entered.
func (s *BaseABSLParserListener) EnterSwitchStatement(ctx *SwitchStatementContext) {}

// ExitSwitchStatement is called when production switchStatement is exited.
func (s *BaseABSLParserListener) ExitSwitchStatement(ctx *SwitchStatementContext) {}

// EnterCaseBlock is called when production caseBlock is entered.
func (s *BaseABSLParserListener) EnterCaseBlock(ctx *CaseBlockContext) {}

// ExitCaseBlock is called when production caseBlock is exited.
func (s *BaseABSLParserListener) ExitCaseBlock(ctx *CaseBlockContext) {}

// EnterCaseClauses is called when production caseClauses is entered.
func (s *BaseABSLParserListener) EnterCaseClauses(ctx *CaseClausesContext) {}

// ExitCaseClauses is called when production caseClauses is exited.
func (s *BaseABSLParserListener) ExitCaseClauses(ctx *CaseClausesContext) {}

// EnterCaseClause is called when production caseClause is entered.
func (s *BaseABSLParserListener) EnterCaseClause(ctx *CaseClauseContext) {}

// ExitCaseClause is called when production caseClause is exited.
func (s *BaseABSLParserListener) ExitCaseClause(ctx *CaseClauseContext) {}

// EnterDefaultClause is called when production defaultClause is entered.
func (s *BaseABSLParserListener) EnterDefaultClause(ctx *DefaultClauseContext) {}

// ExitDefaultClause is called when production defaultClause is exited.
func (s *BaseABSLParserListener) ExitDefaultClause(ctx *DefaultClauseContext) {}

// EnterLabelledStatement is called when production labelledStatement is entered.
func (s *BaseABSLParserListener) EnterLabelledStatement(ctx *LabelledStatementContext) {}

// ExitLabelledStatement is called when production labelledStatement is exited.
func (s *BaseABSLParserListener) ExitLabelledStatement(ctx *LabelledStatementContext) {}

// EnterThrowStatement is called when production throwStatement is entered.
func (s *BaseABSLParserListener) EnterThrowStatement(ctx *ThrowStatementContext) {}

// ExitThrowStatement is called when production throwStatement is exited.
func (s *BaseABSLParserListener) ExitThrowStatement(ctx *ThrowStatementContext) {}

// EnterTryStatement is called when production tryStatement is entered.
func (s *BaseABSLParserListener) EnterTryStatement(ctx *TryStatementContext) {}

// ExitTryStatement is called when production tryStatement is exited.
func (s *BaseABSLParserListener) ExitTryStatement(ctx *TryStatementContext) {}

// EnterCatchProduction is called when production catchProduction is entered.
func (s *BaseABSLParserListener) EnterCatchProduction(ctx *CatchProductionContext) {}

// ExitCatchProduction is called when production catchProduction is exited.
func (s *BaseABSLParserListener) ExitCatchProduction(ctx *CatchProductionContext) {}

// EnterFinallyProduction is called when production finallyProduction is entered.
func (s *BaseABSLParserListener) EnterFinallyProduction(ctx *FinallyProductionContext) {}

// ExitFinallyProduction is called when production finallyProduction is exited.
func (s *BaseABSLParserListener) ExitFinallyProduction(ctx *FinallyProductionContext) {}

// EnterDebuggerStatement is called when production debuggerStatement is entered.
func (s *BaseABSLParserListener) EnterDebuggerStatement(ctx *DebuggerStatementContext) {}

// ExitDebuggerStatement is called when production debuggerStatement is exited.
func (s *BaseABSLParserListener) ExitDebuggerStatement(ctx *DebuggerStatementContext) {}

// EnterFunctionDeclaration is called when production functionDeclaration is entered.
func (s *BaseABSLParserListener) EnterFunctionDeclaration(ctx *FunctionDeclarationContext) {}

// ExitFunctionDeclaration is called when production functionDeclaration is exited.
func (s *BaseABSLParserListener) ExitFunctionDeclaration(ctx *FunctionDeclarationContext) {}

// EnterClassDeclaration is called when production classDeclaration is entered.
func (s *BaseABSLParserListener) EnterClassDeclaration(ctx *ClassDeclarationContext) {}

// ExitClassDeclaration is called when production classDeclaration is exited.
func (s *BaseABSLParserListener) ExitClassDeclaration(ctx *ClassDeclarationContext) {}

// EnterClassTail is called when production classTail is entered.
func (s *BaseABSLParserListener) EnterClassTail(ctx *ClassTailContext) {}

// ExitClassTail is called when production classTail is exited.
func (s *BaseABSLParserListener) ExitClassTail(ctx *ClassTailContext) {}

// EnterClassElement is called when production classElement is entered.
func (s *BaseABSLParserListener) EnterClassElement(ctx *ClassElementContext) {}

// ExitClassElement is called when production classElement is exited.
func (s *BaseABSLParserListener) ExitClassElement(ctx *ClassElementContext) {}

// EnterMethodDefinition is called when production methodDefinition is entered.
func (s *BaseABSLParserListener) EnterMethodDefinition(ctx *MethodDefinitionContext) {}

// ExitMethodDefinition is called when production methodDefinition is exited.
func (s *BaseABSLParserListener) ExitMethodDefinition(ctx *MethodDefinitionContext) {}

// EnterFormalParameterList is called when production formalParameterList is entered.
func (s *BaseABSLParserListener) EnterFormalParameterList(ctx *FormalParameterListContext) {}

// ExitFormalParameterList is called when production formalParameterList is exited.
func (s *BaseABSLParserListener) ExitFormalParameterList(ctx *FormalParameterListContext) {}

// EnterFormalParameterArg is called when production formalParameterArg is entered.
func (s *BaseABSLParserListener) EnterFormalParameterArg(ctx *FormalParameterArgContext) {}

// ExitFormalParameterArg is called when production formalParameterArg is exited.
func (s *BaseABSLParserListener) ExitFormalParameterArg(ctx *FormalParameterArgContext) {}

// EnterLastFormalParameterArg is called when production lastFormalParameterArg is entered.
func (s *BaseABSLParserListener) EnterLastFormalParameterArg(ctx *LastFormalParameterArgContext) {}

// ExitLastFormalParameterArg is called when production lastFormalParameterArg is exited.
func (s *BaseABSLParserListener) ExitLastFormalParameterArg(ctx *LastFormalParameterArgContext) {}

// EnterFunctionBody is called when production functionBody is entered.
func (s *BaseABSLParserListener) EnterFunctionBody(ctx *FunctionBodyContext) {}

// ExitFunctionBody is called when production functionBody is exited.
func (s *BaseABSLParserListener) ExitFunctionBody(ctx *FunctionBodyContext) {}

// EnterSourceElements is called when production sourceElements is entered.
func (s *BaseABSLParserListener) EnterSourceElements(ctx *SourceElementsContext) {}

// ExitSourceElements is called when production sourceElements is exited.
func (s *BaseABSLParserListener) ExitSourceElements(ctx *SourceElementsContext) {}

// EnterArrayLiteral is called when production arrayLiteral is entered.
func (s *BaseABSLParserListener) EnterArrayLiteral(ctx *ArrayLiteralContext) {}

// ExitArrayLiteral is called when production arrayLiteral is exited.
func (s *BaseABSLParserListener) ExitArrayLiteral(ctx *ArrayLiteralContext) {}

// EnterElementList is called when production elementList is entered.
func (s *BaseABSLParserListener) EnterElementList(ctx *ElementListContext) {}

// ExitElementList is called when production elementList is exited.
func (s *BaseABSLParserListener) ExitElementList(ctx *ElementListContext) {}

// EnterArrayElement is called when production arrayElement is entered.
func (s *BaseABSLParserListener) EnterArrayElement(ctx *ArrayElementContext) {}

// ExitArrayElement is called when production arrayElement is exited.
func (s *BaseABSLParserListener) ExitArrayElement(ctx *ArrayElementContext) {}

// EnterObjectLiteral is called when production objectLiteral is entered.
func (s *BaseABSLParserListener) EnterObjectLiteral(ctx *ObjectLiteralContext) {}

// ExitObjectLiteral is called when production objectLiteral is exited.
func (s *BaseABSLParserListener) ExitObjectLiteral(ctx *ObjectLiteralContext) {}

// EnterPropertyExpressionAssignment is called when production PropertyExpressionAssignment is entered.
func (s *BaseABSLParserListener) EnterPropertyExpressionAssignment(ctx *PropertyExpressionAssignmentContext) {
}

// ExitPropertyExpressionAssignment is called when production PropertyExpressionAssignment is exited.
func (s *BaseABSLParserListener) ExitPropertyExpressionAssignment(ctx *PropertyExpressionAssignmentContext) {
}

// EnterComputedPropertyExpressionAssignment is called when production ComputedPropertyExpressionAssignment is entered.
func (s *BaseABSLParserListener) EnterComputedPropertyExpressionAssignment(ctx *ComputedPropertyExpressionAssignmentContext) {
}

// ExitComputedPropertyExpressionAssignment is called when production ComputedPropertyExpressionAssignment is exited.
func (s *BaseABSLParserListener) ExitComputedPropertyExpressionAssignment(ctx *ComputedPropertyExpressionAssignmentContext) {
}

// EnterFunctionProperty is called when production FunctionProperty is entered.
func (s *BaseABSLParserListener) EnterFunctionProperty(ctx *FunctionPropertyContext) {}

// ExitFunctionProperty is called when production FunctionProperty is exited.
func (s *BaseABSLParserListener) ExitFunctionProperty(ctx *FunctionPropertyContext) {}

// EnterPropertyGetter is called when production PropertyGetter is entered.
func (s *BaseABSLParserListener) EnterPropertyGetter(ctx *PropertyGetterContext) {}

// ExitPropertyGetter is called when production PropertyGetter is exited.
func (s *BaseABSLParserListener) ExitPropertyGetter(ctx *PropertyGetterContext) {}

// EnterPropertySetter is called when production PropertySetter is entered.
func (s *BaseABSLParserListener) EnterPropertySetter(ctx *PropertySetterContext) {}

// ExitPropertySetter is called when production PropertySetter is exited.
func (s *BaseABSLParserListener) ExitPropertySetter(ctx *PropertySetterContext) {}

// EnterPropertyShorthand is called when production PropertyShorthand is entered.
func (s *BaseABSLParserListener) EnterPropertyShorthand(ctx *PropertyShorthandContext) {}

// ExitPropertyShorthand is called when production PropertyShorthand is exited.
func (s *BaseABSLParserListener) ExitPropertyShorthand(ctx *PropertyShorthandContext) {}

// EnterPropertyName is called when production propertyName is entered.
func (s *BaseABSLParserListener) EnterPropertyName(ctx *PropertyNameContext) {}

// ExitPropertyName is called when production propertyName is exited.
func (s *BaseABSLParserListener) ExitPropertyName(ctx *PropertyNameContext) {}

// EnterArguments is called when production arguments is entered.
func (s *BaseABSLParserListener) EnterArguments(ctx *ArgumentsContext) {}

// ExitArguments is called when production arguments is exited.
func (s *BaseABSLParserListener) ExitArguments(ctx *ArgumentsContext) {}

// EnterArgument is called when production argument is entered.
func (s *BaseABSLParserListener) EnterArgument(ctx *ArgumentContext) {}

// ExitArgument is called when production argument is exited.
func (s *BaseABSLParserListener) ExitArgument(ctx *ArgumentContext) {}

// EnterExpressionSequence is called when production expressionSequence is entered.
func (s *BaseABSLParserListener) EnterExpressionSequence(ctx *ExpressionSequenceContext) {}

// ExitExpressionSequence is called when production expressionSequence is exited.
func (s *BaseABSLParserListener) ExitExpressionSequence(ctx *ExpressionSequenceContext) {}

// EnterTemplateStringExpression is called when production TemplateStringExpression is entered.
func (s *BaseABSLParserListener) EnterTemplateStringExpression(ctx *TemplateStringExpressionContext) {}

// ExitTemplateStringExpression is called when production TemplateStringExpression is exited.
func (s *BaseABSLParserListener) ExitTemplateStringExpression(ctx *TemplateStringExpressionContext) {}

// EnterTernaryExpression is called when production TernaryExpression is entered.
func (s *BaseABSLParserListener) EnterTernaryExpression(ctx *TernaryExpressionContext) {}

// ExitTernaryExpression is called when production TernaryExpression is exited.
func (s *BaseABSLParserListener) ExitTernaryExpression(ctx *TernaryExpressionContext) {}

// EnterLogicalAndExpression is called when production LogicalAndExpression is entered.
func (s *BaseABSLParserListener) EnterLogicalAndExpression(ctx *LogicalAndExpressionContext) {}

// ExitLogicalAndExpression is called when production LogicalAndExpression is exited.
func (s *BaseABSLParserListener) ExitLogicalAndExpression(ctx *LogicalAndExpressionContext) {}

// EnterPowerExpression is called when production PowerExpression is entered.
func (s *BaseABSLParserListener) EnterPowerExpression(ctx *PowerExpressionContext) {}

// ExitPowerExpression is called when production PowerExpression is exited.
func (s *BaseABSLParserListener) ExitPowerExpression(ctx *PowerExpressionContext) {}

// EnterPreIncrementExpression is called when production PreIncrementExpression is entered.
func (s *BaseABSLParserListener) EnterPreIncrementExpression(ctx *PreIncrementExpressionContext) {}

// ExitPreIncrementExpression is called when production PreIncrementExpression is exited.
func (s *BaseABSLParserListener) ExitPreIncrementExpression(ctx *PreIncrementExpressionContext) {}

// EnterObjectLiteralExpression is called when production ObjectLiteralExpression is entered.
func (s *BaseABSLParserListener) EnterObjectLiteralExpression(ctx *ObjectLiteralExpressionContext) {}

// ExitObjectLiteralExpression is called when production ObjectLiteralExpression is exited.
func (s *BaseABSLParserListener) ExitObjectLiteralExpression(ctx *ObjectLiteralExpressionContext) {}

// EnterMetaExpression is called when production MetaExpression is entered.
func (s *BaseABSLParserListener) EnterMetaExpression(ctx *MetaExpressionContext) {}

// ExitMetaExpression is called when production MetaExpression is exited.
func (s *BaseABSLParserListener) ExitMetaExpression(ctx *MetaExpressionContext) {}

// EnterInExpression is called when production InExpression is entered.
func (s *BaseABSLParserListener) EnterInExpression(ctx *InExpressionContext) {}

// ExitInExpression is called when production InExpression is exited.
func (s *BaseABSLParserListener) ExitInExpression(ctx *InExpressionContext) {}

// EnterLogicalOrExpression is called when production LogicalOrExpression is entered.
func (s *BaseABSLParserListener) EnterLogicalOrExpression(ctx *LogicalOrExpressionContext) {}

// ExitLogicalOrExpression is called when production LogicalOrExpression is exited.
func (s *BaseABSLParserListener) ExitLogicalOrExpression(ctx *LogicalOrExpressionContext) {}

// EnterNotExpression is called when production NotExpression is entered.
func (s *BaseABSLParserListener) EnterNotExpression(ctx *NotExpressionContext) {}

// ExitNotExpression is called when production NotExpression is exited.
func (s *BaseABSLParserListener) ExitNotExpression(ctx *NotExpressionContext) {}

// EnterPreDecreaseExpression is called when production PreDecreaseExpression is entered.
func (s *BaseABSLParserListener) EnterPreDecreaseExpression(ctx *PreDecreaseExpressionContext) {}

// ExitPreDecreaseExpression is called when production PreDecreaseExpression is exited.
func (s *BaseABSLParserListener) ExitPreDecreaseExpression(ctx *PreDecreaseExpressionContext) {}

// EnterArgumentsExpression is called when production ArgumentsExpression is entered.
func (s *BaseABSLParserListener) EnterArgumentsExpression(ctx *ArgumentsExpressionContext) {}

// ExitArgumentsExpression is called when production ArgumentsExpression is exited.
func (s *BaseABSLParserListener) ExitArgumentsExpression(ctx *ArgumentsExpressionContext) {}

// EnterAwaitExpression is called when production AwaitExpression is entered.
func (s *BaseABSLParserListener) EnterAwaitExpression(ctx *AwaitExpressionContext) {}

// ExitAwaitExpression is called when production AwaitExpression is exited.
func (s *BaseABSLParserListener) ExitAwaitExpression(ctx *AwaitExpressionContext) {}

// EnterThisExpression is called when production ThisExpression is entered.
func (s *BaseABSLParserListener) EnterThisExpression(ctx *ThisExpressionContext) {}

// ExitThisExpression is called when production ThisExpression is exited.
func (s *BaseABSLParserListener) ExitThisExpression(ctx *ThisExpressionContext) {}

// EnterFunctionExpression is called when production FunctionExpression is entered.
func (s *BaseABSLParserListener) EnterFunctionExpression(ctx *FunctionExpressionContext) {}

// ExitFunctionExpression is called when production FunctionExpression is exited.
func (s *BaseABSLParserListener) ExitFunctionExpression(ctx *FunctionExpressionContext) {}

// EnterUnaryMinusExpression is called when production UnaryMinusExpression is entered.
func (s *BaseABSLParserListener) EnterUnaryMinusExpression(ctx *UnaryMinusExpressionContext) {}

// ExitUnaryMinusExpression is called when production UnaryMinusExpression is exited.
func (s *BaseABSLParserListener) ExitUnaryMinusExpression(ctx *UnaryMinusExpressionContext) {}

// EnterAssignmentExpression is called when production AssignmentExpression is entered.
func (s *BaseABSLParserListener) EnterAssignmentExpression(ctx *AssignmentExpressionContext) {}

// ExitAssignmentExpression is called when production AssignmentExpression is exited.
func (s *BaseABSLParserListener) ExitAssignmentExpression(ctx *AssignmentExpressionContext) {}

// EnterPostDecreaseExpression is called when production PostDecreaseExpression is entered.
func (s *BaseABSLParserListener) EnterPostDecreaseExpression(ctx *PostDecreaseExpressionContext) {}

// ExitPostDecreaseExpression is called when production PostDecreaseExpression is exited.
func (s *BaseABSLParserListener) ExitPostDecreaseExpression(ctx *PostDecreaseExpressionContext) {}

// EnterTypeofExpression is called when production TypeofExpression is entered.
func (s *BaseABSLParserListener) EnterTypeofExpression(ctx *TypeofExpressionContext) {}

// ExitTypeofExpression is called when production TypeofExpression is exited.
func (s *BaseABSLParserListener) ExitTypeofExpression(ctx *TypeofExpressionContext) {}

// EnterInstanceofExpression is called when production InstanceofExpression is entered.
func (s *BaseABSLParserListener) EnterInstanceofExpression(ctx *InstanceofExpressionContext) {}

// ExitInstanceofExpression is called when production InstanceofExpression is exited.
func (s *BaseABSLParserListener) ExitInstanceofExpression(ctx *InstanceofExpressionContext) {}

// EnterUnaryPlusExpression is called when production UnaryPlusExpression is entered.
func (s *BaseABSLParserListener) EnterUnaryPlusExpression(ctx *UnaryPlusExpressionContext) {}

// ExitUnaryPlusExpression is called when production UnaryPlusExpression is exited.
func (s *BaseABSLParserListener) ExitUnaryPlusExpression(ctx *UnaryPlusExpressionContext) {}

// EnterDeleteExpression is called when production DeleteExpression is entered.
func (s *BaseABSLParserListener) EnterDeleteExpression(ctx *DeleteExpressionContext) {}

// ExitDeleteExpression is called when production DeleteExpression is exited.
func (s *BaseABSLParserListener) ExitDeleteExpression(ctx *DeleteExpressionContext) {}

// EnterImportExpression is called when production ImportExpression is entered.
func (s *BaseABSLParserListener) EnterImportExpression(ctx *ImportExpressionContext) {}

// ExitImportExpression is called when production ImportExpression is exited.
func (s *BaseABSLParserListener) ExitImportExpression(ctx *ImportExpressionContext) {}

// EnterEqualityExpression is called when production EqualityExpression is entered.
func (s *BaseABSLParserListener) EnterEqualityExpression(ctx *EqualityExpressionContext) {}

// ExitEqualityExpression is called when production EqualityExpression is exited.
func (s *BaseABSLParserListener) ExitEqualityExpression(ctx *EqualityExpressionContext) {}

// EnterBitXOrExpression is called when production BitXOrExpression is entered.
func (s *BaseABSLParserListener) EnterBitXOrExpression(ctx *BitXOrExpressionContext) {}

// ExitBitXOrExpression is called when production BitXOrExpression is exited.
func (s *BaseABSLParserListener) ExitBitXOrExpression(ctx *BitXOrExpressionContext) {}

// EnterSuperExpression is called when production SuperExpression is entered.
func (s *BaseABSLParserListener) EnterSuperExpression(ctx *SuperExpressionContext) {}

// ExitSuperExpression is called when production SuperExpression is exited.
func (s *BaseABSLParserListener) ExitSuperExpression(ctx *SuperExpressionContext) {}

// EnterMultiplicativeExpression is called when production MultiplicativeExpression is entered.
func (s *BaseABSLParserListener) EnterMultiplicativeExpression(ctx *MultiplicativeExpressionContext) {}

// ExitMultiplicativeExpression is called when production MultiplicativeExpression is exited.
func (s *BaseABSLParserListener) ExitMultiplicativeExpression(ctx *MultiplicativeExpressionContext) {}

// EnterBitShiftExpression is called when production BitShiftExpression is entered.
func (s *BaseABSLParserListener) EnterBitShiftExpression(ctx *BitShiftExpressionContext) {}

// ExitBitShiftExpression is called when production BitShiftExpression is exited.
func (s *BaseABSLParserListener) ExitBitShiftExpression(ctx *BitShiftExpressionContext) {}

// EnterParenthesizedExpression is called when production ParenthesizedExpression is entered.
func (s *BaseABSLParserListener) EnterParenthesizedExpression(ctx *ParenthesizedExpressionContext) {}

// ExitParenthesizedExpression is called when production ParenthesizedExpression is exited.
func (s *BaseABSLParserListener) ExitParenthesizedExpression(ctx *ParenthesizedExpressionContext) {}

// EnterAdditiveExpression is called when production AdditiveExpression is entered.
func (s *BaseABSLParserListener) EnterAdditiveExpression(ctx *AdditiveExpressionContext) {}

// ExitAdditiveExpression is called when production AdditiveExpression is exited.
func (s *BaseABSLParserListener) ExitAdditiveExpression(ctx *AdditiveExpressionContext) {}

// EnterRelationalExpression is called when production RelationalExpression is entered.
func (s *BaseABSLParserListener) EnterRelationalExpression(ctx *RelationalExpressionContext) {}

// ExitRelationalExpression is called when production RelationalExpression is exited.
func (s *BaseABSLParserListener) ExitRelationalExpression(ctx *RelationalExpressionContext) {}

// EnterPostIncrementExpression is called when production PostIncrementExpression is entered.
func (s *BaseABSLParserListener) EnterPostIncrementExpression(ctx *PostIncrementExpressionContext) {}

// ExitPostIncrementExpression is called when production PostIncrementExpression is exited.
func (s *BaseABSLParserListener) ExitPostIncrementExpression(ctx *PostIncrementExpressionContext) {}

// EnterYieldExpression is called when production YieldExpression is entered.
func (s *BaseABSLParserListener) EnterYieldExpression(ctx *YieldExpressionContext) {}

// ExitYieldExpression is called when production YieldExpression is exited.
func (s *BaseABSLParserListener) ExitYieldExpression(ctx *YieldExpressionContext) {}

// EnterBitNotExpression is called when production BitNotExpression is entered.
func (s *BaseABSLParserListener) EnterBitNotExpression(ctx *BitNotExpressionContext) {}

// ExitBitNotExpression is called when production BitNotExpression is exited.
func (s *BaseABSLParserListener) ExitBitNotExpression(ctx *BitNotExpressionContext) {}

// EnterNewExpression is called when production NewExpression is entered.
func (s *BaseABSLParserListener) EnterNewExpression(ctx *NewExpressionContext) {}

// ExitNewExpression is called when production NewExpression is exited.
func (s *BaseABSLParserListener) ExitNewExpression(ctx *NewExpressionContext) {}

// EnterLiteralExpression is called when production LiteralExpression is entered.
func (s *BaseABSLParserListener) EnterLiteralExpression(ctx *LiteralExpressionContext) {}

// ExitLiteralExpression is called when production LiteralExpression is exited.
func (s *BaseABSLParserListener) ExitLiteralExpression(ctx *LiteralExpressionContext) {}

// EnterArrayLiteralExpression is called when production ArrayLiteralExpression is entered.
func (s *BaseABSLParserListener) EnterArrayLiteralExpression(ctx *ArrayLiteralExpressionContext) {}

// ExitArrayLiteralExpression is called when production ArrayLiteralExpression is exited.
func (s *BaseABSLParserListener) ExitArrayLiteralExpression(ctx *ArrayLiteralExpressionContext) {}

// EnterMemberDotExpression is called when production MemberDotExpression is entered.
func (s *BaseABSLParserListener) EnterMemberDotExpression(ctx *MemberDotExpressionContext) {}

// ExitMemberDotExpression is called when production MemberDotExpression is exited.
func (s *BaseABSLParserListener) ExitMemberDotExpression(ctx *MemberDotExpressionContext) {}

// EnterClassExpression is called when production ClassExpression is entered.
func (s *BaseABSLParserListener) EnterClassExpression(ctx *ClassExpressionContext) {}

// ExitClassExpression is called when production ClassExpression is exited.
func (s *BaseABSLParserListener) ExitClassExpression(ctx *ClassExpressionContext) {}

// EnterMemberIndexExpression is called when production MemberIndexExpression is entered.
func (s *BaseABSLParserListener) EnterMemberIndexExpression(ctx *MemberIndexExpressionContext) {}

// ExitMemberIndexExpression is called when production MemberIndexExpression is exited.
func (s *BaseABSLParserListener) ExitMemberIndexExpression(ctx *MemberIndexExpressionContext) {}

// EnterIdentifierExpression is called when production IdentifierExpression is entered.
func (s *BaseABSLParserListener) EnterIdentifierExpression(ctx *IdentifierExpressionContext) {}

// ExitIdentifierExpression is called when production IdentifierExpression is exited.
func (s *BaseABSLParserListener) ExitIdentifierExpression(ctx *IdentifierExpressionContext) {}

// EnterBitAndExpression is called when production BitAndExpression is entered.
func (s *BaseABSLParserListener) EnterBitAndExpression(ctx *BitAndExpressionContext) {}

// ExitBitAndExpression is called when production BitAndExpression is exited.
func (s *BaseABSLParserListener) ExitBitAndExpression(ctx *BitAndExpressionContext) {}

// EnterBitOrExpression is called when production BitOrExpression is entered.
func (s *BaseABSLParserListener) EnterBitOrExpression(ctx *BitOrExpressionContext) {}

// ExitBitOrExpression is called when production BitOrExpression is exited.
func (s *BaseABSLParserListener) ExitBitOrExpression(ctx *BitOrExpressionContext) {}

// EnterAssignmentOperatorExpression is called when production AssignmentOperatorExpression is entered.
func (s *BaseABSLParserListener) EnterAssignmentOperatorExpression(ctx *AssignmentOperatorExpressionContext) {
}

// ExitAssignmentOperatorExpression is called when production AssignmentOperatorExpression is exited.
func (s *BaseABSLParserListener) ExitAssignmentOperatorExpression(ctx *AssignmentOperatorExpressionContext) {
}

// EnterVoidExpression is called when production VoidExpression is entered.
func (s *BaseABSLParserListener) EnterVoidExpression(ctx *VoidExpressionContext) {}

// ExitVoidExpression is called when production VoidExpression is exited.
func (s *BaseABSLParserListener) ExitVoidExpression(ctx *VoidExpressionContext) {}

// EnterCoalesceExpression is called when production CoalesceExpression is entered.
func (s *BaseABSLParserListener) EnterCoalesceExpression(ctx *CoalesceExpressionContext) {}

// ExitCoalesceExpression is called when production CoalesceExpression is exited.
func (s *BaseABSLParserListener) ExitCoalesceExpression(ctx *CoalesceExpressionContext) {}

// EnterAssignable is called when production assignable is entered.
func (s *BaseABSLParserListener) EnterAssignable(ctx *AssignableContext) {}

// ExitAssignable is called when production assignable is exited.
func (s *BaseABSLParserListener) ExitAssignable(ctx *AssignableContext) {}

// EnterFunctionDecl is called when production FunctionDecl is entered.
func (s *BaseABSLParserListener) EnterFunctionDecl(ctx *FunctionDeclContext) {}

// ExitFunctionDecl is called when production FunctionDecl is exited.
func (s *BaseABSLParserListener) ExitFunctionDecl(ctx *FunctionDeclContext) {}

// EnterAnoymousFunctionDecl is called when production AnoymousFunctionDecl is entered.
func (s *BaseABSLParserListener) EnterAnoymousFunctionDecl(ctx *AnoymousFunctionDeclContext) {}

// ExitAnoymousFunctionDecl is called when production AnoymousFunctionDecl is exited.
func (s *BaseABSLParserListener) ExitAnoymousFunctionDecl(ctx *AnoymousFunctionDeclContext) {}

// EnterArrowFunction is called when production ArrowFunction is entered.
func (s *BaseABSLParserListener) EnterArrowFunction(ctx *ArrowFunctionContext) {}

// ExitArrowFunction is called when production ArrowFunction is exited.
func (s *BaseABSLParserListener) ExitArrowFunction(ctx *ArrowFunctionContext) {}

// EnterArrowFunctionParameters is called when production arrowFunctionParameters is entered.
func (s *BaseABSLParserListener) EnterArrowFunctionParameters(ctx *ArrowFunctionParametersContext) {}

// ExitArrowFunctionParameters is called when production arrowFunctionParameters is exited.
func (s *BaseABSLParserListener) ExitArrowFunctionParameters(ctx *ArrowFunctionParametersContext) {}

// EnterArrowFunctionBody is called when production arrowFunctionBody is entered.
func (s *BaseABSLParserListener) EnterArrowFunctionBody(ctx *ArrowFunctionBodyContext) {}

// ExitArrowFunctionBody is called when production arrowFunctionBody is exited.
func (s *BaseABSLParserListener) ExitArrowFunctionBody(ctx *ArrowFunctionBodyContext) {}

// EnterAssignmentOperator is called when production assignmentOperator is entered.
func (s *BaseABSLParserListener) EnterAssignmentOperator(ctx *AssignmentOperatorContext) {}

// ExitAssignmentOperator is called when production assignmentOperator is exited.
func (s *BaseABSLParserListener) ExitAssignmentOperator(ctx *AssignmentOperatorContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BaseABSLParserListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BaseABSLParserListener) ExitLiteral(ctx *LiteralContext) {}

// EnterNumericLiteral is called when production numericLiteral is entered.
func (s *BaseABSLParserListener) EnterNumericLiteral(ctx *NumericLiteralContext) {}

// ExitNumericLiteral is called when production numericLiteral is exited.
func (s *BaseABSLParserListener) ExitNumericLiteral(ctx *NumericLiteralContext) {}

// EnterBigintLiteral is called when production bigintLiteral is entered.
func (s *BaseABSLParserListener) EnterBigintLiteral(ctx *BigintLiteralContext) {}

// ExitBigintLiteral is called when production bigintLiteral is exited.
func (s *BaseABSLParserListener) ExitBigintLiteral(ctx *BigintLiteralContext) {}

// EnterIdentifierName is called when production identifierName is entered.
func (s *BaseABSLParserListener) EnterIdentifierName(ctx *IdentifierNameContext) {}

// ExitIdentifierName is called when production identifierName is exited.
func (s *BaseABSLParserListener) ExitIdentifierName(ctx *IdentifierNameContext) {}

// EnterReservedWord is called when production reservedWord is entered.
func (s *BaseABSLParserListener) EnterReservedWord(ctx *ReservedWordContext) {}

// ExitReservedWord is called when production reservedWord is exited.
func (s *BaseABSLParserListener) ExitReservedWord(ctx *ReservedWordContext) {}

// EnterKeyword is called when production keyword is entered.
func (s *BaseABSLParserListener) EnterKeyword(ctx *KeywordContext) {}

// ExitKeyword is called when production keyword is exited.
func (s *BaseABSLParserListener) ExitKeyword(ctx *KeywordContext) {}

// EnterGetter is called when production getter is entered.
func (s *BaseABSLParserListener) EnterGetter(ctx *GetterContext) {}

// ExitGetter is called when production getter is exited.
func (s *BaseABSLParserListener) ExitGetter(ctx *GetterContext) {}

// EnterSetter is called when production setter is entered.
func (s *BaseABSLParserListener) EnterSetter(ctx *SetterContext) {}

// ExitSetter is called when production setter is exited.
func (s *BaseABSLParserListener) ExitSetter(ctx *SetterContext) {}

// EnterEos is called when production eos is entered.
func (s *BaseABSLParserListener) EnterEos(ctx *EosContext) {}

// ExitEos is called when production eos is exited.
func (s *BaseABSLParserListener) ExitEos(ctx *EosContext) {}

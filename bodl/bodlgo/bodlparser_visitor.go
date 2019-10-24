// Code generated from ./BODLParser.g4 by ANTLR 4.7.2. DO NOT EDIT.

package bodlgo // BODLParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by BODLParser.
type BODLParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by BODLParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by BODLParser#statements.
	VisitStatements(ctx *StatementsContext) interface{}

	// Visit a parse tree produced by BODLParser#importStatement.
	VisitImportStatement(ctx *ImportStatementContext) interface{}

	// Visit a parse tree produced by BODLParser#definitions.
	VisitDefinitions(ctx *DefinitionsContext) interface{}

	// Visit a parse tree produced by BODLParser#definition.
	VisitDefinition(ctx *DefinitionContext) interface{}

	// Visit a parse tree produced by BODLParser#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by BODLParser#itemList.
	VisitItemList(ctx *ItemListContext) interface{}

	// Visit a parse tree produced by BODLParser#element.
	VisitElement(ctx *ElementContext) interface{}

	// Visit a parse tree produced by BODLParser#boAction.
	VisitBoAction(ctx *BoActionContext) interface{}

	// Visit a parse tree produced by BODLParser#message.
	VisitMessage(ctx *MessageContext) interface{}

	// Visit a parse tree produced by BODLParser#node.
	VisitNode(ctx *NodeContext) interface{}

	// Visit a parse tree produced by BODLParser#association.
	VisitAssociation(ctx *AssociationContext) interface{}

	// Visit a parse tree produced by BODLParser#associationUsingDefinition.
	VisitAssociationUsingDefinition(ctx *AssociationUsingDefinitionContext) interface{}

	// Visit a parse tree produced by BODLParser#valuationDefinition.
	VisitValuationDefinition(ctx *ValuationDefinitionContext) interface{}

	// Visit a parse tree produced by BODLParser#valutaionExpressionList.
	VisitValutaionExpressionList(ctx *ValutaionExpressionListContext) interface{}

	// Visit a parse tree produced by BODLParser#valutaionExpression.
	VisitValutaionExpression(ctx *ValutaionExpressionContext) interface{}

	// Visit a parse tree produced by BODLParser#raiseMessage.
	VisitRaiseMessage(ctx *RaiseMessageContext) interface{}

	// Visit a parse tree produced by BODLParser#annotations.
	VisitAnnotations(ctx *AnnotationsContext) interface{}

	// Visit a parse tree produced by BODLParser#annotation.
	VisitAnnotation(ctx *AnnotationContext) interface{}

	// Visit a parse tree produced by BODLParser#typeList.
	VisitTypeList(ctx *TypeListContext) interface{}

	// Visit a parse tree produced by BODLParser#typeDeclaration.
	VisitTypeDeclaration(ctx *TypeDeclarationContext) interface{}

	// Visit a parse tree produced by BODLParser#typeDefaultValue.
	VisitTypeDefaultValue(ctx *TypeDefaultValueContext) interface{}

	// Visit a parse tree produced by BODLParser#valueAssignList.
	VisitValueAssignList(ctx *ValueAssignListContext) interface{}

	// Visit a parse tree produced by BODLParser#valueAssign.
	VisitValueAssign(ctx *ValueAssignContext) interface{}

	// Visit a parse tree produced by BODLParser#multiplicity.
	VisitMultiplicity(ctx *MultiplicityContext) interface{}

	// Visit a parse tree produced by BODLParser#memberExpression.
	VisitMemberExpression(ctx *MemberExpressionContext) interface{}

	// Visit a parse tree produced by BODLParser#identifierList.
	VisitIdentifierList(ctx *IdentifierListContext) interface{}

	// Visit a parse tree produced by BODLParser#keyword.
	VisitKeyword(ctx *KeywordContext) interface{}

	// Visit a parse tree produced by BODLParser#literal.
	VisitLiteral(ctx *LiteralContext) interface{}

	// Visit a parse tree produced by BODLParser#comments.
	VisitComments(ctx *CommentsContext) interface{}

	// Visit a parse tree produced by BODLParser#compareOperator.
	VisitCompareOperator(ctx *CompareOperatorContext) interface{}

	// Visit a parse tree produced by BODLParser#logicOperator.
	VisitLogicOperator(ctx *LogicOperatorContext) interface{}

	// Visit a parse tree produced by BODLParser#identifier.
	VisitIdentifier(ctx *IdentifierContext) interface{}
}

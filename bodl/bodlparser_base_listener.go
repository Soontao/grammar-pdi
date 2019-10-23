// Code generated from ../BODLParser.g4 by ANTLR 4.7.2. DO NOT EDIT.

package bodl // BODLParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseBODLParserListener is a complete listener for a parse tree produced by BODLParser.
type BaseBODLParserListener struct{}

var _ BODLParserListener = &BaseBODLParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseBODLParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseBODLParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseBODLParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseBODLParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseBODLParserListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseBODLParserListener) ExitProgram(ctx *ProgramContext) {}

// EnterImportStatment is called when production importStatment is entered.
func (s *BaseBODLParserListener) EnterImportStatment(ctx *ImportStatmentContext) {}

// ExitImportStatment is called when production importStatment is exited.
func (s *BaseBODLParserListener) ExitImportStatment(ctx *ImportStatmentContext) {}

// EnterDefinition is called when production definition is entered.
func (s *BaseBODLParserListener) EnterDefinition(ctx *DefinitionContext) {}

// ExitDefinition is called when production definition is exited.
func (s *BaseBODLParserListener) ExitDefinition(ctx *DefinitionContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseBODLParserListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseBODLParserListener) ExitBlock(ctx *BlockContext) {}

// EnterItemList is called when production itemList is entered.
func (s *BaseBODLParserListener) EnterItemList(ctx *ItemListContext) {}

// ExitItemList is called when production itemList is exited.
func (s *BaseBODLParserListener) ExitItemList(ctx *ItemListContext) {}

// EnterElement is called when production element is entered.
func (s *BaseBODLParserListener) EnterElement(ctx *ElementContext) {}

// ExitElement is called when production element is exited.
func (s *BaseBODLParserListener) ExitElement(ctx *ElementContext) {}

// EnterBoAction is called when production boAction is entered.
func (s *BaseBODLParserListener) EnterBoAction(ctx *BoActionContext) {}

// ExitBoAction is called when production boAction is exited.
func (s *BaseBODLParserListener) ExitBoAction(ctx *BoActionContext) {}

// EnterMessage is called when production message is entered.
func (s *BaseBODLParserListener) EnterMessage(ctx *MessageContext) {}

// ExitMessage is called when production message is exited.
func (s *BaseBODLParserListener) ExitMessage(ctx *MessageContext) {}

// EnterNode is called when production node is entered.
func (s *BaseBODLParserListener) EnterNode(ctx *NodeContext) {}

// ExitNode is called when production node is exited.
func (s *BaseBODLParserListener) ExitNode(ctx *NodeContext) {}

// EnterAssociation is called when production association is entered.
func (s *BaseBODLParserListener) EnterAssociation(ctx *AssociationContext) {}

// ExitAssociation is called when production association is exited.
func (s *BaseBODLParserListener) ExitAssociation(ctx *AssociationContext) {}

// EnterAssociationUsingDefinition is called when production associationUsingDefinition is entered.
func (s *BaseBODLParserListener) EnterAssociationUsingDefinition(ctx *AssociationUsingDefinitionContext) {
}

// ExitAssociationUsingDefinition is called when production associationUsingDefinition is exited.
func (s *BaseBODLParserListener) ExitAssociationUsingDefinition(ctx *AssociationUsingDefinitionContext) {
}

// EnterValuationDefinition is called when production valuationDefinition is entered.
func (s *BaseBODLParserListener) EnterValuationDefinition(ctx *ValuationDefinitionContext) {}

// ExitValuationDefinition is called when production valuationDefinition is exited.
func (s *BaseBODLParserListener) ExitValuationDefinition(ctx *ValuationDefinitionContext) {}

// EnterValutaionExpressionList is called when production valutaionExpressionList is entered.
func (s *BaseBODLParserListener) EnterValutaionExpressionList(ctx *ValutaionExpressionListContext) {}

// ExitValutaionExpressionList is called when production valutaionExpressionList is exited.
func (s *BaseBODLParserListener) ExitValutaionExpressionList(ctx *ValutaionExpressionListContext) {}

// EnterValutaionExpression is called when production valutaionExpression is entered.
func (s *BaseBODLParserListener) EnterValutaionExpression(ctx *ValutaionExpressionContext) {}

// ExitValutaionExpression is called when production valutaionExpression is exited.
func (s *BaseBODLParserListener) ExitValutaionExpression(ctx *ValutaionExpressionContext) {}

// EnterRaiseMessage is called when production raiseMessage is entered.
func (s *BaseBODLParserListener) EnterRaiseMessage(ctx *RaiseMessageContext) {}

// ExitRaiseMessage is called when production raiseMessage is exited.
func (s *BaseBODLParserListener) ExitRaiseMessage(ctx *RaiseMessageContext) {}

// EnterAnnotationList is called when production annotationList is entered.
func (s *BaseBODLParserListener) EnterAnnotationList(ctx *AnnotationListContext) {}

// ExitAnnotationList is called when production annotationList is exited.
func (s *BaseBODLParserListener) ExitAnnotationList(ctx *AnnotationListContext) {}

// EnterAnnotation is called when production annotation is entered.
func (s *BaseBODLParserListener) EnterAnnotation(ctx *AnnotationContext) {}

// ExitAnnotation is called when production annotation is exited.
func (s *BaseBODLParserListener) ExitAnnotation(ctx *AnnotationContext) {}

// EnterTypeList is called when production typeList is entered.
func (s *BaseBODLParserListener) EnterTypeList(ctx *TypeListContext) {}

// ExitTypeList is called when production typeList is exited.
func (s *BaseBODLParserListener) ExitTypeList(ctx *TypeListContext) {}

// EnterType is called when production type is entered.
func (s *BaseBODLParserListener) EnterType(ctx *TypeContext) {}

// ExitType is called when production type is exited.
func (s *BaseBODLParserListener) ExitType(ctx *TypeContext) {}

// EnterTypeDefaultValue is called when production typeDefaultValue is entered.
func (s *BaseBODLParserListener) EnterTypeDefaultValue(ctx *TypeDefaultValueContext) {}

// ExitTypeDefaultValue is called when production typeDefaultValue is exited.
func (s *BaseBODLParserListener) ExitTypeDefaultValue(ctx *TypeDefaultValueContext) {}

// EnterValueAssignList is called when production valueAssignList is entered.
func (s *BaseBODLParserListener) EnterValueAssignList(ctx *ValueAssignListContext) {}

// ExitValueAssignList is called when production valueAssignList is exited.
func (s *BaseBODLParserListener) ExitValueAssignList(ctx *ValueAssignListContext) {}

// EnterValueAssign is called when production valueAssign is entered.
func (s *BaseBODLParserListener) EnterValueAssign(ctx *ValueAssignContext) {}

// ExitValueAssign is called when production valueAssign is exited.
func (s *BaseBODLParserListener) ExitValueAssign(ctx *ValueAssignContext) {}

// EnterMultiplicity is called when production multiplicity is entered.
func (s *BaseBODLParserListener) EnterMultiplicity(ctx *MultiplicityContext) {}

// ExitMultiplicity is called when production multiplicity is exited.
func (s *BaseBODLParserListener) ExitMultiplicity(ctx *MultiplicityContext) {}

// EnterMemberExpression is called when production memberExpression is entered.
func (s *BaseBODLParserListener) EnterMemberExpression(ctx *MemberExpressionContext) {}

// ExitMemberExpression is called when production memberExpression is exited.
func (s *BaseBODLParserListener) ExitMemberExpression(ctx *MemberExpressionContext) {}

// EnterIdentifierList is called when production identifierList is entered.
func (s *BaseBODLParserListener) EnterIdentifierList(ctx *IdentifierListContext) {}

// ExitIdentifierList is called when production identifierList is exited.
func (s *BaseBODLParserListener) ExitIdentifierList(ctx *IdentifierListContext) {}

// EnterKeyword is called when production keyword is entered.
func (s *BaseBODLParserListener) EnterKeyword(ctx *KeywordContext) {}

// ExitKeyword is called when production keyword is exited.
func (s *BaseBODLParserListener) ExitKeyword(ctx *KeywordContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BaseBODLParserListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BaseBODLParserListener) ExitLiteral(ctx *LiteralContext) {}

// EnterComments is called when production comments is entered.
func (s *BaseBODLParserListener) EnterComments(ctx *CommentsContext) {}

// ExitComments is called when production comments is exited.
func (s *BaseBODLParserListener) ExitComments(ctx *CommentsContext) {}

// EnterCompareOperator is called when production compareOperator is entered.
func (s *BaseBODLParserListener) EnterCompareOperator(ctx *CompareOperatorContext) {}

// ExitCompareOperator is called when production compareOperator is exited.
func (s *BaseBODLParserListener) ExitCompareOperator(ctx *CompareOperatorContext) {}

// EnterLogicOperator is called when production logicOperator is entered.
func (s *BaseBODLParserListener) EnterLogicOperator(ctx *LogicOperatorContext) {}

// ExitLogicOperator is called when production logicOperator is exited.
func (s *BaseBODLParserListener) ExitLogicOperator(ctx *LogicOperatorContext) {}

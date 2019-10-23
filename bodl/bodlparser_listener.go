// Code generated from ../BODLParser.g4 by ANTLR 4.7.2. DO NOT EDIT.

package bodl // BODLParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BODLParserListener is a complete listener for a parse tree produced by BODLParser.
type BODLParserListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterImportStatment is called when entering the importStatment production.
	EnterImportStatment(c *ImportStatmentContext)

	// EnterDefinition is called when entering the definition production.
	EnterDefinition(c *DefinitionContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterItemList is called when entering the itemList production.
	EnterItemList(c *ItemListContext)

	// EnterElement is called when entering the element production.
	EnterElement(c *ElementContext)

	// EnterBoAction is called when entering the boAction production.
	EnterBoAction(c *BoActionContext)

	// EnterMessage is called when entering the message production.
	EnterMessage(c *MessageContext)

	// EnterNode is called when entering the node production.
	EnterNode(c *NodeContext)

	// EnterAssociation is called when entering the association production.
	EnterAssociation(c *AssociationContext)

	// EnterAssociationUsingDefinition is called when entering the associationUsingDefinition production.
	EnterAssociationUsingDefinition(c *AssociationUsingDefinitionContext)

	// EnterValuationDefinition is called when entering the valuationDefinition production.
	EnterValuationDefinition(c *ValuationDefinitionContext)

	// EnterValutaionExpressionList is called when entering the valutaionExpressionList production.
	EnterValutaionExpressionList(c *ValutaionExpressionListContext)

	// EnterValutaionExpression is called when entering the valutaionExpression production.
	EnterValutaionExpression(c *ValutaionExpressionContext)

	// EnterRaiseMessage is called when entering the raiseMessage production.
	EnterRaiseMessage(c *RaiseMessageContext)

	// EnterAnnotationList is called when entering the annotationList production.
	EnterAnnotationList(c *AnnotationListContext)

	// EnterAnnotation is called when entering the annotation production.
	EnterAnnotation(c *AnnotationContext)

	// EnterTypeList is called when entering the typeList production.
	EnterTypeList(c *TypeListContext)

	// EnterType is called when entering the type production.
	EnterType(c *TypeContext)

	// EnterTypeDefaultValue is called when entering the typeDefaultValue production.
	EnterTypeDefaultValue(c *TypeDefaultValueContext)

	// EnterValueAssignList is called when entering the valueAssignList production.
	EnterValueAssignList(c *ValueAssignListContext)

	// EnterValueAssign is called when entering the valueAssign production.
	EnterValueAssign(c *ValueAssignContext)

	// EnterMultiplicity is called when entering the multiplicity production.
	EnterMultiplicity(c *MultiplicityContext)

	// EnterMemberExpression is called when entering the memberExpression production.
	EnterMemberExpression(c *MemberExpressionContext)

	// EnterIdentifierList is called when entering the identifierList production.
	EnterIdentifierList(c *IdentifierListContext)

	// EnterKeyword is called when entering the keyword production.
	EnterKeyword(c *KeywordContext)

	// EnterLiteral is called when entering the literal production.
	EnterLiteral(c *LiteralContext)

	// EnterCopyright is called when entering the copyright production.
	EnterCopyright(c *CopyrightContext)

	// EnterComments is called when entering the comments production.
	EnterComments(c *CommentsContext)

	// EnterCompareOperator is called when entering the compareOperator production.
	EnterCompareOperator(c *CompareOperatorContext)

	// EnterLogicOperator is called when entering the logicOperator production.
	EnterLogicOperator(c *LogicOperatorContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitImportStatment is called when exiting the importStatment production.
	ExitImportStatment(c *ImportStatmentContext)

	// ExitDefinition is called when exiting the definition production.
	ExitDefinition(c *DefinitionContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitItemList is called when exiting the itemList production.
	ExitItemList(c *ItemListContext)

	// ExitElement is called when exiting the element production.
	ExitElement(c *ElementContext)

	// ExitBoAction is called when exiting the boAction production.
	ExitBoAction(c *BoActionContext)

	// ExitMessage is called when exiting the message production.
	ExitMessage(c *MessageContext)

	// ExitNode is called when exiting the node production.
	ExitNode(c *NodeContext)

	// ExitAssociation is called when exiting the association production.
	ExitAssociation(c *AssociationContext)

	// ExitAssociationUsingDefinition is called when exiting the associationUsingDefinition production.
	ExitAssociationUsingDefinition(c *AssociationUsingDefinitionContext)

	// ExitValuationDefinition is called when exiting the valuationDefinition production.
	ExitValuationDefinition(c *ValuationDefinitionContext)

	// ExitValutaionExpressionList is called when exiting the valutaionExpressionList production.
	ExitValutaionExpressionList(c *ValutaionExpressionListContext)

	// ExitValutaionExpression is called when exiting the valutaionExpression production.
	ExitValutaionExpression(c *ValutaionExpressionContext)

	// ExitRaiseMessage is called when exiting the raiseMessage production.
	ExitRaiseMessage(c *RaiseMessageContext)

	// ExitAnnotationList is called when exiting the annotationList production.
	ExitAnnotationList(c *AnnotationListContext)

	// ExitAnnotation is called when exiting the annotation production.
	ExitAnnotation(c *AnnotationContext)

	// ExitTypeList is called when exiting the typeList production.
	ExitTypeList(c *TypeListContext)

	// ExitType is called when exiting the type production.
	ExitType(c *TypeContext)

	// ExitTypeDefaultValue is called when exiting the typeDefaultValue production.
	ExitTypeDefaultValue(c *TypeDefaultValueContext)

	// ExitValueAssignList is called when exiting the valueAssignList production.
	ExitValueAssignList(c *ValueAssignListContext)

	// ExitValueAssign is called when exiting the valueAssign production.
	ExitValueAssign(c *ValueAssignContext)

	// ExitMultiplicity is called when exiting the multiplicity production.
	ExitMultiplicity(c *MultiplicityContext)

	// ExitMemberExpression is called when exiting the memberExpression production.
	ExitMemberExpression(c *MemberExpressionContext)

	// ExitIdentifierList is called when exiting the identifierList production.
	ExitIdentifierList(c *IdentifierListContext)

	// ExitKeyword is called when exiting the keyword production.
	ExitKeyword(c *KeywordContext)

	// ExitLiteral is called when exiting the literal production.
	ExitLiteral(c *LiteralContext)

	// ExitCopyright is called when exiting the copyright production.
	ExitCopyright(c *CopyrightContext)

	// ExitComments is called when exiting the comments production.
	ExitComments(c *CommentsContext)

	// ExitCompareOperator is called when exiting the compareOperator production.
	ExitCompareOperator(c *CompareOperatorContext)

	// ExitLogicOperator is called when exiting the logicOperator production.
	ExitLogicOperator(c *LogicOperatorContext)
}

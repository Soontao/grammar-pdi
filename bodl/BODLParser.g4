parser grammar BODLParser;

options {
	tokenVocab = BODLLexer;
}

program: comments? statments? businessObjectDefinition EOF;

statments: importStatment*;

importStatment:
	IMPORT memberExpression SemiColon
	| IMPORT memberExpression AS Identifier SemiColon;

businessObjectDefinition:
	annotationList? BUSINESSOBJECT Identifier messageRaiseDefinition? block;

block: OpenBrace itemList CloseBrace;

itemList: (element | message | node | action | association)*;

element:
	comments? annotationList? ELEMENT Identifier Colon type SemiColon;

action:
	comments? ACTION Identifier messageRaiseDefinition? SemiColon;

message:
	comments? MESSAGE Identifier TEXT StringLiteral Colon typeList SemiColon;

node:
	comments? annotationList? NODE Identifier multiplicity? messageRaiseDefinition? block;

association:
	comments? annotationList? ASSOCIATION Identifier multiplicity? TO Identifier
		associationUsingDefinition? valuationDefinition? SemiColon;

associationUsingDefinition: USING Identifier;

valuationDefinition:
	VALUATION OpenParen valutaionExpressionList CloseParen;

valutaionExpressionList:
	valutaionExpression
	| valutaionExpression logicOperator valutaionExpressionList;

valutaionExpression: Identifier compareOperator literal;

messageRaiseDefinition: RAISES? identifierList?;

annotationList: annotation*;

annotation:
	OpenBracket Identifier CloseBracket
	| OpenBracket Identifier OpenParen Identifier CloseParen CloseBracket
	| OpenBracket Identifier OpenParen literal CloseParen CloseBracket;

typeList: type | type Comma typeList;

type:
	Identifier typeDefaultValue?
	| memberExpression Colon Identifier typeDefaultValue?;

typeDefaultValue:
	Assign literal
	| Assign OpenBrace valueAssignList CloseBrace;

valueAssignList:
	valueAssign
	| valueAssign Comma valueAssignList;

valueAssign: Identifier Assign literal;

multiplicity:
	OpenBracket literal Comma literal CloseBracket
	| OpenBracket literal Comma N CloseBracket;

memberExpression: Identifier | Identifier Dot memberExpression;

identifierList: Identifier | Identifier Comma identifierList;

keyword:
	IMPORT
	| BUSINESSOBJECT
	| TO
	| ASSOCIATION
	| ELEMENT
	| NODE
	| ACTION
	| MESSAGE
	| RAISES
	| USING
	| TEXT
	| AS;

literal: DecimalLiteral | BooleanLiteral | StringLiteral;

comments: (SingleLineComment | MultiLineComment)*;

compareOperator:
	Equals_
	| NotEquals
	| MoreThan
	| GreaterThanEquals
	| LessThan
	| LessThanEquals;

logicOperator: And | Or;
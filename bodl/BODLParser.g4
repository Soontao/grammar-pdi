parser grammar BODLParser;

options {
	tokenVocab = BODLLexer;
}

program: EOF | (importStatment)* definition* EOF;

importStatment:
	comments? IMPORT memberExpression SemiColon
	| comments? IMPORT memberExpression AS Identifier SemiColon;

definition:
	comments? annotationList? BUSINESSOBJECT (Identifier | type) raiseMessage? block;

block: OpenBrace itemList comments CloseBrace;

itemList: (element | message | node | boAction | association)*;

element:
	comments? annotationList? ELEMENT Identifier Colon type SemiColon;

boAction: comments? ACTION Identifier raiseMessage? SemiColon;

message:
	comments? MESSAGE Identifier TEXT StringLiteral Colon typeList SemiColon;

node:
	comments? annotationList? NODE Identifier multiplicity? raiseMessage? block?;

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

raiseMessage: RAISES? identifierList?;

annotationList: annotation*;

annotation:
	CustomAnnotationStart? OpenBracket Identifier CloseBracket
	| CustomAnnotationStart? OpenBracket Identifier OpenParen Identifier CloseParen CloseBracket
	| CustomAnnotationStart? OpenBracket Identifier OpenParen literal CloseParen CloseBracket;

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

comments: ( SingleLineComment | MultiLineComment)*;

compareOperator:
	Equals_
	| NotEquals
	| MoreThan
	| GreaterThanEquals
	| LessThan
	| LessThanEquals;

logicOperator: And | Or;

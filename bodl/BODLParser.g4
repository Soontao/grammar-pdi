parser grammar BODLParser;

options {
	tokenVocab = BODLLexer;
}

program: EOF | statements definitions EOF;

statements: (importStatement)*;

importStatement:
	comments? IMPORT memberExpression SemiColon
	| comments? IMPORT memberExpression AS identifier SemiColon;

definitions: definition*;

definition:
	comments? annotations? BUSINESSOBJECT (
		identifier
		| typeDeclaration
	) raiseMessage? block;

block: OpenBrace itemList comments CloseBrace;

itemList: (element | message | node | boAction | association)*;

element:
	comments? annotations? ELEMENT identifier Colon typeDeclaration SemiColon;

boAction:
	comments? annotations? ACTION identifier raiseMessage? SemiColon;

message:
	comments? annotations? MESSAGE identifier TEXT StringLiteral Colon typeList SemiColon;

node:
	comments? annotations? NODE identifier multiplicity? raiseMessage? block? SemiColon?;

association:
	comments? annotations? ASSOCIATION identifier multiplicity? TO memberExpression
		associationUsingDefinition? valuationDefinition? SemiColon;

associationUsingDefinition: USING identifier;

valuationDefinition:
	VALUATION OpenParen valutaionExpressionList CloseParen;

valutaionExpressionList:
	valutaionExpression
	| valutaionExpression logicOperator valutaionExpressionList;

valutaionExpression: identifier compareOperator literal;

raiseMessage: RAISES? identifierList?;

annotations: annotation*;

annotation:
	CustomAnnotationStart? OpenBracket identifier CloseBracket
	| CustomAnnotationStart? OpenBracket identifier OpenParen identifier CloseParen CloseBracket
	| CustomAnnotationStart? OpenBracket identifier OpenParen literal CloseParen CloseBracket;

typeList: typeDeclaration | typeDeclaration Comma typeList;

typeDeclaration:
	identifier typeDefaultValue?
	| memberExpression Colon identifier typeDefaultValue?;

typeDefaultValue:
	Assign literal
	| Assign OpenBrace valueAssignList CloseBrace;

valueAssignList:
	valueAssign
	| valueAssign Comma valueAssignList;

valueAssign: identifier Assign literal;

multiplicity:
	OpenBracket literal Comma literal CloseBracket
	| OpenBracket literal Comma N CloseBracket;

memberExpression: identifier ( Dot memberExpression)*;

identifierList: identifier (Comma identifier)*;

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

identifier: Identifier;
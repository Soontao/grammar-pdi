parser grammar ABSLParser;

options {
	tokenVocab = ABSLLexer;
}

program: HashBangLine? sourceElements? EOF;

sourceElement: statement;

statement:
	block
	| variableStatement
	| importStatement
	| exportStatement
	| programEmptyStatement
	| classDeclaration
	| expressionStatement
	| ifStatement
	| iterationStatement
	| continueStatement
	| breakStatement
	| returnStatement
	| yieldStatement
	| withStatement
	| labelledStatement
	| switchStatement
	| throwStatement
	| tryStatement
	| debuggerStatement
	| functionDeclaration;

block: '{' statementList? '}';

statementList: statement+;

importStatement:
	Import importFromBlock
	| Import singleExpression As? Identifier?;

importFromBlock:
	importDefault? (importNamespace | moduleItems) importFrom eos
	| StringLiteral eos;

moduleItems: '{' (aliasName ',')* (aliasName ','?)? '}';

importDefault: aliasName ',';

importNamespace: '*' (As identifierName)?;

importFrom: From StringLiteral;

aliasName: identifierName (As identifierName)?;

exportStatement:
	Export (exportFromBlock | declaration) eos	# ExportDeclaration
	| Export Default singleExpression eos		# ExportDefaultDeclaration;

exportFromBlock:
	importNamespace importFrom eos
	| moduleItems importFrom? eos;

declaration:
	variableStatement
	| classDeclaration
	| functionDeclaration;

variableStatement: varModifier variableDeclarationList eos;

variableDeclarationList:
	variableDeclaration (',' variableDeclaration)*;

variableDeclaration:
	assignable (':' variableType)? ('=' singleExpression)?; // ECMAScript 6: Array & Object Matching

variableType: CollectionOf? ElementOf? singleExpression;

programEmptyStatement: SemiColon;

expressionStatement: expressionSequence eos;

ifStatement:
	If '(' expressionSequence ')' statement (Else statement)?;

iterationStatement:
	Do statement While '(' expressionSequence ')' eos	# DoStatement
	| While '(' expressionSequence ')' statement		# WhileStatement
	| For '(' (expressionSequence | variableStatement)? ';' expressionSequence? ';'
		expressionSequence? ')' statement													# ForStatement
	| For '(' (singleExpression | variableStatement) In expressionSequence ')' statement	#
		ForInStatement
	// strange, 'of' is an identifier. and this.p("of") not work in sometime.
	| For Await? '(' (singleExpression | variableStatement) Identifier expressionSequence ')'
		statement														# ForOfStatement
	| Foreach '(' Var Identifier 'in' singleExpression ')' statement	# ForeachStatement;

varModifier: Var | Let | Const; // let, const - ECMAScript 6

continueStatement: Continue ( Identifier)? eos;

breakStatement: Break ( Identifier)? eos;

returnStatement: Return ( expressionSequence)? eos;

yieldStatement: Yield ( expressionSequence)? eos;

withStatement: With '(' expressionSequence ')' statement;

switchStatement: Switch '(' expressionSequence ')' caseBlock;

caseBlock: '{' caseClauses? (defaultClause caseClauses?)? '}';

caseClauses: caseClause+;

caseClause:
	Case expressionSequence ':' statementList?
	| Case '('? singleExpression ')'? statementList?;

defaultClause: Default ':'? statementList?;

labelledStatement: Identifier ':' statement;

throwStatement:
	Throw expressionSequence eos
	| Raise expressionSequence eos;

tryStatement:
	Try block (
		catchProduction finallyProduction?
		| finallyProduction
	);

catchProduction: Catch ('(' assignable? ')')? block;

finallyProduction: Finally block;

debuggerStatement: Debugger eos;

functionDeclaration:
	Async? Function '*'? Identifier '(' formalParameterList? ')' '{' functionBody '}';

classDeclaration: Class Identifier classTail;

classTail: (Extends singleExpression)? '{' classElement* '}';

classElement: (Static | Identifier | Async)* methodDefinition
	| programEmptyStatement
	| '#'? propertyName '=' singleExpression;

methodDefinition:
	'*'? '#'? propertyName '(' formalParameterList? ')' '{' functionBody '}'
	| '*'? '#'? getter '(' ')' '{' functionBody '}'
	| '*'? '#'? setter '(' formalParameterList? ')' '{' functionBody '}';

formalParameterList:
	formalParameterArg (',' formalParameterArg)* (
		',' lastFormalParameterArg
	)?
	| lastFormalParameterArg;

formalParameterArg:
	assignable ('=' singleExpression)?; // ECMAScript 6: Initialization

lastFormalParameterArg: // ECMAScript 6: Rest Parameter
	Ellipsis singleExpression;

functionBody: sourceElements?;

sourceElements: sourceElement+;

arrayLiteral: ('[' elementList ']');

elementList:
	','* arrayElement? (','+ arrayElement)* ','*; // Yes, everything is optional

arrayElement: Ellipsis? singleExpression;

objectLiteral:
	'{' (propertyAssignment (',' propertyAssignment)*)? ','? '}';

propertyAssignment:
	propertyName ':' singleExpression												# PropertyExpressionAssignment
	| '[' singleExpression ']' ':' singleExpression									# ComputedPropertyExpressionAssignment
	| Async? '*'? propertyName '(' formalParameterList? ')' '{' functionBody '}'	# FunctionProperty
	| getter '(' ')' '{' functionBody '}'											# PropertyGetter
	| setter '(' formalParameterArg ')' '{' functionBody '}'						# PropertySetter
	| Ellipsis? singleExpression													# PropertyShorthand;

propertyName:
	identifierName
	| StringLiteral
	| numericLiteral
	| '[' singleExpression ']';

arguments: '(' (argument (',' argument)* ','?)? ')';

argument: Ellipsis? (singleExpression | Identifier);

expressionSequence: singleExpression (',' singleExpression)*;

singleExpression:
	anoymousFunction														# FunctionExpression
	| Class Identifier? classTail											# ClassExpression
	| singleExpression '[' expressionSequence ']'							# MemberIndexExpression
	| singleExpression '?'? '.' '#'? identifierName							# MemberDotExpression
	| singleExpression arguments											# ArgumentsExpression
	| New singleExpression arguments?										# NewExpression
	| New '.' Identifier													# MetaExpression // new.target
	| singleExpression '++'													# PostIncrementExpression
	| singleExpression '--'													# PostDecreaseExpression
	| Delete singleExpression												# DeleteExpression
	| Void singleExpression													# VoidExpression
	| Typeof singleExpression												# TypeofExpression
	| '++' singleExpression													# PreIncrementExpression
	| '--' singleExpression													# PreDecreaseExpression
	| '+' singleExpression													# UnaryPlusExpression
	| '-' singleExpression													# UnaryMinusExpression
	| '~' singleExpression													# BitNotExpression
	| '!' singleExpression													# NotExpression
	| Await singleExpression												# AwaitExpression
	| <assoc = right> singleExpression '**' singleExpression				# PowerExpression
	| singleExpression ('*' | '/' | '%') singleExpression					# MultiplicativeExpression
	| singleExpression ('+' | '-') singleExpression							# AdditiveExpression
	| singleExpression '??' singleExpression								# CoalesceExpression
	| singleExpression ('<<' | '>>' | '>>>') singleExpression				# BitShiftExpression
	| singleExpression ('<' | '>' | '<=' | '>=') singleExpression			# RelationalExpression
	| singleExpression Instanceof singleExpression							# InstanceofExpression
	| singleExpression In singleExpression									# InExpression
	| singleExpression ('==' | '!=' | '===' | '!==') singleExpression		# EqualityExpression
	| singleExpression '&' singleExpression									# BitAndExpression
	| singleExpression '^' singleExpression									# BitXOrExpression
	| singleExpression '|' singleExpression									# BitOrExpression
	| singleExpression '&&' singleExpression								# LogicalAndExpression
	| singleExpression '||' singleExpression								# LogicalOrExpression
	| singleExpression '?' singleExpression ':' singleExpression			# TernaryExpression
	| <assoc = right> singleExpression '=' singleExpression					# AssignmentExpression
	| <assoc = right> singleExpression assignmentOperator singleExpression	#
		AssignmentOperatorExpression
	| Import '(' singleExpression ')'			# ImportExpression
	| singleExpression TemplateStringLiteral	# TemplateStringExpression // ECMAScript 6
	| yieldStatement							# YieldExpression // ECMAScript 6
	| This										# ThisExpression
	| Identifier								# IdentifierExpression
	| Super										# SuperExpression
	| literal									# LiteralExpression
	| arrayLiteral								# ArrayLiteralExpression
	| objectLiteral								# ObjectLiteralExpression
	| '(' expressionSequence ')'				# ParenthesizedExpression;

assignable: Identifier | arrayLiteral | objectLiteral;

anoymousFunction:
	functionDeclaration															# FunctionDecl
	| Async? Function '*'? '(' formalParameterList? ')' '{' functionBody '}'	# AnoymousFunctionDecl
	| Async? arrowFunctionParameters '=>' arrowFunctionBody						# ArrowFunction;

arrowFunctionParameters:
	Identifier
	| '(' formalParameterList? ')';

arrowFunctionBody: singleExpression | '{' functionBody '}';

assignmentOperator:
	'*='
	| '/='
	| '%='
	| '+='
	| '-='
	| '<<='
	| '>>='
	| '>>>='
	| '&='
	| '^='
	| '|='
	| '**=';

literal:
	NullLiteral
	| BooleanLiteral
	| StringLiteral
	| TemplateStringLiteral
	| RegularExpressionLiteral
	| numericLiteral
	| bigintLiteral;

numericLiteral:
	DecimalLiteral
	| HexIntegerLiteral
	| OctalIntegerLiteral
	| OctalIntegerLiteral2
	| BinaryIntegerLiteral;

bigintLiteral:
	BigDecimalIntegerLiteral
	| BigHexIntegerLiteral
	| BigOctalIntegerLiteral
	| BigBinaryIntegerLiteral;

identifierName: Identifier | reservedWord;

reservedWord: keyword | NullLiteral | BooleanLiteral;

keyword:
	Break
	| Do
	| Instanceof
	| Typeof
	| Case
	| Else
	| New
	| Var
	| Catch
	| Finally
	| Return
	| Void
	| Continue
	| For
	| Switch
	| While
	| Debugger
	| Function
	| This
	| With
	| Default
	| If
	| Throw
	| Delete
	| In
	| Try
	| Class
	| Enum
	| Extends
	| Super
	| Const
	| Export
	| Import
	| Implements
	| Let
	| Private
	| Public
	| Interface
	| Package
	| Protected
	| Static
	| Yield
	| Async
	| Await
	| From
	| As;

getter: Identifier propertyName;

setter: Identifier propertyName;

eos: SemiColon | EOF;
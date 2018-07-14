grammar CMINUS ;

// Palavras Reservadas
IF : 'if' ;
ELSE : 'else' ;
INT : 'int' ;
VOID : 'void' ;
RETURN : 'return' ;
WHILE : 'while' ;
ASSIGN : '=' ;
EQUAL : '==' ;
DIFFERENT : '!=' ;
SMALLER : '<' ;
LARGER : '>' ;
SMALLERANDEQUAL : '<=' ;
LARGERANDEQUAL : '>=' ;
PLUS : '+' ;
MINUS : '-' ;
TIMES : '*' ;
DIVISION : '/' ;
LBRACKET : '[' ;
RBRACKET : ']' ;
LPARENT : '(' ;
RPARENT : ')' ;
SEMICOLON : ';' ;
COMMA : ',' ;
LKEY : '{' ;
RKEY : '}' ;

// Lexer Rules
NUMBER : [0-9]+ ;
IDENTIFIER : [a-zA-Z][a-zA-Z0-9]* ;
WHITESPACE : [ \t\r\n]+ -> skip ;
COMMENT     : '/*' .*? '*/' -> skip ; 

//Parser Rules
program : declarationList;

declarationList : 
	declarationList declaration 
	| 
	declaration
	;

declaration :
	varDeclaration 
	| 
	funDeclaration
	;

varDeclaration : 
	INT ident SEMICOLON 
	| 
	INT ident LBRACKET NUMBER RBRACKET SEMICOLON
	;

funDeclaration :
	INT ident LPARENT params RPARENT compoundDecl
	|
	VOID ident LPARENT params RPARENT compoundDecl
	;

params : 
	paramList
	|
	VOID
	;

paramList : 
	paramList COMMA param 
	| 
	param 
	;

param : 
	INT ident 
	| 
	INT ident LBRACKET RBRACKET 
	;

compoundDecl : 
	LKEY localDeclarations statementList RKEY 
	|
	LKEY localDeclarations RKEY 
	|
	LKEY statementList RKEY
	|
	LKEY RKEY
	;

localDeclarations : 
	localDeclarations varDeclaration 
	| 
	varDeclaration
	;

statementList : 
	statementList statement 
	| 
	statement 
	;

statement :
	expressionDecl
	|
	compoundDecl
	|
	selectionDecl
	|
	iteratorDecl
	|
	returnDecl
	;

expressionDecl : 
	expression SEMICOLON 
	| 
	SEMICOLON 
	;

selectionDecl : 
	IF LPARENT expression RPARENT statement
	|
	IF LPARENT expression RPARENT statement ELSE statement
	;

iteratorDecl : WHILE LPARENT expression RPARENT statement ;

returnDecl : 
	RETURN SEMICOLON 
	| 
	RETURN expression SEMICOLON 
	;

expression : 
	vari ASSIGN expression 
	| 
	simpleExpression 
	;

vari : 
	ident 
	| 
	ident LBRACKET expression RBRACKET 
	;

simpleExpression : 
	sumExpression relational sumExpression
	|
	sumExpression
	;

relational : 
	EQUAL 
	| 
	DIFFERENT 
	| 
	SMALLER 
	| 
	SMALLERANDEQUAL 
	| 
	LARGER 
	| 
	LARGERANDEQUAL 
	;

sumExpression : 
	sumExpression sum term 
	| 
	term 
	;

sum : 
	PLUS 
	| 
	MINUS 
	;

term : 
	term mult factor 
	| 
	factor 
	;

mult : 
	TIMES 
	| 
	DIVISION 
	;

factor : 
	LPARENT expression RPARENT 
	| 
	vari 
	| 
	activation 
	| 
	num 
	;

activation : 
	ident LPARENT argList RPARENT 
	| 
	ident LPARENT RPARENT
	;

argList : 
	argList COMMA expression 
	| 
	expression 
	;

ident : IDENTIFIER ;

num : NUMBER ;


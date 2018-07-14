package ast

type Node struct{
	line int
}

type Program struct {
	Node
	*DeclarationList 
}

type DeclarationList struct {
	Node
	*Declaration
	*DeclarationList
}

type Declaration struct {
	Node 
	*VarDeclaration
	*FunDeclaration
}

type VarDeclaration struct {
	Node
	int_ string
	semicolon_ string
	lbracket_ string
	rbracket_ string
	number_ string
	*Ident
}

type FunDeclaration struct {
	Node
	int_ string
	void_ string
	lparent_ string
	rparent_ string
	*Ident
	*Params
	*CompoundDecl 
}

type Params struct {
	Node
	void_ string
	*ParamList
}

type ParamList struct {
	Node 
	comma_ string
	*Param 
	*ParamList
}

type Param struct {
	Node 
	int_ string 
	lbracket_ string 
	rbracket_ string
	*Ident
}

type CompoundDecl struct {
	Node 
	lkey_ string
	rkey_ string
	*LocalDeclarations
	*StatementList
}

type LocalDeclarations struct {
	Node
	*LocalDeclarations
	*VarDeclaration
}

type StatementList struct {
	Node 
	*StatementList
	*Statement 
}

type Statement struct {
	Node
	*ExpressionDecl
	*CompoundDecl
	*SelectionDecl
	*IteratorDecl
	*ReturnDecl
}

type ExpressionDecl struct {
	Node
	*Expression
	semicolon_ string
}

type SelectionDecl struct {
	Node 
	if_ string
	else_ string
	lparent_ string 
	rparent_ string 
	*Expression
	statement0_ *Statement
	statement1_ *Statement
}

type IteratorDecl struct {
	Node
	while_ string 
	lparent_ string 
	rparent_ string 
	*Expression
	*Statement
}

type ReturnDecl struct {
	Node
	return_ string
	semicolon_ string 
	*Expression	
}

type Expression struct {
	Node
	assign_ string
	*Vari 
	*Expression
	*SimpleExpression
}

type Vari struct {
	Node
	lbracket_ string
 	rbracket_ string
	*Ident
	*Expression
}

type SimpleExpression struct {
	Node
	sumexpression0_ *SumExpression
	sumexpression1_ *SumExpression
	*Relational
}

type Relational struct {
	Node 
	relation_ string 
}

type SumExpression struct {
	Node
	*SumExpression
	*Sum
	*Term
}

type Sum struct {
	Node
	operation_ string
}

type Term struct {
	Node 
	*Term
	*Mult
	*Factor
}

type Mult struct {
	Node 
	operation_ string
}

type Factor struct {
	Node 
	lparent_ string
	rparent_ string 
	*Expression
	*Vari
	*Activation
	*Num 
}

type Activation struct {
	Node 
	lparent_ string
	rparent_ string 
	*Ident
	*ArgList
}

type ArgList struct {
	Node 
	comma_ string 
	*ArgList
	*Expression
}

type Ident struct {
	Node 
	ID_ string
}

type Num struct {
	Node 
	NUMBER_ string
}

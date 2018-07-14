package semantic

import(
	"../ast"
)

type Symbol struct {
	ID string
	line int 
	typeID string
	scope string 
	dataType string
}
var table map[int]Symbol 
var contLines int

func SymbolTableBuilder (self *Program) {
	table = make(map[int]Symbol)
	contLines = 0
	SymbolProgram(self)
}

func SymbolProgram (self *ast.Program) {
	if(self == nil) {
		return
	}

	SymbolDeclList(self.DeclarationList)
}

func SymbolDeclList (self *DeclarationList) {
	if(self == nil) {
		return
	}

	SymbolDeclList(self.DeclarationList)
	SymbolDeclaration(self.Declaration)
}

func SymbolDeclaration (self *Declaration) {
	if(self == nil) {
		return
	}

	SymbolVarDeclaration(self.VarDeclaration)
	SymbolFunDeclaration(self.FunDeclaration)
}

func SymbolVarDeclaration (self *VarDeclaration) {
	if(self == nil) {
		return
	}
	t := self.int_
	id, line := SymbolIdent(self.Ident)
	symbol_ := Symbol{
		ID : id,
		line : line,
		typeID : "variable",
		scope : "",
		dataType: t}
	table[contLines] = symbol_
	contLines++
}

func SymbolFunDeclaration (self *FunDeclaration) {
	if(self == nil) {
		return
	}

	var funType string

	if(self.int_ == "") {
		funType = "void"
	} else {
		funType = "int"
	}

	id, line := SymbolIdent(self.Ident)
	symbol_ := Symbol{
		ID : id,
		line : line,
		typeID : "function",
		scope : "",
		dataType: funType}
	table[contLines] = symbol_
	contLines++


	SymbolParams(self.Params, id)
	SymbolCompoundDecl(self.CompoundDecl, id)
}

func SymbolParams (self *Params, idFun string) {
	if(self == nil) {
		return 
	}

	SymbolParamList(self.ParamList, idFun)
}

func SymbolParamList (self *ParamList, idFun string) {
	if(self == nil) {
		return
	}

	SymbolParam(self.Param, idFun)
	SymbolParamList(self.ParamList, idFun)
}

func SymbolParam (self *Param, idFun string) {
	if(self == nil) {
		return
	}

	id, line := SymbolIdent(self.Ident)
	symbol_ := Symbol{
		ID : id,
		line : line,
		typeID : "variable",
		scope : idFun,
		dataType: "int"}
	table[contLines] = symbol_
	contLines++
}

func SymbolCompoundDecl (self *CompoundDecl, idFun string) {
	if(self == nil) {
		return
	}

	SymbolLocalDeclarations(self.LocalDeclarations, idFun)
	SymbolStatementList(self.StatementList, idFun)
}

func SymbolLocalDeclarations (self *LocalDeclarations, idFun string) {
	if(self == nil) {
		return
	}

	SymbolLocalDeclarations(self.LocalDeclarations, idFun)
	SymbolVarDeclaration(self.VarDeclaration, idFun)
}

func SymbolStatementList (self *StatementList, idFun string) {
	if(self == nil) {
		return
	}

	SymbolStatementList(self.StatementList, idFun)	
	SymbolStatement(self.Statement, idFun)
}

func SymbolStatement (self *Statement, idFun string) {
	if(self == nil) {
		return
	}

	SymbolExpressionDecl(self.ExpressionDecl, idFun)
	SymbolCompoundDecl(self.CompoundDecl, idFun)
	SymbolSelectionDecl(self.SelectionDecl, idFun)
	SymbolIteratorDecl(self.IteratorDecl, idFun)
	SymbolReturnDecl(self.ReturnDecl, idFun)
}

func SymbolExpressionDecl (self *ExpressionDecl, idFun string) {
	if(self == nil) {
		return
	}

	SymbolExpression(self.Expression, idFun)
}

func SymbolSelectionDecl (self *SelectionDecl, idFun string) {
	if(self == nil) {
		return
	}

	SymbolExpression(self.Expression, idFun)
	SymbolStatement(self.statement0_, idFun)
	SymbolStatement(self.statement1_, idFun)
}

func SymbolIteratorDecl (self *IteratorDecl, idFun string) {
	if(self == nil) {
		return
	}

	SymbolExpression(self.Expression, idFun)
	SymbolStatement(self.Statement, idFun)
}

func SymbolReturnDecl (self *ReturnDecl, idFun string) {
	if(self == nil) {
		return
	}

	SymbolExpression(self.Expression, idFun)
}

func SymbolExpression (self *Expression, idFun string) {
	if(self == nil) {
		return
	}

	SymbolVari(self.Vari, idFun)
	SymbolExpression(self.Expression, idFun)
	SymbolSimpleExpression(self.SimpleExpression, idFun)
}

func SymbolVari (self *Vari, idFun string) {
	if(self == nil) {
		return
	}

	id, line := SymbolIdent(self.Ident)
	symbol_ := Symbol{
		ID : id,
		line : line,
		typeID : "variable",
		scope : idFun,
		dataType: ""}
	table[contLines] = symbol_
	contLines++
}

func SymbolSimpleExpression (self *SimpleExpression, idFun string) {
	if(self == nil) {
		return
	}

	SymbolSimpleExpression(self.sumexpression0_, idFun)
	SymbolSimpleExpression(self.sumexpression1_, idFun)
	//SymbolRelational(self.Relational, idFun)
}
/*
func SymbolRelational (self *Relatinal, idFun string) {
*/

func SymbolSumExpression (self *SumExpression, idFun string) {
	if(self == nil) {
		return
	}

	SymbolSumExpression(self.SumExpression, idFun)
	//SymbolSum(self.Sum, idFun)
	SymbolTerm(self.Term, idFun)
}

//func SymbolSum()

func SymbolTerm (self *Term, idFun string) {
	if(self == nil) {
		return
	}

	SymbolTerm(self.Term, idFun)
	//SymbolMult
	SymbolFactor(self.Factor, idFun)
}

//func SymbolMult()

func SymbolFactor (self *Factor, idFun string) {
	if(self == nil) {
		return
	}

	SymbolExpression(self.Expression, idFun)
	SymbolVari(self.Vari, idFun)
	//SymbolNum
	SymbolActivation(self.Activation, idFun)
}

func SymbolActivation (self *Activation, idFun string) {
	if(self == nil) {
		return
	}

	id, line := SymbolIdent(self.Ident)
	symbol_ := Symbol{
		ID : id,
		line : line,
		typeID : "variable",
		scope : idFun,
		dataType: ""}
	table[contLines] = symbol_
	contLines++

	SymbolArgList(self.ArgList, idFun)
}

func SymbolArgList (self *ArgList, idFun string) {
	if(self == nil) {
		return
	}

	SymbolArgList(self.ArgList, idFun)
	SymbolExpression(self.Expression, idFun)
}

func SymbolIdent (self *Ident) (string, int){
	if(self == nil) {
		return ""
	}
	node := self.Node 
	return self.ID_, node.line
}

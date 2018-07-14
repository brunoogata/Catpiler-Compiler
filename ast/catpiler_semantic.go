package ast

import (
	"fmt"
	"os"
	"strconv"
)

var erro bool

type Symbol struct {
	ID string
	line []int 
	typeID string
	scope string 
	dataType string
	vector string
	spacemem int
}
var table map[int]Symbol 
var contLines int

func SymbolTableBuilder (self *Program) {
	table = make(map[int]Symbol)
	contLines = 0
	erro = false
	SymbolProgram(self)
	/*
	Rule1()
	Rule4()

	*/
	//Rule7()
	Rule6()
	PrintSymbolTable()

	if(erro){
		os.Exit(1)
	}
}

func SymbolProgram (self *Program) {
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

	SymbolVarDeclaration(self.VarDeclaration, "")
	SymbolFunDeclaration(self.FunDeclaration)
}

func SymbolVarDeclaration (self *VarDeclaration, idFun string) {
	if(self == nil) {
		return
	}
	t := self.int_

	id, line_ := SymbolIdent(self.Ident)

	for _, symbol2_ := range table {
		if(id == symbol2_.ID){
			if(idFun == symbol2_.scope || symbol2_.scope == "global"){
				erro = true
				fmt.Println("ERRO: Declaracao invalida - ", id, "- nome de variavel ja declarado. Linha: ", line_)
				break
			} else if (symbol2_.scope == "") {
				erro = true
				fmt.Println("ERRO: Declaracao invalida - ", id, "- nome de variavel ja declarado como funcao. Linha: ", line_)
				break
			}

		}
	}

	symbol_ := Symbol{
		ID : id,
		typeID : "variable",
		scope : idFun,
		dataType: t}

	if(self.lbracket_ == "[" && self.rbracket_ == "]") {
		symbol_.vector = "yes"
		vl, _ := strconv.Atoi(self.number_)
		symbol_.spacemem = vl
	}

	symbol_.line = append(symbol_.line,line_)
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

	id, line_ := SymbolIdent(self.Ident)
	symbol_ := Symbol{
		ID : id,
		typeID : "function",
		scope : "",
		dataType: funType}

	symbol_.line = append(symbol_.line, line_)
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

	id, line_ := SymbolIdent(self.Ident)
	symbol_ := Symbol{
		ID : id,
		typeID : "variable",
		scope : idFun,
		dataType: "int"}

	symbol_.line = append(symbol_.line,line_)
	if(self.lbracket_ == "[" && self.rbracket_ == "]") {
		symbol_.vector = "yes"
	}
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
	flag := true
	id, line := SymbolIdent(self.Ident)

	for key, symbol2_ := range table {
		if(id == symbol2_.ID){
			if(symbol2_.scope == idFun || symbol2_.scope == "") {
				symbol2_.line = append(symbol2_.line, line)
				table[key] = symbol2_
				flag = false
			}
		}
	}

	if(flag){
		erro = true
		fmt.Println("ERRO: Variavel ", id, " nao declarada. Linha: ", line)
	}
	SymbolExpression(self.Expression, idFun)
	/*	
	symbol_ := Symbol{
		ID : id,
		line : line,
		typeID : "variable",
		scope : idFun,
		dataType: ""}
	table[contLines] = symbol_
	contLines++
	*/
}

func SymbolSimpleExpression (self *SimpleExpression, idFun string) {
	if(self == nil) {
		return
	}

	SymbolSumExpression(self.sumexpression0_, idFun)
	SymbolSumExpression(self.sumexpression1_, idFun)
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

	flag := true

	id, line_ := SymbolIdent(self.Ident)
	
	for key, symbol2_ := range table {
		if(id == symbol2_.ID){
			if(symbol2_.typeID == "variable") {
				fmt.Println("ERRO: Declaracao invalida - ", id, "- nome de variavel ja declarado. Linha: ", line_)
				erro = true
			} else {
				symbol2_.line = append(symbol2_.line, line_)
				table[key] = symbol2_
				flag = false
			}
		}
	}

	if(flag){
		fmt.Println("ERRO: Funcao ", id, " nao declarada. Linha: ", line_)
	}

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
		return "", 0
	}
	node := self.Node 
	return self.ID_, node.line
}
/*
func AnalyzeSymbolTable() bool{
	if(Rule1() || Rule2() || Rule3() || Rule4() || Rule5() || Rule6() || Rule7()){
		return true
	} else {
		return false
	}
}
*/
/*
func Rule1 () bool {
	erro := false 
	var aux Symbol 



	for key, symbol_ := range table {
		if(symbol_.dataType == "") {
			for key2, symbol2_ := range table {
				if(key != key2){
					if(symbol2_.ID == symbol_.ID){
						if(symbol2_.scope == symbol_.scope || symbol2_.scope == "global"){
							if(symbol2_.dataType != ""){
								symbol_.dataType = "DECLARED"
								aux = symbol_
								table[key] = aux
								break
							}
						}
					}
				}
			}
			if(symbol_.dataType == ""){
				fmt.Println("ERRO: Variavel ", symbol_.ID, " nao declarada. Linha: ", symbol_.line)
				symbol_.dataType = "NON DECLARED"
				aux = symbol_
				table[key] = aux
				erro = true 
			}
		}
	}
	return erro
}
*/
func Rule2 () bool {
	return true
}
/*
func Rule3 () bool {
	var erro bool
	for _, symbol_ := range table {
		if (symbol_.dataType == "void"){
			if(symbol_.typeID == "variable"){
				fmt.Println("ERRO: Declaracao invalida de variavel - Variavel nao pode assumir tipo VOID. Linha: ", symbol_.line)
				erro = true
			}
		}
	}
	return erro
}
*/

func Rule4 () bool {
	erro := false
	for key, symbol_ := range table {
		for key2, symbol2_ := range table {
			if(key != key2){
				if(symbol2_.ID == symbol_.ID && symbol2_.dataType != "DECLARED" && symbol_.dataType != "DECLARED" && symbol_.scope == symbol2_.scope) {
					fmt.Println("ERRO: ", symbol2_.typeID, " com nome igual ao de ", symbol_.typeID, ": ", symbol_.ID, "Linha: ", symbol2_.line)
					erro = true
				}
			}
		}
	}
	return erro
}
/*
func Rule5 () bool {

}
*/
func Rule6 () bool {
	erro := true
	for _, symbol_ := range table {
		if(symbol_.ID == "main" && symbol_.typeID == "function" && symbol_.dataType != ""){
			erro = false
		}
	}
	if(erro){
		fmt.Println("ERRO: Funcao main nao declarada")
	}
	return erro
}
/*
func Rule7 () bool {

}
*/
func PrintSymbolTable() {
	fmt.Println("\n|--------------------------------Catpiler SymbolTable----------------------------------|")
	fmt.Println("|Num|    \t|ID|    \t|Type|    \t|Scope|    \t|DataType|\t|Line|")
	for key, symbol_ := range table {
		PrintLineSymbolTable(key, symbol_)
	}
	fmt.Println("|--------------------------------------------------------------------------------------|\n")
}

func PrintLineSymbolTable(num int, self Symbol){
	var i int
	fmt.Print(num, "\t\t")
	fmt.Print(self.ID, "\t\t")
	fmt.Print(self.typeID, "\t")
	fmt.Print(self.scope, "\t\t")
	fmt.Print(self.dataType, "\t\t")
	for i = 0; i < len(self.line); i++ {
		fmt.Print(self.line[i], ",")
	}
	fmt.Println()
}

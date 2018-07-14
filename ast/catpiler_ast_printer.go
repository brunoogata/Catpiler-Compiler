package ast

import "fmt"

func (ast *AstVisitor) PrintAstCatpiler (prNode *Program) {
	fmt.Println("\n|-------------------------------------Catpiler AST-------------------------------------|")
	ast.PrintProgram(prNode)
	fmt.Println("|--------------------------------------------------------------------------------------|\n")

}

func (ast *AstVisitor) PrintProgram (prNode *Program) {
	if(prNode == nil){
		return
	}

	cont_tab := 0
	fmt.Println("| Program")
	cont_tab++

	PrintDeclarationList(prNode.DeclarationList, cont_tab)
}

func PrintDeclarationList (declListNode *DeclarationList, cont_tab int) {
	if(declListNode == nil){
		return
	}

	PrintTab(cont_tab)
	fmt.Println("| DeclarationList")
	cont_tab++

	PrintDeclarationList(declListNode.DeclarationList, cont_tab)
	PrintDeclaration(declListNode.Declaration, cont_tab)
}

func PrintDeclaration (declNode *Declaration, cont_tab int) {
	if(declNode == nil){
		return
	}

	PrintTab(cont_tab)
	fmt.Println("| Declaration")
	cont_tab++

	PrintVarDeclaration(declNode.VarDeclaration, cont_tab)
	PrintFunDeclaration(declNode.FunDeclaration, cont_tab)
}

func PrintVarDeclaration (varDeclNode *VarDeclaration, cont_tab int) {
	if(varDeclNode == nil){
		return
	}

	PrintTab(cont_tab)
	fmt.Println("| VarDeclaration")
	cont_tab++
	PrintTab(cont_tab)
	fmt.Println("| TYPE: ", varDeclNode.int_)

	PrintIdent(varDeclNode.Ident, cont_tab)
}

func PrintFunDeclaration (funDeclNode *FunDeclaration, cont_tab int) {
	if(funDeclNode == nil){
		return
	}

	var funType string
	if(funDeclNode.int_ == ""){
		funType = "void"
	} else {
		funType = "int"
	}

	PrintTab(cont_tab)
	fmt.Println("| FunDeclaration")
	cont_tab++
	PrintTab(cont_tab)
	fmt.Println("| TYPE: ", funType)

	PrintIdent(funDeclNode.Ident, cont_tab)
	PrintParams(funDeclNode.Params, cont_tab)
	PrintCompoundDecl(funDeclNode.CompoundDecl, cont_tab)
}

func PrintParams (paramsNode *Params, cont_tab int) {
	if(paramsNode == nil){
		return
	}

	PrintTab(cont_tab)
	fmt.Println("| Params")
	cont_tab++

	PrintParamList(paramsNode.ParamList, cont_tab)
}

func PrintParamList (paramListNode *ParamList, cont_tab int) {
	if(paramListNode == nil){
		return
	}

	PrintTab(cont_tab)
	fmt.Println("| ParamList")
	cont_tab++

	PrintParam(paramListNode.Param, cont_tab)
	PrintParamList(paramListNode.ParamList, cont_tab)
}

func PrintParam (paramNode *Param, cont_tab int) {
	if(paramNode == nil){
		return
	}

	PrintTab(cont_tab)
	fmt.Println("| Param")
	cont_tab++
	PrintTab(cont_tab)
	fmt.Println("| TYPE: ", paramNode.int_)

	PrintIdent(paramNode.Ident, cont_tab)
}

func PrintCompoundDecl (compDeclNode *CompoundDecl, cont_tab int) {
	if(compDeclNode == nil){
		return
	}

	PrintTab(cont_tab)
	fmt.Println("| CompoundDecl")
	cont_tab++

	PrintLocalDeclarations(compDeclNode.LocalDeclarations, cont_tab)
	PrintStatementList(compDeclNode.StatementList, cont_tab)	
}

func PrintLocalDeclarations (localDeclNode *LocalDeclarations, cont_tab int) {
	if(localDeclNode == nil){
		return
	}

	PrintTab(cont_tab)
	fmt.Println("| LocalDeclarations")
	cont_tab++

	PrintLocalDeclarations(localDeclNode.LocalDeclarations, cont_tab)
	PrintVarDeclaration(localDeclNode.VarDeclaration, cont_tab)
}

func PrintStatementList (statListNode *StatementList, cont_tab int) {
	if(statListNode == nil){
		return 
	}

	PrintTab(cont_tab)
	fmt.Println("| StatementList")
	cont_tab++

	PrintStatementList(statListNode.StatementList, cont_tab)
	PrintStatement(statListNode.Statement, cont_tab)
}

func PrintStatement (statNode *Statement, cont_tab int) {
	if(statNode == nil){
		return
	}

	PrintTab(cont_tab)
	fmt.Println("| Statement")
	cont_tab++

	PrintExpressionDecl(statNode.ExpressionDecl, cont_tab)
	PrintCompoundDecl(statNode.CompoundDecl, cont_tab)
	PrintSelectionDecl(statNode.SelectionDecl, cont_tab)
	PrintIteratorDecl(statNode.IteratorDecl, cont_tab)
	PrintReturnDecl(statNode.ReturnDecl, cont_tab)
}

func PrintExpressionDecl (expDeclNode *ExpressionDecl, cont_tab int) {
	if(expDeclNode == nil){
		return
	}

	PrintTab(cont_tab)
	fmt.Println("| ExpressionDecl")
	cont_tab++

	PrintExpression(expDeclNode.Expression, cont_tab)
}

func PrintSelectionDecl (selDeclNode *SelectionDecl, cont_tab int) {
	if(selDeclNode == nil){
		return 
	}

	PrintTab(cont_tab)
	fmt.Println("| SelectionDecl")
	cont_tab++

	PrintExpression(selDeclNode.Expression, cont_tab)
	PrintStatement(selDeclNode.statement0_, cont_tab)
	PrintStatement(selDeclNode.statement1_, cont_tab)
}

func PrintIteratorDecl (itDeclNode *IteratorDecl, cont_tab int) {
	if(itDeclNode == nil){
		return 
	}

	PrintTab(cont_tab)
	fmt.Println("| IteratorDecl")
	cont_tab++

	PrintExpression(itDeclNode.Expression, cont_tab)
	PrintStatement(itDeclNode.Statement, cont_tab)
}

func PrintReturnDecl (retDeclNode *ReturnDecl, cont_tab int) {
	if(retDeclNode == nil){
		return 
	}

	PrintTab(cont_tab)
	fmt.Println(" ReturnDecl")
	cont_tab++

	PrintExpression(retDeclNode.Expression, cont_tab)
}

func PrintExpression (expNode *Expression, cont_tab int) {
	if(expNode == nil){
		return 
	}

	PrintTab(cont_tab)
	fmt.Println("| Expression")
	cont_tab++

	PrintVari(expNode.Vari, cont_tab)
	PrintExpression(expNode.Expression, cont_tab)
	PrintSimpleExpression(expNode.SimpleExpression, cont_tab)
}

func PrintVari (variNode *Vari, cont_tab int) {
	if(variNode == nil){
		return
	}

	PrintTab(cont_tab)
	fmt.Println("| Vari")
	cont_tab++

	PrintIdent(variNode.Ident, cont_tab)
	PrintExpression(variNode.Expression, cont_tab)
}

func PrintSimpleExpression (simpExpNode *SimpleExpression, cont_tab int) {
	if(simpExpNode == nil){
		return 
	}

	PrintTab(cont_tab)
	fmt.Println("| SimpleExpression")
	cont_tab++

	PrintSumExpression(simpExpNode.sumexpression0_, cont_tab)
	PrintSumExpression(simpExpNode.sumexpression1_, cont_tab)
	PrintRelational(simpExpNode.Relational, cont_tab)
}

func PrintSumExpression (sumExpNode *SumExpression, cont_tab int) {
	if(sumExpNode == nil){
		return 
	}

	PrintTab(cont_tab)
	fmt.Println("| SumExpression")
	cont_tab++

	PrintSumExpression(sumExpNode.SumExpression, cont_tab)
	PrintSum(sumExpNode.Sum, cont_tab)
	PrintTerm(sumExpNode.Term, cont_tab)
}

func PrintTerm (termNode *Term, cont_tab int) {
	if(termNode == nil){
		return 
	}

	PrintTab(cont_tab)
	fmt.Println("| Term")
	cont_tab++

	PrintTerm(termNode.Term, cont_tab)
	PrintMult(termNode.Mult, cont_tab)
	PrintFactor(termNode.Factor, cont_tab)
}

func PrintFactor (factNode *Factor, cont_tab int) {
	if(factNode == nil){
		return 
	}

	PrintTab(cont_tab)
	fmt.Println("| Factor")
	cont_tab++

	PrintExpression(factNode.Expression, cont_tab)
	PrintVari(factNode.Vari, cont_tab)
	PrintActivation(factNode.Activation, cont_tab)
	PrintNum(factNode.Num, cont_tab)
}

func PrintActivation (actNode *Activation, cont_tab int) {
	if(actNode == nil){
		return
	}

	PrintTab(cont_tab)
	fmt.Println("| Activation")
	cont_tab++

	PrintIdent(actNode.Ident, cont_tab)
	PrintArgList(actNode.ArgList, cont_tab)
}

func PrintArgList (argListNode *ArgList, cont_tab int) {
	if(argListNode == nil){
		return
	}

	PrintTab(cont_tab)
	fmt.Println("| ArgList")
	cont_tab++

	PrintArgList(argListNode.ArgList, cont_tab)
	PrintExpression(argListNode.Expression, cont_tab)
}

func PrintRelational (relNode *Relational, cont_tab int) {
	if(relNode == nil){
		return 
	}

	PrintTab(cont_tab)
	fmt.Println("| RELATIONAL: ", relNode.relation_)
}

func PrintSum (sumNode *Sum, cont_tab int) {
	if(sumNode == nil){
		return 
	}

	PrintTab(cont_tab)
	fmt.Println("| SUM: ", sumNode.operation_)
}

func PrintMult (multNode *Mult, cont_tab int) {
	if(multNode == nil){
		return 
	}

	PrintTab(cont_tab)
	fmt.Println("| MULT: ", multNode.operation_)
}

func PrintIdent (identNode *Ident, cont_tab int) {
	if(identNode == nil){
		return 
	}

	PrintTab(cont_tab)
	fmt.Println("| IDENT: ", identNode.ID_)
}

func PrintNum (numNode *Num, cont_tab int) {
	if(numNode == nil){
		return 
	}

	PrintTab(cont_tab)
	fmt.Println("| NUM: ", numNode.NUMBER_)
}

func PrintTab(cont_tab int) {
	for i := 0; i < cont_tab; i++ {
		fmt.Print("--")
	}
}
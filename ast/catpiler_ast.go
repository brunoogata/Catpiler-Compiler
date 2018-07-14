package ast

import (
	"../parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)


type AstVisitor struct {
	parser.CMINUSVisitor
}

func (ast *AstVisitor) VisitProgram(ctx *parser.ProgramContext) *Program{
	if(ctx == nil){
		return nil
	}

	decllist_ctx, _ := ctx.DeclarationList().(*parser.DeclarationListContext)
	return &Program {
		Node: Node{line: ctx.GetStart().GetLine()},
		DeclarationList: ast.VisitDeclarationList(decllist_ctx)}
}
 
func (ast *AstVisitor) VisitDeclarationList(ctx *parser.DeclarationListContext) *DeclarationList{
	if(ctx == nil){
		return nil
	}

	decllist_ctx, _ := ctx.DeclarationList().(*parser.DeclarationListContext)
	decl_ctx, _ := ctx.Declaration().(*parser.DeclarationContext)
	return &DeclarationList {
		Node: Node{ctx.GetStart().GetLine()},
		DeclarationList: ast.VisitDeclarationList(decllist_ctx),
		Declaration: ast.VisitDeclaration(decl_ctx)}
}

func (ast *AstVisitor) VisitDeclaration(ctx *parser.DeclarationContext) *Declaration{
	if(ctx == nil){
		return nil
	}
	
	fundecl_ctx, _ := ctx.FunDeclaration().(*parser.FunDeclarationContext)	
	vardecl_ctx, _ := ctx.VarDeclaration().(*parser.VarDeclarationContext) 
	
	return &Declaration {
		Node: Node{line: ctx.GetStart().GetLine()},
		FunDeclaration: ast.VisitFunDeclaration(fundecl_ctx),
		VarDeclaration: ast.VisitVarDeclaration(vardecl_ctx)}
}

func (ast *AstVisitor) VisitVarDeclaration(ctx *parser.VarDeclarationContext) *VarDeclaration{
	if(ctx == nil){
		return nil
	}
	
	ident_ctx, _ := ctx.Ident().(*parser.IdentContext)
	return &VarDeclaration {
		Node: Node{line: ctx.GetStart().GetLine()},
		int_: GetTokenValue(ctx.INT()),
		semicolon_: GetTokenValue(ctx.SEMICOLON()),
		lbracket_: GetTokenValue(ctx.LBRACKET()),
		rbracket_: GetTokenValue(ctx.RBRACKET()),
		number_: GetTokenValue(ctx.NUMBER()),
		Ident: ast.VisitIdent(ident_ctx)}
}

func (ast *AstVisitor) VisitFunDeclaration(ctx *parser.FunDeclarationContext) *FunDeclaration{
	if(ctx == nil){
		return nil
	}
	
	ident_ctx, _ := ctx.Ident().(*parser.IdentContext)
	params_ctx, _ := ctx.Params().(*parser.ParamsContext)
	compounddecl_ctx, _ := ctx.CompoundDecl().(*parser.CompoundDeclContext)

	return &FunDeclaration {
		Node: Node{line: ctx.GetStart().GetLine()},
		int_: GetTokenValue(ctx.INT()),
		void_: GetTokenValue(ctx.VOID()),
		lparent_: GetTokenValue(ctx.LPARENT()),
		rparent_: GetTokenValue(ctx.RPARENT()),
		Ident: ast.VisitIdent(ident_ctx),
		Params: ast.VisitParams(params_ctx),
		CompoundDecl: ast.VisitCompoundDecl(compounddecl_ctx)}
}

func (ast *AstVisitor) VisitParams(ctx *parser.ParamsContext) *Params{
	if(ctx == nil){
		return nil
	}

	paramlist_ctx, _ := ctx.ParamList().(*parser.ParamListContext)
	return &Params {
		Node: Node{line: ctx.GetStart().GetLine()},
		void_: GetTokenValue(ctx.VOID()),
		ParamList: ast.VisitParamList(paramlist_ctx)}
}

func (ast *AstVisitor) VisitParamList(ctx *parser.ParamListContext) *ParamList{
	if(ctx == nil){
		return nil
	}

	param_ctx, _ := ctx.Param().(*parser.ParamContext)
	paramlist_ctx, _ := ctx.ParamList().(*parser.ParamListContext)
	return &ParamList {
		Node: Node{line: ctx.GetStart().GetLine()},
		comma_: GetTokenValue(ctx.COMMA()),
		Param: ast.VisitParam(param_ctx),
		ParamList: ast.VisitParamList(paramlist_ctx)}
}

func (ast *AstVisitor) VisitParam(ctx *parser.ParamContext) *Param{
	if(ctx == nil){
		return nil
	}

	ident_ctx, _ := ctx.Ident().(*parser.IdentContext)
	return &Param {
		Node: Node{line: ctx.GetStart().GetLine()},
		int_: GetTokenValue(ctx.INT()),
		lbracket_: GetTokenValue(ctx.LBRACKET()),
		rbracket_: GetTokenValue(ctx.RBRACKET()),
		Ident: ast.VisitIdent(ident_ctx)}
}

func (ast *AstVisitor) VisitCompoundDecl(ctx *parser.CompoundDeclContext) *CompoundDecl{
	if(ctx == nil){
		return nil
	}

	localdecl_ctx, _ := ctx.LocalDeclarations().(*parser.LocalDeclarationsContext)
	statementlist_ctx, _ := ctx.StatementList().(*parser.StatementListContext)
	return &CompoundDecl {
		Node: Node{line: ctx.GetStart().GetLine()},
		lkey_: GetTokenValue(ctx.LKEY()),
		rkey_: GetTokenValue(ctx.RKEY()),
		LocalDeclarations: ast.VisitLocalDeclarations(localdecl_ctx),
		StatementList: ast.VisitStatementList(statementlist_ctx)}
}

func (ast *AstVisitor) VisitLocalDeclarations(ctx *parser.LocalDeclarationsContext) *LocalDeclarations{
	if(ctx == nil){
		return nil
	}

	localdecl_ctx, _ := ctx.LocalDeclarations().(*parser.LocalDeclarationsContext)
	vardecl_ctx, _ := ctx.VarDeclaration().(*parser.VarDeclarationContext)
	return &LocalDeclarations {
		Node: Node{line: ctx.GetStart().GetLine()},
		LocalDeclarations: ast.VisitLocalDeclarations(localdecl_ctx),
		VarDeclaration: ast.VisitVarDeclaration(vardecl_ctx)}
}

func (ast *AstVisitor) VisitStatementList(ctx *parser.StatementListContext) *StatementList{
	if(ctx == nil){
		return nil
	}

	statementlist_ctx, _ := ctx.StatementList().(*parser.StatementListContext)
	statement_ctx, _ := ctx.Statement().(*parser.StatementContext)
	return &StatementList {
		Node: Node{line: ctx.GetStart().GetLine()},
		StatementList: ast.VisitStatementList(statementlist_ctx),
		Statement: ast.VisitStatement(statement_ctx)}
}

func (ast *AstVisitor) VisitStatement(ctx *parser.StatementContext) *Statement{
	if(ctx == nil){
		return nil
	}

	expressiondecl_ctx, _ := ctx.ExpressionDecl().(*parser.ExpressionDeclContext)
	compounddecl_ctx, _ := ctx.CompoundDecl().(*parser.CompoundDeclContext)
	selectiondecl_ctx, _ := ctx.SelectionDecl().(*parser.SelectionDeclContext)
	iteratordecl_ctx, _ := ctx.IteratorDecl().(*parser.IteratorDeclContext)
	returndecl_ctx, _ := ctx.ReturnDecl().(*parser.ReturnDeclContext)
	return &Statement {
		Node: Node{line: ctx.GetStart().GetLine()},
		ExpressionDecl: ast.VisitExpressionDecl(expressiondecl_ctx),
		CompoundDecl: ast.VisitCompoundDecl(compounddecl_ctx),
		SelectionDecl: ast.VisitSelectionDecl(selectiondecl_ctx),
		IteratorDecl: ast.VisitIteratorDecl(iteratordecl_ctx),
		ReturnDecl: ast.VisitReturnDecl(returndecl_ctx)}
}

func (ast *AstVisitor) VisitExpressionDecl(ctx *parser.ExpressionDeclContext) *ExpressionDecl{
	if(ctx == nil){
		return nil
	}

	expression_ctx, _ := ctx.Expression().(*parser.ExpressionContext)
	return &ExpressionDecl {
		Node: Node{line: ctx.GetStart().GetLine()},
		Expression: ast.VisitExpression(expression_ctx),
		semicolon_: GetTokenValue(ctx.SEMICOLON())}
}

func (ast *AstVisitor) VisitSelectionDecl(ctx *parser.SelectionDeclContext) *SelectionDecl{
	if(ctx == nil){
		return nil
	}

	expression_ctx, _ := ctx.Expression().(*parser.ExpressionContext)
	statement0_ctx, _ := ctx.Statement(0).(*parser.StatementContext)
	statement1_ctx, _ := ctx.Statement(1).(*parser.StatementContext)
	return &SelectionDecl {
		Node: Node{line: ctx.GetStart().GetLine()},
		if_: GetTokenValue(ctx.IF()),
		else_: GetTokenValue(ctx.ELSE()),
		lparent_: GetTokenValue(ctx.LPARENT()),
		rparent_: GetTokenValue(ctx.RPARENT()),
		Expression: ast.VisitExpression(expression_ctx),
		statement0_: ast.VisitStatement(statement0_ctx),
		statement1_: ast.VisitStatement(statement1_ctx)}
}

func (ast *AstVisitor) VisitIteratorDecl(ctx *parser.IteratorDeclContext) *IteratorDecl{
	if(ctx == nil){
		return nil
	}

	expression_ctx, _ := ctx.Expression().(*parser.ExpressionContext)
	statement_ctx, _ := ctx.Statement().(*parser.StatementContext)
	return &IteratorDecl {
		Node: Node{line: ctx.GetStart().GetLine()},
		while_: GetTokenValue(ctx.WHILE()),
		lparent_: GetTokenValue(ctx.LPARENT()),
		rparent_: GetTokenValue(ctx.RPARENT()),
		Expression: ast.VisitExpression(expression_ctx),
		Statement: ast.VisitStatement(statement_ctx)}
}

func (ast *AstVisitor) VisitReturnDecl(ctx *parser.ReturnDeclContext) *ReturnDecl{
	if(ctx == nil){
		return nil
	}

	expression_ctx, _ := ctx.Expression().(*parser.ExpressionContext)
	return &ReturnDecl {
		Node: Node{line: ctx.GetStart().GetLine()},
		return_: GetTokenValue(ctx.RETURN()),
		semicolon_: GetTokenValue(ctx.SEMICOLON()),
		Expression: ast.VisitExpression(expression_ctx)}
}

func (ast *AstVisitor) VisitExpression(ctx *parser.ExpressionContext) *Expression{
	if(ctx == nil){
		return nil
	}

	vari_ctx, _ := ctx.Vari().(*parser.VariContext)
	expression_ctx, _ := ctx.Expression().(*parser.ExpressionContext)
	simpleexpression_ctx, _ := ctx.SimpleExpression().(*parser.SimpleExpressionContext)
	return &Expression {
		Node: Node{line: ctx.GetStart().GetLine()},
		assign_: GetTokenValue(ctx.ASSIGN()),
		Vari: ast.VisitVari(vari_ctx),
		Expression: ast.VisitExpression(expression_ctx),
		SimpleExpression: ast.VisitSimpleExpression(simpleexpression_ctx)}
}

func (ast *AstVisitor) VisitVari(ctx *parser.VariContext) *Vari{
	if(ctx == nil){
		return nil
	}

	ident_ctx, _ := ctx.Ident().(*parser.IdentContext)
	expression_ctx, _ := ctx.Expression().(*parser.ExpressionContext)
	return &Vari {
		Node: Node{line: ctx.GetStart().GetLine()},
		lbracket_: GetTokenValue(ctx.LBRACKET()),
		rbracket_: GetTokenValue(ctx.RBRACKET()),
		Ident: ast.VisitIdent(ident_ctx),
		Expression: ast.VisitExpression(expression_ctx)}
}

func (ast *AstVisitor) VisitSimpleExpression(ctx *parser.SimpleExpressionContext) *SimpleExpression{
	if(ctx == nil){
		return nil
	}

	sumexpression0_ctx, _ := ctx.SumExpression(0).(*parser.SumExpressionContext)
	sumexpression1_ctx, _ := ctx.SumExpression(1).(*parser.SumExpressionContext)
	relational_ctx, _ := ctx.Relational().(*parser.RelationalContext)
	return &SimpleExpression {
		Node: Node{line: ctx.GetStart().GetLine()},
		sumexpression0_: ast.VisitSumExpression(sumexpression0_ctx),
		sumexpression1_: ast.VisitSumExpression(sumexpression1_ctx),
		Relational: ast.VisitRelational(relational_ctx)}
}

func (ast *AstVisitor) VisitRelational(ctx *parser.RelationalContext) *Relational{
	if(ctx == nil){
		return nil
	}

	return &Relational {
		Node: Node{line: ctx.GetStart().GetLine()},
		relation_: GetRelation(ctx)}
}

func (ast *AstVisitor) VisitSumExpression(ctx *parser.SumExpressionContext) *SumExpression{
	if(ctx == nil){
		return nil
	}

	sumexpression_ctx, _ := ctx.SumExpression().(*parser.SumExpressionContext)
	sum_ctx, _ := ctx.Sum().(*parser.SumContext)
	term_ctx, _ := ctx.Term().(*parser.TermContext)
	return &SumExpression {
		Node: Node{line: ctx.GetStart().GetLine()},
		SumExpression: ast.VisitSumExpression(sumexpression_ctx),
		Sum: ast.VisitSum(sum_ctx),
		Term: ast.VisitTerm(term_ctx)}
}

func (ast *AstVisitor) VisitSum(ctx *parser.SumContext) *Sum{
	if(ctx == nil){
		return nil
	}

	return &Sum {
		Node: Node{line: ctx.GetStart().GetLine()},
		operation_: GetOperationSum(ctx)}
}

func (ast *AstVisitor) VisitTerm(ctx *parser.TermContext) *Term{
	if(ctx == nil){
		return nil
	}

	term_ctx, _ := ctx.Term().(*parser.TermContext)
	mult_ctx, _ := ctx.Mult().(*parser.MultContext)
	factor_ctx, _ := ctx.Factor().(*parser.FactorContext)
	return &Term {
		Node: Node{line: ctx.GetStart().GetLine()},
		Term: ast.VisitTerm(term_ctx),
		Mult: ast.VisitMult(mult_ctx),
		Factor: ast.VisitFactor(factor_ctx)}
}

func (ast *AstVisitor) VisitMult(ctx *parser.MultContext) *Mult{
	if(ctx == nil){
		return nil
	}

	return &Mult {
		Node: Node{line: ctx.GetStart().GetLine()},
		operation_: GetOperationMult(ctx)}
}

func (ast *AstVisitor) VisitFactor(ctx *parser.FactorContext) *Factor{
	if(ctx == nil){
		return nil
	}

	expression_ctx, _ := ctx.Expression().(*parser.ExpressionContext)
	vari_ctx, _ := ctx.Vari().(*parser.VariContext)
	activation_ctx, _ := ctx.Activation().(*parser.ActivationContext)
	num_ctx, _ := ctx.Num().(*parser.NumContext)

	return &Factor {
		Node: Node{line: ctx.GetStart().GetLine()},
		lparent_: GetTokenValue(ctx.LPARENT()),
		rparent_: GetTokenValue(ctx.RPARENT()),
		Expression: ast.VisitExpression(expression_ctx),
		Vari: ast.VisitVari(vari_ctx),
		Activation: ast.VisitActivation(activation_ctx),
		Num: ast.VisitNum(num_ctx)}
}

func (ast *AstVisitor) VisitActivation(ctx *parser.ActivationContext) *Activation{
	if(ctx == nil){
		return nil
	}

	ident_ctx, _ := ctx.Ident().(*parser.IdentContext)
	arglist_ctx, _ := ctx.ArgList().(*parser.ArgListContext)
	return &Activation {
		Node: Node{line: ctx.GetStart().GetLine()},
		lparent_: GetTokenValue(ctx.LPARENT()),
		rparent_: GetTokenValue(ctx.RPARENT()),
		Ident: ast.VisitIdent(ident_ctx),
		ArgList: ast.VisitArgList(arglist_ctx)}
}

func (ast *AstVisitor) VisitArgList(ctx *parser.ArgListContext) *ArgList{
	if(ctx == nil){
		return nil
	}

	arglist_ctx, _ := ctx.ArgList().(*parser.ArgListContext)
	expression_ctx, _ := ctx.Expression().(*parser.ExpressionContext)
	return &ArgList {
		Node: Node{line: ctx.GetStart().GetLine()},
		comma_: GetTokenValue(ctx.COMMA()),
		ArgList: ast.VisitArgList(arglist_ctx),
		Expression: ast.VisitExpression(expression_ctx)}
}

func (ast *AstVisitor) VisitIdent(ctx *parser.IdentContext) *Ident{
	if(ctx == nil){
		return nil
	}

	return &Ident {
		Node: Node{line: ctx.GetStart().GetLine()},
		ID_: GetTokenValue(ctx.IDENTIFIER())}
}

func (ast *AstVisitor) VisitNum(ctx *parser.NumContext) *Num{
	if(ctx == nil){
		return nil
	}

	return &Num {
		Node: Node{line: ctx.GetStart().GetLine()},
		NUMBER_: GetTokenValue(ctx.NUMBER())}
}

func GetRelation(ctx *parser.RelationalContext) string {
	if (ctx == nil){
		return ""
	}

	if (ctx.EQUAL() != nil) {
		return ctx.EQUAL().GetText()
	}
	if (ctx.DIFFERENT() != nil){ 
		return ctx.DIFFERENT().GetText()
	}	
	if (ctx.SMALLER() != nil){
		return ctx.SMALLER().GetText()
	}
	if (ctx.SMALLERANDEQUAL() != nil){
		return ctx.SMALLERANDEQUAL().GetText()
	}
	if (ctx.LARGER() != nil){
		return ctx.LARGER().GetText()
	}
	if (ctx.LARGERANDEQUAL() != nil){
		return ctx.LARGERANDEQUAL().GetText()
	}

	return ""
}

func GetOperationMult(ctx *parser.MultContext) string {
	if (ctx == nil){
		return ""
	}

	if (ctx.TIMES() != nil){
		return ctx.TIMES().GetText()
	}
	if (ctx.DIVISION() != nil){
		return ctx.DIVISION().GetText()
	}
	
	return ""
}

func GetOperationSum(ctx *parser.SumContext) string {
	if (ctx == nil){
		return ""
	}

	if (ctx.PLUS() != nil) {
		return ctx.PLUS().GetText()
	}

	if (ctx.MINUS() != nil) {
		return ctx.MINUS().GetText()
	}
	
	return ""
}

func GetTokenValue(token antlr.TerminalNode) string {
	if (token == nil){
		return ""
	} else {
		return token.GetText()
	}
}

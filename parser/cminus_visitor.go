// Code generated from CMINUS.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // CMINUS

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by CMINUSParser.
type CMINUSVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by CMINUSParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by CMINUSParser#declarationList.
	VisitDeclarationList(ctx *DeclarationListContext) interface{}

	// Visit a parse tree produced by CMINUSParser#declaration.
	VisitDeclaration(ctx *DeclarationContext) interface{}

	// Visit a parse tree produced by CMINUSParser#varDeclaration.
	VisitVarDeclaration(ctx *VarDeclarationContext) interface{}

	// Visit a parse tree produced by CMINUSParser#funDeclaration.
	VisitFunDeclaration(ctx *FunDeclarationContext) interface{}

	// Visit a parse tree produced by CMINUSParser#params.
	VisitParams(ctx *ParamsContext) interface{}

	// Visit a parse tree produced by CMINUSParser#paramList.
	VisitParamList(ctx *ParamListContext) interface{}

	// Visit a parse tree produced by CMINUSParser#param.
	VisitParam(ctx *ParamContext) interface{}

	// Visit a parse tree produced by CMINUSParser#compoundDecl.
	VisitCompoundDecl(ctx *CompoundDeclContext) interface{}

	// Visit a parse tree produced by CMINUSParser#localDeclarations.
	VisitLocalDeclarations(ctx *LocalDeclarationsContext) interface{}

	// Visit a parse tree produced by CMINUSParser#statementList.
	VisitStatementList(ctx *StatementListContext) interface{}

	// Visit a parse tree produced by CMINUSParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by CMINUSParser#expressionDecl.
	VisitExpressionDecl(ctx *ExpressionDeclContext) interface{}

	// Visit a parse tree produced by CMINUSParser#selectionDecl.
	VisitSelectionDecl(ctx *SelectionDeclContext) interface{}

	// Visit a parse tree produced by CMINUSParser#iteratorDecl.
	VisitIteratorDecl(ctx *IteratorDeclContext) interface{}

	// Visit a parse tree produced by CMINUSParser#returnDecl.
	VisitReturnDecl(ctx *ReturnDeclContext) interface{}

	// Visit a parse tree produced by CMINUSParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by CMINUSParser#vari.
	VisitVari(ctx *VariContext) interface{}

	// Visit a parse tree produced by CMINUSParser#simpleExpression.
	VisitSimpleExpression(ctx *SimpleExpressionContext) interface{}

	// Visit a parse tree produced by CMINUSParser#relational.
	VisitRelational(ctx *RelationalContext) interface{}

	// Visit a parse tree produced by CMINUSParser#sumExpression.
	VisitSumExpression(ctx *SumExpressionContext) interface{}

	// Visit a parse tree produced by CMINUSParser#sum.
	VisitSum(ctx *SumContext) interface{}

	// Visit a parse tree produced by CMINUSParser#term.
	VisitTerm(ctx *TermContext) interface{}

	// Visit a parse tree produced by CMINUSParser#mult.
	VisitMult(ctx *MultContext) interface{}

	// Visit a parse tree produced by CMINUSParser#factor.
	VisitFactor(ctx *FactorContext) interface{}

	// Visit a parse tree produced by CMINUSParser#activation.
	VisitActivation(ctx *ActivationContext) interface{}

	// Visit a parse tree produced by CMINUSParser#argList.
	VisitArgList(ctx *ArgListContext) interface{}

	// Visit a parse tree produced by CMINUSParser#ident.
	VisitIdent(ctx *IdentContext) interface{}

	// Visit a parse tree produced by CMINUSParser#num.
	VisitNum(ctx *NumContext) interface{}
}

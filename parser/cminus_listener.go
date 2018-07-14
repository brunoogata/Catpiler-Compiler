// Code generated from CMINUS.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // CMINUS

import "github.com/antlr/antlr4/runtime/Go/antlr"

// CMINUSListener is a complete listener for a parse tree produced by CMINUSParser.
type CMINUSListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterDeclarationList is called when entering the declarationList production.
	EnterDeclarationList(c *DeclarationListContext)

	// EnterDeclaration is called when entering the declaration production.
	EnterDeclaration(c *DeclarationContext)

	// EnterVarDeclaration is called when entering the varDeclaration production.
	EnterVarDeclaration(c *VarDeclarationContext)

	// EnterFunDeclaration is called when entering the funDeclaration production.
	EnterFunDeclaration(c *FunDeclarationContext)

	// EnterParams is called when entering the params production.
	EnterParams(c *ParamsContext)

	// EnterParamList is called when entering the paramList production.
	EnterParamList(c *ParamListContext)

	// EnterParam is called when entering the param production.
	EnterParam(c *ParamContext)

	// EnterCompoundDecl is called when entering the compoundDecl production.
	EnterCompoundDecl(c *CompoundDeclContext)

	// EnterLocalDeclarations is called when entering the localDeclarations production.
	EnterLocalDeclarations(c *LocalDeclarationsContext)

	// EnterStatementList is called when entering the statementList production.
	EnterStatementList(c *StatementListContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterExpressionDecl is called when entering the expressionDecl production.
	EnterExpressionDecl(c *ExpressionDeclContext)

	// EnterSelectionDecl is called when entering the selectionDecl production.
	EnterSelectionDecl(c *SelectionDeclContext)

	// EnterIteratorDecl is called when entering the iteratorDecl production.
	EnterIteratorDecl(c *IteratorDeclContext)

	// EnterReturnDecl is called when entering the returnDecl production.
	EnterReturnDecl(c *ReturnDeclContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterVari is called when entering the vari production.
	EnterVari(c *VariContext)

	// EnterSimpleExpression is called when entering the simpleExpression production.
	EnterSimpleExpression(c *SimpleExpressionContext)

	// EnterRelational is called when entering the relational production.
	EnterRelational(c *RelationalContext)

	// EnterSumExpression is called when entering the sumExpression production.
	EnterSumExpression(c *SumExpressionContext)

	// EnterSum is called when entering the sum production.
	EnterSum(c *SumContext)

	// EnterTerm is called when entering the term production.
	EnterTerm(c *TermContext)

	// EnterMult is called when entering the mult production.
	EnterMult(c *MultContext)

	// EnterFactor is called when entering the factor production.
	EnterFactor(c *FactorContext)

	// EnterActivation is called when entering the activation production.
	EnterActivation(c *ActivationContext)

	// EnterArgList is called when entering the argList production.
	EnterArgList(c *ArgListContext)

	// EnterIdent is called when entering the ident production.
	EnterIdent(c *IdentContext)

	// EnterNum is called when entering the num production.
	EnterNum(c *NumContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitDeclarationList is called when exiting the declarationList production.
	ExitDeclarationList(c *DeclarationListContext)

	// ExitDeclaration is called when exiting the declaration production.
	ExitDeclaration(c *DeclarationContext)

	// ExitVarDeclaration is called when exiting the varDeclaration production.
	ExitVarDeclaration(c *VarDeclarationContext)

	// ExitFunDeclaration is called when exiting the funDeclaration production.
	ExitFunDeclaration(c *FunDeclarationContext)

	// ExitParams is called when exiting the params production.
	ExitParams(c *ParamsContext)

	// ExitParamList is called when exiting the paramList production.
	ExitParamList(c *ParamListContext)

	// ExitParam is called when exiting the param production.
	ExitParam(c *ParamContext)

	// ExitCompoundDecl is called when exiting the compoundDecl production.
	ExitCompoundDecl(c *CompoundDeclContext)

	// ExitLocalDeclarations is called when exiting the localDeclarations production.
	ExitLocalDeclarations(c *LocalDeclarationsContext)

	// ExitStatementList is called when exiting the statementList production.
	ExitStatementList(c *StatementListContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitExpressionDecl is called when exiting the expressionDecl production.
	ExitExpressionDecl(c *ExpressionDeclContext)

	// ExitSelectionDecl is called when exiting the selectionDecl production.
	ExitSelectionDecl(c *SelectionDeclContext)

	// ExitIteratorDecl is called when exiting the iteratorDecl production.
	ExitIteratorDecl(c *IteratorDeclContext)

	// ExitReturnDecl is called when exiting the returnDecl production.
	ExitReturnDecl(c *ReturnDeclContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitVari is called when exiting the vari production.
	ExitVari(c *VariContext)

	// ExitSimpleExpression is called when exiting the simpleExpression production.
	ExitSimpleExpression(c *SimpleExpressionContext)

	// ExitRelational is called when exiting the relational production.
	ExitRelational(c *RelationalContext)

	// ExitSumExpression is called when exiting the sumExpression production.
	ExitSumExpression(c *SumExpressionContext)

	// ExitSum is called when exiting the sum production.
	ExitSum(c *SumContext)

	// ExitTerm is called when exiting the term production.
	ExitTerm(c *TermContext)

	// ExitMult is called when exiting the mult production.
	ExitMult(c *MultContext)

	// ExitFactor is called when exiting the factor production.
	ExitFactor(c *FactorContext)

	// ExitActivation is called when exiting the activation production.
	ExitActivation(c *ActivationContext)

	// ExitArgList is called when exiting the argList production.
	ExitArgList(c *ArgListContext)

	// ExitIdent is called when exiting the ident production.
	ExitIdent(c *IdentContext)

	// ExitNum is called when exiting the num production.
	ExitNum(c *NumContext)
}

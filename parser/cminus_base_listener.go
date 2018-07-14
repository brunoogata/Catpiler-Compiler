// Code generated from CMINUS.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // CMINUS

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseCMINUSListener is a complete listener for a parse tree produced by CMINUSParser.
type BaseCMINUSListener struct{}

var _ CMINUSListener = &BaseCMINUSListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseCMINUSListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseCMINUSListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseCMINUSListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseCMINUSListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseCMINUSListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseCMINUSListener) ExitProgram(ctx *ProgramContext) {}

// EnterDeclarationList is called when production declarationList is entered.
func (s *BaseCMINUSListener) EnterDeclarationList(ctx *DeclarationListContext) {}

// ExitDeclarationList is called when production declarationList is exited.
func (s *BaseCMINUSListener) ExitDeclarationList(ctx *DeclarationListContext) {}

// EnterDeclaration is called when production declaration is entered.
func (s *BaseCMINUSListener) EnterDeclaration(ctx *DeclarationContext) {}

// ExitDeclaration is called when production declaration is exited.
func (s *BaseCMINUSListener) ExitDeclaration(ctx *DeclarationContext) {}

// EnterVarDeclaration is called when production varDeclaration is entered.
func (s *BaseCMINUSListener) EnterVarDeclaration(ctx *VarDeclarationContext) {}

// ExitVarDeclaration is called when production varDeclaration is exited.
func (s *BaseCMINUSListener) ExitVarDeclaration(ctx *VarDeclarationContext) {}

// EnterFunDeclaration is called when production funDeclaration is entered.
func (s *BaseCMINUSListener) EnterFunDeclaration(ctx *FunDeclarationContext) {}

// ExitFunDeclaration is called when production funDeclaration is exited.
func (s *BaseCMINUSListener) ExitFunDeclaration(ctx *FunDeclarationContext) {}

// EnterParams is called when production params is entered.
func (s *BaseCMINUSListener) EnterParams(ctx *ParamsContext) {}

// ExitParams is called when production params is exited.
func (s *BaseCMINUSListener) ExitParams(ctx *ParamsContext) {}

// EnterParamList is called when production paramList is entered.
func (s *BaseCMINUSListener) EnterParamList(ctx *ParamListContext) {}

// ExitParamList is called when production paramList is exited.
func (s *BaseCMINUSListener) ExitParamList(ctx *ParamListContext) {}

// EnterParam is called when production param is entered.
func (s *BaseCMINUSListener) EnterParam(ctx *ParamContext) {}

// ExitParam is called when production param is exited.
func (s *BaseCMINUSListener) ExitParam(ctx *ParamContext) {}

// EnterCompoundDecl is called when production compoundDecl is entered.
func (s *BaseCMINUSListener) EnterCompoundDecl(ctx *CompoundDeclContext) {}

// ExitCompoundDecl is called when production compoundDecl is exited.
func (s *BaseCMINUSListener) ExitCompoundDecl(ctx *CompoundDeclContext) {}

// EnterLocalDeclarations is called when production localDeclarations is entered.
func (s *BaseCMINUSListener) EnterLocalDeclarations(ctx *LocalDeclarationsContext) {}

// ExitLocalDeclarations is called when production localDeclarations is exited.
func (s *BaseCMINUSListener) ExitLocalDeclarations(ctx *LocalDeclarationsContext) {}

// EnterStatementList is called when production statementList is entered.
func (s *BaseCMINUSListener) EnterStatementList(ctx *StatementListContext) {}

// ExitStatementList is called when production statementList is exited.
func (s *BaseCMINUSListener) ExitStatementList(ctx *StatementListContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseCMINUSListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseCMINUSListener) ExitStatement(ctx *StatementContext) {}

// EnterExpressionDecl is called when production expressionDecl is entered.
func (s *BaseCMINUSListener) EnterExpressionDecl(ctx *ExpressionDeclContext) {}

// ExitExpressionDecl is called when production expressionDecl is exited.
func (s *BaseCMINUSListener) ExitExpressionDecl(ctx *ExpressionDeclContext) {}

// EnterSelectionDecl is called when production selectionDecl is entered.
func (s *BaseCMINUSListener) EnterSelectionDecl(ctx *SelectionDeclContext) {}

// ExitSelectionDecl is called when production selectionDecl is exited.
func (s *BaseCMINUSListener) ExitSelectionDecl(ctx *SelectionDeclContext) {}

// EnterIteratorDecl is called when production iteratorDecl is entered.
func (s *BaseCMINUSListener) EnterIteratorDecl(ctx *IteratorDeclContext) {}

// ExitIteratorDecl is called when production iteratorDecl is exited.
func (s *BaseCMINUSListener) ExitIteratorDecl(ctx *IteratorDeclContext) {}

// EnterReturnDecl is called when production returnDecl is entered.
func (s *BaseCMINUSListener) EnterReturnDecl(ctx *ReturnDeclContext) {}

// ExitReturnDecl is called when production returnDecl is exited.
func (s *BaseCMINUSListener) ExitReturnDecl(ctx *ReturnDeclContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseCMINUSListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseCMINUSListener) ExitExpression(ctx *ExpressionContext) {}

// EnterVari is called when production vari is entered.
func (s *BaseCMINUSListener) EnterVari(ctx *VariContext) {}

// ExitVari is called when production vari is exited.
func (s *BaseCMINUSListener) ExitVari(ctx *VariContext) {}

// EnterSimpleExpression is called when production simpleExpression is entered.
func (s *BaseCMINUSListener) EnterSimpleExpression(ctx *SimpleExpressionContext) {}

// ExitSimpleExpression is called when production simpleExpression is exited.
func (s *BaseCMINUSListener) ExitSimpleExpression(ctx *SimpleExpressionContext) {}

// EnterRelational is called when production relational is entered.
func (s *BaseCMINUSListener) EnterRelational(ctx *RelationalContext) {}

// ExitRelational is called when production relational is exited.
func (s *BaseCMINUSListener) ExitRelational(ctx *RelationalContext) {}

// EnterSumExpression is called when production sumExpression is entered.
func (s *BaseCMINUSListener) EnterSumExpression(ctx *SumExpressionContext) {}

// ExitSumExpression is called when production sumExpression is exited.
func (s *BaseCMINUSListener) ExitSumExpression(ctx *SumExpressionContext) {}

// EnterSum is called when production sum is entered.
func (s *BaseCMINUSListener) EnterSum(ctx *SumContext) {}

// ExitSum is called when production sum is exited.
func (s *BaseCMINUSListener) ExitSum(ctx *SumContext) {}

// EnterTerm is called when production term is entered.
func (s *BaseCMINUSListener) EnterTerm(ctx *TermContext) {}

// ExitTerm is called when production term is exited.
func (s *BaseCMINUSListener) ExitTerm(ctx *TermContext) {}

// EnterMult is called when production mult is entered.
func (s *BaseCMINUSListener) EnterMult(ctx *MultContext) {}

// ExitMult is called when production mult is exited.
func (s *BaseCMINUSListener) ExitMult(ctx *MultContext) {}

// EnterFactor is called when production factor is entered.
func (s *BaseCMINUSListener) EnterFactor(ctx *FactorContext) {}

// ExitFactor is called when production factor is exited.
func (s *BaseCMINUSListener) ExitFactor(ctx *FactorContext) {}

// EnterActivation is called when production activation is entered.
func (s *BaseCMINUSListener) EnterActivation(ctx *ActivationContext) {}

// ExitActivation is called when production activation is exited.
func (s *BaseCMINUSListener) ExitActivation(ctx *ActivationContext) {}

// EnterArgList is called when production argList is entered.
func (s *BaseCMINUSListener) EnterArgList(ctx *ArgListContext) {}

// ExitArgList is called when production argList is exited.
func (s *BaseCMINUSListener) ExitArgList(ctx *ArgListContext) {}

// EnterIdent is called when production ident is entered.
func (s *BaseCMINUSListener) EnterIdent(ctx *IdentContext) {}

// ExitIdent is called when production ident is exited.
func (s *BaseCMINUSListener) ExitIdent(ctx *IdentContext) {}

// EnterNum is called when production num is entered.
func (s *BaseCMINUSListener) EnterNum(ctx *NumContext) {}

// ExitNum is called when production num is exited.
func (s *BaseCMINUSListener) ExitNum(ctx *NumContext) {}

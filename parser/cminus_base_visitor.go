// Code generated from CMINUS.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // CMINUS

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseCMINUSVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseCMINUSVisitor) VisitProgram(ctx *ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCMINUSVisitor) VisitDeclarationList(ctx *DeclarationListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCMINUSVisitor) VisitDeclaration(ctx *DeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCMINUSVisitor) VisitVarDeclaration(ctx *VarDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCMINUSVisitor) VisitFunDeclaration(ctx *FunDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCMINUSVisitor) VisitParams(ctx *ParamsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCMINUSVisitor) VisitParamList(ctx *ParamListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCMINUSVisitor) VisitParam(ctx *ParamContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCMINUSVisitor) VisitCompoundDecl(ctx *CompoundDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCMINUSVisitor) VisitLocalDeclarations(ctx *LocalDeclarationsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCMINUSVisitor) VisitStatementList(ctx *StatementListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCMINUSVisitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCMINUSVisitor) VisitExpressionDecl(ctx *ExpressionDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCMINUSVisitor) VisitSelectionDecl(ctx *SelectionDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCMINUSVisitor) VisitIteratorDecl(ctx *IteratorDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCMINUSVisitor) VisitReturnDecl(ctx *ReturnDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCMINUSVisitor) VisitExpression(ctx *ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCMINUSVisitor) VisitVari(ctx *VariContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCMINUSVisitor) VisitSimpleExpression(ctx *SimpleExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCMINUSVisitor) VisitRelational(ctx *RelationalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCMINUSVisitor) VisitSumExpression(ctx *SumExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCMINUSVisitor) VisitSum(ctx *SumContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCMINUSVisitor) VisitTerm(ctx *TermContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCMINUSVisitor) VisitMult(ctx *MultContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCMINUSVisitor) VisitFactor(ctx *FactorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCMINUSVisitor) VisitActivation(ctx *ActivationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCMINUSVisitor) VisitArgList(ctx *ArgListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCMINUSVisitor) VisitIdent(ctx *IdentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCMINUSVisitor) VisitNum(ctx *NumContext) interface{} {
	return v.VisitChildren(ctx)
}

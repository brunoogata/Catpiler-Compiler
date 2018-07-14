package ast

import (
	"fmt"
	"container/list"
	"strings"
)

var contTemp int
var contLabel int
var contAssign int 
var contValue int
var contAdd int
var contParam int
var mainLine int

var N int 

var indexOp2 Operand

type Operand struct {
	ID string
	Scope string
	Type string
}

var emptyOp Operand

type Quad struct {
	Inst string
	Op1 Operand
	Op2 Operand
	Op3 Operand
}

func SyntAnalysis (prog *Program) {
	list_ := list.New()
	
	contTemp = 0
	contLabel = 0
	contAssign = 0
	contValue = 0
	contAdd = 0
	contParam = 0

	var selfOp Operand
	selfOp.ID = "main"

	var selfQuad Quad
	selfQuad.Inst = "goto"
	selfQuad.Op1 = selfOp

	list_.PushBack(selfQuad)	
	
	N = 0
	ProgramK(prog, list_)
	PrintIntermediate(list_)
	PrintAssembly(list_)
}

func ProgramK (self *Program, list_ *list.List) {
	if(self == nil) {
		return
	}

	DeclListK(self.DeclarationList, list_)
}

func DeclListK (self *DeclarationList, list_ *list.List) {
	if(self == nil) {
		return
	}

	DeclListK(self.DeclarationList, list_)
	DeclarationK(self.Declaration, list_)
}

func DeclarationK (self *Declaration, list_ *list.List) {
	if(self == nil) {
		return
	}

	VarDeclarationK(self.VarDeclaration, list_)
	FunDeclarationK(self.FunDeclaration, list_)
}

func VarDeclarationK (self *VarDeclaration, list_ *list.List) {
	if(self == nil) {
		return
	}
	_ = IdentK(self.Ident)
}

func FunDeclarationK (self *FunDeclaration, list_ *list.List) {
	if(self == nil) {
		return
	}

	var selfQuad Quad
	var selfOp Operand
	selfOp = IdentK(self.Ident)
	selfQuad.Inst = "function"
	selfQuad.Op1 = selfOp

	list_.PushBack(selfQuad)

	ParamsK(self.Params, list_, selfOp.ID)
	CompoundDeclK(self.CompoundDecl, list_, self.void_)
}

func ParamsK (self *Params, list_ *list.List, scope_ string) {
	if(self == nil) {
		return 
	}

	ParamListK(self.ParamList, list_, self.void_, scope_)
}

func ParamListK (self *ParamList, list_ *list.List, void_ string, scope_ string) {
	if(self == nil) {
		return
	}

	if(void_ != "void"){
		ParamK(self.Param, list_, scope_)	
	}
	
	ParamListK(self.ParamList, list_, void_, scope_)
}

func ParamK (self *Param, list_ *list.List, scope_ string) {
	if(self == nil) {
		return
	}
	
	var selfQuad Quad
	var selfOp Operand
	selfOp2 := Operand{ID: scope_}

	selfOp = IdentK(self.Ident)
	if(self.lbracket_ == "[" && self.rbracket_ == "]"){
		selfOp.Type = "vector"
	}

	selfQuad.Inst = "param"
	selfQuad.Op1 = selfOp
	selfQuad.Op2 = selfOp2
	list_.PushBack(selfQuad)
}

func CompoundDeclK (self *CompoundDecl, list_ *list.List, type_ string) {
	if(self == nil) {
		return
	}

	LocalDeclarationsK(self.LocalDeclarations, list_)
	StatementListK(self.StatementList, list_)

	if(type_ == "void") {
		var selfQuad Quad
		selfQuad.Inst = "return_void"
		list_.PushBack(selfQuad)
	}
}

func LocalDeclarationsK (self *LocalDeclarations, list_ *list.List) {
	if(self == nil) {
		return
	}

	LocalDeclarationsK(self.LocalDeclarations, list_)
	VarDeclarationK(self.VarDeclaration, list_)
}

func StatementListK (self *StatementList, list_ *list.List) {
	if(self == nil) {
		return
	}

	StatementListK(self.StatementList, list_)	
	StatementK(self.Statement, list_)
}

func StatementK (self *Statement, list_ *list.List) {
	if(self == nil) {
		return
	}

	ExpressionDeclK(self.ExpressionDecl, list_)
	CompoundDeclK(self.CompoundDecl, list_, "")
	SelectionDeclK(self.SelectionDecl, list_)
	IteratorDeclK(self.IteratorDecl, list_)
	ReturnDeclK(self.ReturnDecl, list_)
}

func ExpressionDeclK (self *ExpressionDecl, list_ *list.List) {
	if(self == nil) {
		return
	}

	_ = ExpressionK(self.Expression, list_)
}

func SelectionDeclK (self *SelectionDecl, list_ *list.List) {
	if(self == nil) {
		return
	}  
	
	var selfQuad Quad
	var selfOp Operand

	if(self.else_ == ""){
		selfQuad.Inst = "jump_if_false"
		selfQuad.Op1 = ExpressionK(self.Expression, list_)
		selfOp.ID = GenLabel()
		selfQuad.Op2 = selfOp
		selfQuad.Op3 = emptyOp
		list_.PushBack(selfQuad)

		StatementK(self.statement0_, list_)

		selfQuad.Inst = "label"
		selfQuad.Op1 = Operand{ID: selfOp.ID}
		selfQuad.Op2 = emptyOp
		selfQuad.Op3 = emptyOp
		list_.PushBack(selfQuad)
	} else {
		selfQuad.Inst = "jump_if_false"
		selfQuad.Op1 = ExpressionK(self.Expression, list_)
		selfOp.ID = GenLabel()
		selfQuad.Op2 = selfOp
		selfQuad.Op3 = emptyOp
		list_.PushBack(selfQuad)

		StatementK(self.statement0_, list_)

		selfQuad.Inst = "label"
		selfQuad.Op1 = Operand{ID: selfOp.ID}
		selfQuad.Op2 = emptyOp
		selfQuad.Op3 = emptyOp
		list_.PushBack(selfQuad)

		StatementK(self.statement1_, list_)
	}
}

func IteratorDeclK (self *IteratorDecl, list_ *list.List) {
	if(self == nil) {
		return
	}

	selfOp := Operand{ID: GenLabel()}
	selfQuad := Quad{Inst: "label", Op1: selfOp, Op2: emptyOp, Op3: emptyOp}
	list_.PushBack(selfQuad)

	selfOp2 := Operand{ID: GenLabel()}
	selfQuad.Inst = "jump_if_false"
	selfQuad.Op1 = ExpressionK(self.Expression, list_)
	selfQuad.Op2 = selfOp2
	selfQuad.Op3 = emptyOp
	list_.PushBack(selfQuad)

	StatementK(self.Statement, list_)

	selfQuad.Inst = "goto"
	selfQuad.Op1 = selfOp
	selfQuad.Op2 = emptyOp
	selfQuad.Op3 = emptyOp
	list_.PushBack(selfQuad)

	selfQuad = Quad{Inst: "label", Op1: selfOp2, Op2: emptyOp, Op3: emptyOp}
	list_.PushBack(selfQuad)
}

func ReturnDeclK (self *ReturnDecl, list_ *list.List) {
	if(self == nil) {
		return
	}

	var selfQuad Quad
	selfQuad.Inst = "return_int"
	selfQuad.Op1 = ExpressionK(self.Expression, list_)

	list_.PushBack(selfQuad) 
}

func ExpressionK (self *Expression, list_ *list.List) Operand {
	var selfOp Operand
	
	if(self == nil) {
		return selfOp
	}

	if(self.Vari != nil){
		var selfQuad Quad
		var selfQuad2 Quad
		var opAux Operand
		var opAux2 Operand
		var selfOp Operand
		var selfOp2 Operand
		var ida Operand

		opAux, ida = VariK(self.Vari, list_, false)
		opAux2 = ExpressionK(self.Expression, list_)

		if(opAux.Type == "vector") {
			selfOp.ID = GenTemp()
			selfOp.Type = "temp"
			//tempPackageAux.Var = IdentK(self.Ident)
			//tempPackageAux.Index = ExpressionK(self.Expression, list_)
			selfQuad2.Inst = "vectortotemp"
			selfQuad2.Op1 = selfOp
			selfQuad2.Op2 = opAux//tempPackageAux.Var
			selfQuad2.Op3 = ida
			list_.PushBack(selfQuad2)

			if(opAux2.Type == "vector"){
				selfOp2.ID = GenTemp()
				selfOp2.Type = "temp"

				selfQuad2.Inst = "vectortotemp"
				selfQuad2.Op1 = selfOp2
				selfQuad2.Op2 = opAux2
				selfQuad2.Op3 = indexOp2
				list_.PushBack(selfQuad2)

				selfQuad.Inst = "assign"
				selfQuad.Op1 = selfOp
				selfQuad.Op2 = selfOp2
				list_.PushBack(selfQuad)
			} else {
				selfQuad.Inst = "assign"
				selfQuad.Op1 = selfOp
				selfQuad.Op2 = opAux2
				list_.PushBack(selfQuad)
			}
					
			selfQuad2.Inst = "temptovector"
			selfQuad2.Op1 = opAux
			selfQuad2.Op2 = ida
			selfQuad2.Op3 = selfOp
			list_.PushBack(selfQuad2)
		} else if(opAux2.Type == "vector"){

			selfOp.ID = GenTemp()
			selfOp.Type = "temp"
			
			selfQuad2.Inst = "vectortotemp"
			selfQuad2.Op1 = selfOp
			selfQuad2.Op2 = opAux2
			selfQuad2.Op3 = indexOp2
			list_.PushBack(selfQuad2)

			/*
			if(opAux.Type == "vector"){
				selfOp2.ID = GenTemp()
				selfOp2.Type = "temp"

				selfQuad2.Inst = "vectortotemp"
				selfQuad2.Op1 = selfOp2
				selfQuad2.Op2 = opAux//tempPackageAux.Var
				selfQuad2.Op3 = ida
				list_.PushBack(selfQuad2)

				selfQuad.Inst = "assign"
				selfQuad.Op1 = selfOp2
				selfQuad.Op2 = selfOp 
				list_.PushBack(selfQuad)

				selfQuad2.Inst = "temptovector"
				selfQuad2.Op1 = opAux
				selfQuad2.Op2 = ida
				selfQuad2.Op3 = selfOp2
				list_.PushBack(selfQuad2)
			} else {*/
				selfQuad.Inst = "assign"
				selfQuad.Op1 = opAux
				selfQuad.Op2 = selfOp
				list_.PushBack(selfQuad)
			//}
		} else {
			selfQuad.Inst = "assign"
			selfQuad.Op1 = opAux
			selfQuad.Op2 = opAux2
			list_.PushBack(selfQuad) 
		}
		return selfOp
	} else{
		return SimpleExpressionK(self.SimpleExpression, list_)
	}
}

func VariK (self *Vari, list_ *list.List, flag bool) (Operand, Operand) {
	var selfOp Operand
	var indexOp Operand

	if(self == nil) {
		return selfOp, selfOp
	}

	if(self.lbracket_ == "[" && self.rbracket_ == "]"){
		/*var selfQuad Quad
		selfOp.ID = GenTemp()
		selfOp.Type = "temp"
		tempPackageAux.Var = IdentK(self.Ident)
		tempPackageAux.Index = ExpressionK(self.Expression, list_)*/
		selfOp = IdentK(self.Ident)
		selfOp.Type = "vector"

		indexOp = ExpressionK(self.Expression, list_)
		if(flag){
			indexOp2 = indexOp
		}
		return selfOp, indexOp
	} else {
		return IdentK(self.Ident), selfOp
	}
}

func SimpleExpressionK (self *SimpleExpression, list_ *list.List) Operand{
	var selfOp Operand
	var selfOp1 Operand
	var selfOp2 Operand 
	if(self == nil) {
		return selfOp
	}

	if(self.Relational == nil){
		return SumExpressionK(self.sumexpression0_, list_)
	} else {
		selfOp.ID = GenTemp()
		selfOp.Type = "temp"
		var selfQuad Quad 
		opAux := SumExpressionK(self.sumexpression0_, list_)  
		

		if(opAux.Type == "vector"){
			var selfQuad2 Quad
			selfOp1.ID = GenTemp()
			selfOp1.Type = "temp"

			selfQuad2.Inst = "vectortotemp"
			selfQuad2.Op1 = selfOp1
			selfQuad2.Op2 = opAux
			selfQuad2.Op3 = indexOp2
			list_.PushBack(selfQuad2)
			opAux = selfOp1
		} 

		opAux2 := SumExpressionK(self.sumexpression1_, list_)

		if(opAux2.Type == "vector"){
			var selfQuad2 Quad 
			selfOp2.ID = GenTemp()
			selfOp2.Type = "temp"
			selfQuad2.Inst = "vectortotemp"
			selfQuad2.Op1 = selfOp2
			selfQuad2.Op2 = opAux2
			selfQuad2.Op3 = indexOp2
			list_.PushBack(selfQuad2)
			opAux2 = selfOp2
		}

		selfQuad.Inst = RelationalK(self.Relational)
		selfQuad.Op2 = opAux
		selfQuad.Op3 = opAux2
		selfQuad.Op1 = selfOp
		list_.PushBack(selfQuad) 

		return selfOp
	}
}

func RelationalK (self *Relational) string{
	if(self == nil) {
		return ""
	}

	if(self.relation_ == "=="){
		return "equal"
	} else if (self.relation_ == "!="){
		return "different"
	} else if (self.relation_ == "<"){
		return "smaller"
	} else if (self.relation_ == "<="){
		return "smaller_and_equal"
	} else if (self.relation_ == ">"){
		return "greater"
	} else if (self.relation_ == ">="){
		return "greater_and_equal"
	}

	return ""
}

func SumExpressionK (self *SumExpression, list_ *list.List) Operand{
	var selfOp Operand
	
	if(self == nil) {
		return selfOp
	}

	if(self.Sum != nil){
		var selfQuad Quad
		selfOp.ID = GenTemp()
		selfOp.Type = "temp"
		selfQuad.Inst = SumK(self.Sum)
		selfQuad.Op1 = selfOp
		selfQuad.Op2 = SumExpressionK(self.SumExpression, list_)
		selfQuad.Op3 = TermK(self.Term, list_)

		list_.PushBack(selfQuad)
	} else {
		selfOp = TermK(self.Term, list_)
	}
	return selfOp
}

func SumK(self *Sum) string{
	if(self == nil) {
		return ""
	}

	if(self.operation_ == "+"){
		return "addition"
	} else {
		return "subtraction"
	}

	return ""
}

func TermK (self *Term, list_ *list.List) Operand {
	var selfOp Operand
	
	if(self == nil) {
		return selfOp
	}

	if(self.Mult != nil){
		var selfQuad Quad 
		selfOp.ID = GenTemp()
		selfOp.Type = "temp"
		selfQuad.Inst = MultK(self.Mult)
		selfQuad.Op1 = selfOp
		selfQuad.Op2 = TermK(self.Term, list_)
		selfQuad.Op3 = FactorK(self.Factor, list_)

		list_.PushBack(selfQuad)
	} else {
		selfOp = FactorK(self.Factor, list_)
	}
	return selfOp
}

func MultK(self *Mult) string{
	if(self == nil) {
		return ""
	}

	if(self.operation_ == "*"){
		return "multiplication"
	} else {
		return "division"
	}

	return ""
}

func FactorK (self *Factor, list_ *list.List) Operand{
	var selfOp Operand
	
	if(self == nil) {
		return selfOp
	}

	if(self.Expression != nil){
		return ExpressionK(self.Expression, list_)
	} else if(self.Vari != nil){
		opAux, _ := VariK(self.Vari, list_, true) 
		return opAux
	} else if(self.Num != nil){
		return NumK(self.Num)
	} else {
		return ActivationK(self.Activation, list_)
	}
}

func ActivationK (self *Activation, list_ *list.List) Operand{
	var selfOp Operand
	
	if(self == nil) {
		return selfOp
	}

	var selfQuad Quad 
	selfOp2 := IdentK(self.Ident)
	x := ArgListK(self.ArgList, list_, selfOp2.ID)
	if(selfOp2.ID == "output" || selfOp2.ID == "input"){
		selfQuad.Inst = selfOp2.ID
		if(selfOp2.ID == "output"){
			selfQuad.Op1 = x
		} else {
			selfOp.ID = GenTemp()
			selfOp.Type = "temp"
			selfQuad.Op1 = selfOp
		}
		
	} else {
		selfOp.ID = GenTemp()
		selfOp.Type = "temp"
		selfQuad.Inst = "call"
		selfQuad.Op1 = selfOp2
		selfQuad.Op2 = selfOp //vazio //ArgListK(self.ArgList, list_)
		selfQuad.Op3 = selfOp //vazio
	}

	list_.PushBack(selfQuad)

	return selfOp
}

func ArgListK (self *ArgList, list_ *list.List, scope_ string) Operand{
	if(self == nil) {
		return Operand{}
	}
	selfOp2 := Operand{ID: scope_}
	var selfQuad Quad 
	var selfOp Operand

	selfOp = ExpressionK(self.Expression, list_)

	if(scope_ == "output"){
		return selfOp 	
	}

/*
	if(selfOp.Type == "vector"){
		selfQuad.Inst = "param_vector"
		selfQuad.Op1 = selfOp
		list_.PushBack(selfQuad)
	} else {*/
		selfQuad.Inst = "argument"
		selfQuad.Op1 = selfOp
		selfQuad.Op2 = selfOp2
		list_.PushBack(selfQuad)
	//}

	ArgListK(self.ArgList, list_, scope_)
	return Operand{}
}

func IdentK (self *Ident) Operand{
	var selfOp Operand
	
	if(self == nil) {
		return selfOp
	}

	selfOp.ID = self.ID_
	return selfOp
}

func NumK (self *Num) Operand{
	var selfOp Operand
	
	if(self == nil) {
		return selfOp
	}

	selfOp.ID = self.NUMBER_

	return selfOp
}

func PrintIntermediate(list_ *list.List) {
	fmt.Println("\n|-----------------------------Catpiler Intermediate Code-------------------------------|")
	for quad_ := list_.Front(); quad_ != nil; quad_ = quad_.Next() {
		selfQuad, _ := quad_.Value.(Quad)
		op1, op2, op3 := OpText(selfQuad)
		fmt.Println("(", selfQuad.Inst, ",", op1, ",", op2, ",", op3, ")")
	}
	fmt.Println("|--------------------------------------------------------------------------------------|\n")
}

func OpText(quad_ Quad) (string, string, string){
	var op1_, op2_, op3_ string
	aux1 := quad_.Op1
	aux2 := quad_.Op2
	aux3 := quad_.Op3

	if(aux1.ID == ""){
		op1_ = "_"
	} else {
		op1_ = aux1.ID
	}
	if(aux2.ID == ""){
		op2_ = "_"
	} else {
		op2_ = aux2.ID
	}
	if(aux3.ID == ""){
		op3_ = "_"
	} else {
		op3_ = aux3.ID
	}

	return op1_, op2_, op3_
}

func GenTemp() string {
	var a []string
	a = append(a, "t", fmt.Sprint(contTemp))
	contTemp++
	return strings.Join(a, "")
}

func GenLabel() string {
	var a []string
	a = append(a, "L", fmt.Sprint(contLabel))
	contLabel++
	return strings.Join(a, "")
}

func GenAssign() string {
	var a []string 
	a = append(a, "s", fmt.Sprint(contAssign))
	contAssign++
	return strings.Join(a, "")
}

func GenValue() string {
	var a []string 
	a = append(a, "v", fmt.Sprint(contValue))
	contValue++
	return strings.Join(a, "")
}

func GenAddress() string {
	var a []string 
	a = append(a, "adr", fmt.Sprint(contAdd))
	contAdd++
	return strings.Join(a, "")
}

func GenParam() string {
	var a []string 
	a = append(a, "p", fmt.Sprint(contParam))
	contParam++
	return strings.Join(a, "")
}

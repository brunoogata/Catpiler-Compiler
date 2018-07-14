package ast

import (
	"fmt"
	"strconv"
	"container/list"
)

type VarPos struct {
	posicao int 
	memspace int
	scope string
	auxvector bool
	auxparam bool
	auxreg string
}

type QuadAssembly struct {
	inst string
	op1 string 
	op2 string
	op3 string 
}

func PrintAssembly(list_ *list.List) {
	listAssembly := list.New()
	var quadA QuadAssembly
	
	var nVariable map[string]int 
	nVariable = make(map[string]int)
	CountVari(nVariable)

	fmt.Println("\n|----------------------------------Catpiler Assembly-----------------------------------|")
	fmt.Println("\tNOP")
	quadA = QuadAssembly{inst: "NOP"}
	listAssembly.PushBack(quadA)
	
	m := make(map[string]VarPos)
	countGlobal := 0

	ml := make(map[string]*list.List)

	for _, symbol_ := range table {
		if(symbol_.scope == "" && symbol_.typeID == "variable"){
				/*m[symbol_.ID]*/aux := VarPos{
					posicao: countGlobal,
					scope: ""}
				if(symbol_.vector == "yes"){
					aux.memspace = symbol_.spacemem
					m[symbol_.ID] = aux
					if(symbol_.spacemem == 0){
						countGlobal += 1
					} else {
						countGlobal += symbol_.spacemem
					}
				} else {
					m[symbol_.ID] = aux
					countGlobal += 1
				}
		}
	}

	Nvar := 0

	for quad_ := list_.Front(); quad_ != nil; quad_ = quad_.Next() {
		selfQuad, _ := quad_.Value.(Quad)
		op1, op2, op3 := OpText(selfQuad)
		//fmt.Println("(", selfQuad.Inst, ",", op1, ",", op2, ",", op3, ")")
		switch selfQuad.Inst {
		case "function":
			fmt.Println(op1, ":")
			quadA = QuadAssembly{inst: "FUNCTION", op1: op1}
			listAssembly.PushBack(quadA)

			for key_, symbol_ := range m{
				if(symbol_.scope != ""){
					delete(m, key_)
				}
			}

			Nvar = CountNVar(op1, m)
			// adiciona um N (numero de variaveis) no $sp

			ml[op1] = list.New()
			
			fmt.Println("\tSUBI", "$sp", "$sp", Nvar)
			quadA = QuadAssembly{inst: "SUBI", op1: "sp", op2: "sp", op3: strconv.Itoa(Nvar)}
			listAssembly.PushBack(quadA)
			
			
		case "label":
			fmt.Println(op1, ":")
			quadA = QuadAssembly{inst: "LABEL", op1: op1}
			listAssembly.PushBack(quadA)
		case "assign":
			_, err := strconv.Atoi(op2)
			var pos int
			// verificar posicao da pilha da variravel em questao
			if(err == nil) {
				opAux := GenValue()

				fmt.Println("\tLOADI", "$" + opAux, op2)
				quadA = QuadAssembly{inst: "LOADI", op1: opAux, op2: op2}
				listAssembly.PushBack(quadA)
				if(selfQuad.Op1.Type == "temp"){
					fmt.Println("\tMOV", "$" + op1, "$" + opAux)
					quadA = QuadAssembly{inst: "MOV", op1: op1, op2: opAux}
					listAssembly.PushBack(quadA)
				} else {
					pos = m[op1].posicao
					if(m[op1].scope != ""){
						fmt.Println("\tSW", /*posicao da pilha*/pos, "( $sp )", "$" + opAux)
						quadA = QuadAssembly{inst: "SW", op1: strconv.Itoa(pos), op2: "sp", op3: opAux}
						listAssembly.PushBack(quadA)
					} else if(selfQuad.Op1.Type == "temp"){
						fmt.Println("\tMOV", "$" + op1, "$" + opAux)
						quadA = QuadAssembly{inst: "MOV", op1: op1, op2: opAux}
						listAssembly.PushBack(quadA)
					} else {
						fmt.Println("\tSW", /*posicao da pilha*/pos, "( $gp )", "$" + opAux)
						quadA = QuadAssembly{inst: "SW", op1: strconv.Itoa(pos), op2: "gp", op3: opAux}
						listAssembly.PushBack(quadA)
					}
				}
				
			} else {
				// TODO
				if(selfQuad.Op2.Type == "temp"){
					if(selfQuad.Op1.Type == "temp"){
						fmt.Println("\tMOV", "$" + op1, "$" + op2)
						quadA = QuadAssembly{inst: "MOV", op1: op1, op2: op2}
						listAssembly.PushBack(quadA)
					}else {
						pos = m[op1].posicao
						if(m[op1].scope != ""){
							fmt.Println("\tSW", /*posicao da pilha*/pos, "( $sp )", "$" + op2)
							quadA = QuadAssembly{inst: "SW", op1: strconv.Itoa(pos), op2: "sp", op3: op2}
							listAssembly.PushBack(quadA)
						} else {
							fmt.Println("\tSW", /*posicao da pilha*/pos, "( $gp )", "$" + op2)
							quadA = QuadAssembly{inst: "SW", op1: strconv.Itoa(pos), op2: "gp", op3: op2}
							listAssembly.PushBack(quadA)
						}
					}	
				} else {
					// verificar a posicao na pilha
					opAux := GenAssign()
					

					pos = m[op2].posicao
					if(m[op2].scope != ""){
						fmt.Println("\tLW", "$" + opAux, pos /*posicao da pilha*/, "( $sp )")
						quadA = QuadAssembly{inst: "LW", op1: opAux, op2: strconv.Itoa(pos), op3: "sp"}
						listAssembly.PushBack(quadA)
					} else {
						fmt.Println("\tLW", "$" + opAux, pos /*posicao da pilha*/, "( $gp )")
						quadA = QuadAssembly{inst: "LW", op1: opAux, op2: strconv.Itoa(pos), op3: "gp"}
						listAssembly.PushBack(quadA)
					}
					
					if(selfQuad.Op1.Type == "temp"){
						fmt.Println("\tMOV", "$" + op1, "$" + opAux)
						quadA = QuadAssembly{inst: "MOV", op1: op1, op2: opAux}
						listAssembly.PushBack(quadA)
					} else {
						pos = m[op1].posicao
						if(m[op1].scope != ""){
							fmt.Println("\tSW", /*posicao da pilha*/pos, "( $sp )", "$" + opAux)
							quadA = QuadAssembly{inst: "SW", op1: strconv.Itoa(pos), op2: "sp", op3: opAux}
							listAssembly.PushBack(quadA)
						} else {
							fmt.Println("\tSW", /*posicao da pilha*/pos, "( $gp )", "$" + opAux)
							quadA = QuadAssembly{inst: "SW", op1: strconv.Itoa(pos), op2: "gp", op3: opAux}
							listAssembly.PushBack(quadA)
						}
					}
				}
			} 
			
		case "goto":
			fmt.Println("\tJ", op1)
			quadA = QuadAssembly{inst: "J", op1: op1}
			listAssembly.PushBack(quadA)
		case "addition":
			//if(selfQuad.Op1.Type == "temp"){
			_, err1 := strconv.Atoi(op2)
			_, err2 := strconv.Atoi(op3)

			var pos int 

			if(err1 == nil && err2 == nil){
				opAux := GenValue()

				fmt.Println("\tLOADI", "$" + opAux, op2)
				quadA = QuadAssembly{inst: "LOADI", op1: opAux, op2: op2}
				listAssembly.PushBack(quadA)

				fmt.Println("\tADDI", "$" + op1, "$" + opAux, op3)
				quadA = QuadAssembly{inst: "ADDI", op1: op1, op2: opAux, op3: op3}
				listAssembly.PushBack(quadA)
			} else if(err1 == nil){
				if(selfQuad.Op3.Type == "temp"){
					fmt.Println("\tADDI", "$" + op1, "$" + op3, op2)
					quadA = QuadAssembly{inst: "ADDI", op1: op1, op2: op3, op3: op2}
					listAssembly.PushBack(quadA)
				} else {
					opAux := GenAssign()
					pos = m[op3].posicao
					if(m[op3].scope != ""){
						fmt.Println("\tLW", "$" + opAux, pos, "( $sp )")
						quadA = QuadAssembly{inst: "LW", op1: opAux, op2: strconv.Itoa(pos), op3: "sp"}
						listAssembly.PushBack(quadA)
					} else {
						fmt.Println("\tLW", "$" + opAux, pos, "( $gp )")
						quadA = QuadAssembly{inst: "LW", op1: opAux, op2: strconv.Itoa(pos), op3: "gp"}
						listAssembly.PushBack(quadA)
					}
					fmt.Println("\tADDI", "$" + op1, "$" + opAux, op2)
					quadA = QuadAssembly{inst: "ADDI", op1: op1, op2: opAux, op3: op2}
					listAssembly.PushBack(quadA)
				}
				
			} else if(err2 == nil) {
				if(selfQuad.Op2.Type == "temp"){
					fmt.Println("\tADDI", "$" + op1, "$" + op2, op3)
					quadA = QuadAssembly{inst: "ADDI", op1: op1, op2: op2, op3: op3}
					listAssembly.PushBack(quadA)
				} else {
					opAux := GenAssign() 
					pos = m[op2].posicao
					if(m[op2].scope != ""){
						fmt.Println("\tLW", "$" + opAux, pos, "( $sp )")
						quadA = QuadAssembly{inst: "LW", op1: opAux, op2: strconv.Itoa(pos), op3: "sp"}
						listAssembly.PushBack(quadA)
					} else {
						fmt.Println("\tLW", "$" + opAux, pos, "( $gp )")
						quadA = QuadAssembly{inst: "LW", op1: opAux, op2: strconv.Itoa(pos), op3: "gp"}
						listAssembly.PushBack(quadA)
					}
					fmt.Println("\tADDI", "$" + op1, "$" + opAux, op3)
					quadA = QuadAssembly{inst: "ADDI", op1: op1, op2: opAux, op3: op3}
					listAssembly.PushBack(quadA)
				}
			} else {
				var opAux1 string
				var opAux2 string 

				if(selfQuad.Op2.Type != "temp"){
					opAux1 = GenAssign()
					pos = m[op2].posicao
					if(m[op2].scope != ""){
						fmt.Println("\tLW", "$" + opAux1, pos, "( $sp )")
						quadA = QuadAssembly{inst: "LW", op1: opAux1, op2: strconv.Itoa(pos), op3: "sp"}
						listAssembly.PushBack(quadA)
					} else {
						fmt.Println("\tLW", "$" + opAux1, pos, "( $gp )")
						quadA = QuadAssembly{inst: "LW", op1: opAux1, op2: strconv.Itoa(pos), op3: "gp"}
						listAssembly.PushBack(quadA)
					}
				} else {
					opAux1 = op2
				}

				if(selfQuad.Op3.Type != "temp"){
					opAux2 = GenAssign()
					pos = m[op3].posicao
					if(m[op3].scope != ""){
						fmt.Println("\tLW", "$" + opAux2, pos, "( $sp )")
						quadA = QuadAssembly{inst: "LW", op1: opAux2, op2: strconv.Itoa(pos), op3: "sp"}
						listAssembly.PushBack(quadA)
					} else {
						fmt.Println("\tLW", "$" + opAux2, pos, "( $gp )")
						quadA = QuadAssembly{inst: "LW", op1: opAux2, op2: strconv.Itoa(pos), op3: "gp"}
						listAssembly.PushBack(quadA)
					}
				} else {
					opAux2 = op3
				}
				fmt.Println("\tADD", "$" + op1, "$" + opAux1, "$" + opAux2)
				quadA = QuadAssembly{inst: "ADD", op1: op1, op2: opAux1, op3: opAux2}
				listAssembly.PushBack(quadA)
			}
		case "subtraction":
			_, err1 := strconv.Atoi(op2)
			_, err2 := strconv.Atoi(op3)

			var pos int 

			if(err1 == nil && err2 == nil){
				opAux := GenValue()
				fmt.Println("\tLOADI", "$" + opAux, op2)
				quadA = QuadAssembly{inst: "LOADI", op1: opAux, op2: op2}
				listAssembly.PushBack(quadA)
				fmt.Println("\tSUBI", "$" + op1, "$" + opAux, op3)
				quadA = QuadAssembly{inst: "SUBI", op1: op1, op2: opAux, op3: op3}
				listAssembly.PushBack(quadA)
			} else if(err1 == nil){
				if(selfQuad.Op3.Type == "temp"){
					fmt.Println("\tSUBI", "$" + op1, "$" + op3, op2)
					quadA = QuadAssembly{inst: "SUBI", op1: op1, op2: op3, op3: op2}
					listAssembly.PushBack(quadA)
				} else {
					opAux := GenAssign()
					pos = m[op3].posicao
					if(m[op3].scope != ""){
						fmt.Println("\tLW", "$" + opAux, pos, "( $sp )")
						quadA = QuadAssembly{inst: "LW", op1: opAux, op2: strconv.Itoa(pos), op3: "sp"}
						listAssembly.PushBack(quadA)
					} else {
						fmt.Println("\tLW", "$" + opAux, pos, "( $gp )")
						quadA = QuadAssembly{inst: "LW", op1: opAux, op2: strconv.Itoa(pos), op3: "gp"}
						listAssembly.PushBack(quadA)
					}
					fmt.Println("\tSUBI", "$" + op1, "$" + opAux, op2)
					quadA = QuadAssembly{inst: "SUBI", op1: op1, op2: opAux, op3: op2}
					listAssembly.PushBack(quadA)
				}
				
			} else if(err2 == nil) {
				if(selfQuad.Op2.Type == "temp"){
					fmt.Println("\tSUBI", "$" + op1, "$" + op2, op3)
					quadA = QuadAssembly{inst: "SUBI", op1: op1, op2: op2, op3: op3}
					listAssembly.PushBack(quadA)
				} else {
					opAux := GenAssign() 
					pos = m[op2].posicao
					if(m[op2].scope != ""){
						fmt.Println("\tLW", "$" + opAux, pos, "( $sp )")
						quadA = QuadAssembly{inst: "LW", op1: opAux, op2: strconv.Itoa(pos), op3: "sp"}
						listAssembly.PushBack(quadA)
					} else {
						fmt.Println("\tLW", "$" + opAux, pos, "( $gp )")
						quadA = QuadAssembly{inst: "LW", op1: opAux, op2: strconv.Itoa(pos), op3: "gp"}
						listAssembly.PushBack(quadA)
					}
					fmt.Println("\tSUBI", "$" + op1, "$" + opAux, op3)
					quadA = QuadAssembly{inst: "SUBI", op1: op1, op2: opAux, op3: op3}
					listAssembly.PushBack(quadA)
				}
			} else {
				var opAux1 string
				var opAux2 string 

				if(selfQuad.Op2.Type != "temp"){
					opAux1 = GenAssign()
					pos = m[op2].posicao
					if(m[op2].scope != ""){
						fmt.Println("\tLW", "$" + opAux1, pos, "( $sp )")
						quadA = QuadAssembly{inst: "LW", op1: opAux1, op2: strconv.Itoa(pos), op3: "sp"}
						listAssembly.PushBack(quadA)
					} else {
						fmt.Println("\tLW", "$" + opAux1, pos, "( $gp )")
						quadA = QuadAssembly{inst: "LW", op1: opAux1, op2: strconv.Itoa(pos), op3: "gp"}
						listAssembly.PushBack(quadA)
					}
				} else {
					opAux1 = op2
				}

				if(selfQuad.Op3.Type != "temp"){
					opAux2 = GenAssign()
					pos = m[op3].posicao
					if(m[op3].scope != ""){
						fmt.Println("\tLW", "$" + opAux2, pos, "( $sp )")
						quadA = QuadAssembly{inst: "LW", op1: opAux2, op2: strconv.Itoa(pos), op3: "sp"}
						listAssembly.PushBack(quadA)
					} else {
						fmt.Println("\tLW", "$" + opAux2, pos, "( $gp )")
						quadA = QuadAssembly{inst: "LW", op1: opAux2, op2: strconv.Itoa(pos), op3: "gp"}
						listAssembly.PushBack(quadA)
					}
				} else {
					opAux2 = op3
				}
				fmt.Println("\tSUB", "$" + op1, "$" + opAux1, "$" + opAux2)
				quadA = QuadAssembly{inst: "SUB", op1: op1, op2: opAux1, op3: opAux2}
				listAssembly.PushBack(quadA)
			}
		case "multiplication":
			_, err1 := strconv.Atoi(op2)
			_, err2 := strconv.Atoi(op3)

			var pos int 

			if(err1 == nil && err2 == nil){
				opAux := GenValue()
				fmt.Println("\tLOADI", "$" + opAux, op2)
				quadA = QuadAssembly{inst: "LOADI", op1: opAux, op2: op2}
				listAssembly.PushBack(quadA)
				fmt.Println("\tMULTI", "$" + op1, "$" + opAux, op3)
				quadA = QuadAssembly{inst: "MULTI", op1: op1, op2: opAux, op3: op3}
				listAssembly.PushBack(quadA)
			} else if(err1 == nil){
				if(selfQuad.Op3.Type == "temp"){
					fmt.Println("\tMULTI", "$" + op1, "$" + op3, op2)
					quadA = QuadAssembly{inst: "MULTI", op1: op1, op2: op3, op3: op2}
					listAssembly.PushBack(quadA)
				} else {
					opAux := GenAssign()
					pos = m[op3].posicao
					if(m[op3].scope != ""){
						fmt.Println("\tLW", "$" + opAux, pos, "( $sp )")
						quadA = QuadAssembly{inst: "LW", op1: opAux, op2: strconv.Itoa(pos), op3: "sp"}
						listAssembly.PushBack(quadA)
					} else {
						fmt.Println("\tLW", "$" + opAux, pos, "( $gp )")
						quadA = QuadAssembly{inst: "LW", op1: opAux, op2: strconv.Itoa(pos), op3: "gp"}
						listAssembly.PushBack(quadA)
					}
					fmt.Println("\tMULTI", "$" + op1, "$" + opAux, op2)
					quadA = QuadAssembly{inst: "MULTI", op1: op1, op2: opAux, op3: op2}
					listAssembly.PushBack(quadA)
				}
				
			} else if(err2 == nil) {
				if(selfQuad.Op2.Type == "temp"){
					fmt.Println("\tMULTI", "$" + op1, "$" + op2, op3)
					quadA = QuadAssembly{inst: "MULTI", op1: op1, op2: op2, op3: op3}
					listAssembly.PushBack(quadA)
				} else {
					opAux := GenAssign() 
					pos = m[op2].posicao
					if(m[op2].scope != ""){
						fmt.Println("\tLW", "$" + opAux, pos, "( $sp )")
						quadA = QuadAssembly{inst: "LW", op1: opAux, op2: strconv.Itoa(pos), op3: "sp"}
						listAssembly.PushBack(quadA)
					} else {
						fmt.Println("\tLW", "$" + opAux, pos, "( $gp )")
						quadA = QuadAssembly{inst: "LW", op1: opAux, op2: strconv.Itoa(pos), op3: "gp"}
						listAssembly.PushBack(quadA)
					}
					fmt.Println("\tMULTI", "$" + op1, "$" + opAux, op3)
					quadA = QuadAssembly{inst: "MULTI", op1: op1, op2: opAux, op3: op3}
					listAssembly.PushBack(quadA)
				}
			} else {
				var opAux1 string
				var opAux2 string 

				if(selfQuad.Op2.Type != "temp"){
					opAux1 = GenAssign()
					pos = m[op2].posicao
					if(m[op2].scope != ""){
						fmt.Println("\tLW", "$" + opAux1, pos, "( $sp )")
						quadA = QuadAssembly{inst: "LW", op1: opAux1, op2: strconv.Itoa(pos), op3: "sp"}
						listAssembly.PushBack(quadA)
					} else {
						fmt.Println("\tLW", "$" + opAux1, pos, "( $gp )")
						quadA = QuadAssembly{inst: "LW", op1: opAux1, op2: strconv.Itoa(pos), op3: "gp"}
						listAssembly.PushBack(quadA)
					}
				} else {
					opAux1 = op2
				}

				if(selfQuad.Op3.Type != "temp"){
					opAux2 = GenAssign()
					pos = m[op3].posicao
					if(m[op3].scope != ""){
						fmt.Println("\tLW", "$" + opAux2, pos, "( $sp )")
						quadA = QuadAssembly{inst: "LW", op1: opAux2, op2: strconv.Itoa(pos), op3: "sp"}
						listAssembly.PushBack(quadA)
					} else {
						fmt.Println("\tLW", "$" + opAux2, pos, "( $gp )")
						quadA = QuadAssembly{inst: "LW", op1: opAux2, op2: strconv.Itoa(pos), op3: "gp"}
						listAssembly.PushBack(quadA)
					}
				} else {
					opAux2 = op3
				}
				fmt.Println("\tMULT", "$" + op1, "$" + opAux1, "$" + opAux2)
				quadA = QuadAssembly{inst: "MULT", op1: op1, op2: opAux1, op3: opAux2}
				listAssembly.PushBack(quadA)
			}
		case "division":
			_, err1 := strconv.Atoi(op2)
			_, err2 := strconv.Atoi(op3)

			var pos int 

			if(err1 == nil && err2 == nil){
				opAux := GenValue()
				fmt.Println("\tLOADI", "$" + opAux, op2)
				quadA = QuadAssembly{inst: "LOADI", op1: opAux, op2: op2}
				listAssembly.PushBack(quadA)
				fmt.Println("\tDIVI", "$" + op1, "$" + opAux, op3)
				quadA = QuadAssembly{inst: "DIVI", op1: op1, op2: opAux, op3: op3}
				listAssembly.PushBack(quadA)
			} else if(err1 == nil){
				if(selfQuad.Op3.Type == "temp"){
					fmt.Println("\tDIVI", "$" + op1, "$" + op3, op2)
					quadA = QuadAssembly{inst: "DIVI", op1: op1, op2: op3, op3: op2}
					listAssembly.PushBack(quadA)
				} else {
					opAux := GenAssign()
					pos = m[op3].posicao
					if(m[op3].scope != ""){
						fmt.Println("\tLW", "$" + opAux, pos, "( $sp )")
						quadA = QuadAssembly{inst: "LW", op1: opAux, op2: strconv.Itoa(pos), op3: "sp"}
						listAssembly.PushBack(quadA)
					} else {
						fmt.Println("\tLW", "$" + opAux, pos, "( $gp )")
						quadA = QuadAssembly{inst: "LW", op1: opAux, op2: strconv.Itoa(pos), op3: "gp"}
						listAssembly.PushBack(quadA)
					}
					fmt.Println("\tDIVI", "$" + op1, "$" + opAux, op2)
					quadA = QuadAssembly{inst: "DIVI", op1: op1, op2: opAux, op3: op2}
					listAssembly.PushBack(quadA)
				}
				
			} else if(err2 == nil) {
				if(selfQuad.Op2.Type == "temp"){
					fmt.Println("\tDIVI", "$" + op1, "$" + op2, op3)
					quadA = QuadAssembly{inst: "DIVI", op1: op1, op2: op2, op3: op3}
					listAssembly.PushBack(quadA)
				} else {
					opAux := GenAssign() 
					pos = m[op2].posicao
					if(m[op2].scope != ""){
						fmt.Println("\tLW", "$" + opAux, pos, "( $sp )")
						quadA = QuadAssembly{inst: "LW", op1: opAux, op2: strconv.Itoa(pos), op3: "sp"}
						listAssembly.PushBack(quadA)
					} else {
						fmt.Println("\tLW", "$" + opAux, pos, "( $gp )")
						quadA = QuadAssembly{inst: "LW", op1: opAux, op2: strconv.Itoa(pos), op3: "gp"}
						listAssembly.PushBack(quadA)
					}
					fmt.Println("\tDIVI", "$" + op1, "$" + opAux, op3)
					quadA = QuadAssembly{inst: "DIVI", op1: op1, op2: opAux, op3: op3}
					listAssembly.PushBack(quadA)
				}
			} else {
				var opAux1 string
				var opAux2 string 

				if(selfQuad.Op2.Type != "temp"){
					opAux1 = GenAssign()
					pos = m[op2].posicao
					if(m[op2].scope != ""){
						fmt.Println("\tLW", "$" + opAux1, pos, "( $sp )")
						quadA = QuadAssembly{inst: "LW", op1: opAux1, op2: strconv.Itoa(pos), op3: "sp"}
						listAssembly.PushBack(quadA)
					} else {
						fmt.Println("\tLW", "$" + opAux1, pos, "( $gp )")
						quadA = QuadAssembly{inst: "LW", op1: opAux1, op2: strconv.Itoa(pos), op3: "gp"}
						listAssembly.PushBack(quadA)
					}
				} else {
					opAux1 = op2
				}

				if(selfQuad.Op3.Type != "temp"){
					opAux2 = GenAssign()
					pos = m[op3].posicao
					if(m[op3].scope != ""){
						fmt.Println("\tLW", "$" + opAux2, pos, "( $sp )")
						quadA = QuadAssembly{inst: "LW", op1: opAux2, op2: strconv.Itoa(pos), op3: "sp"}
						listAssembly.PushBack(quadA)
					} else {
						fmt.Println("\tLW", "$" + opAux2, pos, "( $gp )")
						quadA = QuadAssembly{inst: "LW", op1: opAux2, op2: strconv.Itoa(pos), op3: "gp"}
						listAssembly.PushBack(quadA)
					}
				} else {
					opAux2 = op3
				}
				fmt.Println("\tDIV", "$" + op1, "$" + opAux1, "$" + opAux2)
				quadA = QuadAssembly{inst: "DIV", op1: op1, op2: opAux1, op3: opAux2}
				listAssembly.PushBack(quadA)
			}
		case "smaller_and_equal":
			_, err := strconv.Atoi(op3)

			var pos int 

			if(err == nil){
				opAux := GenValue()
				

				fmt.Println("\tLOADI", "$" + opAux, op3)
				quadA = QuadAssembly{inst: "LOADI", op1: opAux, op2: op3}
				listAssembly.PushBack(quadA)

				if(selfQuad.Op2.Type == "temp"){
					fmt.Println("\tSMLEQ", "$" + op1, "$" + op2, "$" + opAux)
					quadA = QuadAssembly{inst: "SMLEQ", op1: op1, op2: op2, op3: opAux}
					listAssembly.PushBack(quadA)
				} else {
					opAux2 := GenAssign()
					pos = m[op2].posicao
					fmt.Println("\tLW", "$" + opAux2, pos/*posicao pilha*/, "( $sp )")
					quadA = QuadAssembly{inst: "LW", op1: opAux2, op2: strconv.Itoa(pos), op3: "sp"}
					listAssembly.PushBack(quadA)

					fmt.Println("\tSMLEQ", "$" + op1, "$" + opAux2, "$" + opAux)
					quadA = QuadAssembly{inst: "SMLEQ", op1: op1, op2: opAux2, op3: opAux}
					listAssembly.PushBack(quadA)
				}
				
			} else {
				if(selfQuad.Op3.Type == "temp"){
					opAux := GenAssign()

					if(selfQuad.Op2.Type == "temp"){
						fmt.Println("\tSMLEQ", "$" + op1, "$" + op2, "$" + op3)
						quadA = QuadAssembly{inst: "SMLEQ", op1: op1, op2: op2, op3: op3}
						listAssembly.PushBack(quadA)
					} else {
						pos = m[op2].posicao

						fmt.Println("\tLW", "$" + opAux, pos/*posicao pilha*/, "( $sp )")
						quadA = QuadAssembly{inst: "LW", op1: opAux, op2: strconv.Itoa(pos), op3: "sp"}
						listAssembly.PushBack(quadA)

						fmt.Println("\tSMLEQ", "$" + op1, "$" + opAux, "$" + op3)
						quadA = QuadAssembly{inst: "SMLEQ", op1: op1, op2: opAux, op3: op3}
						listAssembly.PushBack(quadA)
					}
					
				} else {
					opAux := GenAssign()
					opAux2 := GenAssign()

					pos = m[op3].posicao
					fmt.Println("\tLW", "$" + opAux2, pos/*posicao pilha*/, "( $sp )") // op2
					quadA = QuadAssembly{inst: "LW", op1: opAux2, op2: strconv.Itoa(pos), op3: "sp"}
					listAssembly.PushBack(quadA)

					if(selfQuad.Op2.Type == "temp"){
						fmt.Println("\tSMLEQ", "$" + op1, "$" + op2, "$" + opAux2)
						quadA = QuadAssembly{inst: "SMLEQ", op1: op1, op2: op2, op3: opAux2}
						listAssembly.PushBack(quadA)
					}else {
						pos = m[op2].posicao
						fmt.Println("\tLW", "$" + opAux, pos/*posicao pilha*/, "( $sp )") // op1
						quadA = QuadAssembly{inst: "LW", op1: opAux, op2: strconv.Itoa(pos), op3: "sp"}
						listAssembly.PushBack(quadA)

						fmt.Println("\tSMLEQ", "$" + op1, "$" + opAux, "$" + opAux2)
						quadA = QuadAssembly{inst: "SMLEQ", op1: op1, op2: opAux, op3: opAux2}
						listAssembly.PushBack(quadA)
					}
				
				}
			}
		case "smaller":
			_, err := strconv.Atoi(op3)

			var pos int 

			if(err == nil){
				opAux := GenValue()
				

				fmt.Println("\tLOADI", "$" + opAux, op3)
				quadA = QuadAssembly{inst: "LOADI", op1: opAux, op2: op3}
				listAssembly.PushBack(quadA)

				if(selfQuad.Op2.Type == "temp"){
					fmt.Println("\tSML", "$" + op1, "$" + op2, "$" + opAux)
					quadA = QuadAssembly{inst: "SML", op1: op1, op2: op2, op3: opAux}
					listAssembly.PushBack(quadA)
				} else {
					opAux2 := GenAssign()
					pos = m[op2].posicao
					fmt.Println("\tLW", "$" + opAux2, pos/*posicao pilha*/, "( $sp )")
					quadA = QuadAssembly{inst: "LW", op1: opAux2, op2: strconv.Itoa(pos), op3: "sp"}
					listAssembly.PushBack(quadA)

					fmt.Println("\tSML", "$" + op1, "$" + opAux2, "$" + opAux)
					quadA = QuadAssembly{inst: "SML", op1: op1, op2: opAux2, op3: opAux}
					listAssembly.PushBack(quadA)
				}
				
			} else {
				if(selfQuad.Op3.Type == "temp"){
					opAux := GenAssign()

					if(selfQuad.Op2.Type == "temp"){
						fmt.Println("\tSML", "$" + op1, "$" + op2, "$" + op3)
						quadA = QuadAssembly{inst: "SML", op1: op1, op2: op2, op3: op3}
						listAssembly.PushBack(quadA)
					} else {
						pos = m[op2].posicao

						fmt.Println("\tLW", "$" + opAux, pos/*posicao pilha*/, "( $sp )")
						quadA = QuadAssembly{inst: "LW", op1: opAux, op2: strconv.Itoa(pos), op3: "sp"}
						listAssembly.PushBack(quadA)

						fmt.Println("\tSML", "$" + op1, "$" + opAux, "$" + op3)
						quadA = QuadAssembly{inst: "SML", op1: op1, op2: opAux, op3: op3}
						listAssembly.PushBack(quadA)
					}
					
				} else {
					opAux := GenAssign()
					opAux2 := GenAssign()

					pos = m[op3].posicao
					fmt.Println("\tLW", "$" + opAux2, pos/*posicao pilha*/, "( $sp )") // op2
					quadA = QuadAssembly{inst: "LW", op1: opAux2, op2: strconv.Itoa(pos), op3: "sp"}
					listAssembly.PushBack(quadA)

					if(selfQuad.Op2.Type == "temp"){
						fmt.Println("\tSML", "$" + op1, "$" + op2, "$" + opAux2)
						quadA = QuadAssembly{inst: "SML", op1: op1, op2: op2, op3: opAux2}
						listAssembly.PushBack(quadA)
					}else {
						pos = m[op2].posicao
						fmt.Println("\tLW", "$" + opAux, pos/*posicao pilha*/, "( $sp )") // op1
						quadA = QuadAssembly{inst: "LW", op1: opAux, op2: strconv.Itoa(pos), op3: "sp"}
						listAssembly.PushBack(quadA)

						fmt.Println("\tSML", "$" + op1, "$" + opAux, "$" + opAux2)
						quadA = QuadAssembly{inst: "SML", op1: op1, op2: opAux, op3: opAux2}
						listAssembly.PushBack(quadA)
					}
				
				}
			}
		case "greater_and_equal":
			_, err := strconv.Atoi(op3)

			var pos int 

			if(err == nil){
				opAux := GenValue()
				

				fmt.Println("\tLOADI", "$" + opAux, op3)
				quadA = QuadAssembly{inst: "LOADI", op1: opAux, op2: op3}
				listAssembly.PushBack(quadA)

				if(selfQuad.Op2.Type == "temp"){
					fmt.Println("\tLGREQ", "$" + op1, "$" + op2, "$" + opAux)
					quadA = QuadAssembly{inst: "LGREQ", op1: op1, op2: op2, op3: opAux}
					listAssembly.PushBack(quadA)
				} else {
					opAux2 := GenAssign()
					pos = m[op2].posicao
					fmt.Println("\tLW", "$" + opAux2, pos/*posicao pilha*/, "( $sp )")
					quadA = QuadAssembly{inst: "LW", op1: opAux2, op2: strconv.Itoa(pos), op3: "sp"}
					listAssembly.PushBack(quadA)

					fmt.Println("\tLGREQ", "$" + op1, "$" + opAux2, "$" + opAux)
					quadA = QuadAssembly{inst: "LGREQ", op1: op1, op2: opAux2, op3: opAux}
					listAssembly.PushBack(quadA)
				}
				
			} else {
				if(selfQuad.Op3.Type == "temp"){
					opAux := GenAssign()

					if(selfQuad.Op2.Type == "temp"){
						fmt.Println("\tLGREQ", "$" + op1, "$" + op2, "$" + op3)
						quadA = QuadAssembly{inst: "LGREQ", op1: op1, op2: op2, op3: op3}
						listAssembly.PushBack(quadA)
					} else {
						pos = m[op2].posicao

						fmt.Println("\tLW", "$" + opAux, pos/*posicao pilha*/, "( $sp )")
						quadA = QuadAssembly{inst: "LW", op1: opAux, op2: strconv.Itoa(pos), op3: "sp"}
						listAssembly.PushBack(quadA)

						fmt.Println("\tLGREQ", "$" + op1, "$" + opAux, "$" + op3)
						quadA = QuadAssembly{inst: "LGREQ", op1: op1, op2: opAux, op3: op3}
						listAssembly.PushBack(quadA)
					}
					
				} else {
					opAux := GenAssign()
					opAux2 := GenAssign()

					pos = m[op3].posicao
					fmt.Println("\tLW", "$" + opAux2, pos/*posicao pilha*/, "( $sp )") // op2
					quadA = QuadAssembly{inst: "LW", op1: opAux2, op2: strconv.Itoa(pos), op3: "sp"}
					listAssembly.PushBack(quadA)

					if(selfQuad.Op2.Type == "temp"){
						fmt.Println("\tLGREQ", "$" + op1, "$" + op2, "$" + opAux2)
						quadA = QuadAssembly{inst: "LGREQ", op1: op1, op2: op2, op3: opAux2}
						listAssembly.PushBack(quadA)
					}else {
						pos = m[op2].posicao
						fmt.Println("\tLW", "$" + opAux, pos/*posicao pilha*/, "( $sp )") // op1
						quadA = QuadAssembly{inst: "LW", op1: opAux, op2: strconv.Itoa(pos), op3: "sp"}
						listAssembly.PushBack(quadA)

						fmt.Println("\tLGREQ", "$" + op1, "$" + opAux, "$" + opAux2)
						quadA = QuadAssembly{inst: "LGREQ", op1: op1, op2: opAux, op3: opAux2}
						listAssembly.PushBack(quadA)
					}
				
				}
			}
		case "greater":
			_, err := strconv.Atoi(op3)

			var pos int 

			if(err == nil){
				opAux := GenValue()
				

				fmt.Println("\tLOADI", "$" + opAux, op3)
				quadA = QuadAssembly{inst: "LOADI", op1: opAux, op2: op3}
				listAssembly.PushBack(quadA)

				if(selfQuad.Op2.Type == "temp"){
					fmt.Println("\tLGR", "$" + op1, "$" + op2, "$" + opAux)
					quadA = QuadAssembly{inst: "LGR", op1: op1, op2: op2, op3: opAux}
					listAssembly.PushBack(quadA)
				} else {
					opAux2 := GenAssign()
					pos = m[op2].posicao
					fmt.Println("\tLW", "$" + opAux2, pos/*posicao pilha*/, "( $sp )")
					quadA = QuadAssembly{inst: "LW", op1: opAux2, op2: strconv.Itoa(pos), op3: "sp"}
					listAssembly.PushBack(quadA)

					fmt.Println("\tLGR", "$" + op1, "$" + opAux2, "$" + opAux)
					quadA = QuadAssembly{inst: "LGR", op1: op1, op2: opAux2, op3: opAux}
					listAssembly.PushBack(quadA)
				}
				
			} else {
				if(selfQuad.Op3.Type == "temp"){
					opAux := GenAssign()

					if(selfQuad.Op2.Type == "temp"){
						fmt.Println("\tLGR", "$" + op1, "$" + op2, "$" + op3)
						quadA = QuadAssembly{inst: "LGR", op1: op1, op2: op2, op3: op3}
						listAssembly.PushBack(quadA)
					} else {
						pos = m[op2].posicao

						fmt.Println("\tLW", "$" + opAux, pos/*posicao pilha*/, "( $sp )")
						quadA = QuadAssembly{inst: "LW", op1: opAux, op2: strconv.Itoa(pos), op3: "sp"}
						listAssembly.PushBack(quadA)

						fmt.Println("\tLGR", "$" + op1, "$" + opAux, "$" + op3)
						quadA = QuadAssembly{inst: "LGR", op1: op1, op2: opAux, op3: op3}
						listAssembly.PushBack(quadA)
					}
					
				} else {
					opAux := GenAssign()
					opAux2 := GenAssign()

					pos = m[op3].posicao
					fmt.Println("\tLW", "$" + opAux2, pos/*posicao pilha*/, "( $sp )") // op2
					quadA = QuadAssembly{inst: "LW", op1: opAux2, op2: strconv.Itoa(pos), op3: "sp"}
					listAssembly.PushBack(quadA)

					if(selfQuad.Op2.Type == "temp"){
						fmt.Println("\tLGR", "$" + op1, "$" + op2, "$" + opAux2)
						quadA = QuadAssembly{inst: "LGR", op1: op1, op2: op2, op3: opAux2}
						listAssembly.PushBack(quadA)
					}else {
						pos = m[op2].posicao
						fmt.Println("\tLW", "$" + opAux, pos/*posicao pilha*/, "( $sp )") // op1
						quadA = QuadAssembly{inst: "LW", op1: opAux, op2: strconv.Itoa(pos), op3: "sp"}
						listAssembly.PushBack(quadA)

						fmt.Println("\tLGR", "$" + op1, "$" + opAux, "$" + opAux2)
						quadA = QuadAssembly{inst: "LGR", op1: op1, op2: opAux, op3: opAux2}
						listAssembly.PushBack(quadA)
					}
				
				}
			}
		case "equal":
			_, err := strconv.Atoi(op3)

			var pos int 

			if(err == nil){
				opAux := GenValue()
				

				fmt.Println("\tLOADI", "$" + opAux, op3)
				quadA = QuadAssembly{inst: "LOADI", op1: opAux, op2: op3}
				listAssembly.PushBack(quadA)

				if(selfQuad.Op2.Type == "temp"){
					fmt.Println("\tEQ", "$" + op1, "$" + op2, "$" + opAux)
					quadA = QuadAssembly{inst: "EQ", op1: op1, op2: op2, op3: opAux}
					listAssembly.PushBack(quadA)
				} else {
					opAux2 := GenAssign()
					pos = m[op2].posicao
					fmt.Println("\tLW", "$" + opAux2, pos/*posicao pilha*/, "( $sp )")
					quadA = QuadAssembly{inst: "LW", op1: opAux2, op2: strconv.Itoa(pos), op3: "sp"}
					listAssembly.PushBack(quadA)

					fmt.Println("\tEQ", "$" + op1, "$" + opAux2, "$" + opAux)
					quadA = QuadAssembly{inst: "EQ", op1: op1, op2: opAux2, op3: opAux}
					listAssembly.PushBack(quadA)
				}
				
			} else {
				if(selfQuad.Op3.Type == "temp"){
					opAux := GenAssign()

					if(selfQuad.Op2.Type == "temp"){
						fmt.Println("\tEQ", "$" + op1, "$" + op2, "$" + op3)
						quadA = QuadAssembly{inst: "EQ", op1: op1, op2: op2, op3: op3}
						listAssembly.PushBack(quadA)
					} else {
						pos = m[op2].posicao

						fmt.Println("\tLW", "$" + opAux, pos/*posicao pilha*/, "( $sp )")
						quadA = QuadAssembly{inst: "LW", op1: opAux, op2: strconv.Itoa(pos), op3: "sp"}
						listAssembly.PushBack(quadA)

						fmt.Println("\tEQ", "$" + op1, "$" + opAux, "$" + op3)
						quadA = QuadAssembly{inst: "EQ", op1: op1, op2: opAux, op3: op3}
						listAssembly.PushBack(quadA)
					}
					
				} else {
					opAux := GenAssign()
					opAux2 := GenAssign()

					pos = m[op3].posicao
					fmt.Println("\tLW", "$" + opAux2, pos/*posicao pilha*/, "( $sp )") // op2
					quadA = QuadAssembly{inst: "LW", op1: opAux2, op2: strconv.Itoa(pos), op3: "sp"}
					listAssembly.PushBack(quadA)

					if(selfQuad.Op2.Type == "temp"){
						fmt.Println("\tEQ", "$" + op1, "$" + op2, "$" + opAux2)
						quadA = QuadAssembly{inst: "EQ", op1: op1, op2: op2, op3: opAux2}
						listAssembly.PushBack(quadA)
					}else {
						pos = m[op2].posicao
						fmt.Println("\tLW", "$" + opAux, pos/*posicao pilha*/, "( $sp )") // op1
						quadA = QuadAssembly{inst: "LW", op1: opAux, op2: strconv.Itoa(pos), op3: "sp"}
						listAssembly.PushBack(quadA)

						fmt.Println("\tEQ", "$" + op1, "$" + opAux, "$" + opAux2)
						quadA = QuadAssembly{inst: "EQ", op1: op1, op2: opAux, op3: opAux2}
						listAssembly.PushBack(quadA)
					}
				
				}
			}
		case "jump_if_false":
			fmt.Println("\tBNE", "$" + op1, "$RTRUE", op2)
			quadA = QuadAssembly{inst: "BNE", op1: op1, op2: "RTRUE", op3: op2}
			listAssembly.PushBack(quadA)
		case "vectortotemp":
			var pos int
			_, err := strconv.Atoi(op3)
			if(err == nil){
				opAux := GenAddress()
				opAux2 := GenValue()

				if(m[op2].auxparam){
					fmt.Println("\tADDI", "$" + opAux, "$" + m[op2].auxreg, op3)
					quadA = QuadAssembly{inst: "ADDI", op1: opAux, op2: m[op2].auxreg, op3: op3}
					listAssembly.PushBack(quadA)
				} else {
					fmt.Println("\tLOADI", "$" + opAux2, op3)
					quadA = QuadAssembly{inst: "LOADI", op1: opAux2, op2: op3}
					listAssembly.PushBack(quadA)

					pos = m[op2].posicao
					fmt.Println("\tADDI", "$" + opAux, "$" + opAux2, pos)
					quadA = QuadAssembly{inst: "ADDI", op1: opAux, op2: opAux2, op3: strconv.Itoa(pos)}
					listAssembly.PushBack(quadA)
				}

				fmt.Println("\tLWR", "$" + op1, "$" + opAux, "( $sp )"/*op2*/)
				quadA = QuadAssembly{inst: "LWR", op1: op1, op2: opAux, op3: "sp"}
				listAssembly.PushBack(quadA)
			} else{
				if(selfQuad.Op3.Type == "temp"){
					opAux := GenAddress()
					if(m[op2].auxparam){
						fmt.Println("\tADD", "$" + opAux, "$" + op3, "$" + m[op2].auxreg)
						quadA = QuadAssembly{inst: "ADD", op1: opAux, op2: op3, op3: m[op2].auxreg}
						listAssembly.PushBack(quadA)
					}else {
						pos = m[op2].posicao
						fmt.Println("\tADDI", "$" + opAux, "$" + op3, pos)
						quadA = QuadAssembly{inst: "ADDI", op1: opAux, op2: op3, op3: strconv.Itoa(pos)}
						listAssembly.PushBack(quadA)
					}
					fmt.Println("\tLWR", "$" + op1, "$" + opAux, "( $sp )"/*op2*/)
					quadA = QuadAssembly{inst: "LWR", op1: op1, op2: opAux, op3: "sp"}
					listAssembly.PushBack(quadA)
				} else {
					opAux := GenAssign()
					opAux2 := GenAddress()
					pos = m[op3].posicao
					fmt.Println("\tLW", "$" + opAux, pos/*posicao pilha da variavel x*/, "( $sp )")
					quadA = QuadAssembly{inst: "LW", op1: opAux, op2: strconv.Itoa(pos), op3: "sp"}
					listAssembly.PushBack(quadA)

					if(m[op2].auxparam){
						fmt.Println("\tADD", "$" + opAux2, "$" + opAux, "$" + m[op2].auxreg)
						quadA = QuadAssembly{inst: "ADD", op1: opAux2, op2: opAux, op3: m[op2].auxreg}
						listAssembly.PushBack(quadA)
					} else {
						pos = m[op2].posicao
						fmt.Println("\tADDI", "$" + opAux2, "$" + opAux, pos)
						quadA = QuadAssembly{inst: "ADDI", op1: opAux2, op2: opAux, op3: strconv.Itoa(pos)}
						listAssembly.PushBack(quadA)
					}

					fmt.Println("\tLWR", "$" + op1, "$" + opAux2, /*posicao na pilha*/"( $sp )"/*op2*/)
					quadA = QuadAssembly{inst: "LWR", op1: op1, op2: opAux2, op3: "sp"}
					listAssembly.PushBack(quadA)
				}
			}
		case "temptovector":
			var pos int
			_, err := strconv.Atoi(op2)
			if(err == nil) {
				opAux := GenAddress()
				opAux2 := GenValue()

				if(m[op1].auxparam){
					fmt.Println("\tADDI", "$" + opAux, "$" + m[op1].auxreg, op2)
					quadA = QuadAssembly{inst: "ADDI", op1: opAux, op2: m[op1].auxreg, op3: op2}
					listAssembly.PushBack(quadA)
				} else {
					fmt.Println("\tLOADI", "$" + opAux2, op2)
					quadA = QuadAssembly{inst: "LOADI", op1: opAux2, op2: op2}
					listAssembly.PushBack(quadA)
					
					pos = m[op1].posicao
					fmt.Println("\tADDI", "$" + opAux, "$" + opAux2, pos)
					quadA = QuadAssembly{inst: "ADDI", op1: opAux, op2: opAux2, op3: strconv.Itoa(pos)}
					listAssembly.PushBack(quadA)
				}

				fmt.Println("\tSWR", "$" + opAux, "( $sp )"/*op1*/, "$" + op3)
				quadA = QuadAssembly{inst: "SWR", op1: opAux, op2: "sp", op3: op3}
				listAssembly.PushBack(quadA)
			} else {
				if(selfQuad.Op2.Type == "temp") {
					opAux := GenAddress()
					pos = m[op1].posicao

					if(m[op1].auxparam){
						fmt.Println("\tADD", "$" + opAux, "$" + m[op1].auxreg, "$" + op2)
						quadA = QuadAssembly{inst: "ADD", op1: opAux, op2: m[op1].auxreg, op3: op2}
						listAssembly.PushBack(quadA)
					} else {
						fmt.Println("\tADDI", "$" + opAux, "$" + op2, pos)
						quadA = QuadAssembly{inst: "ADDI", op1: opAux, op2: op2, op3: strconv.Itoa(pos)}
						listAssembly.PushBack(quadA)
					}
					
					fmt.Println("\tSWR", "$" + opAux, /*posicao pilha*/"( $sp )"/*op1*/, "$" + op3)
					quadA = QuadAssembly{inst: "SWR", op1: opAux, op2: "sp", op3: op3}
					listAssembly.PushBack(quadA)
				} else {
					opAux := GenAssign()
					opAux2 := GenAddress()
					pos = m[op2].posicao

					fmt.Println("\tLW", "$" + opAux, pos/*posicao pilha da variavel x*/, "( $sp )") // dado do op2
					quadA = QuadAssembly{inst: "LW", op1: opAux, op2: strconv.Itoa(pos), op3: "sp"}
					listAssembly.PushBack(quadA)

					if(m[op1].auxparam){
						fmt.Println("\tADD", "$" + opAux2, "$" + opAux, "$" + m[op1].auxreg)
						quadA = QuadAssembly{inst: "ADD", op1: opAux2, op2: opAux, op3: m[op1].auxreg}
						listAssembly.PushBack(quadA)
					} else {
						pos = m[op1].posicao
						fmt.Println("\tADDI", "$" + opAux2, "$" + opAux, pos)
						quadA = QuadAssembly{inst: "ADDI", op1: opAux2, op2: opAux, op3: strconv.Itoa(pos)}
						listAssembly.PushBack(quadA)
					}


					fmt.Println("\tSWR", "$" + opAux2 , /*posicao na pilha*/"( $sp )"/*op1*/, "$" + op3)
					quadA = QuadAssembly{inst: "SWR", op1: opAux2, op2: "sp", op3: op3}
					listAssembly.PushBack(quadA)
				}
			}
		case "return_int":
			var pos int
			opAux := GenAddress()
			opAux2 := GenAssign()

			_, err := strconv.Atoi(op1)
			if(err == nil){
				fmt.Println("\tLOADI", "$dl", op1)
				quadA = QuadAssembly{inst: "LOADI", op1: "dl", op2: op1}
				listAssembly.PushBack(quadA)

				fmt.Println("\tSUBI", "$ap", "$ap", "1")
				quadA = QuadAssembly{inst: "SUBI", op1: "ap", op2: "ap", op3: "1"}
				listAssembly.PushBack(quadA)

				fmt.Println("\tJR", "$" + opAux)
				quadA = QuadAssembly{inst: "JR", op1: opAux}
				listAssembly.PushBack(quadA)
			} else {
				if(selfQuad.Op1.Type == "temp"){
					pos = m[op1].posicao
					
					fmt.Println("\tMOV", "$dl", "$" + op1)
					quadA = QuadAssembly{inst: "MOV", op1: "dl", op2: op1}
					listAssembly.PushBack(quadA)

					fmt.Println("\tLW", "$" + opAux, "0( $ap )")
					quadA = QuadAssembly{inst: "LW", op1: opAux, op2: "0", op3: "ap"}
					listAssembly.PushBack(quadA)

					fmt.Println("\tSUBI", "$ap", "$ap", "1")
					quadA = QuadAssembly{inst: "SUBI", op1: "ap", op2: "ap", op3: "1"}
					listAssembly.PushBack(quadA)

					fmt.Println("\tJR", "$" + opAux)
					quadA = QuadAssembly{inst: "JR", op1: opAux}
					listAssembly.PushBack(quadA)
				} else {
					pos = m[op1].posicao
					fmt.Println("\tLW", "$" + opAux2, pos, "( $sp )")
					quadA = QuadAssembly{inst: "LW", op1: opAux2, op2: strconv.Itoa(pos), op3: "sp"}
					listAssembly.PushBack(quadA)

					fmt.Println("\tMOV", "$dl", "$" + opAux2)
					quadA = QuadAssembly{inst: "MOV", op1: "dl", op2: opAux2}
					listAssembly.PushBack(quadA)

					fmt.Println("\tLW", "$" + opAux, "0( $ap )")
					quadA = QuadAssembly{inst: "LW", op1: opAux, op2: "0", op3: "ap"}
					listAssembly.PushBack(quadA)

					fmt.Println("\tSUBI", "$ap", "$ap", "1")
					quadA = QuadAssembly{inst: "SUBI", op1: "ap", op2: "ap", op3: "1"}
					listAssembly.PushBack(quadA)

					fmt.Println("\tJR", "$" + opAux)
					quadA = QuadAssembly{inst: "JR", op1: opAux}
					listAssembly.PushBack(quadA)
				}
			}
			
		case "return_void":
			opAux := GenAddress()

			fmt.Println("\tLW", "$" + opAux, "0( $ap )")
			quadA = QuadAssembly{inst: "LW", op1: opAux, op2: "0", op3: "ap"}
			listAssembly.PushBack(quadA)

			fmt.Println("\tSUBI", "$ap", "$ap", "1")
			quadA = QuadAssembly{inst: "SUBI", op1: "ap", op2: "ap", op3: "1"}
			listAssembly.PushBack(quadA)

			fmt.Println("\tJR", "$" + opAux)
			quadA = QuadAssembly{inst: "JR", op1: opAux}
			listAssembly.PushBack(quadA)
		case "call":
			fmt.Println("\tADDI", "$ap", "$ap", "1") // $ap começa no -1
			quadA = QuadAssembly{inst: "ADDI", op1: "ap", op2: "ap", op3: "1"}
			listAssembly.PushBack(quadA)

			fmt.Println("\tJAL", op1) /*pula pro op1 e armazena o endereço de instrução na pilha*/
			quadA = QuadAssembly{inst: "JAL", op1: op1}
			listAssembly.PushBack(quadA)

			fmt.Println("\tMOV", "$" + op2, "$dl")
			quadA = QuadAssembly{inst: "MOV", op1: op2, op2: "dl"}
			listAssembly.PushBack(quadA)

			fmt.Println("\tADDI", "$sp", "$sp", nVariable[op1]) /*numero de variaveis que tinha na funcao*///)
			quadA = QuadAssembly{inst: "ADDI", op1: "sp", op2: "sp", op3: strconv.Itoa(nVariable[op1])}
			listAssembly.PushBack(quadA)
		case "output":
			_, err := strconv.Atoi(op1)
			var opAux2 string
			var pos int
			if(err == nil){
				opAux := GenValue()
				
				fmt.Println("\tLOADI", "$" + opAux, op1)
				quadA = QuadAssembly{inst: "LOADI", op1: opAux, op2: op1}
				opAux2 = opAux
				listAssembly.PushBack(quadA)
			} else if (selfQuad.Op1.Type == "temp"){
				opAux2 = op1
			} else {
				opAux := GenAssign()
				pos = m[op1].posicao
				fmt.Println("\tLW", "$" + opAux, pos/*posicao pilha*/, "( $sp )")
				quadA = QuadAssembly{inst: "LW", op1: opAux, op2: strconv.Itoa(pos), op3: "sp"}
				listAssembly.PushBack(quadA)
				opAux2 = opAux
			}
			fmt.Println("\tOPT", "$" + opAux2)
			quadA = QuadAssembly{inst: "OPT", op1: opAux2}
			listAssembly.PushBack(quadA)
		case "input":
			fmt.Println("\tIPT", "$" + op1)
			quadA = QuadAssembly{inst: "IPT", op1: op1}
			listAssembly.PushBack(quadA)
		case "argument":
			_, err := strconv.Atoi(op1)
			var pos int 
			if(err == nil){
				listAux := ml[op2].Front()
				typeAssert, _ := listAux.Value.(string)
				fmt.Println("\tLOADI", "$" + typeAssert, op1)
				quadA = QuadAssembly{inst: "LOADI", op1: typeAssert, op2: op1}
				listAssembly.PushBack(quadA)
				ml[op2].Remove(listAux)
				ml[op2].PushBack(typeAssert)
			} else {
				if(selfQuad.Op1.Type == "temp"){
					listAux := ml[op2].Front()
					typeAssert, _ := listAux.Value.(string)
					fmt.Println("\tMOV", "$", typeAssert, "$" + op1)
					quadA = QuadAssembly{inst: "MOV", op1: typeAssert, op2: op1}
					listAssembly.PushBack(quadA)
					ml[op2].Remove(listAux)
					ml[op2].PushBack(typeAssert)
				} else if (m[op1].auxvector){
					listAux := ml[op2].Front()
					typeAssert, _ := listAux.Value.(string)
					pos = m[op1].posicao
					fmt.Println("\tLOADI", "$", typeAssert, pos)
					quadA = QuadAssembly{inst: "LOADI", op1: typeAssert, op2: strconv.Itoa(pos)}
					listAssembly.PushBack(quadA)
					ml[op2].Remove(listAux)
					ml[op2].PushBack(typeAssert)
				} else {
					listAux := ml[op2].Front()
					typeAssert, _ := listAux.Value.(string)
					pos = m[op1].posicao
					fmt.Println("\tLW", "$", typeAssert, pos, "( $sp )")
					quadA = QuadAssembly{inst: "LW", op1: typeAssert, op2: strconv.Itoa(pos), op3: "sp"}
					listAssembly.PushBack(quadA)
					ml[op2].Remove(listAux)
					ml[op2].PushBack(typeAssert)
				}
			}
		case "param":
			opAux := GenParam()
			pos := m[op1].posicao 

			if(m[op1].auxvector){
				fmt.Println("\tADDI", "$" + opAux, "$" + opAux, Nvar)
				quadA = QuadAssembly{inst: "ADDI", op1: opAux, op2: opAux, op3: strconv.Itoa(Nvar)}
				auxx := VarPos{
					posicao: m[op1].posicao,
					memspace: m[op1].memspace,
					scope: m[op1].scope,
					auxvector: m[op1].auxvector,
					auxreg: opAux,
					auxparam: true}
				m[op1] = auxx
				listAssembly.PushBack(quadA)
				ml[op2].PushBack(opAux)
			} else {
				fmt.Println("\tSW", pos, "( $sp )", "$" + opAux)
				quadA = QuadAssembly{inst: "SW", op1: strconv.Itoa(pos), op2: "sp", op3: opAux}
				listAssembly.PushBack(quadA)
				ml[op2].PushBack(opAux)
			}
			
		}
	}
	fmt.Println("\tHALT")
	quadA = QuadAssembly{inst: "HALT"}
	listAssembly.PushBack(quadA)
	fmt.Println("|--------------------------------------------------------------------------------------|")

	count := 0

	var funcAndLab map[string]int
	funcAndLab = make(map[string]int)

	for quad22_ := listAssembly.Front(); quad22_ != nil; quad22_ = quad22_.Next() {
		selfQuad22, _ := quad22_.Value.(QuadAssembly)
		if(selfQuad22.inst == "FUNCTION" || selfQuad22.inst == "LABEL") {
			funcAndLab[selfQuad22.op1] = count
		}
		count += 1
	}
	BinaryGen(listAssembly, funcAndLab)
}

func CountNVar(scope_name string, m map[string]VarPos) int {
	count := 0
	for _, symbol_ := range table {
		if(symbol_.scope == scope_name){
			/*m[symbol_.ID]*/aux := VarPos{
				posicao: count,
				scope: scope_name}
			if(symbol_.vector == "yes"){
				aux.memspace = symbol_.spacemem
				aux.auxvector = true
				m[symbol_.ID] = aux
				if(symbol_.spacemem == 0){
					count += 1
				} else {
					count += symbol_.spacemem
				}
			} else {
				aux.auxvector = false
				m[symbol_.ID] = aux
				count += 1
			}
		}
	}
	return count
}

func CountVari(m map[string]int){
	for _, symbol_ := range table{
		if(symbol_.vector == "yes"){
			if(symbol_.spacemem == 0){
				m[symbol_.scope] += 1
			} else {
				m[symbol_.scope] += symbol_.spacemem	
			}
		} else {
			m[symbol_.scope] += 1 
		}	
	}
}
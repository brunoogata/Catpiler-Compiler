package ast

import (
	"fmt"
	"strconv"
	"container/list"
	"strings"
)

var countRegister int
var parameters map[string]int

func BinaryGen(list_ *list.List, funcAndLab map[string]int) {
	fmt.Println("\n|-----------------------------------Catpiler Binary------------------------------------|")
	countRegister = -1
	countLines := 0
	
	var registers map[string]int
	registers = make(map[string]int)
	registers["dl"] = 27
	registers["sp"] = 28
	registers["gp"] = 29
	registers["ap"] = 30
	registers["RTRUE"] = 31
		
	parameters = make(map[string]int)

	for quad_ := list_.Front(); quad_ != nil; quad_ = quad_.Next() {
		var instruction []string 
		
		selfQuad, _ := quad_.Value.(QuadAssembly)

		switch selfQuad.inst {
		//case ("ADD" || "SUB" || "MULT" || "DIV"):
		//	switch selfQuad.inst{
		case "ADD":
			instruction = append(instruction, "100000")
			instruction = append(instruction, TreatRegisters(/*countRegister, */registers, selfQuad.op1))
			instruction = append(instruction, TreatRegisters(/*countRegister, */registers, selfQuad.op2))
			instruction = append(instruction, TreatRegisters(/*countRegister, */registers, selfQuad.op3))
			instruction = append(instruction, "00000000000")
			s := strings.Join(instruction, "")
			fmt.Println("registradores_instrucao[", countLines, "] = ", "32'b" + s, ";") 
			//fmt.Println(countLines, "ADD", ConvertBin(s[6:11], registers), ConvertBin(s[11:16], registers), ConvertBin(s[16:21], registers))
			countLines += 1
		case "SUB":
			instruction = append(instruction, "100001" )	
			instruction = append(instruction, TreatRegisters(/*countRegister, */registers, selfQuad.op1) )
			instruction = append(instruction, TreatRegisters(/*countRegister, */registers, selfQuad.op2) )
			instruction = append(instruction, TreatRegisters(/*countRegister, */registers, selfQuad.op3) )
			instruction = append(instruction, "00000000000")
			s := strings.Join(instruction, "")

			fmt.Println("registradores_instrucao[", countLines, "] = ", "32'b" + s, ";") 
			//fmt.Println(countLines, "SUB", ConvertBin(s[6:11], registers), ConvertBin(s[11:16], registers), ConvertBin(s[16:21], registers))
			countLines += 1
		case "MULT":
			instruction = append(instruction, "101010" )	
			instruction = append(instruction, TreatRegisters(/*countRegister, */registers, selfQuad.op1) )
			instruction = append(instruction, TreatRegisters(/*countRegister, */registers, selfQuad.op2) )
			instruction = append(instruction, TreatRegisters(/*countRegister, */registers, selfQuad.op3) )
			instruction = append(instruction, "00000000000")
			s := strings.Join(instruction, "")

			fmt.Println("registradores_instrucao[", countLines, "] = ", "32'b" + s, ";") 
			//fmt.Println(countLines, "MULT", ConvertBin(s[6:11], registers), ConvertBin(s[11:16], registers), ConvertBin(s[16:21], registers))
			countLines += 1
		case "DIV":
			instruction = append(instruction, "101011" )	
			instruction = append(instruction, TreatRegisters(/*countRegister, */registers, selfQuad.op1) )
			instruction = append(instruction, TreatRegisters(/*countRegister, */registers, selfQuad.op2) )
			instruction = append(instruction, TreatRegisters(/*countRegister, */registers, selfQuad.op3) )
			instruction = append(instruction, "00000000000")
			s := strings.Join(instruction, "")

			fmt.Println("registradores_instrucao[", countLines, "] = ", "32'b" + s, ";") 
			//fmt.Println(countLines, "DIV", ConvertBin(s[6:11], registers), ConvertBin(s[11:16], registers), ConvertBin(s[16:21], registers))
			countLines += 1
		case "LW":
			instruction = append(instruction, "011111" )
			instruction = append(instruction, TreatRegisters(/*countRegister, */registers, selfQuad.op1) )
			instruction = append(instruction, TreatRegisters(/*countRegister, */registers, selfQuad.op3) )
			aux, _ := strconv.Atoi(selfQuad.op2)
			instruction = append(instruction, ConvertToNBits(aux, 16))
			s := strings.Join(instruction, "")
			
			fmt.Println("registradores_instrucao[", countLines, "] = ", "32'b" + s, ";") 
			//fmt.Println(countLines, "LW", ConvertBin(s[6:11], registers), aux, ConvertBin(s[11:16], registers))
			countLines += 1
		case "SW":		/// verificar
			instruction = append(instruction, "011110" )
			instruction = append(instruction, TreatRegisters(/*countRegister, */registers, selfQuad.op3) )
			instruction = append(instruction, TreatRegisters(/*countRegister, */registers, selfQuad.op2) )
			aux, _ := strconv.Atoi(selfQuad.op1)
			instruction = append(instruction, ConvertToNBits(aux, 16))
			s := strings.Join(instruction, "")
			
			fmt.Println("registradores_instrucao[", countLines, "] = ", "32'b" + s, ";") 
			//fmt.Println(countLines, "SW", aux, ConvertBin(s[11:16], registers), ConvertBin(s[6:11], registers))
			countLines += 1
		case "FUNCTION":
			//funcAndLab[selfQuad.op1] = countLines
			instruction = append(instruction, "000000_00000000000000000000000000")
			s := strings.Join(instruction, "")

			fmt.Println("registradores_instrucao[", countLines, "] = ", "32'b" + s, ";") 
			//fmt.Println(countLines, "FUNCTION", selfQuad.op3, s)
			countLines += 1
		case "LABEL":
			//funcAndLab[selfQuad.op1] = countLines
			instruction = append(instruction, "000000_00000000000000000000000000")
			s := strings.Join(instruction, "")

			fmt.Println("registradores_instrucao[", countLines, "] = ", "32'b" + s, ";") 
			//fmt.Println(countLines, "LABEL", selfQuad.op3, s)
			countLines += 1
		case "ADDI":
			instruction = append(instruction, "010000" )
			instruction = append(instruction, TreatRegisters(/*countRegister, */registers, selfQuad.op1) )
			instruction = append(instruction, TreatRegisters(/*countRegister, */registers, selfQuad.op2) )
			aux, _ := strconv.Atoi(selfQuad.op3)
			instruction = append(instruction, ConvertToNBits(aux, 16))
			s := strings.Join(instruction, "")
			
			fmt.Println("registradores_instrucao[", countLines, "] = ", "32'b" + s, ";") 
			//fmt.Println(countLines, "ADDI", ConvertBin(s[6:11], registers), ConvertBin(s[11:16], registers), aux)
			countLines += 1
		case "SUBI":
			instruction = append(instruction, "010001" )
			instruction = append(instruction, TreatRegisters(/*countRegister, */registers, selfQuad.op1) )
			instruction = append(instruction, TreatRegisters(/*countRegister, */registers, selfQuad.op2) )
			aux, _ := strconv.Atoi(selfQuad.op3)
			instruction = append(instruction, ConvertToNBits(aux, 16))
			s := strings.Join(instruction, "")
			
			fmt.Println("registradores_instrucao[", countLines, "] = ", "32'b" + s, ";") 
			//fmt.Println(countLines, "SUBI", ConvertBin(s[6:11], registers), ConvertBin(s[11:16], registers), aux)
			countLines += 1
		case "SMLEQ":
			instruction = append(instruction, "100100" )
			instruction = append(instruction, TreatRegisters(/*countRegister, */registers, selfQuad.op1) )
			instruction = append(instruction, TreatRegisters(/*countRegister, */registers, selfQuad.op2) )
			instruction = append(instruction, TreatRegisters(/*countRegister, */registers, selfQuad.op3) )
			instruction = append(instruction, "00000000000")
			s := strings.Join(instruction, "")

			fmt.Println("registradores_instrucao[", countLines, "] = ", "32'b" + s, ";") 
			//fmt.Println(countLines, "SMLEQ", ConvertBin(s[6:11], registers), ConvertBin(s[11:16], registers), ConvertBin(s[16:21], registers))
			countLines += 1
		case "SML":
			instruction = append(instruction, "100101" )
			instruction = append(instruction, TreatRegisters(/*countRegister, */registers, selfQuad.op1) )
			instruction = append(instruction, TreatRegisters(/*countRegister, */registers, selfQuad.op2) )
			instruction = append(instruction, TreatRegisters(/*countRegister, */registers, selfQuad.op3) )
			instruction = append(instruction, "00000000000")
			s := strings.Join(instruction, "")

			fmt.Println("registradores_instrucao[", countLines, "] = ", "32'b" + s, ";") 
			//fmt.Println(countLines, "SML", ConvertBin(s[6:11], registers), ConvertBin(s[11:16], registers), ConvertBin(s[16:21], registers))
			countLines += 1
		case "LGREQ":
			instruction = append(instruction, "100110" )
			instruction = append(instruction, TreatRegisters(/*countRegister, */registers, selfQuad.op1) )
			instruction = append(instruction, TreatRegisters(/*countRegister, */registers, selfQuad.op2) )
			instruction = append(instruction, TreatRegisters(/*countRegister, */registers, selfQuad.op3) )
			instruction = append(instruction, "00000000000")
			s := strings.Join(instruction, "")

			fmt.Println("registradores_instrucao[", countLines, "] = ", "32'b" + s, ";") 
			//fmt.Println(countLines, "LGREQ", ConvertBin(s[6:11], registers), ConvertBin(s[11:16], registers), ConvertBin(s[16:21], registers))
			countLines += 1
		case "LGR":
			instruction = append(instruction, "100111" )		
			instruction = append(instruction, TreatRegisters(/*countRegister, */registers, selfQuad.op1) )
			instruction = append(instruction, TreatRegisters(/*countRegister, */registers, selfQuad.op2) )
			instruction = append(instruction, TreatRegisters(/*countRegister, */registers, selfQuad.op3) )
			instruction = append(instruction, "00000000000")
			s := strings.Join(instruction, "")

			fmt.Println("registradores_instrucao[", countLines, "] = ", "32'b" + s, ";") 
			//fmt.Println(countLines, "LGR", ConvertBin(s[6:11], registers), ConvertBin(s[11:16], registers), ConvertBin(s[16:21], registers))
			countLines += 1
		case "EQ":
			instruction = append(instruction, "101110" )		
			instruction = append(instruction, TreatRegisters(/*countRegister, */registers, selfQuad.op1) )
			instruction = append(instruction, TreatRegisters(/*countRegister, */registers, selfQuad.op2) )
			instruction = append(instruction, TreatRegisters(/*countRegister, */registers, selfQuad.op3) )
			instruction = append(instruction, "00000000000")
			s := strings.Join(instruction, "")

			fmt.Println("registradores_instrucao[", countLines, "] = ", "32'b" + s, ";") 
			//fmt.Println(countLines, "EQ", ConvertBin(s[6:11], registers), ConvertBin(s[11:16], registers), ConvertBin(s[16:21], registers))
			countLines += 1
		case "LOADI":
			instruction = append(instruction, "011001" )		
			instruction = append(instruction, TreatRegisters(/*countRegister, */registers, selfQuad.op1) )
			instruction = append(instruction, "00000" )
			aux, _ := strconv.Atoi(selfQuad.op2)
			instruction = append(instruction, ConvertToNBits(aux, 16))
			s := strings.Join(instruction, "")

			fmt.Println("registradores_instrucao[", countLines, "] = ", "32'b" + s, ";") 
			//fmt.Println(countLines, "LOADI", ConvertBin(s[6:11], registers), aux, ConvertBin(s[11:16], registers))
			countLines += 1
		case "LWR":
			instruction = append(instruction, "011101" )	
			instruction = append(instruction, TreatRegisters(/*countRegister, */registers, selfQuad.op1) )
			instruction = append(instruction, TreatRegisters(/*countRegister, */registers, selfQuad.op2) )
			instruction = append(instruction, TreatRegisters(/*countRegister, */registers, selfQuad.op3) )
			instruction = append(instruction, "00000000000")
			s := strings.Join(instruction, "")

			fmt.Println("registradores_instrucao[", countLines, "] = ", "32'b" + s, ";") 
			//fmt.Println(countLines, "LWR", ConvertBin(s[6:11], registers), ConvertBin(s[11:16], registers), ConvertBin(s[16:21], registers))
			countLines += 1
		case "SWR":
			instruction = append(instruction, "011100" )
			instruction = append(instruction, TreatRegisters(/*countRegister, */registers, selfQuad.op3) )
			instruction = append(instruction, TreatRegisters(/*countRegister, */registers, selfQuad.op2) )
			instruction = append(instruction, TreatRegisters(/*countRegister, */registers, selfQuad.op1) )
			instruction = append(instruction, "00000000000")
			s := strings.Join(instruction, "")	

			fmt.Println("registradores_instrucao[", countLines, "] = ", "32'b" + s, ";") 
			//fmt.Println(countLines, "SWR", ConvertBin(s[16:21], registers), ConvertBin(s[11:16], registers), ConvertBin(s[6:11], registers))
			countLines += 1
		case "J":
			instruction = append(instruction, "111111" )
			instruction = append(instruction, ConvertToNBits(funcAndLab[selfQuad.op1], 26))
			s := strings.Join(instruction, "")

			fmt.Println("registradores_instrucao[", countLines, "] = ", "32'b" + s, ";") 
			//fmt.Println(countLines, "JUMP", s)
			countLines += 1
		case "OPT":
			instruction = append(instruction, "111000" )
			instruction = append(instruction, TreatRegisters(/*countRegister, */registers, selfQuad.op1) )
			instruction = append(instruction, "000000000000000000000")
			s := strings.Join(instruction, "")

			fmt.Println("registradores_instrucao[", countLines, "] = ", "32'b" + s, ";") 
			//fmt.Println(countLines, "OPT", ConvertBin(s[6:11], registers))
			countLines += 1
		case "IPT":
			instruction = append(instruction, "000111" )
			instruction = append(instruction, TreatRegisters(/*countRegister, */registers, selfQuad.op1) )
			instruction = append(instruction, "000000000000000000000")
			s := strings.Join(instruction, "")

			fmt.Println("registradores_instrucao[", countLines, "] = ", "32'b" + s, ";") 
			//fmt.Println(countLines, "IPT", ConvertBin(s[6:11], registers))
			countLines += 1
		case "BNE":
			instruction = append(instruction, "010111" )
			instruction = append(instruction, TreatRegisters(/*countRegister, */registers, selfQuad.op1) )
			instruction = append(instruction, TreatRegisters(/*countRegister, */registers, selfQuad.op2) )
			instruction = append(instruction, ConvertToNBits(funcAndLab[selfQuad.op3], 16))
			s := strings.Join(instruction, "")

			fmt.Println("registradores_instrucao[", countLines, "] = ", "32'b" + s, ";") 
			//fmt.Println(countLines, "BNE", ConvertBin(s[6:11], registers), ConvertBin(s[11:16], registers), selfQuad.op3)
			countLines += 1
		case "HALT":
			instruction = append(instruction, "111110_00000000000000000000000000")
			s := strings.Join(instruction, "")

			fmt.Println("registradores_instrucao[", countLines, "] = ", "32'b" + s, ";") 
			//fmt.Println(countLines, "HALT", s)
			countLines += 1
		case "NOP":
			instruction = append(instruction, "000000_00000000000000000000000000")
			s := strings.Join(instruction, "")

			fmt.Println("registradores_instrucao[", countLines, "] = ", "32'b" + s, ";") 
			//fmt.Println(countLines, "NOP", s)
			countLines += 1
		case "MOV":
			instruction = append(instruction, "010100" )
			instruction = append(instruction, TreatRegisters(registers, selfQuad.op1) )
			instruction = append(instruction, TreatRegisters(registers, selfQuad.op2) )
			instruction = append(instruction, "0000000000000000")
			s := strings.Join(instruction, "")

			fmt.Println("registradores_instrucao[", countLines, "] = ", "32'b" + s, ";") 
			//fmt.Println(countLines, "MOV", ConvertBin(s[6:11], registers), ConvertBin(s[11:16], registers))
			countLines += 1
		case "JR":
			instruction = append(instruction, "011010" )
			instruction = append(instruction, TreatRegisters(registers, selfQuad.op1) )
			instruction = append(instruction, "000000000000000000000")
			s := strings.Join(instruction, "")

			fmt.Println("registradores_instrucao[", countLines, "] = ", "32'b" + s, ";") 
			//fmt.Println(countLines, "JR", ConvertBin(s[6:11], registers))
			countLines += 1
		case "JAL": 
			instruction = append(instruction, "011000" )
			instruction = append(instruction, ConvertToNBits(funcAndLab[selfQuad.op1], 26))
			s := strings.Join(instruction, "")

			fmt.Println("registradores_instrucao[", countLines, "] = ", "32'b" + s, ";") 
			//fmt.Println(countLines, "JAL", ConvertBin(s[6:11], registers))
			countLines += 1
		}
	}
	fmt.Println("|--------------------------------------------------------------------------------------|")
	/*
	for key_, symbol_ := range registers {
		fmt.Println(key_, symbol_)
	}
	*/
}

func TreatRegisters (/*countRegister int, */registers map[string]int, data string) string{
	if (data == "sp"){
		x := ConvertToNBits(28, 5)
		//fmt.Println(data, registers[data], x)
		return x
	}

	if (data == "gp"){
		x := ConvertToNBits(29, 5)
		//fmt.Println(data, registers[data], x)
		return x
	}
	if (data == "ap"){
		x := ConvertToNBits(30, 5)
		//fmt.Println(data, registers[data], x)
		return x 
	}
	if (data == "dl"){
		x := ConvertToNBits(27, 5)
		//fmt.Println(data, registers[data], x)
		return x 
	}
	if (data == "RTRUE"){
		x := ConvertToNBits(31, 5)
		//fmt.Println(data, registers[data], x)
		return x 
	}

	if _, exists := registers[data]; exists{
		x := ConvertToNBits(registers[data], 5)
		//fmt.Println(data, registers[data], x)
		return x
	}

	if(data[0:1] == "p"){
		if _, exists := parameters[data]; exists{
			x := ConvertToNBits(parameters[data], 5)
			//fmt.Println(data, parameters[data], x)
			return x
		}
	}

	if(countRegister == 26){
		countRegister = 0
	} else {
		countRegister += 1
	}

	if(data[0:1] == "p"){
		if _, exists := parameters[data]; exists{

		} else {
			parameters[data] = countRegister
		}
	}

	for key_, symbol_ := range registers {
		
		if(symbol_ == countRegister){
			if(key_ == "p2"){
				countRegister += 1
				for key2_, symbol2_ := range registers {
					if(symbol2_ == countRegister) {
						delete(registers, key2_)
						break
					}
				}
				break
			} else {
				delete(registers, key_)
				break
			}
		}
	}

	x := ConvertToNBits(countRegister, 5)

	//fmt.Println(data, countRegister, x)

	registers[data] = countRegister
	return x
}

func ConvertToNBits(val int, n int) string{
	var end []string
	aux := fmt.Sprintf("%b", val)
	for i := len(aux) ; i < n ; i++ {
		end = append(end, "0")
	}
	end = append(end, aux)
	return strings.Join(end, "")
}

func ConvertBin(val string, registers map[string]int) string {
	for key_, symbol_ := range registers {
		aux := ConvertToNBits(symbol_, 5)
		if(aux == val){
			return key_
		}
	}
	return ""
}
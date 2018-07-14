package main

import(
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"os"
	"./parser"
	"./ast"	
)

func main() {
	input, _ := antlr.NewFileStream(os.Args[1])
	lexer := parser.NewCMINUSLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewCMINUSParser(stream)
	tree, _ := p.Program().(*parser.ProgramContext)
	a := new(ast.AstVisitor)
	x := a.VisitProgram(tree)

	a.PrintAstCatpiler(x)
	ast.SymbolTableBuilder(x)
	ast.SyntAnalysis(x)
}
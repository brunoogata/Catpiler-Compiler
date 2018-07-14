// Code generated from CMINUS.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // CMINUS

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 31, 297,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 7, 3,
	68, 10, 3, 12, 3, 14, 3, 71, 11, 3, 3, 4, 3, 4, 5, 4, 75, 10, 4, 3, 5,
	3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 88, 10,
	5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3,
	6, 3, 6, 3, 6, 5, 6, 104, 10, 6, 3, 7, 3, 7, 5, 7, 108, 10, 7, 3, 8, 3,
	8, 3, 8, 3, 8, 3, 8, 3, 8, 7, 8, 116, 10, 8, 12, 8, 14, 8, 119, 11, 8,
	3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 5, 9, 128, 10, 9, 3, 10, 3, 10,
	3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3,
	10, 3, 10, 3, 10, 5, 10, 145, 10, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11,
	7, 11, 152, 10, 11, 12, 11, 14, 11, 155, 11, 11, 3, 12, 3, 12, 3, 12, 3,
	12, 3, 12, 7, 12, 162, 10, 12, 12, 12, 14, 12, 165, 11, 12, 3, 13, 3, 13,
	3, 13, 3, 13, 3, 13, 5, 13, 172, 10, 13, 3, 14, 3, 14, 3, 14, 3, 14, 5,
	14, 178, 10, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15,
	3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 5, 15, 194, 10, 15, 3, 16, 3,
	16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17,
	5, 17, 208, 10, 17, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 5, 18, 215, 10,
	18, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 5, 19, 223, 10, 19, 3, 20,
	3, 20, 3, 20, 3, 20, 3, 20, 5, 20, 230, 10, 20, 3, 21, 3, 21, 3, 22, 3,
	22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 7, 22, 241, 10, 22, 12, 22, 14,
	22, 244, 11, 22, 3, 23, 3, 23, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24,
	3, 24, 7, 24, 255, 10, 24, 12, 24, 14, 24, 258, 11, 24, 3, 25, 3, 25, 3,
	26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 5, 26, 269, 10, 26, 3, 27,
	3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 5, 27, 280, 10,
	27, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 7, 28, 288, 10, 28, 12, 28,
	14, 28, 291, 11, 28, 3, 29, 3, 29, 3, 30, 3, 30, 3, 30, 2, 9, 4, 14, 20,
	22, 42, 46, 54, 31, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28,
	30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 2, 5, 3, 2,
	10, 15, 3, 2, 16, 17, 3, 2, 18, 19, 2, 296, 2, 60, 3, 2, 2, 2, 4, 62, 3,
	2, 2, 2, 6, 74, 3, 2, 2, 2, 8, 87, 3, 2, 2, 2, 10, 103, 3, 2, 2, 2, 12,
	107, 3, 2, 2, 2, 14, 109, 3, 2, 2, 2, 16, 127, 3, 2, 2, 2, 18, 144, 3,
	2, 2, 2, 20, 146, 3, 2, 2, 2, 22, 156, 3, 2, 2, 2, 24, 171, 3, 2, 2, 2,
	26, 177, 3, 2, 2, 2, 28, 193, 3, 2, 2, 2, 30, 195, 3, 2, 2, 2, 32, 207,
	3, 2, 2, 2, 34, 214, 3, 2, 2, 2, 36, 222, 3, 2, 2, 2, 38, 229, 3, 2, 2,
	2, 40, 231, 3, 2, 2, 2, 42, 233, 3, 2, 2, 2, 44, 245, 3, 2, 2, 2, 46, 247,
	3, 2, 2, 2, 48, 259, 3, 2, 2, 2, 50, 268, 3, 2, 2, 2, 52, 279, 3, 2, 2,
	2, 54, 281, 3, 2, 2, 2, 56, 292, 3, 2, 2, 2, 58, 294, 3, 2, 2, 2, 60, 61,
	5, 4, 3, 2, 61, 3, 3, 2, 2, 2, 62, 63, 8, 3, 1, 2, 63, 64, 5, 6, 4, 2,
	64, 69, 3, 2, 2, 2, 65, 66, 12, 4, 2, 2, 66, 68, 5, 6, 4, 2, 67, 65, 3,
	2, 2, 2, 68, 71, 3, 2, 2, 2, 69, 67, 3, 2, 2, 2, 69, 70, 3, 2, 2, 2, 70,
	5, 3, 2, 2, 2, 71, 69, 3, 2, 2, 2, 72, 75, 5, 8, 5, 2, 73, 75, 5, 10, 6,
	2, 74, 72, 3, 2, 2, 2, 74, 73, 3, 2, 2, 2, 75, 7, 3, 2, 2, 2, 76, 77, 7,
	5, 2, 2, 77, 78, 5, 56, 29, 2, 78, 79, 7, 24, 2, 2, 79, 88, 3, 2, 2, 2,
	80, 81, 7, 5, 2, 2, 81, 82, 5, 56, 29, 2, 82, 83, 7, 20, 2, 2, 83, 84,
	7, 28, 2, 2, 84, 85, 7, 21, 2, 2, 85, 86, 7, 24, 2, 2, 86, 88, 3, 2, 2,
	2, 87, 76, 3, 2, 2, 2, 87, 80, 3, 2, 2, 2, 88, 9, 3, 2, 2, 2, 89, 90, 7,
	5, 2, 2, 90, 91, 5, 56, 29, 2, 91, 92, 7, 22, 2, 2, 92, 93, 5, 12, 7, 2,
	93, 94, 7, 23, 2, 2, 94, 95, 5, 18, 10, 2, 95, 104, 3, 2, 2, 2, 96, 97,
	7, 6, 2, 2, 97, 98, 5, 56, 29, 2, 98, 99, 7, 22, 2, 2, 99, 100, 5, 12,
	7, 2, 100, 101, 7, 23, 2, 2, 101, 102, 5, 18, 10, 2, 102, 104, 3, 2, 2,
	2, 103, 89, 3, 2, 2, 2, 103, 96, 3, 2, 2, 2, 104, 11, 3, 2, 2, 2, 105,
	108, 5, 14, 8, 2, 106, 108, 7, 6, 2, 2, 107, 105, 3, 2, 2, 2, 107, 106,
	3, 2, 2, 2, 108, 13, 3, 2, 2, 2, 109, 110, 8, 8, 1, 2, 110, 111, 5, 16,
	9, 2, 111, 117, 3, 2, 2, 2, 112, 113, 12, 4, 2, 2, 113, 114, 7, 25, 2,
	2, 114, 116, 5, 16, 9, 2, 115, 112, 3, 2, 2, 2, 116, 119, 3, 2, 2, 2, 117,
	115, 3, 2, 2, 2, 117, 118, 3, 2, 2, 2, 118, 15, 3, 2, 2, 2, 119, 117, 3,
	2, 2, 2, 120, 121, 7, 5, 2, 2, 121, 128, 5, 56, 29, 2, 122, 123, 7, 5,
	2, 2, 123, 124, 5, 56, 29, 2, 124, 125, 7, 20, 2, 2, 125, 126, 7, 21, 2,
	2, 126, 128, 3, 2, 2, 2, 127, 120, 3, 2, 2, 2, 127, 122, 3, 2, 2, 2, 128,
	17, 3, 2, 2, 2, 129, 130, 7, 26, 2, 2, 130, 131, 5, 20, 11, 2, 131, 132,
	5, 22, 12, 2, 132, 133, 7, 27, 2, 2, 133, 145, 3, 2, 2, 2, 134, 135, 7,
	26, 2, 2, 135, 136, 5, 20, 11, 2, 136, 137, 7, 27, 2, 2, 137, 145, 3, 2,
	2, 2, 138, 139, 7, 26, 2, 2, 139, 140, 5, 22, 12, 2, 140, 141, 7, 27, 2,
	2, 141, 145, 3, 2, 2, 2, 142, 143, 7, 26, 2, 2, 143, 145, 7, 27, 2, 2,
	144, 129, 3, 2, 2, 2, 144, 134, 3, 2, 2, 2, 144, 138, 3, 2, 2, 2, 144,
	142, 3, 2, 2, 2, 145, 19, 3, 2, 2, 2, 146, 147, 8, 11, 1, 2, 147, 148,
	5, 8, 5, 2, 148, 153, 3, 2, 2, 2, 149, 150, 12, 4, 2, 2, 150, 152, 5, 8,
	5, 2, 151, 149, 3, 2, 2, 2, 152, 155, 3, 2, 2, 2, 153, 151, 3, 2, 2, 2,
	153, 154, 3, 2, 2, 2, 154, 21, 3, 2, 2, 2, 155, 153, 3, 2, 2, 2, 156, 157,
	8, 12, 1, 2, 157, 158, 5, 24, 13, 2, 158, 163, 3, 2, 2, 2, 159, 160, 12,
	4, 2, 2, 160, 162, 5, 24, 13, 2, 161, 159, 3, 2, 2, 2, 162, 165, 3, 2,
	2, 2, 163, 161, 3, 2, 2, 2, 163, 164, 3, 2, 2, 2, 164, 23, 3, 2, 2, 2,
	165, 163, 3, 2, 2, 2, 166, 172, 5, 26, 14, 2, 167, 172, 5, 18, 10, 2, 168,
	172, 5, 28, 15, 2, 169, 172, 5, 30, 16, 2, 170, 172, 5, 32, 17, 2, 171,
	166, 3, 2, 2, 2, 171, 167, 3, 2, 2, 2, 171, 168, 3, 2, 2, 2, 171, 169,
	3, 2, 2, 2, 171, 170, 3, 2, 2, 2, 172, 25, 3, 2, 2, 2, 173, 174, 5, 34,
	18, 2, 174, 175, 7, 24, 2, 2, 175, 178, 3, 2, 2, 2, 176, 178, 7, 24, 2,
	2, 177, 173, 3, 2, 2, 2, 177, 176, 3, 2, 2, 2, 178, 27, 3, 2, 2, 2, 179,
	180, 7, 3, 2, 2, 180, 181, 7, 22, 2, 2, 181, 182, 5, 34, 18, 2, 182, 183,
	7, 23, 2, 2, 183, 184, 5, 24, 13, 2, 184, 194, 3, 2, 2, 2, 185, 186, 7,
	3, 2, 2, 186, 187, 7, 22, 2, 2, 187, 188, 5, 34, 18, 2, 188, 189, 7, 23,
	2, 2, 189, 190, 5, 24, 13, 2, 190, 191, 7, 4, 2, 2, 191, 192, 5, 24, 13,
	2, 192, 194, 3, 2, 2, 2, 193, 179, 3, 2, 2, 2, 193, 185, 3, 2, 2, 2, 194,
	29, 3, 2, 2, 2, 195, 196, 7, 8, 2, 2, 196, 197, 7, 22, 2, 2, 197, 198,
	5, 34, 18, 2, 198, 199, 7, 23, 2, 2, 199, 200, 5, 24, 13, 2, 200, 31, 3,
	2, 2, 2, 201, 202, 7, 7, 2, 2, 202, 208, 7, 24, 2, 2, 203, 204, 7, 7, 2,
	2, 204, 205, 5, 34, 18, 2, 205, 206, 7, 24, 2, 2, 206, 208, 3, 2, 2, 2,
	207, 201, 3, 2, 2, 2, 207, 203, 3, 2, 2, 2, 208, 33, 3, 2, 2, 2, 209, 210,
	5, 36, 19, 2, 210, 211, 7, 9, 2, 2, 211, 212, 5, 34, 18, 2, 212, 215, 3,
	2, 2, 2, 213, 215, 5, 38, 20, 2, 214, 209, 3, 2, 2, 2, 214, 213, 3, 2,
	2, 2, 215, 35, 3, 2, 2, 2, 216, 223, 5, 56, 29, 2, 217, 218, 5, 56, 29,
	2, 218, 219, 7, 20, 2, 2, 219, 220, 5, 34, 18, 2, 220, 221, 7, 21, 2, 2,
	221, 223, 3, 2, 2, 2, 222, 216, 3, 2, 2, 2, 222, 217, 3, 2, 2, 2, 223,
	37, 3, 2, 2, 2, 224, 225, 5, 42, 22, 2, 225, 226, 5, 40, 21, 2, 226, 227,
	5, 42, 22, 2, 227, 230, 3, 2, 2, 2, 228, 230, 5, 42, 22, 2, 229, 224, 3,
	2, 2, 2, 229, 228, 3, 2, 2, 2, 230, 39, 3, 2, 2, 2, 231, 232, 9, 2, 2,
	2, 232, 41, 3, 2, 2, 2, 233, 234, 8, 22, 1, 2, 234, 235, 5, 46, 24, 2,
	235, 242, 3, 2, 2, 2, 236, 237, 12, 4, 2, 2, 237, 238, 5, 44, 23, 2, 238,
	239, 5, 46, 24, 2, 239, 241, 3, 2, 2, 2, 240, 236, 3, 2, 2, 2, 241, 244,
	3, 2, 2, 2, 242, 240, 3, 2, 2, 2, 242, 243, 3, 2, 2, 2, 243, 43, 3, 2,
	2, 2, 244, 242, 3, 2, 2, 2, 245, 246, 9, 3, 2, 2, 246, 45, 3, 2, 2, 2,
	247, 248, 8, 24, 1, 2, 248, 249, 5, 50, 26, 2, 249, 256, 3, 2, 2, 2, 250,
	251, 12, 4, 2, 2, 251, 252, 5, 48, 25, 2, 252, 253, 5, 50, 26, 2, 253,
	255, 3, 2, 2, 2, 254, 250, 3, 2, 2, 2, 255, 258, 3, 2, 2, 2, 256, 254,
	3, 2, 2, 2, 256, 257, 3, 2, 2, 2, 257, 47, 3, 2, 2, 2, 258, 256, 3, 2,
	2, 2, 259, 260, 9, 4, 2, 2, 260, 49, 3, 2, 2, 2, 261, 262, 7, 22, 2, 2,
	262, 263, 5, 34, 18, 2, 263, 264, 7, 23, 2, 2, 264, 269, 3, 2, 2, 2, 265,
	269, 5, 36, 19, 2, 266, 269, 5, 52, 27, 2, 267, 269, 5, 58, 30, 2, 268,
	261, 3, 2, 2, 2, 268, 265, 3, 2, 2, 2, 268, 266, 3, 2, 2, 2, 268, 267,
	3, 2, 2, 2, 269, 51, 3, 2, 2, 2, 270, 271, 5, 56, 29, 2, 271, 272, 7, 22,
	2, 2, 272, 273, 5, 54, 28, 2, 273, 274, 7, 23, 2, 2, 274, 280, 3, 2, 2,
	2, 275, 276, 5, 56, 29, 2, 276, 277, 7, 22, 2, 2, 277, 278, 7, 23, 2, 2,
	278, 280, 3, 2, 2, 2, 279, 270, 3, 2, 2, 2, 279, 275, 3, 2, 2, 2, 280,
	53, 3, 2, 2, 2, 281, 282, 8, 28, 1, 2, 282, 283, 5, 34, 18, 2, 283, 289,
	3, 2, 2, 2, 284, 285, 12, 4, 2, 2, 285, 286, 7, 25, 2, 2, 286, 288, 5,
	34, 18, 2, 287, 284, 3, 2, 2, 2, 288, 291, 3, 2, 2, 2, 289, 287, 3, 2,
	2, 2, 289, 290, 3, 2, 2, 2, 290, 55, 3, 2, 2, 2, 291, 289, 3, 2, 2, 2,
	292, 293, 7, 29, 2, 2, 293, 57, 3, 2, 2, 2, 294, 295, 7, 28, 2, 2, 295,
	59, 3, 2, 2, 2, 24, 69, 74, 87, 103, 107, 117, 127, 144, 153, 163, 171,
	177, 193, 207, 214, 222, 229, 242, 256, 268, 279, 289,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'if'", "'else'", "'int'", "'void'", "'return'", "'while'", "'='",
	"'=='", "'!='", "'<'", "'>'", "'<='", "'>='", "'+'", "'-'", "'*'", "'/'",
	"'['", "']'", "'('", "')'", "';'", "','", "'{'", "'}'",
}
var symbolicNames = []string{
	"", "IF", "ELSE", "INT", "VOID", "RETURN", "WHILE", "ASSIGN", "EQUAL",
	"DIFFERENT", "SMALLER", "LARGER", "SMALLERANDEQUAL", "LARGERANDEQUAL",
	"PLUS", "MINUS", "TIMES", "DIVISION", "LBRACKET", "RBRACKET", "LPARENT",
	"RPARENT", "SEMICOLON", "COMMA", "LKEY", "RKEY", "NUMBER", "IDENTIFIER",
	"WHITESPACE", "COMMENT",
}

var ruleNames = []string{
	"program", "declarationList", "declaration", "varDeclaration", "funDeclaration",
	"params", "paramList", "param", "compoundDecl", "localDeclarations", "statementList",
	"statement", "expressionDecl", "selectionDecl", "iteratorDecl", "returnDecl",
	"expression", "vari", "simpleExpression", "relational", "sumExpression",
	"sum", "term", "mult", "factor", "activation", "argList", "ident", "num",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type CMINUSParser struct {
	*antlr.BaseParser
}

func NewCMINUSParser(input antlr.TokenStream) *CMINUSParser {
	this := new(CMINUSParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "CMINUS.g4"

	return this
}

// CMINUSParser tokens.
const (
	CMINUSParserEOF             = antlr.TokenEOF
	CMINUSParserIF              = 1
	CMINUSParserELSE            = 2
	CMINUSParserINT             = 3
	CMINUSParserVOID            = 4
	CMINUSParserRETURN          = 5
	CMINUSParserWHILE           = 6
	CMINUSParserASSIGN          = 7
	CMINUSParserEQUAL           = 8
	CMINUSParserDIFFERENT       = 9
	CMINUSParserSMALLER         = 10
	CMINUSParserLARGER          = 11
	CMINUSParserSMALLERANDEQUAL = 12
	CMINUSParserLARGERANDEQUAL  = 13
	CMINUSParserPLUS            = 14
	CMINUSParserMINUS           = 15
	CMINUSParserTIMES           = 16
	CMINUSParserDIVISION        = 17
	CMINUSParserLBRACKET        = 18
	CMINUSParserRBRACKET        = 19
	CMINUSParserLPARENT         = 20
	CMINUSParserRPARENT         = 21
	CMINUSParserSEMICOLON       = 22
	CMINUSParserCOMMA           = 23
	CMINUSParserLKEY            = 24
	CMINUSParserRKEY            = 25
	CMINUSParserNUMBER          = 26
	CMINUSParserIDENTIFIER      = 27
	CMINUSParserWHITESPACE      = 28
	CMINUSParserCOMMENT         = 29
)

// CMINUSParser rules.
const (
	CMINUSParserRULE_program           = 0
	CMINUSParserRULE_declarationList   = 1
	CMINUSParserRULE_declaration       = 2
	CMINUSParserRULE_varDeclaration    = 3
	CMINUSParserRULE_funDeclaration    = 4
	CMINUSParserRULE_params            = 5
	CMINUSParserRULE_paramList         = 6
	CMINUSParserRULE_param             = 7
	CMINUSParserRULE_compoundDecl      = 8
	CMINUSParserRULE_localDeclarations = 9
	CMINUSParserRULE_statementList     = 10
	CMINUSParserRULE_statement         = 11
	CMINUSParserRULE_expressionDecl    = 12
	CMINUSParserRULE_selectionDecl     = 13
	CMINUSParserRULE_iteratorDecl      = 14
	CMINUSParserRULE_returnDecl        = 15
	CMINUSParserRULE_expression        = 16
	CMINUSParserRULE_vari              = 17
	CMINUSParserRULE_simpleExpression  = 18
	CMINUSParserRULE_relational        = 19
	CMINUSParserRULE_sumExpression     = 20
	CMINUSParserRULE_sum               = 21
	CMINUSParserRULE_term              = 22
	CMINUSParserRULE_mult              = 23
	CMINUSParserRULE_factor            = 24
	CMINUSParserRULE_activation        = 25
	CMINUSParserRULE_argList           = 26
	CMINUSParserRULE_ident             = 27
	CMINUSParserRULE_num               = 28
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CMINUSParserRULE_program
	return p
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CMINUSParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) DeclarationList() IDeclarationListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeclarationListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDeclarationListContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CMINUSListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CMINUSListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (s *ProgramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CMINUSVisitor:
		return t.VisitProgram(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CMINUSParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, CMINUSParserRULE_program)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(58)
		p.declarationList(0)
	}

	return localctx
}

// IDeclarationListContext is an interface to support dynamic dispatch.
type IDeclarationListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDeclarationListContext differentiates from other interfaces.
	IsDeclarationListContext()
}

type DeclarationListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclarationListContext() *DeclarationListContext {
	var p = new(DeclarationListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CMINUSParserRULE_declarationList
	return p
}

func (*DeclarationListContext) IsDeclarationListContext() {}

func NewDeclarationListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclarationListContext {
	var p = new(DeclarationListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CMINUSParserRULE_declarationList

	return p
}

func (s *DeclarationListContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclarationListContext) Declaration() IDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDeclarationContext)
}

func (s *DeclarationListContext) DeclarationList() IDeclarationListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeclarationListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDeclarationListContext)
}

func (s *DeclarationListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarationListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclarationListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CMINUSListener); ok {
		listenerT.EnterDeclarationList(s)
	}
}

func (s *DeclarationListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CMINUSListener); ok {
		listenerT.ExitDeclarationList(s)
	}
}

func (s *DeclarationListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CMINUSVisitor:
		return t.VisitDeclarationList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CMINUSParser) DeclarationList() (localctx IDeclarationListContext) {
	return p.declarationList(0)
}

func (p *CMINUSParser) declarationList(_p int) (localctx IDeclarationListContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewDeclarationListContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IDeclarationListContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 2
	p.EnterRecursionRule(localctx, 2, CMINUSParserRULE_declarationList, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(61)
		p.Declaration()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(67)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewDeclarationListContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, CMINUSParserRULE_declarationList)
			p.SetState(63)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(64)
				p.Declaration()
			}

		}
		p.SetState(69)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())
	}

	return localctx
}

// IDeclarationContext is an interface to support dynamic dispatch.
type IDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDeclarationContext differentiates from other interfaces.
	IsDeclarationContext()
}

type DeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclarationContext() *DeclarationContext {
	var p = new(DeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CMINUSParserRULE_declaration
	return p
}

func (*DeclarationContext) IsDeclarationContext() {}

func NewDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclarationContext {
	var p = new(DeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CMINUSParserRULE_declaration

	return p
}

func (s *DeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclarationContext) VarDeclaration() IVarDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVarDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVarDeclarationContext)
}

func (s *DeclarationContext) FunDeclaration() IFunDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunDeclarationContext)
}

func (s *DeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CMINUSListener); ok {
		listenerT.EnterDeclaration(s)
	}
}

func (s *DeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CMINUSListener); ok {
		listenerT.ExitDeclaration(s)
	}
}

func (s *DeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CMINUSVisitor:
		return t.VisitDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CMINUSParser) Declaration() (localctx IDeclarationContext) {
	localctx = NewDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, CMINUSParserRULE_declaration)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(72)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(70)
			p.VarDeclaration()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(71)
			p.FunDeclaration()
		}

	}

	return localctx
}

// IVarDeclarationContext is an interface to support dynamic dispatch.
type IVarDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVarDeclarationContext differentiates from other interfaces.
	IsVarDeclarationContext()
}

type VarDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarDeclarationContext() *VarDeclarationContext {
	var p = new(VarDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CMINUSParserRULE_varDeclaration
	return p
}

func (*VarDeclarationContext) IsVarDeclarationContext() {}

func NewVarDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarDeclarationContext {
	var p = new(VarDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CMINUSParserRULE_varDeclaration

	return p
}

func (s *VarDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *VarDeclarationContext) INT() antlr.TerminalNode {
	return s.GetToken(CMINUSParserINT, 0)
}

func (s *VarDeclarationContext) Ident() IIdentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *VarDeclarationContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(CMINUSParserSEMICOLON, 0)
}

func (s *VarDeclarationContext) LBRACKET() antlr.TerminalNode {
	return s.GetToken(CMINUSParserLBRACKET, 0)
}

func (s *VarDeclarationContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(CMINUSParserNUMBER, 0)
}

func (s *VarDeclarationContext) RBRACKET() antlr.TerminalNode {
	return s.GetToken(CMINUSParserRBRACKET, 0)
}

func (s *VarDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CMINUSListener); ok {
		listenerT.EnterVarDeclaration(s)
	}
}

func (s *VarDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CMINUSListener); ok {
		listenerT.ExitVarDeclaration(s)
	}
}

func (s *VarDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CMINUSVisitor:
		return t.VisitVarDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CMINUSParser) VarDeclaration() (localctx IVarDeclarationContext) {
	localctx = NewVarDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, CMINUSParserRULE_varDeclaration)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(85)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(74)
			p.Match(CMINUSParserINT)
		}
		{
			p.SetState(75)
			p.Ident()
		}
		{
			p.SetState(76)
			p.Match(CMINUSParserSEMICOLON)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(78)
			p.Match(CMINUSParserINT)
		}
		{
			p.SetState(79)
			p.Ident()
		}
		{
			p.SetState(80)
			p.Match(CMINUSParserLBRACKET)
		}
		{
			p.SetState(81)
			p.Match(CMINUSParserNUMBER)
		}
		{
			p.SetState(82)
			p.Match(CMINUSParserRBRACKET)
		}
		{
			p.SetState(83)
			p.Match(CMINUSParserSEMICOLON)
		}

	}

	return localctx
}

// IFunDeclarationContext is an interface to support dynamic dispatch.
type IFunDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunDeclarationContext differentiates from other interfaces.
	IsFunDeclarationContext()
}

type FunDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunDeclarationContext() *FunDeclarationContext {
	var p = new(FunDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CMINUSParserRULE_funDeclaration
	return p
}

func (*FunDeclarationContext) IsFunDeclarationContext() {}

func NewFunDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunDeclarationContext {
	var p = new(FunDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CMINUSParserRULE_funDeclaration

	return p
}

func (s *FunDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *FunDeclarationContext) INT() antlr.TerminalNode {
	return s.GetToken(CMINUSParserINT, 0)
}

func (s *FunDeclarationContext) Ident() IIdentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *FunDeclarationContext) LPARENT() antlr.TerminalNode {
	return s.GetToken(CMINUSParserLPARENT, 0)
}

func (s *FunDeclarationContext) Params() IParamsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParamsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParamsContext)
}

func (s *FunDeclarationContext) RPARENT() antlr.TerminalNode {
	return s.GetToken(CMINUSParserRPARENT, 0)
}

func (s *FunDeclarationContext) CompoundDecl() ICompoundDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICompoundDeclContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICompoundDeclContext)
}

func (s *FunDeclarationContext) VOID() antlr.TerminalNode {
	return s.GetToken(CMINUSParserVOID, 0)
}

func (s *FunDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CMINUSListener); ok {
		listenerT.EnterFunDeclaration(s)
	}
}

func (s *FunDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CMINUSListener); ok {
		listenerT.ExitFunDeclaration(s)
	}
}

func (s *FunDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CMINUSVisitor:
		return t.VisitFunDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CMINUSParser) FunDeclaration() (localctx IFunDeclarationContext) {
	localctx = NewFunDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, CMINUSParserRULE_funDeclaration)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(101)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CMINUSParserINT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(87)
			p.Match(CMINUSParserINT)
		}
		{
			p.SetState(88)
			p.Ident()
		}
		{
			p.SetState(89)
			p.Match(CMINUSParserLPARENT)
		}
		{
			p.SetState(90)
			p.Params()
		}
		{
			p.SetState(91)
			p.Match(CMINUSParserRPARENT)
		}
		{
			p.SetState(92)
			p.CompoundDecl()
		}

	case CMINUSParserVOID:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(94)
			p.Match(CMINUSParserVOID)
		}
		{
			p.SetState(95)
			p.Ident()
		}
		{
			p.SetState(96)
			p.Match(CMINUSParserLPARENT)
		}
		{
			p.SetState(97)
			p.Params()
		}
		{
			p.SetState(98)
			p.Match(CMINUSParserRPARENT)
		}
		{
			p.SetState(99)
			p.CompoundDecl()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IParamsContext is an interface to support dynamic dispatch.
type IParamsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParamsContext differentiates from other interfaces.
	IsParamsContext()
}

type ParamsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParamsContext() *ParamsContext {
	var p = new(ParamsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CMINUSParserRULE_params
	return p
}

func (*ParamsContext) IsParamsContext() {}

func NewParamsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamsContext {
	var p = new(ParamsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CMINUSParserRULE_params

	return p
}

func (s *ParamsContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamsContext) ParamList() IParamListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParamListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParamListContext)
}

func (s *ParamsContext) VOID() antlr.TerminalNode {
	return s.GetToken(CMINUSParserVOID, 0)
}

func (s *ParamsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParamsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CMINUSListener); ok {
		listenerT.EnterParams(s)
	}
}

func (s *ParamsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CMINUSListener); ok {
		listenerT.ExitParams(s)
	}
}

func (s *ParamsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CMINUSVisitor:
		return t.VisitParams(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CMINUSParser) Params() (localctx IParamsContext) {
	localctx = NewParamsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, CMINUSParserRULE_params)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(105)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CMINUSParserINT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(103)
			p.paramList(0)
		}

	case CMINUSParserVOID:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(104)
			p.Match(CMINUSParserVOID)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IParamListContext is an interface to support dynamic dispatch.
type IParamListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParamListContext differentiates from other interfaces.
	IsParamListContext()
}

type ParamListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParamListContext() *ParamListContext {
	var p = new(ParamListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CMINUSParserRULE_paramList
	return p
}

func (*ParamListContext) IsParamListContext() {}

func NewParamListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamListContext {
	var p = new(ParamListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CMINUSParserRULE_paramList

	return p
}

func (s *ParamListContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamListContext) Param() IParamContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParamContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParamContext)
}

func (s *ParamListContext) ParamList() IParamListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParamListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParamListContext)
}

func (s *ParamListContext) COMMA() antlr.TerminalNode {
	return s.GetToken(CMINUSParserCOMMA, 0)
}

func (s *ParamListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParamListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CMINUSListener); ok {
		listenerT.EnterParamList(s)
	}
}

func (s *ParamListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CMINUSListener); ok {
		listenerT.ExitParamList(s)
	}
}

func (s *ParamListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CMINUSVisitor:
		return t.VisitParamList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CMINUSParser) ParamList() (localctx IParamListContext) {
	return p.paramList(0)
}

func (p *CMINUSParser) paramList(_p int) (localctx IParamListContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewParamListContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IParamListContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 12
	p.EnterRecursionRule(localctx, 12, CMINUSParserRULE_paramList, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(108)
		p.Param()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(115)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewParamListContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, CMINUSParserRULE_paramList)
			p.SetState(110)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(111)
				p.Match(CMINUSParserCOMMA)
			}
			{
				p.SetState(112)
				p.Param()
			}

		}
		p.SetState(117)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())
	}

	return localctx
}

// IParamContext is an interface to support dynamic dispatch.
type IParamContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParamContext differentiates from other interfaces.
	IsParamContext()
}

type ParamContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParamContext() *ParamContext {
	var p = new(ParamContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CMINUSParserRULE_param
	return p
}

func (*ParamContext) IsParamContext() {}

func NewParamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamContext {
	var p = new(ParamContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CMINUSParserRULE_param

	return p
}

func (s *ParamContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamContext) INT() antlr.TerminalNode {
	return s.GetToken(CMINUSParserINT, 0)
}

func (s *ParamContext) Ident() IIdentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *ParamContext) LBRACKET() antlr.TerminalNode {
	return s.GetToken(CMINUSParserLBRACKET, 0)
}

func (s *ParamContext) RBRACKET() antlr.TerminalNode {
	return s.GetToken(CMINUSParserRBRACKET, 0)
}

func (s *ParamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParamContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CMINUSListener); ok {
		listenerT.EnterParam(s)
	}
}

func (s *ParamContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CMINUSListener); ok {
		listenerT.ExitParam(s)
	}
}

func (s *ParamContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CMINUSVisitor:
		return t.VisitParam(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CMINUSParser) Param() (localctx IParamContext) {
	localctx = NewParamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, CMINUSParserRULE_param)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(125)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(118)
			p.Match(CMINUSParserINT)
		}
		{
			p.SetState(119)
			p.Ident()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(120)
			p.Match(CMINUSParserINT)
		}
		{
			p.SetState(121)
			p.Ident()
		}
		{
			p.SetState(122)
			p.Match(CMINUSParserLBRACKET)
		}
		{
			p.SetState(123)
			p.Match(CMINUSParserRBRACKET)
		}

	}

	return localctx
}

// ICompoundDeclContext is an interface to support dynamic dispatch.
type ICompoundDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCompoundDeclContext differentiates from other interfaces.
	IsCompoundDeclContext()
}

type CompoundDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCompoundDeclContext() *CompoundDeclContext {
	var p = new(CompoundDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CMINUSParserRULE_compoundDecl
	return p
}

func (*CompoundDeclContext) IsCompoundDeclContext() {}

func NewCompoundDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CompoundDeclContext {
	var p = new(CompoundDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CMINUSParserRULE_compoundDecl

	return p
}

func (s *CompoundDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *CompoundDeclContext) LKEY() antlr.TerminalNode {
	return s.GetToken(CMINUSParserLKEY, 0)
}

func (s *CompoundDeclContext) LocalDeclarations() ILocalDeclarationsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILocalDeclarationsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILocalDeclarationsContext)
}

func (s *CompoundDeclContext) StatementList() IStatementListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementListContext)
}

func (s *CompoundDeclContext) RKEY() antlr.TerminalNode {
	return s.GetToken(CMINUSParserRKEY, 0)
}

func (s *CompoundDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompoundDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CompoundDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CMINUSListener); ok {
		listenerT.EnterCompoundDecl(s)
	}
}

func (s *CompoundDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CMINUSListener); ok {
		listenerT.ExitCompoundDecl(s)
	}
}

func (s *CompoundDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CMINUSVisitor:
		return t.VisitCompoundDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CMINUSParser) CompoundDecl() (localctx ICompoundDeclContext) {
	localctx = NewCompoundDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, CMINUSParserRULE_compoundDecl)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(142)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(127)
			p.Match(CMINUSParserLKEY)
		}
		{
			p.SetState(128)
			p.localDeclarations(0)
		}
		{
			p.SetState(129)
			p.statementList(0)
		}
		{
			p.SetState(130)
			p.Match(CMINUSParserRKEY)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(132)
			p.Match(CMINUSParserLKEY)
		}
		{
			p.SetState(133)
			p.localDeclarations(0)
		}
		{
			p.SetState(134)
			p.Match(CMINUSParserRKEY)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(136)
			p.Match(CMINUSParserLKEY)
		}
		{
			p.SetState(137)
			p.statementList(0)
		}
		{
			p.SetState(138)
			p.Match(CMINUSParserRKEY)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(140)
			p.Match(CMINUSParserLKEY)
		}
		{
			p.SetState(141)
			p.Match(CMINUSParserRKEY)
		}

	}

	return localctx
}

// ILocalDeclarationsContext is an interface to support dynamic dispatch.
type ILocalDeclarationsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLocalDeclarationsContext differentiates from other interfaces.
	IsLocalDeclarationsContext()
}

type LocalDeclarationsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLocalDeclarationsContext() *LocalDeclarationsContext {
	var p = new(LocalDeclarationsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CMINUSParserRULE_localDeclarations
	return p
}

func (*LocalDeclarationsContext) IsLocalDeclarationsContext() {}

func NewLocalDeclarationsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LocalDeclarationsContext {
	var p = new(LocalDeclarationsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CMINUSParserRULE_localDeclarations

	return p
}

func (s *LocalDeclarationsContext) GetParser() antlr.Parser { return s.parser }

func (s *LocalDeclarationsContext) VarDeclaration() IVarDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVarDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVarDeclarationContext)
}

func (s *LocalDeclarationsContext) LocalDeclarations() ILocalDeclarationsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILocalDeclarationsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILocalDeclarationsContext)
}

func (s *LocalDeclarationsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LocalDeclarationsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LocalDeclarationsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CMINUSListener); ok {
		listenerT.EnterLocalDeclarations(s)
	}
}

func (s *LocalDeclarationsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CMINUSListener); ok {
		listenerT.ExitLocalDeclarations(s)
	}
}

func (s *LocalDeclarationsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CMINUSVisitor:
		return t.VisitLocalDeclarations(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CMINUSParser) LocalDeclarations() (localctx ILocalDeclarationsContext) {
	return p.localDeclarations(0)
}

func (p *CMINUSParser) localDeclarations(_p int) (localctx ILocalDeclarationsContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewLocalDeclarationsContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx ILocalDeclarationsContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 18
	p.EnterRecursionRule(localctx, 18, CMINUSParserRULE_localDeclarations, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(145)
		p.VarDeclaration()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(151)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewLocalDeclarationsContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, CMINUSParserRULE_localDeclarations)
			p.SetState(147)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(148)
				p.VarDeclaration()
			}

		}
		p.SetState(153)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())
	}

	return localctx
}

// IStatementListContext is an interface to support dynamic dispatch.
type IStatementListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementListContext differentiates from other interfaces.
	IsStatementListContext()
}

type StatementListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementListContext() *StatementListContext {
	var p = new(StatementListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CMINUSParserRULE_statementList
	return p
}

func (*StatementListContext) IsStatementListContext() {}

func NewStatementListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementListContext {
	var p = new(StatementListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CMINUSParserRULE_statementList

	return p
}

func (s *StatementListContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementListContext) Statement() IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *StatementListContext) StatementList() IStatementListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementListContext)
}

func (s *StatementListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CMINUSListener); ok {
		listenerT.EnterStatementList(s)
	}
}

func (s *StatementListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CMINUSListener); ok {
		listenerT.ExitStatementList(s)
	}
}

func (s *StatementListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CMINUSVisitor:
		return t.VisitStatementList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CMINUSParser) StatementList() (localctx IStatementListContext) {
	return p.statementList(0)
}

func (p *CMINUSParser) statementList(_p int) (localctx IStatementListContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewStatementListContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IStatementListContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 20
	p.EnterRecursionRule(localctx, 20, CMINUSParserRULE_statementList, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(155)
		p.Statement()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(161)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewStatementListContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, CMINUSParserRULE_statementList)
			p.SetState(157)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(158)
				p.Statement()
			}

		}
		p.SetState(163)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext())
	}

	return localctx
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CMINUSParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CMINUSParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) ExpressionDecl() IExpressionDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionDeclContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionDeclContext)
}

func (s *StatementContext) CompoundDecl() ICompoundDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICompoundDeclContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICompoundDeclContext)
}

func (s *StatementContext) SelectionDecl() ISelectionDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISelectionDeclContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISelectionDeclContext)
}

func (s *StatementContext) IteratorDecl() IIteratorDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIteratorDeclContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIteratorDeclContext)
}

func (s *StatementContext) ReturnDecl() IReturnDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReturnDeclContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IReturnDeclContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CMINUSListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CMINUSListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (s *StatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CMINUSVisitor:
		return t.VisitStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CMINUSParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, CMINUSParserRULE_statement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(169)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CMINUSParserLPARENT, CMINUSParserSEMICOLON, CMINUSParserNUMBER, CMINUSParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(164)
			p.ExpressionDecl()
		}

	case CMINUSParserLKEY:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(165)
			p.CompoundDecl()
		}

	case CMINUSParserIF:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(166)
			p.SelectionDecl()
		}

	case CMINUSParserWHILE:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(167)
			p.IteratorDecl()
		}

	case CMINUSParserRETURN:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(168)
			p.ReturnDecl()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IExpressionDeclContext is an interface to support dynamic dispatch.
type IExpressionDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionDeclContext differentiates from other interfaces.
	IsExpressionDeclContext()
}

type ExpressionDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionDeclContext() *ExpressionDeclContext {
	var p = new(ExpressionDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CMINUSParserRULE_expressionDecl
	return p
}

func (*ExpressionDeclContext) IsExpressionDeclContext() {}

func NewExpressionDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionDeclContext {
	var p = new(ExpressionDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CMINUSParserRULE_expressionDecl

	return p
}

func (s *ExpressionDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionDeclContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionDeclContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(CMINUSParserSEMICOLON, 0)
}

func (s *ExpressionDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CMINUSListener); ok {
		listenerT.EnterExpressionDecl(s)
	}
}

func (s *ExpressionDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CMINUSListener); ok {
		listenerT.ExitExpressionDecl(s)
	}
}

func (s *ExpressionDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CMINUSVisitor:
		return t.VisitExpressionDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CMINUSParser) ExpressionDecl() (localctx IExpressionDeclContext) {
	localctx = NewExpressionDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, CMINUSParserRULE_expressionDecl)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(175)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CMINUSParserLPARENT, CMINUSParserNUMBER, CMINUSParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(171)
			p.Expression()
		}
		{
			p.SetState(172)
			p.Match(CMINUSParserSEMICOLON)
		}

	case CMINUSParserSEMICOLON:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(174)
			p.Match(CMINUSParserSEMICOLON)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ISelectionDeclContext is an interface to support dynamic dispatch.
type ISelectionDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSelectionDeclContext differentiates from other interfaces.
	IsSelectionDeclContext()
}

type SelectionDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelectionDeclContext() *SelectionDeclContext {
	var p = new(SelectionDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CMINUSParserRULE_selectionDecl
	return p
}

func (*SelectionDeclContext) IsSelectionDeclContext() {}

func NewSelectionDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectionDeclContext {
	var p = new(SelectionDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CMINUSParserRULE_selectionDecl

	return p
}

func (s *SelectionDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectionDeclContext) IF() antlr.TerminalNode {
	return s.GetToken(CMINUSParserIF, 0)
}

func (s *SelectionDeclContext) LPARENT() antlr.TerminalNode {
	return s.GetToken(CMINUSParserLPARENT, 0)
}

func (s *SelectionDeclContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *SelectionDeclContext) RPARENT() antlr.TerminalNode {
	return s.GetToken(CMINUSParserRPARENT, 0)
}

func (s *SelectionDeclContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *SelectionDeclContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *SelectionDeclContext) ELSE() antlr.TerminalNode {
	return s.GetToken(CMINUSParserELSE, 0)
}

func (s *SelectionDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectionDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SelectionDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CMINUSListener); ok {
		listenerT.EnterSelectionDecl(s)
	}
}

func (s *SelectionDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CMINUSListener); ok {
		listenerT.ExitSelectionDecl(s)
	}
}

func (s *SelectionDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CMINUSVisitor:
		return t.VisitSelectionDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CMINUSParser) SelectionDecl() (localctx ISelectionDeclContext) {
	localctx = NewSelectionDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, CMINUSParserRULE_selectionDecl)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(191)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(177)
			p.Match(CMINUSParserIF)
		}
		{
			p.SetState(178)
			p.Match(CMINUSParserLPARENT)
		}
		{
			p.SetState(179)
			p.Expression()
		}
		{
			p.SetState(180)
			p.Match(CMINUSParserRPARENT)
		}
		{
			p.SetState(181)
			p.Statement()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(183)
			p.Match(CMINUSParserIF)
		}
		{
			p.SetState(184)
			p.Match(CMINUSParserLPARENT)
		}
		{
			p.SetState(185)
			p.Expression()
		}
		{
			p.SetState(186)
			p.Match(CMINUSParserRPARENT)
		}
		{
			p.SetState(187)
			p.Statement()
		}
		{
			p.SetState(188)
			p.Match(CMINUSParserELSE)
		}
		{
			p.SetState(189)
			p.Statement()
		}

	}

	return localctx
}

// IIteratorDeclContext is an interface to support dynamic dispatch.
type IIteratorDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIteratorDeclContext differentiates from other interfaces.
	IsIteratorDeclContext()
}

type IteratorDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIteratorDeclContext() *IteratorDeclContext {
	var p = new(IteratorDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CMINUSParserRULE_iteratorDecl
	return p
}

func (*IteratorDeclContext) IsIteratorDeclContext() {}

func NewIteratorDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IteratorDeclContext {
	var p = new(IteratorDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CMINUSParserRULE_iteratorDecl

	return p
}

func (s *IteratorDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *IteratorDeclContext) WHILE() antlr.TerminalNode {
	return s.GetToken(CMINUSParserWHILE, 0)
}

func (s *IteratorDeclContext) LPARENT() antlr.TerminalNode {
	return s.GetToken(CMINUSParserLPARENT, 0)
}

func (s *IteratorDeclContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IteratorDeclContext) RPARENT() antlr.TerminalNode {
	return s.GetToken(CMINUSParserRPARENT, 0)
}

func (s *IteratorDeclContext) Statement() IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *IteratorDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IteratorDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IteratorDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CMINUSListener); ok {
		listenerT.EnterIteratorDecl(s)
	}
}

func (s *IteratorDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CMINUSListener); ok {
		listenerT.ExitIteratorDecl(s)
	}
}

func (s *IteratorDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CMINUSVisitor:
		return t.VisitIteratorDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CMINUSParser) IteratorDecl() (localctx IIteratorDeclContext) {
	localctx = NewIteratorDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, CMINUSParserRULE_iteratorDecl)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(193)
		p.Match(CMINUSParserWHILE)
	}
	{
		p.SetState(194)
		p.Match(CMINUSParserLPARENT)
	}
	{
		p.SetState(195)
		p.Expression()
	}
	{
		p.SetState(196)
		p.Match(CMINUSParserRPARENT)
	}
	{
		p.SetState(197)
		p.Statement()
	}

	return localctx
}

// IReturnDeclContext is an interface to support dynamic dispatch.
type IReturnDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsReturnDeclContext differentiates from other interfaces.
	IsReturnDeclContext()
}

type ReturnDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturnDeclContext() *ReturnDeclContext {
	var p = new(ReturnDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CMINUSParserRULE_returnDecl
	return p
}

func (*ReturnDeclContext) IsReturnDeclContext() {}

func NewReturnDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturnDeclContext {
	var p = new(ReturnDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CMINUSParserRULE_returnDecl

	return p
}

func (s *ReturnDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *ReturnDeclContext) RETURN() antlr.TerminalNode {
	return s.GetToken(CMINUSParserRETURN, 0)
}

func (s *ReturnDeclContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(CMINUSParserSEMICOLON, 0)
}

func (s *ReturnDeclContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ReturnDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReturnDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CMINUSListener); ok {
		listenerT.EnterReturnDecl(s)
	}
}

func (s *ReturnDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CMINUSListener); ok {
		listenerT.ExitReturnDecl(s)
	}
}

func (s *ReturnDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CMINUSVisitor:
		return t.VisitReturnDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CMINUSParser) ReturnDecl() (localctx IReturnDeclContext) {
	localctx = NewReturnDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, CMINUSParserRULE_returnDecl)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(205)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(199)
			p.Match(CMINUSParserRETURN)
		}
		{
			p.SetState(200)
			p.Match(CMINUSParserSEMICOLON)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(201)
			p.Match(CMINUSParserRETURN)
		}
		{
			p.SetState(202)
			p.Expression()
		}
		{
			p.SetState(203)
			p.Match(CMINUSParserSEMICOLON)
		}

	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CMINUSParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CMINUSParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Vari() IVariContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariContext)
}

func (s *ExpressionContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(CMINUSParserASSIGN, 0)
}

func (s *ExpressionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionContext) SimpleExpression() ISimpleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISimpleExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISimpleExpressionContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CMINUSListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CMINUSListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (s *ExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CMINUSVisitor:
		return t.VisitExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CMINUSParser) Expression() (localctx IExpressionContext) {
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, CMINUSParserRULE_expression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(212)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(207)
			p.Vari()
		}
		{
			p.SetState(208)
			p.Match(CMINUSParserASSIGN)
		}
		{
			p.SetState(209)
			p.Expression()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(211)
			p.SimpleExpression()
		}

	}

	return localctx
}

// IVariContext is an interface to support dynamic dispatch.
type IVariContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariContext differentiates from other interfaces.
	IsVariContext()
}

type VariContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariContext() *VariContext {
	var p = new(VariContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CMINUSParserRULE_vari
	return p
}

func (*VariContext) IsVariContext() {}

func NewVariContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariContext {
	var p = new(VariContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CMINUSParserRULE_vari

	return p
}

func (s *VariContext) GetParser() antlr.Parser { return s.parser }

func (s *VariContext) Ident() IIdentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *VariContext) LBRACKET() antlr.TerminalNode {
	return s.GetToken(CMINUSParserLBRACKET, 0)
}

func (s *VariContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *VariContext) RBRACKET() antlr.TerminalNode {
	return s.GetToken(CMINUSParserRBRACKET, 0)
}

func (s *VariContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CMINUSListener); ok {
		listenerT.EnterVari(s)
	}
}

func (s *VariContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CMINUSListener); ok {
		listenerT.ExitVari(s)
	}
}

func (s *VariContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CMINUSVisitor:
		return t.VisitVari(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CMINUSParser) Vari() (localctx IVariContext) {
	localctx = NewVariContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, CMINUSParserRULE_vari)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(220)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(214)
			p.Ident()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(215)
			p.Ident()
		}
		{
			p.SetState(216)
			p.Match(CMINUSParserLBRACKET)
		}
		{
			p.SetState(217)
			p.Expression()
		}
		{
			p.SetState(218)
			p.Match(CMINUSParserRBRACKET)
		}

	}

	return localctx
}

// ISimpleExpressionContext is an interface to support dynamic dispatch.
type ISimpleExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSimpleExpressionContext differentiates from other interfaces.
	IsSimpleExpressionContext()
}

type SimpleExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySimpleExpressionContext() *SimpleExpressionContext {
	var p = new(SimpleExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CMINUSParserRULE_simpleExpression
	return p
}

func (*SimpleExpressionContext) IsSimpleExpressionContext() {}

func NewSimpleExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SimpleExpressionContext {
	var p = new(SimpleExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CMINUSParserRULE_simpleExpression

	return p
}

func (s *SimpleExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *SimpleExpressionContext) AllSumExpression() []ISumExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISumExpressionContext)(nil)).Elem())
	var tst = make([]ISumExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISumExpressionContext)
		}
	}

	return tst
}

func (s *SimpleExpressionContext) SumExpression(i int) ISumExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISumExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISumExpressionContext)
}

func (s *SimpleExpressionContext) Relational() IRelationalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRelationalContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRelationalContext)
}

func (s *SimpleExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SimpleExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SimpleExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CMINUSListener); ok {
		listenerT.EnterSimpleExpression(s)
	}
}

func (s *SimpleExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CMINUSListener); ok {
		listenerT.ExitSimpleExpression(s)
	}
}

func (s *SimpleExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CMINUSVisitor:
		return t.VisitSimpleExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CMINUSParser) SimpleExpression() (localctx ISimpleExpressionContext) {
	localctx = NewSimpleExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, CMINUSParserRULE_simpleExpression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(227)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(222)
			p.sumExpression(0)
		}
		{
			p.SetState(223)
			p.Relational()
		}
		{
			p.SetState(224)
			p.sumExpression(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(226)
			p.sumExpression(0)
		}

	}

	return localctx
}

// IRelationalContext is an interface to support dynamic dispatch.
type IRelationalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRelationalContext differentiates from other interfaces.
	IsRelationalContext()
}

type RelationalContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRelationalContext() *RelationalContext {
	var p = new(RelationalContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CMINUSParserRULE_relational
	return p
}

func (*RelationalContext) IsRelationalContext() {}

func NewRelationalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RelationalContext {
	var p = new(RelationalContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CMINUSParserRULE_relational

	return p
}

func (s *RelationalContext) GetParser() antlr.Parser { return s.parser }

func (s *RelationalContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(CMINUSParserEQUAL, 0)
}

func (s *RelationalContext) DIFFERENT() antlr.TerminalNode {
	return s.GetToken(CMINUSParserDIFFERENT, 0)
}

func (s *RelationalContext) SMALLER() antlr.TerminalNode {
	return s.GetToken(CMINUSParserSMALLER, 0)
}

func (s *RelationalContext) SMALLERANDEQUAL() antlr.TerminalNode {
	return s.GetToken(CMINUSParserSMALLERANDEQUAL, 0)
}

func (s *RelationalContext) LARGER() antlr.TerminalNode {
	return s.GetToken(CMINUSParserLARGER, 0)
}

func (s *RelationalContext) LARGERANDEQUAL() antlr.TerminalNode {
	return s.GetToken(CMINUSParserLARGERANDEQUAL, 0)
}

func (s *RelationalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelationalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RelationalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CMINUSListener); ok {
		listenerT.EnterRelational(s)
	}
}

func (s *RelationalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CMINUSListener); ok {
		listenerT.ExitRelational(s)
	}
}

func (s *RelationalContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CMINUSVisitor:
		return t.VisitRelational(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CMINUSParser) Relational() (localctx IRelationalContext) {
	localctx = NewRelationalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, CMINUSParserRULE_relational)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(229)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<CMINUSParserEQUAL)|(1<<CMINUSParserDIFFERENT)|(1<<CMINUSParserSMALLER)|(1<<CMINUSParserLARGER)|(1<<CMINUSParserSMALLERANDEQUAL)|(1<<CMINUSParserLARGERANDEQUAL))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ISumExpressionContext is an interface to support dynamic dispatch.
type ISumExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSumExpressionContext differentiates from other interfaces.
	IsSumExpressionContext()
}

type SumExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySumExpressionContext() *SumExpressionContext {
	var p = new(SumExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CMINUSParserRULE_sumExpression
	return p
}

func (*SumExpressionContext) IsSumExpressionContext() {}

func NewSumExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SumExpressionContext {
	var p = new(SumExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CMINUSParserRULE_sumExpression

	return p
}

func (s *SumExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *SumExpressionContext) Term() ITermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *SumExpressionContext) SumExpression() ISumExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISumExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISumExpressionContext)
}

func (s *SumExpressionContext) Sum() ISumContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISumContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISumContext)
}

func (s *SumExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SumExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SumExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CMINUSListener); ok {
		listenerT.EnterSumExpression(s)
	}
}

func (s *SumExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CMINUSListener); ok {
		listenerT.ExitSumExpression(s)
	}
}

func (s *SumExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CMINUSVisitor:
		return t.VisitSumExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CMINUSParser) SumExpression() (localctx ISumExpressionContext) {
	return p.sumExpression(0)
}

func (p *CMINUSParser) sumExpression(_p int) (localctx ISumExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewSumExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx ISumExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 40
	p.EnterRecursionRule(localctx, 40, CMINUSParserRULE_sumExpression, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(232)
		p.term(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(240)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewSumExpressionContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, CMINUSParserRULE_sumExpression)
			p.SetState(234)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(235)
				p.Sum()
			}
			{
				p.SetState(236)
				p.term(0)
			}

		}
		p.SetState(242)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext())
	}

	return localctx
}

// ISumContext is an interface to support dynamic dispatch.
type ISumContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSumContext differentiates from other interfaces.
	IsSumContext()
}

type SumContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySumContext() *SumContext {
	var p = new(SumContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CMINUSParserRULE_sum
	return p
}

func (*SumContext) IsSumContext() {}

func NewSumContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SumContext {
	var p = new(SumContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CMINUSParserRULE_sum

	return p
}

func (s *SumContext) GetParser() antlr.Parser { return s.parser }

func (s *SumContext) PLUS() antlr.TerminalNode {
	return s.GetToken(CMINUSParserPLUS, 0)
}

func (s *SumContext) MINUS() antlr.TerminalNode {
	return s.GetToken(CMINUSParserMINUS, 0)
}

func (s *SumContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SumContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SumContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CMINUSListener); ok {
		listenerT.EnterSum(s)
	}
}

func (s *SumContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CMINUSListener); ok {
		listenerT.ExitSum(s)
	}
}

func (s *SumContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CMINUSVisitor:
		return t.VisitSum(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CMINUSParser) Sum() (localctx ISumContext) {
	localctx = NewSumContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, CMINUSParserRULE_sum)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(243)
		_la = p.GetTokenStream().LA(1)

		if !(_la == CMINUSParserPLUS || _la == CMINUSParserMINUS) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ITermContext is an interface to support dynamic dispatch.
type ITermContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTermContext differentiates from other interfaces.
	IsTermContext()
}

type TermContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTermContext() *TermContext {
	var p = new(TermContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CMINUSParserRULE_term
	return p
}

func (*TermContext) IsTermContext() {}

func NewTermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TermContext {
	var p = new(TermContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CMINUSParserRULE_term

	return p
}

func (s *TermContext) GetParser() antlr.Parser { return s.parser }

func (s *TermContext) Factor() IFactorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFactorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFactorContext)
}

func (s *TermContext) Term() ITermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *TermContext) Mult() IMultContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMultContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMultContext)
}

func (s *TermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CMINUSListener); ok {
		listenerT.EnterTerm(s)
	}
}

func (s *TermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CMINUSListener); ok {
		listenerT.ExitTerm(s)
	}
}

func (s *TermContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CMINUSVisitor:
		return t.VisitTerm(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CMINUSParser) Term() (localctx ITermContext) {
	return p.term(0)
}

func (p *CMINUSParser) term(_p int) (localctx ITermContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewTermContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx ITermContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 44
	p.EnterRecursionRule(localctx, 44, CMINUSParserRULE_term, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(246)
		p.Factor()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(254)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewTermContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, CMINUSParserRULE_term)
			p.SetState(248)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(249)
				p.Mult()
			}
			{
				p.SetState(250)
				p.Factor()
			}

		}
		p.SetState(256)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext())
	}

	return localctx
}

// IMultContext is an interface to support dynamic dispatch.
type IMultContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMultContext differentiates from other interfaces.
	IsMultContext()
}

type MultContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMultContext() *MultContext {
	var p = new(MultContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CMINUSParserRULE_mult
	return p
}

func (*MultContext) IsMultContext() {}

func NewMultContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MultContext {
	var p = new(MultContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CMINUSParserRULE_mult

	return p
}

func (s *MultContext) GetParser() antlr.Parser { return s.parser }

func (s *MultContext) TIMES() antlr.TerminalNode {
	return s.GetToken(CMINUSParserTIMES, 0)
}

func (s *MultContext) DIVISION() antlr.TerminalNode {
	return s.GetToken(CMINUSParserDIVISION, 0)
}

func (s *MultContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MultContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CMINUSListener); ok {
		listenerT.EnterMult(s)
	}
}

func (s *MultContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CMINUSListener); ok {
		listenerT.ExitMult(s)
	}
}

func (s *MultContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CMINUSVisitor:
		return t.VisitMult(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CMINUSParser) Mult() (localctx IMultContext) {
	localctx = NewMultContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, CMINUSParserRULE_mult)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(257)
		_la = p.GetTokenStream().LA(1)

		if !(_la == CMINUSParserTIMES || _la == CMINUSParserDIVISION) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IFactorContext is an interface to support dynamic dispatch.
type IFactorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFactorContext differentiates from other interfaces.
	IsFactorContext()
}

type FactorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFactorContext() *FactorContext {
	var p = new(FactorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CMINUSParserRULE_factor
	return p
}

func (*FactorContext) IsFactorContext() {}

func NewFactorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FactorContext {
	var p = new(FactorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CMINUSParserRULE_factor

	return p
}

func (s *FactorContext) GetParser() antlr.Parser { return s.parser }

func (s *FactorContext) LPARENT() antlr.TerminalNode {
	return s.GetToken(CMINUSParserLPARENT, 0)
}

func (s *FactorContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *FactorContext) RPARENT() antlr.TerminalNode {
	return s.GetToken(CMINUSParserRPARENT, 0)
}

func (s *FactorContext) Vari() IVariContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariContext)
}

func (s *FactorContext) Activation() IActivationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IActivationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IActivationContext)
}

func (s *FactorContext) Num() INumContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INumContext)
}

func (s *FactorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FactorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FactorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CMINUSListener); ok {
		listenerT.EnterFactor(s)
	}
}

func (s *FactorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CMINUSListener); ok {
		listenerT.ExitFactor(s)
	}
}

func (s *FactorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CMINUSVisitor:
		return t.VisitFactor(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CMINUSParser) Factor() (localctx IFactorContext) {
	localctx = NewFactorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, CMINUSParserRULE_factor)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(266)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(259)
			p.Match(CMINUSParserLPARENT)
		}
		{
			p.SetState(260)
			p.Expression()
		}
		{
			p.SetState(261)
			p.Match(CMINUSParserRPARENT)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(263)
			p.Vari()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(264)
			p.Activation()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(265)
			p.Num()
		}

	}

	return localctx
}

// IActivationContext is an interface to support dynamic dispatch.
type IActivationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsActivationContext differentiates from other interfaces.
	IsActivationContext()
}

type ActivationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyActivationContext() *ActivationContext {
	var p = new(ActivationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CMINUSParserRULE_activation
	return p
}

func (*ActivationContext) IsActivationContext() {}

func NewActivationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ActivationContext {
	var p = new(ActivationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CMINUSParserRULE_activation

	return p
}

func (s *ActivationContext) GetParser() antlr.Parser { return s.parser }

func (s *ActivationContext) Ident() IIdentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *ActivationContext) LPARENT() antlr.TerminalNode {
	return s.GetToken(CMINUSParserLPARENT, 0)
}

func (s *ActivationContext) ArgList() IArgListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArgListContext)
}

func (s *ActivationContext) RPARENT() antlr.TerminalNode {
	return s.GetToken(CMINUSParserRPARENT, 0)
}

func (s *ActivationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ActivationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ActivationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CMINUSListener); ok {
		listenerT.EnterActivation(s)
	}
}

func (s *ActivationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CMINUSListener); ok {
		listenerT.ExitActivation(s)
	}
}

func (s *ActivationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CMINUSVisitor:
		return t.VisitActivation(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CMINUSParser) Activation() (localctx IActivationContext) {
	localctx = NewActivationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, CMINUSParserRULE_activation)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(277)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(268)
			p.Ident()
		}
		{
			p.SetState(269)
			p.Match(CMINUSParserLPARENT)
		}
		{
			p.SetState(270)
			p.argList(0)
		}
		{
			p.SetState(271)
			p.Match(CMINUSParserRPARENT)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(273)
			p.Ident()
		}
		{
			p.SetState(274)
			p.Match(CMINUSParserLPARENT)
		}
		{
			p.SetState(275)
			p.Match(CMINUSParserRPARENT)
		}

	}

	return localctx
}

// IArgListContext is an interface to support dynamic dispatch.
type IArgListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArgListContext differentiates from other interfaces.
	IsArgListContext()
}

type ArgListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgListContext() *ArgListContext {
	var p = new(ArgListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CMINUSParserRULE_argList
	return p
}

func (*ArgListContext) IsArgListContext() {}

func NewArgListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgListContext {
	var p = new(ArgListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CMINUSParserRULE_argList

	return p
}

func (s *ArgListContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgListContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ArgListContext) ArgList() IArgListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArgListContext)
}

func (s *ArgListContext) COMMA() antlr.TerminalNode {
	return s.GetToken(CMINUSParserCOMMA, 0)
}

func (s *ArgListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CMINUSListener); ok {
		listenerT.EnterArgList(s)
	}
}

func (s *ArgListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CMINUSListener); ok {
		listenerT.ExitArgList(s)
	}
}

func (s *ArgListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CMINUSVisitor:
		return t.VisitArgList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CMINUSParser) ArgList() (localctx IArgListContext) {
	return p.argList(0)
}

func (p *CMINUSParser) argList(_p int) (localctx IArgListContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewArgListContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IArgListContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 52
	p.EnterRecursionRule(localctx, 52, CMINUSParserRULE_argList, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(280)
		p.Expression()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(287)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewArgListContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, CMINUSParserRULE_argList)
			p.SetState(282)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(283)
				p.Match(CMINUSParserCOMMA)
			}
			{
				p.SetState(284)
				p.Expression()
			}

		}
		p.SetState(289)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext())
	}

	return localctx
}

// IIdentContext is an interface to support dynamic dispatch.
type IIdentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIdentContext differentiates from other interfaces.
	IsIdentContext()
}

type IdentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentContext() *IdentContext {
	var p = new(IdentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CMINUSParserRULE_ident
	return p
}

func (*IdentContext) IsIdentContext() {}

func NewIdentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentContext {
	var p = new(IdentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CMINUSParserRULE_ident

	return p
}

func (s *IdentContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(CMINUSParserIDENTIFIER, 0)
}

func (s *IdentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CMINUSListener); ok {
		listenerT.EnterIdent(s)
	}
}

func (s *IdentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CMINUSListener); ok {
		listenerT.ExitIdent(s)
	}
}

func (s *IdentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CMINUSVisitor:
		return t.VisitIdent(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CMINUSParser) Ident() (localctx IIdentContext) {
	localctx = NewIdentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, CMINUSParserRULE_ident)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(290)
		p.Match(CMINUSParserIDENTIFIER)
	}

	return localctx
}

// INumContext is an interface to support dynamic dispatch.
type INumContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNumContext differentiates from other interfaces.
	IsNumContext()
}

type NumContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumContext() *NumContext {
	var p = new(NumContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CMINUSParserRULE_num
	return p
}

func (*NumContext) IsNumContext() {}

func NewNumContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumContext {
	var p = new(NumContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CMINUSParserRULE_num

	return p
}

func (s *NumContext) GetParser() antlr.Parser { return s.parser }

func (s *NumContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(CMINUSParserNUMBER, 0)
}

func (s *NumContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CMINUSListener); ok {
		listenerT.EnterNum(s)
	}
}

func (s *NumContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CMINUSListener); ok {
		listenerT.ExitNum(s)
	}
}

func (s *NumContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CMINUSVisitor:
		return t.VisitNum(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CMINUSParser) Num() (localctx INumContext) {
	localctx = NewNumContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, CMINUSParserRULE_num)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(292)
		p.Match(CMINUSParserNUMBER)
	}

	return localctx
}

func (p *CMINUSParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 1:
		var t *DeclarationListContext = nil
		if localctx != nil {
			t = localctx.(*DeclarationListContext)
		}
		return p.DeclarationList_Sempred(t, predIndex)

	case 6:
		var t *ParamListContext = nil
		if localctx != nil {
			t = localctx.(*ParamListContext)
		}
		return p.ParamList_Sempred(t, predIndex)

	case 9:
		var t *LocalDeclarationsContext = nil
		if localctx != nil {
			t = localctx.(*LocalDeclarationsContext)
		}
		return p.LocalDeclarations_Sempred(t, predIndex)

	case 10:
		var t *StatementListContext = nil
		if localctx != nil {
			t = localctx.(*StatementListContext)
		}
		return p.StatementList_Sempred(t, predIndex)

	case 20:
		var t *SumExpressionContext = nil
		if localctx != nil {
			t = localctx.(*SumExpressionContext)
		}
		return p.SumExpression_Sempred(t, predIndex)

	case 22:
		var t *TermContext = nil
		if localctx != nil {
			t = localctx.(*TermContext)
		}
		return p.Term_Sempred(t, predIndex)

	case 26:
		var t *ArgListContext = nil
		if localctx != nil {
			t = localctx.(*ArgListContext)
		}
		return p.ArgList_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *CMINUSParser) DeclarationList_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *CMINUSParser) ParamList_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 1:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *CMINUSParser) LocalDeclarations_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 2:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *CMINUSParser) StatementList_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 3:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *CMINUSParser) SumExpression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 4:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *CMINUSParser) Term_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 5:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *CMINUSParser) ArgList_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 6:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

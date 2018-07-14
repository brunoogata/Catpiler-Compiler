// Code generated from CMINUS.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 31, 166,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 4, 30, 9, 30, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3,
	6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3,
	8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 12, 3,
	12, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3, 16, 3, 16,
	3, 17, 3, 17, 3, 18, 3, 18, 3, 19, 3, 19, 3, 20, 3, 20, 3, 21, 3, 21, 3,
	22, 3, 22, 3, 23, 3, 23, 3, 24, 3, 24, 3, 25, 3, 25, 3, 26, 3, 26, 3, 27,
	6, 27, 135, 10, 27, 13, 27, 14, 27, 136, 3, 28, 3, 28, 7, 28, 141, 10,
	28, 12, 28, 14, 28, 144, 11, 28, 3, 29, 6, 29, 147, 10, 29, 13, 29, 14,
	29, 148, 3, 29, 3, 29, 3, 30, 3, 30, 3, 30, 3, 30, 7, 30, 157, 10, 30,
	12, 30, 14, 30, 160, 11, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 158,
	2, 31, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21,
	12, 23, 13, 25, 14, 27, 15, 29, 16, 31, 17, 33, 18, 35, 19, 37, 20, 39,
	21, 41, 22, 43, 23, 45, 24, 47, 25, 49, 26, 51, 27, 53, 28, 55, 29, 57,
	30, 59, 31, 3, 2, 6, 3, 2, 50, 59, 4, 2, 67, 92, 99, 124, 5, 2, 50, 59,
	67, 92, 99, 124, 5, 2, 11, 12, 15, 15, 34, 34, 2, 169, 2, 3, 3, 2, 2, 2,
	2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2,
	2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2,
	2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2,
	2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3,
	2, 2, 2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2, 2, 43,
	3, 2, 2, 2, 2, 45, 3, 2, 2, 2, 2, 47, 3, 2, 2, 2, 2, 49, 3, 2, 2, 2, 2,
	51, 3, 2, 2, 2, 2, 53, 3, 2, 2, 2, 2, 55, 3, 2, 2, 2, 2, 57, 3, 2, 2, 2,
	2, 59, 3, 2, 2, 2, 3, 61, 3, 2, 2, 2, 5, 64, 3, 2, 2, 2, 7, 69, 3, 2, 2,
	2, 9, 73, 3, 2, 2, 2, 11, 78, 3, 2, 2, 2, 13, 85, 3, 2, 2, 2, 15, 91, 3,
	2, 2, 2, 17, 93, 3, 2, 2, 2, 19, 96, 3, 2, 2, 2, 21, 99, 3, 2, 2, 2, 23,
	101, 3, 2, 2, 2, 25, 103, 3, 2, 2, 2, 27, 106, 3, 2, 2, 2, 29, 109, 3,
	2, 2, 2, 31, 111, 3, 2, 2, 2, 33, 113, 3, 2, 2, 2, 35, 115, 3, 2, 2, 2,
	37, 117, 3, 2, 2, 2, 39, 119, 3, 2, 2, 2, 41, 121, 3, 2, 2, 2, 43, 123,
	3, 2, 2, 2, 45, 125, 3, 2, 2, 2, 47, 127, 3, 2, 2, 2, 49, 129, 3, 2, 2,
	2, 51, 131, 3, 2, 2, 2, 53, 134, 3, 2, 2, 2, 55, 138, 3, 2, 2, 2, 57, 146,
	3, 2, 2, 2, 59, 152, 3, 2, 2, 2, 61, 62, 7, 107, 2, 2, 62, 63, 7, 104,
	2, 2, 63, 4, 3, 2, 2, 2, 64, 65, 7, 103, 2, 2, 65, 66, 7, 110, 2, 2, 66,
	67, 7, 117, 2, 2, 67, 68, 7, 103, 2, 2, 68, 6, 3, 2, 2, 2, 69, 70, 7, 107,
	2, 2, 70, 71, 7, 112, 2, 2, 71, 72, 7, 118, 2, 2, 72, 8, 3, 2, 2, 2, 73,
	74, 7, 120, 2, 2, 74, 75, 7, 113, 2, 2, 75, 76, 7, 107, 2, 2, 76, 77, 7,
	102, 2, 2, 77, 10, 3, 2, 2, 2, 78, 79, 7, 116, 2, 2, 79, 80, 7, 103, 2,
	2, 80, 81, 7, 118, 2, 2, 81, 82, 7, 119, 2, 2, 82, 83, 7, 116, 2, 2, 83,
	84, 7, 112, 2, 2, 84, 12, 3, 2, 2, 2, 85, 86, 7, 121, 2, 2, 86, 87, 7,
	106, 2, 2, 87, 88, 7, 107, 2, 2, 88, 89, 7, 110, 2, 2, 89, 90, 7, 103,
	2, 2, 90, 14, 3, 2, 2, 2, 91, 92, 7, 63, 2, 2, 92, 16, 3, 2, 2, 2, 93,
	94, 7, 63, 2, 2, 94, 95, 7, 63, 2, 2, 95, 18, 3, 2, 2, 2, 96, 97, 7, 35,
	2, 2, 97, 98, 7, 63, 2, 2, 98, 20, 3, 2, 2, 2, 99, 100, 7, 62, 2, 2, 100,
	22, 3, 2, 2, 2, 101, 102, 7, 64, 2, 2, 102, 24, 3, 2, 2, 2, 103, 104, 7,
	62, 2, 2, 104, 105, 7, 63, 2, 2, 105, 26, 3, 2, 2, 2, 106, 107, 7, 64,
	2, 2, 107, 108, 7, 63, 2, 2, 108, 28, 3, 2, 2, 2, 109, 110, 7, 45, 2, 2,
	110, 30, 3, 2, 2, 2, 111, 112, 7, 47, 2, 2, 112, 32, 3, 2, 2, 2, 113, 114,
	7, 44, 2, 2, 114, 34, 3, 2, 2, 2, 115, 116, 7, 49, 2, 2, 116, 36, 3, 2,
	2, 2, 117, 118, 7, 93, 2, 2, 118, 38, 3, 2, 2, 2, 119, 120, 7, 95, 2, 2,
	120, 40, 3, 2, 2, 2, 121, 122, 7, 42, 2, 2, 122, 42, 3, 2, 2, 2, 123, 124,
	7, 43, 2, 2, 124, 44, 3, 2, 2, 2, 125, 126, 7, 61, 2, 2, 126, 46, 3, 2,
	2, 2, 127, 128, 7, 46, 2, 2, 128, 48, 3, 2, 2, 2, 129, 130, 7, 125, 2,
	2, 130, 50, 3, 2, 2, 2, 131, 132, 7, 127, 2, 2, 132, 52, 3, 2, 2, 2, 133,
	135, 9, 2, 2, 2, 134, 133, 3, 2, 2, 2, 135, 136, 3, 2, 2, 2, 136, 134,
	3, 2, 2, 2, 136, 137, 3, 2, 2, 2, 137, 54, 3, 2, 2, 2, 138, 142, 9, 3,
	2, 2, 139, 141, 9, 4, 2, 2, 140, 139, 3, 2, 2, 2, 141, 144, 3, 2, 2, 2,
	142, 140, 3, 2, 2, 2, 142, 143, 3, 2, 2, 2, 143, 56, 3, 2, 2, 2, 144, 142,
	3, 2, 2, 2, 145, 147, 9, 5, 2, 2, 146, 145, 3, 2, 2, 2, 147, 148, 3, 2,
	2, 2, 148, 146, 3, 2, 2, 2, 148, 149, 3, 2, 2, 2, 149, 150, 3, 2, 2, 2,
	150, 151, 8, 29, 2, 2, 151, 58, 3, 2, 2, 2, 152, 153, 7, 49, 2, 2, 153,
	154, 7, 44, 2, 2, 154, 158, 3, 2, 2, 2, 155, 157, 11, 2, 2, 2, 156, 155,
	3, 2, 2, 2, 157, 160, 3, 2, 2, 2, 158, 159, 3, 2, 2, 2, 158, 156, 3, 2,
	2, 2, 159, 161, 3, 2, 2, 2, 160, 158, 3, 2, 2, 2, 161, 162, 7, 44, 2, 2,
	162, 163, 7, 49, 2, 2, 163, 164, 3, 2, 2, 2, 164, 165, 8, 30, 2, 2, 165,
	60, 3, 2, 2, 2, 7, 2, 136, 142, 148, 158, 3, 8, 2, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'if'", "'else'", "'int'", "'void'", "'return'", "'while'", "'='",
	"'=='", "'!='", "'<'", "'>'", "'<='", "'>='", "'+'", "'-'", "'*'", "'/'",
	"'['", "']'", "'('", "')'", "';'", "','", "'{'", "'}'",
}

var lexerSymbolicNames = []string{
	"", "IF", "ELSE", "INT", "VOID", "RETURN", "WHILE", "ASSIGN", "EQUAL",
	"DIFFERENT", "SMALLER", "LARGER", "SMALLERANDEQUAL", "LARGERANDEQUAL",
	"PLUS", "MINUS", "TIMES", "DIVISION", "LBRACKET", "RBRACKET", "LPARENT",
	"RPARENT", "SEMICOLON", "COMMA", "LKEY", "RKEY", "NUMBER", "IDENTIFIER",
	"WHITESPACE", "COMMENT",
}

var lexerRuleNames = []string{
	"IF", "ELSE", "INT", "VOID", "RETURN", "WHILE", "ASSIGN", "EQUAL", "DIFFERENT",
	"SMALLER", "LARGER", "SMALLERANDEQUAL", "LARGERANDEQUAL", "PLUS", "MINUS",
	"TIMES", "DIVISION", "LBRACKET", "RBRACKET", "LPARENT", "RPARENT", "SEMICOLON",
	"COMMA", "LKEY", "RKEY", "NUMBER", "IDENTIFIER", "WHITESPACE", "COMMENT",
}

type CMINUSLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewCMINUSLexer(input antlr.CharStream) *CMINUSLexer {

	l := new(CMINUSLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "CMINUS.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// CMINUSLexer tokens.
const (
	CMINUSLexerIF              = 1
	CMINUSLexerELSE            = 2
	CMINUSLexerINT             = 3
	CMINUSLexerVOID            = 4
	CMINUSLexerRETURN          = 5
	CMINUSLexerWHILE           = 6
	CMINUSLexerASSIGN          = 7
	CMINUSLexerEQUAL           = 8
	CMINUSLexerDIFFERENT       = 9
	CMINUSLexerSMALLER         = 10
	CMINUSLexerLARGER          = 11
	CMINUSLexerSMALLERANDEQUAL = 12
	CMINUSLexerLARGERANDEQUAL  = 13
	CMINUSLexerPLUS            = 14
	CMINUSLexerMINUS           = 15
	CMINUSLexerTIMES           = 16
	CMINUSLexerDIVISION        = 17
	CMINUSLexerLBRACKET        = 18
	CMINUSLexerRBRACKET        = 19
	CMINUSLexerLPARENT         = 20
	CMINUSLexerRPARENT         = 21
	CMINUSLexerSEMICOLON       = 22
	CMINUSLexerCOMMA           = 23
	CMINUSLexerLKEY            = 24
	CMINUSLexerRKEY            = 25
	CMINUSLexerNUMBER          = 26
	CMINUSLexerIDENTIFIER      = 27
	CMINUSLexerWHITESPACE      = 28
	CMINUSLexerCOMMENT         = 29
)

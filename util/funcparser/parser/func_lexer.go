// Code generated from Func.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type FuncLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var FuncLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func funclexerLexerInit() {
	staticData := &FuncLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'('", "')'", "','", "'-'", "'true'", "'false'", "'['", "']'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "NAME", "INT", "HEX", "STRING",
		"BOOL", "DOMAIN", "WS",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "NAME",
		"INT", "HEX", "STRING", "BOOL", "DOMAIN", "ENSCHAR", "TRUE", "FALSE",
		"DOUBLEQUOTEDCHAR", "SINGLEQUOTEDCHAR", "NAMESTART", "NAMEPART", "DIGIT",
		"HEXDIGIT", "LETTER", "WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 15, 180, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 1, 0, 1, 0,
		1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5,
		1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 5, 8,
		77, 8, 8, 10, 8, 12, 8, 80, 9, 8, 1, 9, 4, 9, 83, 8, 9, 11, 9, 12, 9, 84,
		1, 10, 1, 10, 1, 10, 1, 10, 4, 10, 91, 8, 10, 11, 10, 12, 10, 92, 1, 11,
		1, 11, 5, 11, 97, 8, 11, 10, 11, 12, 11, 100, 9, 11, 1, 11, 1, 11, 1, 11,
		5, 11, 105, 8, 11, 10, 11, 12, 11, 108, 9, 11, 1, 11, 3, 11, 111, 8, 11,
		1, 12, 1, 12, 3, 12, 115, 8, 12, 1, 13, 1, 13, 4, 13, 119, 8, 13, 11, 13,
		12, 13, 120, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1,
		15, 1, 15, 3, 15, 133, 8, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16,
		1, 16, 1, 16, 1, 16, 1, 16, 3, 16, 145, 8, 16, 1, 17, 1, 17, 1, 17, 3,
		17, 150, 8, 17, 1, 18, 1, 18, 1, 18, 3, 18, 155, 8, 18, 1, 19, 1, 19, 3,
		19, 159, 8, 19, 1, 20, 1, 20, 1, 20, 3, 20, 164, 8, 20, 1, 21, 1, 21, 1,
		22, 1, 22, 3, 22, 170, 8, 22, 1, 23, 1, 23, 1, 24, 4, 24, 175, 8, 24, 11,
		24, 12, 24, 176, 1, 24, 1, 24, 0, 0, 25, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5,
		11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29,
		0, 31, 0, 33, 0, 35, 0, 37, 0, 39, 0, 41, 0, 43, 0, 45, 0, 47, 0, 49, 15,
		1, 0, 8, 2, 0, 41, 41, 44, 44, 3, 0, 48, 57, 65, 90, 97, 122, 4, 0, 10,
		10, 13, 13, 34, 34, 92, 92, 4, 0, 10, 10, 13, 13, 39, 39, 92, 92, 2, 0,
		36, 36, 95, 95, 2, 0, 65, 70, 97, 102, 2, 0, 65, 90, 97, 122, 3, 0, 9,
		10, 12, 13, 32, 32, 186, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0,
		0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1,
		0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21,
		1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0,
		49, 1, 0, 0, 0, 1, 51, 1, 0, 0, 0, 3, 53, 1, 0, 0, 0, 5, 55, 1, 0, 0, 0,
		7, 57, 1, 0, 0, 0, 9, 59, 1, 0, 0, 0, 11, 64, 1, 0, 0, 0, 13, 70, 1, 0,
		0, 0, 15, 72, 1, 0, 0, 0, 17, 74, 1, 0, 0, 0, 19, 82, 1, 0, 0, 0, 21, 86,
		1, 0, 0, 0, 23, 110, 1, 0, 0, 0, 25, 114, 1, 0, 0, 0, 27, 116, 1, 0, 0,
		0, 29, 122, 1, 0, 0, 0, 31, 132, 1, 0, 0, 0, 33, 144, 1, 0, 0, 0, 35, 149,
		1, 0, 0, 0, 37, 154, 1, 0, 0, 0, 39, 158, 1, 0, 0, 0, 41, 163, 1, 0, 0,
		0, 43, 165, 1, 0, 0, 0, 45, 169, 1, 0, 0, 0, 47, 171, 1, 0, 0, 0, 49, 174,
		1, 0, 0, 0, 51, 52, 5, 40, 0, 0, 52, 2, 1, 0, 0, 0, 53, 54, 5, 41, 0, 0,
		54, 4, 1, 0, 0, 0, 55, 56, 5, 44, 0, 0, 56, 6, 1, 0, 0, 0, 57, 58, 5, 45,
		0, 0, 58, 8, 1, 0, 0, 0, 59, 60, 5, 116, 0, 0, 60, 61, 5, 114, 0, 0, 61,
		62, 5, 117, 0, 0, 62, 63, 5, 101, 0, 0, 63, 10, 1, 0, 0, 0, 64, 65, 5,
		102, 0, 0, 65, 66, 5, 97, 0, 0, 66, 67, 5, 108, 0, 0, 67, 68, 5, 115, 0,
		0, 68, 69, 5, 101, 0, 0, 69, 12, 1, 0, 0, 0, 70, 71, 5, 91, 0, 0, 71, 14,
		1, 0, 0, 0, 72, 73, 5, 93, 0, 0, 73, 16, 1, 0, 0, 0, 74, 78, 3, 39, 19,
		0, 75, 77, 3, 41, 20, 0, 76, 75, 1, 0, 0, 0, 77, 80, 1, 0, 0, 0, 78, 76,
		1, 0, 0, 0, 78, 79, 1, 0, 0, 0, 79, 18, 1, 0, 0, 0, 80, 78, 1, 0, 0, 0,
		81, 83, 3, 43, 21, 0, 82, 81, 1, 0, 0, 0, 83, 84, 1, 0, 0, 0, 84, 82, 1,
		0, 0, 0, 84, 85, 1, 0, 0, 0, 85, 20, 1, 0, 0, 0, 86, 87, 5, 48, 0, 0, 87,
		88, 5, 120, 0, 0, 88, 90, 1, 0, 0, 0, 89, 91, 3, 45, 22, 0, 90, 89, 1,
		0, 0, 0, 91, 92, 1, 0, 0, 0, 92, 90, 1, 0, 0, 0, 92, 93, 1, 0, 0, 0, 93,
		22, 1, 0, 0, 0, 94, 98, 5, 34, 0, 0, 95, 97, 3, 35, 17, 0, 96, 95, 1, 0,
		0, 0, 97, 100, 1, 0, 0, 0, 98, 96, 1, 0, 0, 0, 98, 99, 1, 0, 0, 0, 99,
		101, 1, 0, 0, 0, 100, 98, 1, 0, 0, 0, 101, 111, 5, 34, 0, 0, 102, 106,
		5, 39, 0, 0, 103, 105, 3, 37, 18, 0, 104, 103, 1, 0, 0, 0, 105, 108, 1,
		0, 0, 0, 106, 104, 1, 0, 0, 0, 106, 107, 1, 0, 0, 0, 107, 109, 1, 0, 0,
		0, 108, 106, 1, 0, 0, 0, 109, 111, 5, 39, 0, 0, 110, 94, 1, 0, 0, 0, 110,
		102, 1, 0, 0, 0, 111, 24, 1, 0, 0, 0, 112, 115, 3, 31, 15, 0, 113, 115,
		3, 33, 16, 0, 114, 112, 1, 0, 0, 0, 114, 113, 1, 0, 0, 0, 115, 26, 1, 0,
		0, 0, 116, 118, 5, 64, 0, 0, 117, 119, 8, 0, 0, 0, 118, 117, 1, 0, 0, 0,
		119, 120, 1, 0, 0, 0, 120, 118, 1, 0, 0, 0, 120, 121, 1, 0, 0, 0, 121,
		28, 1, 0, 0, 0, 122, 123, 7, 1, 0, 0, 123, 30, 1, 0, 0, 0, 124, 125, 5,
		116, 0, 0, 125, 126, 5, 114, 0, 0, 126, 127, 5, 117, 0, 0, 127, 133, 5,
		101, 0, 0, 128, 129, 5, 84, 0, 0, 129, 130, 5, 114, 0, 0, 130, 131, 5,
		117, 0, 0, 131, 133, 5, 101, 0, 0, 132, 124, 1, 0, 0, 0, 132, 128, 1, 0,
		0, 0, 133, 32, 1, 0, 0, 0, 134, 135, 5, 102, 0, 0, 135, 136, 5, 97, 0,
		0, 136, 137, 5, 108, 0, 0, 137, 138, 5, 115, 0, 0, 138, 145, 5, 101, 0,
		0, 139, 140, 5, 70, 0, 0, 140, 141, 5, 97, 0, 0, 141, 142, 5, 108, 0, 0,
		142, 143, 5, 115, 0, 0, 143, 145, 5, 101, 0, 0, 144, 134, 1, 0, 0, 0, 144,
		139, 1, 0, 0, 0, 145, 34, 1, 0, 0, 0, 146, 150, 8, 2, 0, 0, 147, 148, 5,
		92, 0, 0, 148, 150, 9, 0, 0, 0, 149, 146, 1, 0, 0, 0, 149, 147, 1, 0, 0,
		0, 150, 36, 1, 0, 0, 0, 151, 155, 8, 3, 0, 0, 152, 153, 5, 92, 0, 0, 153,
		155, 9, 0, 0, 0, 154, 151, 1, 0, 0, 0, 154, 152, 1, 0, 0, 0, 155, 38, 1,
		0, 0, 0, 156, 159, 3, 47, 23, 0, 157, 159, 7, 4, 0, 0, 158, 156, 1, 0,
		0, 0, 158, 157, 1, 0, 0, 0, 159, 40, 1, 0, 0, 0, 160, 164, 3, 47, 23, 0,
		161, 164, 7, 4, 0, 0, 162, 164, 3, 43, 21, 0, 163, 160, 1, 0, 0, 0, 163,
		161, 1, 0, 0, 0, 163, 162, 1, 0, 0, 0, 164, 42, 1, 0, 0, 0, 165, 166, 2,
		48, 57, 0, 166, 44, 1, 0, 0, 0, 167, 170, 3, 43, 21, 0, 168, 170, 7, 5,
		0, 0, 169, 167, 1, 0, 0, 0, 169, 168, 1, 0, 0, 0, 170, 46, 1, 0, 0, 0,
		171, 172, 7, 6, 0, 0, 172, 48, 1, 0, 0, 0, 173, 175, 7, 7, 0, 0, 174, 173,
		1, 0, 0, 0, 175, 176, 1, 0, 0, 0, 176, 174, 1, 0, 0, 0, 176, 177, 1, 0,
		0, 0, 177, 178, 1, 0, 0, 0, 178, 179, 6, 24, 0, 0, 179, 50, 1, 0, 0, 0,
		17, 0, 78, 84, 92, 98, 106, 110, 114, 120, 132, 144, 149, 154, 158, 163,
		169, 176, 1, 6, 0, 0,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// FuncLexerInit initializes any static state used to implement FuncLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewFuncLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func FuncLexerInit() {
	staticData := &FuncLexerLexerStaticData
	staticData.once.Do(funclexerLexerInit)
}

// NewFuncLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewFuncLexer(input antlr.CharStream) *FuncLexer {
	FuncLexerInit()
	l := new(FuncLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &FuncLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "Func.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// FuncLexer tokens.
const (
	FuncLexerT__0   = 1
	FuncLexerT__1   = 2
	FuncLexerT__2   = 3
	FuncLexerT__3   = 4
	FuncLexerT__4   = 5
	FuncLexerT__5   = 6
	FuncLexerT__6   = 7
	FuncLexerT__7   = 8
	FuncLexerNAME   = 9
	FuncLexerINT    = 10
	FuncLexerHEX    = 11
	FuncLexerSTRING = 12
	FuncLexerBOOL   = 13
	FuncLexerDOMAIN = 14
	FuncLexerWS     = 15
)

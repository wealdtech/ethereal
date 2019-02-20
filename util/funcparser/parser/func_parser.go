// Code generated from Func.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Func

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 16, 61, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4,
	8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2,
	3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 7, 4, 32, 10, 4, 12, 4, 14, 4, 35, 11, 4,
	5, 4, 37, 10, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 44, 10, 5, 3, 6, 5,
	6, 47, 10, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3,
	10, 3, 10, 3, 10, 3, 10, 2, 2, 11, 2, 4, 6, 8, 10, 12, 14, 16, 18, 2, 3,
	3, 2, 7, 8, 2, 58, 2, 20, 3, 2, 2, 2, 4, 26, 3, 2, 2, 2, 6, 36, 3, 2, 2,
	2, 8, 43, 3, 2, 2, 2, 10, 46, 3, 2, 2, 2, 12, 50, 3, 2, 2, 2, 14, 52, 3,
	2, 2, 2, 16, 54, 3, 2, 2, 2, 18, 56, 3, 2, 2, 2, 20, 21, 5, 4, 3, 2, 21,
	22, 7, 3, 2, 2, 22, 23, 5, 6, 4, 2, 23, 24, 7, 4, 2, 2, 24, 25, 7, 2, 2,
	3, 25, 3, 3, 2, 2, 2, 26, 27, 7, 11, 2, 2, 27, 5, 3, 2, 2, 2, 28, 33, 5,
	8, 5, 2, 29, 30, 7, 5, 2, 2, 30, 32, 5, 8, 5, 2, 31, 29, 3, 2, 2, 2, 32,
	35, 3, 2, 2, 2, 33, 31, 3, 2, 2, 2, 33, 34, 3, 2, 2, 2, 34, 37, 3, 2, 2,
	2, 35, 33, 3, 2, 2, 2, 36, 28, 3, 2, 2, 2, 36, 37, 3, 2, 2, 2, 37, 7, 3,
	2, 2, 2, 38, 44, 5, 10, 6, 2, 39, 44, 5, 12, 7, 2, 40, 44, 5, 14, 8, 2,
	41, 44, 5, 16, 9, 2, 42, 44, 5, 18, 10, 2, 43, 38, 3, 2, 2, 2, 43, 39,
	3, 2, 2, 2, 43, 40, 3, 2, 2, 2, 43, 41, 3, 2, 2, 2, 43, 42, 3, 2, 2, 2,
	44, 9, 3, 2, 2, 2, 45, 47, 7, 6, 2, 2, 46, 45, 3, 2, 2, 2, 46, 47, 3, 2,
	2, 2, 47, 48, 3, 2, 2, 2, 48, 49, 7, 12, 2, 2, 49, 11, 3, 2, 2, 2, 50,
	51, 7, 13, 2, 2, 51, 13, 3, 2, 2, 2, 52, 53, 7, 14, 2, 2, 53, 15, 3, 2,
	2, 2, 54, 55, 9, 2, 2, 2, 55, 17, 3, 2, 2, 2, 56, 57, 7, 9, 2, 2, 57, 58,
	5, 6, 4, 2, 58, 59, 7, 10, 2, 2, 59, 19, 3, 2, 2, 2, 6, 33, 36, 43, 46,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'('", "')'", "','", "'-'", "'true'", "'false'", "'['", "']'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "NAME", "INT", "HEX", "STRING", "BOOL",
	"WS",
}

var ruleNames = []string{
	"start", "funcName", "funcArgs", "arg", "intArg", "hexArg", "stringArg",
	"boolArg", "arrayArg",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type FuncParser struct {
	*antlr.BaseParser
}

func NewFuncParser(input antlr.TokenStream) *FuncParser {
	this := new(FuncParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Func.g4"

	return this
}

// FuncParser tokens.
const (
	FuncParserEOF    = antlr.TokenEOF
	FuncParserT__0   = 1
	FuncParserT__1   = 2
	FuncParserT__2   = 3
	FuncParserT__3   = 4
	FuncParserT__4   = 5
	FuncParserT__5   = 6
	FuncParserT__6   = 7
	FuncParserT__7   = 8
	FuncParserNAME   = 9
	FuncParserINT    = 10
	FuncParserHEX    = 11
	FuncParserSTRING = 12
	FuncParserBOOL   = 13
	FuncParserWS     = 14
)

// FuncParser rules.
const (
	FuncParserRULE_start     = 0
	FuncParserRULE_funcName  = 1
	FuncParserRULE_funcArgs  = 2
	FuncParserRULE_arg       = 3
	FuncParserRULE_intArg    = 4
	FuncParserRULE_hexArg    = 5
	FuncParserRULE_stringArg = 6
	FuncParserRULE_boolArg   = 7
	FuncParserRULE_arrayArg  = 8
)

// IStartContext is an interface to support dynamic dispatch.
type IStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStartContext differentiates from other interfaces.
	IsStartContext()
}

type StartContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStartContext() *StartContext {
	var p = new(StartContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FuncParserRULE_start
	return p
}

func (*StartContext) IsStartContext() {}

func NewStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartContext {
	var p = new(StartContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FuncParserRULE_start

	return p
}

func (s *StartContext) GetParser() antlr.Parser { return s.parser }

func (s *StartContext) FuncName() IFuncNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncNameContext)
}

func (s *StartContext) FuncArgs() IFuncArgsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncArgsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncArgsContext)
}

func (s *StartContext) EOF() antlr.TerminalNode {
	return s.GetToken(FuncParserEOF, 0)
}

func (s *StartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StartContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncListener); ok {
		listenerT.EnterStart(s)
	}
}

func (s *StartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncListener); ok {
		listenerT.ExitStart(s)
	}
}

func (p *FuncParser) Start() (localctx IStartContext) {
	localctx = NewStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, FuncParserRULE_start)

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
		p.SetState(18)
		p.FuncName()
	}
	{
		p.SetState(19)
		p.Match(FuncParserT__0)
	}
	{
		p.SetState(20)
		p.FuncArgs()
	}
	{
		p.SetState(21)
		p.Match(FuncParserT__1)
	}
	{
		p.SetState(22)
		p.Match(FuncParserEOF)
	}

	return localctx
}

// IFuncNameContext is an interface to support dynamic dispatch.
type IFuncNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncNameContext differentiates from other interfaces.
	IsFuncNameContext()
}

type FuncNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncNameContext() *FuncNameContext {
	var p = new(FuncNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FuncParserRULE_funcName
	return p
}

func (*FuncNameContext) IsFuncNameContext() {}

func NewFuncNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncNameContext {
	var p = new(FuncNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FuncParserRULE_funcName

	return p
}

func (s *FuncNameContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncNameContext) NAME() antlr.TerminalNode {
	return s.GetToken(FuncParserNAME, 0)
}

func (s *FuncNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncListener); ok {
		listenerT.EnterFuncName(s)
	}
}

func (s *FuncNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncListener); ok {
		listenerT.ExitFuncName(s)
	}
}

func (p *FuncParser) FuncName() (localctx IFuncNameContext) {
	localctx = NewFuncNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, FuncParserRULE_funcName)

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
		p.SetState(24)
		p.Match(FuncParserNAME)
	}

	return localctx
}

// IFuncArgsContext is an interface to support dynamic dispatch.
type IFuncArgsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncArgsContext differentiates from other interfaces.
	IsFuncArgsContext()
}

type FuncArgsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncArgsContext() *FuncArgsContext {
	var p = new(FuncArgsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FuncParserRULE_funcArgs
	return p
}

func (*FuncArgsContext) IsFuncArgsContext() {}

func NewFuncArgsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncArgsContext {
	var p = new(FuncArgsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FuncParserRULE_funcArgs

	return p
}

func (s *FuncArgsContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncArgsContext) AllArg() []IArgContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IArgContext)(nil)).Elem())
	var tst = make([]IArgContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IArgContext)
		}
	}

	return tst
}

func (s *FuncArgsContext) Arg(i int) IArgContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IArgContext)
}

func (s *FuncArgsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncArgsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncArgsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncListener); ok {
		listenerT.EnterFuncArgs(s)
	}
}

func (s *FuncArgsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncListener); ok {
		listenerT.ExitFuncArgs(s)
	}
}

func (p *FuncParser) FuncArgs() (localctx IFuncArgsContext) {
	localctx = NewFuncArgsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, FuncParserRULE_funcArgs)
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
	p.SetState(34)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<FuncParserT__3)|(1<<FuncParserT__4)|(1<<FuncParserT__5)|(1<<FuncParserT__6)|(1<<FuncParserINT)|(1<<FuncParserHEX)|(1<<FuncParserSTRING))) != 0 {
		{
			p.SetState(26)
			p.Arg()
		}
		p.SetState(31)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == FuncParserT__2 {
			{
				p.SetState(27)
				p.Match(FuncParserT__2)
			}
			{
				p.SetState(28)
				p.Arg()
			}

			p.SetState(33)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}

	return localctx
}

// IArgContext is an interface to support dynamic dispatch.
type IArgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArgContext differentiates from other interfaces.
	IsArgContext()
}

type ArgContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgContext() *ArgContext {
	var p = new(ArgContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FuncParserRULE_arg
	return p
}

func (*ArgContext) IsArgContext() {}

func NewArgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgContext {
	var p = new(ArgContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FuncParserRULE_arg

	return p
}

func (s *ArgContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgContext) IntArg() IIntArgContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIntArgContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIntArgContext)
}

func (s *ArgContext) HexArg() IHexArgContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IHexArgContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IHexArgContext)
}

func (s *ArgContext) StringArg() IStringArgContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStringArgContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStringArgContext)
}

func (s *ArgContext) BoolArg() IBoolArgContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBoolArgContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBoolArgContext)
}

func (s *ArgContext) ArrayArg() IArrayArgContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArrayArgContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArrayArgContext)
}

func (s *ArgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncListener); ok {
		listenerT.EnterArg(s)
	}
}

func (s *ArgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncListener); ok {
		listenerT.ExitArg(s)
	}
}

func (p *FuncParser) Arg() (localctx IArgContext) {
	localctx = NewArgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, FuncParserRULE_arg)

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

	p.SetState(41)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case FuncParserT__3, FuncParserINT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(36)
			p.IntArg()
		}

	case FuncParserHEX:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(37)
			p.HexArg()
		}

	case FuncParserSTRING:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(38)
			p.StringArg()
		}

	case FuncParserT__4, FuncParserT__5:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(39)
			p.BoolArg()
		}

	case FuncParserT__6:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(40)
			p.ArrayArg()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IIntArgContext is an interface to support dynamic dispatch.
type IIntArgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIntArgContext differentiates from other interfaces.
	IsIntArgContext()
}

type IntArgContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIntArgContext() *IntArgContext {
	var p = new(IntArgContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FuncParserRULE_intArg
	return p
}

func (*IntArgContext) IsIntArgContext() {}

func NewIntArgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntArgContext {
	var p = new(IntArgContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FuncParserRULE_intArg

	return p
}

func (s *IntArgContext) GetParser() antlr.Parser { return s.parser }

func (s *IntArgContext) INT() antlr.TerminalNode {
	return s.GetToken(FuncParserINT, 0)
}

func (s *IntArgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntArgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IntArgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncListener); ok {
		listenerT.EnterIntArg(s)
	}
}

func (s *IntArgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncListener); ok {
		listenerT.ExitIntArg(s)
	}
}

func (p *FuncParser) IntArg() (localctx IIntArgContext) {
	localctx = NewIntArgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, FuncParserRULE_intArg)
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
	p.SetState(44)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == FuncParserT__3 {
		{
			p.SetState(43)
			p.Match(FuncParserT__3)
		}

	}
	{
		p.SetState(46)
		p.Match(FuncParserINT)
	}

	return localctx
}

// IHexArgContext is an interface to support dynamic dispatch.
type IHexArgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsHexArgContext differentiates from other interfaces.
	IsHexArgContext()
}

type HexArgContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHexArgContext() *HexArgContext {
	var p = new(HexArgContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FuncParserRULE_hexArg
	return p
}

func (*HexArgContext) IsHexArgContext() {}

func NewHexArgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HexArgContext {
	var p = new(HexArgContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FuncParserRULE_hexArg

	return p
}

func (s *HexArgContext) GetParser() antlr.Parser { return s.parser }

func (s *HexArgContext) HEX() antlr.TerminalNode {
	return s.GetToken(FuncParserHEX, 0)
}

func (s *HexArgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HexArgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HexArgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncListener); ok {
		listenerT.EnterHexArg(s)
	}
}

func (s *HexArgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncListener); ok {
		listenerT.ExitHexArg(s)
	}
}

func (p *FuncParser) HexArg() (localctx IHexArgContext) {
	localctx = NewHexArgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, FuncParserRULE_hexArg)

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
		p.SetState(48)
		p.Match(FuncParserHEX)
	}

	return localctx
}

// IStringArgContext is an interface to support dynamic dispatch.
type IStringArgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStringArgContext differentiates from other interfaces.
	IsStringArgContext()
}

type StringArgContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStringArgContext() *StringArgContext {
	var p = new(StringArgContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FuncParserRULE_stringArg
	return p
}

func (*StringArgContext) IsStringArgContext() {}

func NewStringArgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StringArgContext {
	var p = new(StringArgContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FuncParserRULE_stringArg

	return p
}

func (s *StringArgContext) GetParser() antlr.Parser { return s.parser }

func (s *StringArgContext) STRING() antlr.TerminalNode {
	return s.GetToken(FuncParserSTRING, 0)
}

func (s *StringArgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringArgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StringArgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncListener); ok {
		listenerT.EnterStringArg(s)
	}
}

func (s *StringArgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncListener); ok {
		listenerT.ExitStringArg(s)
	}
}

func (p *FuncParser) StringArg() (localctx IStringArgContext) {
	localctx = NewStringArgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, FuncParserRULE_stringArg)

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
		p.SetState(50)
		p.Match(FuncParserSTRING)
	}

	return localctx
}

// IBoolArgContext is an interface to support dynamic dispatch.
type IBoolArgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBoolArgContext differentiates from other interfaces.
	IsBoolArgContext()
}

type BoolArgContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBoolArgContext() *BoolArgContext {
	var p = new(BoolArgContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FuncParserRULE_boolArg
	return p
}

func (*BoolArgContext) IsBoolArgContext() {}

func NewBoolArgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BoolArgContext {
	var p = new(BoolArgContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FuncParserRULE_boolArg

	return p
}

func (s *BoolArgContext) GetParser() antlr.Parser { return s.parser }
func (s *BoolArgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolArgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BoolArgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncListener); ok {
		listenerT.EnterBoolArg(s)
	}
}

func (s *BoolArgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncListener); ok {
		listenerT.ExitBoolArg(s)
	}
}

func (p *FuncParser) BoolArg() (localctx IBoolArgContext) {
	localctx = NewBoolArgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, FuncParserRULE_boolArg)
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
		p.SetState(52)
		_la = p.GetTokenStream().LA(1)

		if !(_la == FuncParserT__4 || _la == FuncParserT__5) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IArrayArgContext is an interface to support dynamic dispatch.
type IArrayArgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArrayArgContext differentiates from other interfaces.
	IsArrayArgContext()
}

type ArrayArgContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayArgContext() *ArrayArgContext {
	var p = new(ArrayArgContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FuncParserRULE_arrayArg
	return p
}

func (*ArrayArgContext) IsArrayArgContext() {}

func NewArrayArgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayArgContext {
	var p = new(ArrayArgContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FuncParserRULE_arrayArg

	return p
}

func (s *ArrayArgContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayArgContext) FuncArgs() IFuncArgsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncArgsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncArgsContext)
}

func (s *ArrayArgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayArgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayArgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncListener); ok {
		listenerT.EnterArrayArg(s)
	}
}

func (s *ArrayArgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncListener); ok {
		listenerT.ExitArrayArg(s)
	}
}

func (p *FuncParser) ArrayArg() (localctx IArrayArgContext) {
	localctx = NewArrayArgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, FuncParserRULE_arrayArg)

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
		p.SetState(54)
		p.Match(FuncParserT__6)
	}
	{
		p.SetState(55)
		p.FuncArgs()
	}
	{
		p.SetState(56)
		p.Match(FuncParserT__7)
	}

	return localctx
}

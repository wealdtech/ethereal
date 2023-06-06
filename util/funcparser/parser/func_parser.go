// Code generated from Func.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // Func
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type FuncParser struct {
	*antlr.BaseParser
}

var FuncParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func funcParserInit() {
	staticData := &FuncParserStaticData
	staticData.LiteralNames = []string{
		"", "'('", "')'", "','", "'-'", "'true'", "'false'", "'['", "']'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "NAME", "INT", "HEX", "STRING",
		"BOOL", "DOMAIN", "WS",
	}
	staticData.RuleNames = []string{
		"start", "funcName", "funcArgs", "arg", "intArg", "hexArg", "stringArg",
		"boolArg", "domainArg", "arrayArg",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 15, 64, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 1, 0, 1,
		0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 5, 2, 32, 8, 2,
		10, 2, 12, 2, 35, 9, 2, 3, 2, 37, 8, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1,
		3, 3, 3, 45, 8, 3, 1, 4, 3, 4, 48, 8, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6,
		1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 0, 0, 10, 0,
		2, 4, 6, 8, 10, 12, 14, 16, 18, 0, 1, 1, 0, 5, 6, 61, 0, 20, 1, 0, 0, 0,
		2, 26, 1, 0, 0, 0, 4, 36, 1, 0, 0, 0, 6, 44, 1, 0, 0, 0, 8, 47, 1, 0, 0,
		0, 10, 51, 1, 0, 0, 0, 12, 53, 1, 0, 0, 0, 14, 55, 1, 0, 0, 0, 16, 57,
		1, 0, 0, 0, 18, 59, 1, 0, 0, 0, 20, 21, 3, 2, 1, 0, 21, 22, 5, 1, 0, 0,
		22, 23, 3, 4, 2, 0, 23, 24, 5, 2, 0, 0, 24, 25, 5, 0, 0, 1, 25, 1, 1, 0,
		0, 0, 26, 27, 5, 9, 0, 0, 27, 3, 1, 0, 0, 0, 28, 33, 3, 6, 3, 0, 29, 30,
		5, 3, 0, 0, 30, 32, 3, 6, 3, 0, 31, 29, 1, 0, 0, 0, 32, 35, 1, 0, 0, 0,
		33, 31, 1, 0, 0, 0, 33, 34, 1, 0, 0, 0, 34, 37, 1, 0, 0, 0, 35, 33, 1,
		0, 0, 0, 36, 28, 1, 0, 0, 0, 36, 37, 1, 0, 0, 0, 37, 5, 1, 0, 0, 0, 38,
		45, 3, 8, 4, 0, 39, 45, 3, 10, 5, 0, 40, 45, 3, 12, 6, 0, 41, 45, 3, 14,
		7, 0, 42, 45, 3, 16, 8, 0, 43, 45, 3, 18, 9, 0, 44, 38, 1, 0, 0, 0, 44,
		39, 1, 0, 0, 0, 44, 40, 1, 0, 0, 0, 44, 41, 1, 0, 0, 0, 44, 42, 1, 0, 0,
		0, 44, 43, 1, 0, 0, 0, 45, 7, 1, 0, 0, 0, 46, 48, 5, 4, 0, 0, 47, 46, 1,
		0, 0, 0, 47, 48, 1, 0, 0, 0, 48, 49, 1, 0, 0, 0, 49, 50, 5, 10, 0, 0, 50,
		9, 1, 0, 0, 0, 51, 52, 5, 11, 0, 0, 52, 11, 1, 0, 0, 0, 53, 54, 5, 12,
		0, 0, 54, 13, 1, 0, 0, 0, 55, 56, 7, 0, 0, 0, 56, 15, 1, 0, 0, 0, 57, 58,
		5, 14, 0, 0, 58, 17, 1, 0, 0, 0, 59, 60, 5, 7, 0, 0, 60, 61, 3, 4, 2, 0,
		61, 62, 5, 8, 0, 0, 62, 19, 1, 0, 0, 0, 4, 33, 36, 44, 47,
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

// FuncParserInit initializes any static state used to implement FuncParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewFuncParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func FuncParserInit() {
	staticData := &FuncParserStaticData
	staticData.once.Do(funcParserInit)
}

// NewFuncParser produces a new parser instance for the optional input antlr.TokenStream.
func NewFuncParser(input antlr.TokenStream) *FuncParser {
	FuncParserInit()
	this := new(FuncParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &FuncParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
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
	FuncParserDOMAIN = 14
	FuncParserWS     = 15
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
	FuncParserRULE_domainArg = 8
	FuncParserRULE_arrayArg  = 9
)

// IStartContext is an interface to support dynamic dispatch.
type IStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FuncName() IFuncNameContext
	FuncArgs() IFuncArgsContext
	EOF() antlr.TerminalNode

	// IsStartContext differentiates from other interfaces.
	IsStartContext()
}

type StartContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStartContext() *StartContext {
	var p = new(StartContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncParserRULE_start
	return p
}

func InitEmptyStartContext(p *StartContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncParserRULE_start
}

func (*StartContext) IsStartContext() {}

func NewStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartContext {
	var p = new(StartContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FuncParserRULE_start

	return p
}

func (s *StartContext) GetParser() antlr.Parser { return s.parser }

func (s *StartContext) FuncName() IFuncNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncNameContext)
}

func (s *StartContext) FuncArgs() IFuncArgsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncArgsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

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

func (p *FuncParser) Start_() (localctx IStartContext) {
	localctx = NewStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, FuncParserRULE_start)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(20)
		p.FuncName()
	}
	{
		p.SetState(21)
		p.Match(FuncParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(22)
		p.FuncArgs()
	}
	{
		p.SetState(23)
		p.Match(FuncParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(24)
		p.Match(FuncParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFuncNameContext is an interface to support dynamic dispatch.
type IFuncNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NAME() antlr.TerminalNode

	// IsFuncNameContext differentiates from other interfaces.
	IsFuncNameContext()
}

type FuncNameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncNameContext() *FuncNameContext {
	var p = new(FuncNameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncParserRULE_funcName
	return p
}

func InitEmptyFuncNameContext(p *FuncNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncParserRULE_funcName
}

func (*FuncNameContext) IsFuncNameContext() {}

func NewFuncNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncNameContext {
	var p = new(FuncNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

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
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(26)
		p.Match(FuncParserNAME)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFuncArgsContext is an interface to support dynamic dispatch.
type IFuncArgsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllArg() []IArgContext
	Arg(i int) IArgContext

	// IsFuncArgsContext differentiates from other interfaces.
	IsFuncArgsContext()
}

type FuncArgsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncArgsContext() *FuncArgsContext {
	var p = new(FuncArgsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncParserRULE_funcArgs
	return p
}

func InitEmptyFuncArgsContext(p *FuncArgsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncParserRULE_funcArgs
}

func (*FuncArgsContext) IsFuncArgsContext() {}

func NewFuncArgsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncArgsContext {
	var p = new(FuncArgsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FuncParserRULE_funcArgs

	return p
}

func (s *FuncArgsContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncArgsContext) AllArg() []IArgContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IArgContext); ok {
			len++
		}
	}

	tst := make([]IArgContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IArgContext); ok {
			tst[i] = t.(IArgContext)
			i++
		}
	}

	return tst
}

func (s *FuncArgsContext) Arg(i int) IArgContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

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

	p.EnterOuterAlt(localctx, 1)
	p.SetState(36)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&23792) != 0 {
		{
			p.SetState(28)
			p.Arg()
		}
		p.SetState(33)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == FuncParserT__2 {
			{
				p.SetState(29)
				p.Match(FuncParserT__2)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(30)
				p.Arg()
			}

			p.SetState(35)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArgContext is an interface to support dynamic dispatch.
type IArgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IntArg() IIntArgContext
	HexArg() IHexArgContext
	StringArg() IStringArgContext
	BoolArg() IBoolArgContext
	DomainArg() IDomainArgContext
	ArrayArg() IArrayArgContext

	// IsArgContext differentiates from other interfaces.
	IsArgContext()
}

type ArgContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgContext() *ArgContext {
	var p = new(ArgContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncParserRULE_arg
	return p
}

func InitEmptyArgContext(p *ArgContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncParserRULE_arg
}

func (*ArgContext) IsArgContext() {}

func NewArgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgContext {
	var p = new(ArgContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FuncParserRULE_arg

	return p
}

func (s *ArgContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgContext) IntArg() IIntArgContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIntArgContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIntArgContext)
}

func (s *ArgContext) HexArg() IHexArgContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IHexArgContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IHexArgContext)
}

func (s *ArgContext) StringArg() IStringArgContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStringArgContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStringArgContext)
}

func (s *ArgContext) BoolArg() IBoolArgContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBoolArgContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBoolArgContext)
}

func (s *ArgContext) DomainArg() IDomainArgContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDomainArgContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDomainArgContext)
}

func (s *ArgContext) ArrayArg() IArrayArgContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayArgContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

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
	p.SetState(44)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case FuncParserT__3, FuncParserINT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(38)
			p.IntArg()
		}

	case FuncParserHEX:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(39)
			p.HexArg()
		}

	case FuncParserSTRING:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(40)
			p.StringArg()
		}

	case FuncParserT__4, FuncParserT__5:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(41)
			p.BoolArg()
		}

	case FuncParserDOMAIN:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(42)
			p.DomainArg()
		}

	case FuncParserT__6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(43)
			p.ArrayArg()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIntArgContext is an interface to support dynamic dispatch.
type IIntArgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INT() antlr.TerminalNode

	// IsIntArgContext differentiates from other interfaces.
	IsIntArgContext()
}

type IntArgContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIntArgContext() *IntArgContext {
	var p = new(IntArgContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncParserRULE_intArg
	return p
}

func InitEmptyIntArgContext(p *IntArgContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncParserRULE_intArg
}

func (*IntArgContext) IsIntArgContext() {}

func NewIntArgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntArgContext {
	var p = new(IntArgContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

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

	p.EnterOuterAlt(localctx, 1)
	p.SetState(47)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FuncParserT__3 {
		{
			p.SetState(46)
			p.Match(FuncParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(49)
		p.Match(FuncParserINT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IHexArgContext is an interface to support dynamic dispatch.
type IHexArgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	HEX() antlr.TerminalNode

	// IsHexArgContext differentiates from other interfaces.
	IsHexArgContext()
}

type HexArgContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHexArgContext() *HexArgContext {
	var p = new(HexArgContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncParserRULE_hexArg
	return p
}

func InitEmptyHexArgContext(p *HexArgContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncParserRULE_hexArg
}

func (*HexArgContext) IsHexArgContext() {}

func NewHexArgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HexArgContext {
	var p = new(HexArgContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

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
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(51)
		p.Match(FuncParserHEX)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStringArgContext is an interface to support dynamic dispatch.
type IStringArgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRING() antlr.TerminalNode

	// IsStringArgContext differentiates from other interfaces.
	IsStringArgContext()
}

type StringArgContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStringArgContext() *StringArgContext {
	var p = new(StringArgContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncParserRULE_stringArg
	return p
}

func InitEmptyStringArgContext(p *StringArgContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncParserRULE_stringArg
}

func (*StringArgContext) IsStringArgContext() {}

func NewStringArgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StringArgContext {
	var p = new(StringArgContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

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
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(53)
		p.Match(FuncParserSTRING)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
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
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBoolArgContext() *BoolArgContext {
	var p = new(BoolArgContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncParserRULE_boolArg
	return p
}

func InitEmptyBoolArgContext(p *BoolArgContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncParserRULE_boolArg
}

func (*BoolArgContext) IsBoolArgContext() {}

func NewBoolArgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BoolArgContext {
	var p = new(BoolArgContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

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

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(55)
		_la = p.GetTokenStream().LA(1)

		if !(_la == FuncParserT__4 || _la == FuncParserT__5) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDomainArgContext is an interface to support dynamic dispatch.
type IDomainArgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DOMAIN() antlr.TerminalNode

	// IsDomainArgContext differentiates from other interfaces.
	IsDomainArgContext()
}

type DomainArgContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDomainArgContext() *DomainArgContext {
	var p = new(DomainArgContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncParserRULE_domainArg
	return p
}

func InitEmptyDomainArgContext(p *DomainArgContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncParserRULE_domainArg
}

func (*DomainArgContext) IsDomainArgContext() {}

func NewDomainArgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DomainArgContext {
	var p = new(DomainArgContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FuncParserRULE_domainArg

	return p
}

func (s *DomainArgContext) GetParser() antlr.Parser { return s.parser }

func (s *DomainArgContext) DOMAIN() antlr.TerminalNode {
	return s.GetToken(FuncParserDOMAIN, 0)
}

func (s *DomainArgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DomainArgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DomainArgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncListener); ok {
		listenerT.EnterDomainArg(s)
	}
}

func (s *DomainArgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncListener); ok {
		listenerT.ExitDomainArg(s)
	}
}

func (p *FuncParser) DomainArg() (localctx IDomainArgContext) {
	localctx = NewDomainArgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, FuncParserRULE_domainArg)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(57)
		p.Match(FuncParserDOMAIN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArrayArgContext is an interface to support dynamic dispatch.
type IArrayArgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FuncArgs() IFuncArgsContext

	// IsArrayArgContext differentiates from other interfaces.
	IsArrayArgContext()
}

type ArrayArgContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayArgContext() *ArrayArgContext {
	var p = new(ArrayArgContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncParserRULE_arrayArg
	return p
}

func InitEmptyArrayArgContext(p *ArrayArgContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FuncParserRULE_arrayArg
}

func (*ArrayArgContext) IsArrayArgContext() {}

func NewArrayArgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayArgContext {
	var p = new(ArrayArgContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FuncParserRULE_arrayArg

	return p
}

func (s *ArrayArgContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayArgContext) FuncArgs() IFuncArgsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncArgsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

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
	p.EnterRule(localctx, 18, FuncParserRULE_arrayArg)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(59)
		p.Match(FuncParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(60)
		p.FuncArgs()
	}
	{
		p.SetState(61)
		p.Match(FuncParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// Code generated from Func.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // Func
import "github.com/antlr4-go/antlr/v4"

// BaseFuncListener is a complete listener for a parse tree produced by FuncParser.
type BaseFuncListener struct{}

var _ FuncListener = &BaseFuncListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseFuncListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseFuncListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseFuncListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseFuncListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStart is called when production start is entered.
func (s *BaseFuncListener) EnterStart(ctx *StartContext) {}

// ExitStart is called when production start is exited.
func (s *BaseFuncListener) ExitStart(ctx *StartContext) {}

// EnterFuncName is called when production funcName is entered.
func (s *BaseFuncListener) EnterFuncName(ctx *FuncNameContext) {}

// ExitFuncName is called when production funcName is exited.
func (s *BaseFuncListener) ExitFuncName(ctx *FuncNameContext) {}

// EnterFuncArgs is called when production funcArgs is entered.
func (s *BaseFuncListener) EnterFuncArgs(ctx *FuncArgsContext) {}

// ExitFuncArgs is called when production funcArgs is exited.
func (s *BaseFuncListener) ExitFuncArgs(ctx *FuncArgsContext) {}

// EnterArg is called when production arg is entered.
func (s *BaseFuncListener) EnterArg(ctx *ArgContext) {}

// ExitArg is called when production arg is exited.
func (s *BaseFuncListener) ExitArg(ctx *ArgContext) {}

// EnterIntArg is called when production intArg is entered.
func (s *BaseFuncListener) EnterIntArg(ctx *IntArgContext) {}

// ExitIntArg is called when production intArg is exited.
func (s *BaseFuncListener) ExitIntArg(ctx *IntArgContext) {}

// EnterHexArg is called when production hexArg is entered.
func (s *BaseFuncListener) EnterHexArg(ctx *HexArgContext) {}

// ExitHexArg is called when production hexArg is exited.
func (s *BaseFuncListener) ExitHexArg(ctx *HexArgContext) {}

// EnterStringArg is called when production stringArg is entered.
func (s *BaseFuncListener) EnterStringArg(ctx *StringArgContext) {}

// ExitStringArg is called when production stringArg is exited.
func (s *BaseFuncListener) ExitStringArg(ctx *StringArgContext) {}

// EnterBoolArg is called when production boolArg is entered.
func (s *BaseFuncListener) EnterBoolArg(ctx *BoolArgContext) {}

// ExitBoolArg is called when production boolArg is exited.
func (s *BaseFuncListener) ExitBoolArg(ctx *BoolArgContext) {}

// EnterDomainArg is called when production domainArg is entered.
func (s *BaseFuncListener) EnterDomainArg(ctx *DomainArgContext) {}

// ExitDomainArg is called when production domainArg is exited.
func (s *BaseFuncListener) ExitDomainArg(ctx *DomainArgContext) {}

// EnterArrayArg is called when production arrayArg is entered.
func (s *BaseFuncListener) EnterArrayArg(ctx *ArrayArgContext) {}

// ExitArrayArg is called when production arrayArg is exited.
func (s *BaseFuncListener) ExitArrayArg(ctx *ArrayArgContext) {}

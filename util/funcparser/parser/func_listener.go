// Code generated from Func.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // Func
import "github.com/antlr4-go/antlr/v4"

// FuncListener is a complete listener for a parse tree produced by FuncParser.
type FuncListener interface {
	antlr.ParseTreeListener

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

	// EnterFuncName is called when entering the funcName production.
	EnterFuncName(c *FuncNameContext)

	// EnterFuncArgs is called when entering the funcArgs production.
	EnterFuncArgs(c *FuncArgsContext)

	// EnterArg is called when entering the arg production.
	EnterArg(c *ArgContext)

	// EnterIntArg is called when entering the intArg production.
	EnterIntArg(c *IntArgContext)

	// EnterHexArg is called when entering the hexArg production.
	EnterHexArg(c *HexArgContext)

	// EnterStringArg is called when entering the stringArg production.
	EnterStringArg(c *StringArgContext)

	// EnterBoolArg is called when entering the boolArg production.
	EnterBoolArg(c *BoolArgContext)

	// EnterDomainArg is called when entering the domainArg production.
	EnterDomainArg(c *DomainArgContext)

	// EnterArrayArg is called when entering the arrayArg production.
	EnterArrayArg(c *ArrayArgContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitFuncName is called when exiting the funcName production.
	ExitFuncName(c *FuncNameContext)

	// ExitFuncArgs is called when exiting the funcArgs production.
	ExitFuncArgs(c *FuncArgsContext)

	// ExitArg is called when exiting the arg production.
	ExitArg(c *ArgContext)

	// ExitIntArg is called when exiting the intArg production.
	ExitIntArg(c *IntArgContext)

	// ExitHexArg is called when exiting the hexArg production.
	ExitHexArg(c *HexArgContext)

	// ExitStringArg is called when exiting the stringArg production.
	ExitStringArg(c *StringArgContext)

	// ExitBoolArg is called when exiting the boolArg production.
	ExitBoolArg(c *BoolArgContext)

	// ExitDomainArg is called when exiting the domainArg production.
	ExitDomainArg(c *DomainArgContext)

	// ExitArrayArg is called when exiting the arrayArg production.
	ExitArrayArg(c *ArrayArgContext)
}

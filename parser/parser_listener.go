// Code generated from Parser.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser // Parser
import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// ParserListener is a complete listener for a parse tree produced by ParserParser.
type ParserListener interface {
	antlr.ParseTreeListener

	// EnterPrintExpr is called when entering the printExpr production.
	EnterPrintExpr(c *PrintExprContext)

	// EnterAssign is called when entering the assign production.
	EnterAssign(c *AssignContext)

	// EnterBlank is called when entering the blank production.
	EnterBlank(c *BlankContext)

	// EnterMultiDiv is called when entering the multiDiv production.
	EnterMultiDiv(c *MultiDivContext)

	// EnterParen is called when entering the paren production.
	EnterParen(c *ParenContext)

	// EnterAddSub is called when entering the addSub production.
	EnterAddSub(c *AddSubContext)

	// EnterId is called when entering the id production.
	EnterId(c *IdContext)

	// EnterInt is called when entering the int production.
	EnterInt(c *IntContext)

	// ExitPrintExpr is called when exiting the printExpr production.
	ExitPrintExpr(c *PrintExprContext)

	// ExitAssign is called when exiting the assign production.
	ExitAssign(c *AssignContext)

	// ExitBlank is called when exiting the blank production.
	ExitBlank(c *BlankContext)

	// ExitMultiDiv is called when exiting the multiDiv production.
	ExitMultiDiv(c *MultiDivContext)

	// ExitParen is called when exiting the paren production.
	ExitParen(c *ParenContext)

	// ExitAddSub is called when exiting the addSub production.
	ExitAddSub(c *AddSubContext)

	// ExitId is called when exiting the id production.
	ExitId(c *IdContext)

	// ExitInt is called when exiting the int production.
	ExitInt(c *IntContext)
}

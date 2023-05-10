// Code generated from Parser.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser // Parser
import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// BaseParserListener is a complete listener for a parse tree produced by ParserParser.
type BaseParserListener struct{}

var _ ParserListener = &BaseParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterPrintExpr is called when production printExpr is entered.
func (s *BaseParserListener) EnterPrintExpr(ctx *PrintExprContext) {}

// ExitPrintExpr is called when production printExpr is exited.
func (s *BaseParserListener) ExitPrintExpr(ctx *PrintExprContext) {}

// EnterAssign is called when production assign is entered.
func (s *BaseParserListener) EnterAssign(ctx *AssignContext) {}

// ExitAssign is called when production assign is exited.
func (s *BaseParserListener) ExitAssign(ctx *AssignContext) {}

// EnterBlank is called when production blank is entered.
func (s *BaseParserListener) EnterBlank(ctx *BlankContext) {}

// ExitBlank is called when production blank is exited.
func (s *BaseParserListener) ExitBlank(ctx *BlankContext) {}

// EnterMultiDiv is called when production multiDiv is entered.
func (s *BaseParserListener) EnterMultiDiv(ctx *MultiDivContext) {}

// ExitMultiDiv is called when production multiDiv is exited.
func (s *BaseParserListener) ExitMultiDiv(ctx *MultiDivContext) {}

// EnterParen is called when production paren is entered.
func (s *BaseParserListener) EnterParen(ctx *ParenContext) {}

// ExitParen is called when production paren is exited.
func (s *BaseParserListener) ExitParen(ctx *ParenContext) {}

// EnterAddSub is called when production addSub is entered.
func (s *BaseParserListener) EnterAddSub(ctx *AddSubContext) {}

// ExitAddSub is called when production addSub is exited.
func (s *BaseParserListener) ExitAddSub(ctx *AddSubContext) {}

// EnterId is called when production id is entered.
func (s *BaseParserListener) EnterId(ctx *IdContext) {}

// ExitId is called when production id is exited.
func (s *BaseParserListener) ExitId(ctx *IdContext) {}

// EnterInt is called when production int is entered.
func (s *BaseParserListener) EnterInt(ctx *IntContext) {}

// ExitInt is called when production int is exited.
func (s *BaseParserListener) ExitInt(ctx *IntContext) {}

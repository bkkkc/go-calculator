package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/bkkkc/go-calculator/parser"
)

type TreeShapeListener struct {
	*parser.BaseParserListener
	id map[string]int
	stack   []int
	lastErr error
}

func NewTreeShapeListener() *TreeShapeListener {
	t := new(TreeShapeListener)
	t.id = make(map[string]int)
	return t
}

func (t *TreeShapeListener) clearStack() {
	if t.lastErr != nil {
		fmt.Println(t.lastErr)
		t.lastErr = nil
	}
	t.stack = t.stack[:0]
}

func (t *TreeShapeListener) ExitPrintExpr(ctx *parser.PrintExprContext) {
	fmt.Println(t.stack[0])
}

func (t *TreeShapeListener) ExitAssign(ctx *parser.AssignContext) {
	t.id[ctx.ID().GetText()] = t.stack[len(t.stack)-1]
	t.stack = t.stack[:len(t.stack)-1]
}

func (t *TreeShapeListener) ExitInt(ctx *parser.IntContext) {
	i, err := strconv.Atoi(ctx.GetText())
	if err != nil {
		t.lastErr = err
		return
	}
	t.stack = append(t.stack, i)
}

func (t *TreeShapeListener) ExitId(ctx *parser.IdContext) {
	data, ok := t.id[ctx.GetText()]
	if !ok {
		panic(fmt.Sprintf("%s not define", ctx.GetText()))
	}
	t.stack = append(t.stack, data)
}

func (t *TreeShapeListener) ExitAddSub(ctx *parser.AddSubContext) {
	l, r := t.stack[len(t.stack)-2], t.stack[len(t.stack)-1]
	t.stack = t.stack[:len(t.stack)-2]
	var tmp int

	if ctx.ADD() != nil {
		tmp = l + r
	} else {
		tmp = l - r
	}
	t.stack = append(t.stack, tmp)
}

func (t *TreeShapeListener) ExitMultiDiv(ctx *parser.MultiDivContext) {
	l, r := t.stack[len(t.stack)-2], t.stack[len(t.stack)-1]
	t.stack = t.stack[:len(t.stack)-2]
	var tmp int
	if ctx.MUL() != nil {
		tmp = l * r
	} else {
		tmp = l / r
	}
	t.stack = append(t.stack, tmp)
}

type MyErrorListeners struct {
	antlr.DefaultErrorListener
}

func (m *MyErrorListeners) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	p := recognizer.(antlr.Parser)
	stack := p.GetRuleInvocationStack(p.GetParserRuleContext())
	info := fmt.Sprintln("rule stack:", stack, "line", line, ":", line, "at", offendingSymbol, ":", msg)
	panic(info)
}

// func main() {
// 	var lang = "a=a2*2\r\n"
// 	input := antlr.NewInputStream(lang)
// 	lexer := parser.NewParserLexer(input)
// 	stream := antlr.NewCommonTokenStream(lexer, 0)
// 	p := parser.NewParserParser(stream)
// 	p.RemoveErrorListeners()
// 	p.AddErrorListener(new(MyErrorListeners))
// 	//p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
// 	p.BuildParseTrees = true
// 	tree := p.Stat()
// 	antlr.ParseTreeWalkerDefault.Walk(NewTreeShapeListener(), tree)
// }

func main() {
	p := parser.NewParserParser(nil)
	p.RemoveErrorListeners()
	p.AddErrorListener(new(MyErrorListeners))
	p.BuildParseTrees = true

	cal := NewTreeShapeListener()

	fmt.Print(">>> ")
	// tmpdata := `1+2
	// a=3
	// a
	// 4+7`
	// input := bytes.NewBuffer([]byte(tmpdata))
	console := bufio.NewReader(os.Stdin)
	// console := bufio.NewReader(input)
	for {
		func() {
			defer func() {
				info := recover()
				if info != nil {
					fmt.Println(info)
				}
				cal.clearStack()
				fmt.Print(">>> ")
			}()

			expr, err := console.ReadString('\n')
			if err != nil {
				return
			}
			stream := antlr.NewInputStream(expr + "\n")
			lexer := parser.NewParserLexer(stream)

			tokens := antlr.NewCommonTokenStream(lexer, 0)
			p.SetInputStream(tokens)
			tree := p.Stat()
			antlr.ParseTreeWalkerDefault.Walk(cal, tree)

		}()
	}
}

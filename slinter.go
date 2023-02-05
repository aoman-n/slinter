package slinter

import (
	"fmt"
	"go/ast"
	"go/types"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = "slinter is ..."

// Analyzer is ...
var Analyzer = &analysis.Analyzer{
	Name: "slinter",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

const (
	defaultMaxLines = 6
	defaultMaxArgs  = 4
)

var (
	maxLines = defaultMaxLines
	maxArgs  = defaultMaxArgs
)

func createOverLinesMessage(funcName string, lines int) string {
	return fmt.Sprintf(
		"func `%s` has `%d` lines, but no more than %d lines are recommended",
		funcName, lines, maxLines)
}

func createOverArgsMessage(funcName string, args int) string {
	return fmt.Sprintf(
		"func `%s` has `%d` lines, but no more than %d args are recommended",
		funcName, args, maxArgs)
}

func run(pass *analysis.Pass) (any, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.FuncDecl)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		switch n := n.(type) {
		// メソッドも*ast.FuncDecl
		case *ast.FuncDecl:
			argsLen := getArgs(pass, n)
			if argsLen > maxArgs {
				pass.Reportf(n.Pos(), createOverArgsMessage(n.Name.Name, argsLen))
			}

			linesLen := getLines(pass, n)
			if linesLen > maxLines {
				pass.Reportf(n.Pos(), createOverLinesMessage(n.Name.Name, linesLen))
			}
		}
	})

	return nil, nil
}

// 関数の行数を返す。関数宣言と`{}`自体の行も含めた`{`から`}`まで行数。
func getLines(pass *analysis.Pass, n *ast.FuncDecl) int {
	// pos := pass.Fset.Position(n.Pos())
	// end := pass.Fset.Position(n.End())

	lBracePos := pass.Fset.Position(n.Body.Lbrace)
	rBracePos := pass.Fset.Position(n.Body.Rbrace)

	return rBracePos.Line - lBracePos.Line + 1
}

func getArgs(pass *analysis.Pass, n *ast.FuncDecl) int {
	obj := pass.TypesInfo.ObjectOf(n.Name) // n.NameはIdent
	typ := obj.Type()
	signature := typ.(*types.Signature)
	// Params().Len()は可変長引数も1としてカウント
	return signature.Params().Len()
}

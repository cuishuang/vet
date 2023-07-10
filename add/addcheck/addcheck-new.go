package addcheck

import (
	"fmt"
	"go/ast"
	"go/token"
	"go/types"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var Analyzer2 = &analysis.Analyzer{
	Name:     "addlint",
	Doc:      "reports integer additions",
	Run:      run2,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

func run2(pass *analysis.Pass) (interface{}, error) {
	// get the inspector. This will not panic because inspect.Analyzer is part
	// of `Requires`. go/analysis will populate the `pass.ResultOf` map with
	// the prerequisite analyzers.
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	// the inspector has a `filter` feature that enables type-based filtering
	// The anonymous function will be only called for the ast nodes whose type
	// matches an element in the filter
	nodeFilter := []ast.Node{
		(*ast.BinaryExpr)(nil),
	}

	fmt.Println(111)
	// this is basically the same as ast.Inspect(), only we don't return a
	// boolean anymore as it'll visit all the nodes based on the filter.
	inspect.Preorder(nodeFilter, func(n ast.Node) {

		fmt.Println(222)
		be := n.(*ast.BinaryExpr)
		if be.Op != token.ADD {
			return
		}
		if _, ok := be.X.(*ast.BasicLit); !ok {
			return
		}
		if _, ok := be.Y.(*ast.BasicLit); !ok {
			return
		}
		isInteger := func(expr ast.Expr) bool {
			t := pass.TypesInfo.TypeOf(expr)
			if t == nil {
				return false
			}
			bt, ok := t.Underlying().(*types.Basic)
			if !ok {
				return false
			}
			if (bt.Info() & types.IsInteger) == 0 {
				return false
			}
			return true
		}
		// check that both left and right hand side are integers
		if !isInteger(be.X) || !isInteger(be.Y) {
			return
		}
		pass.Reportf(be.Pos(), "integer addition found %q",
			render(pass.Fset, be))
	})
	return nil, nil
}

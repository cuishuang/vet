// Package appendcheck defines an Analyzer that reports xxxxxx
package appendcheck

import (
	"go/ast"
	"go/types"

	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "appendc",
	Doc:  "reports 666666",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	//return nil, errors.New("not implemented yet")

	for _, node := range pass.Files {
		ast.Inspect(node, func(n ast.Node) bool {
			if call, ok := n.(*ast.CallExpr); ok {
				if ident, ok := call.Fun.(*ast.Ident); ok && ident.Name == "append" {
					if _, ok := pass.TypesInfo.Uses[ident].(*types.Builtin); ok {
						if len(call.Args) == 1 {
							pass.ReportRangef(call, "append with no values啊啊啊！！！")
						}
					}

				}
			}
			return true
		})
	}

	return nil, nil
}

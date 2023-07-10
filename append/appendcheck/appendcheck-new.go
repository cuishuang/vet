package appendcheck

import (
	//"fmt"
	"go/ast"
	"go/types"

	//"github.com/davecgh/go-spew/spew"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var Analyzer2 = &analysis.Analyzer{
	Name:     "appendcheck",
	Doc:      "reports integer 8888888",
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
	//nodeFilter := []ast.Node{
	//	(*ast.BinaryExpr)(nil),
	//}
	nodeFilter := []ast.Node{
		(*ast.CallExpr)(nil),
	}
	// this is basically the same as ast.Inspect(), only we don't return a
	// boolean anymore as it'll visit all the nodes based on the filter.

	//	fmt.Println(1111)

	//spew.Dump()

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		//fmt.Println(2222)
		call := n.(*ast.CallExpr)
		//if call, ok := n.(*ast.CallExpr); ok {
		if ident, ok := call.Fun.(*ast.Ident); ok && ident.Name == "append" {
			if _, ok := pass.TypesInfo.Uses[ident].(*types.Builtin); ok {
				if len(call.Args) == 1 {
					pass.ReportRangef(call, "append with no values啊啊啊！！！")
				}
			}
		}
		//}
	})

	//inspect.Preorder(nodeFilter, func(n ast.Node) {
	//	for _, node := range pass.Files {
	//		ast.Inspect(node, func(n ast.Node) bool {
	//			if call, ok := n.(*ast.CallExpr); ok {
	//				if ident, ok := call.Fun.(*ast.Ident); ok && ident.Name == "append" {
	//					if _, ok := pass.TypesInfo.Uses[ident].(*types.Builtin); ok {
	//						if len(call.Args) == 1 {
	//							pass.ReportRangef(call, "append with no values啊啊啊！！！")
	//						}
	//					}
	//
	//				}
	//			}
	//			return true
	//		})
	//	}
	//})
	return nil, nil
}

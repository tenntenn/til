package pointer

import (
	"fmt"
	"go/ast"
	"go/types"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/buildssa"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
	"golang.org/x/tools/go/pointer"
	"golang.org/x/tools/go/ssa"
)

var Analyzer = &analysis.Analyzer{
	Name: "pointer",
	Doc:  Doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
		buildssa.Analyzer,
	},
}

const Doc = ``

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	pkg := pass.ResultOf[buildssa.Analyzer].(*buildssa.SSA).Pkg

	config := &pointer.Config{
		Mains: []*ssa.Package{pkg},
	}

	nodeFilter := []ast.Node{
		(*ast.Ident)(nil),
	}

	inspect.WithStack(nodeFilter, func(n ast.Node, push bool, stack []ast.Node) bool {
		if !push {
			return false
		}

		switch n := n.(type) {
		case *ast.Ident:
			obj, ok := pass.TypesInfo.Defs[n].(*types.Var)
			if !ok {
				return false
			}
			ref := getRef(stack)
			v, _ := pkg.Prog.VarValue(obj, pkg, ref)
			if pointer.CanPoint(v.Type()) {
				config.AddQuery(v)
			}
		}

		return false
	})

	result, err := pointer.Analyze(config)
	if err != nil {
		return nil, err
	}

	//fmt.Printf("%#v\n", result)

	for _, ptr := range result.Queries {
		for _, l := range ptr.PointsTo().Labels() {
			fmt.Println(pass.Fset.Position(l.Pos()), l)
		}
	}

	return nil, nil
}

func getRef(stack []ast.Node) []ast.Node {
	ref := make([]ast.Node, len(stack))
	for i := range stack {
		ref[len(ref)-i-1] = stack[i]
	}
	return ref
}

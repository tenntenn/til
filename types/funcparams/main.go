package main

import (
	"fmt"
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"log"
)

const src = `package main
import "fmt"

func main() {
	fmt.Println("hello")
}`

func main() {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "main.go", src, 0)
	if err != nil {
		log.Fatal(err)
	}

	info := &types.Info{
		Types: map[ast.Expr]types.TypeAndValue{},
	}

	config := &types.Config{
		Importer: importer.Default(),
	}

	_, err = config.Check("main", fset, []*ast.File{f}, info)
	if err != nil {
		log.Fatal(err)
	}

	ast.Inspect(f, func(n ast.Node) bool {
		switch n := n.(type) {
		case *ast.CallExpr:
			if sig, ok := info.TypeOf(n.Fun).(*types.Signature); ok {
				params := sig.Params()
				for i := 0; i < params.Len(); i++ {
					fmt.Println(params.At(i).Type())
				}
			}
		}
		return true
	})
}

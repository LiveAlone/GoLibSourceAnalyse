package gopatch

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

var srcCode = `
package hello

import "fmt"

func greet() {
    var msg = "Hello World!"
    fmt.Println(msg)
}

func foo() {
    var msg = "123"
	fmt.Println(msg)
}
`

func TestAST(t *testing.T) {
	var err error
	fileSet := token.NewFileSet()
	f, err := parser.ParseFile(fileSet, "main.go", srcCode, 0)
	if err != nil {
		fmt.Printf("err = %s", err)
	}

	// 输出AST 结构体
	//err = ast.Print(fileSet, f)
	//fmt.Println(err)

	ast.Inspect(f, func(n ast.Node) bool {
		// Called recursively.
		//ast.Print(fset, n)
		funcDecl, ok := n.(*ast.FuncDecl)
		if !ok {
			return true
		}
		fmt.Println(funcDecl.Name)
		return true
	})
}

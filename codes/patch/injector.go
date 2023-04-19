package patch

import (
	"github.com/pkg/errors"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"strings"
)

const INJECT = `
package main

//import "fmt"

func a(){
	__traceStack()
}
`

type Injector struct {
	MyImport *ast.ImportSpec
	MyStmt   ast.Stmt
}

func NewInjector() *Injector {
	i := &Injector{}
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", INJECT, 0)
	if err != nil {
		panic(err)
	}

	for _, d := range f.Decls {
		if fd, ok := d.(*ast.FuncDecl); ok {
			i.MyStmt = fd.Body.List[0]
		}
	}
	return i
}

func (i *Injector) InjectFunc(f ast.Decl) error {
	fd, ok := f.(*ast.FuncDecl)
	if !ok {
		return nil
	}

	// 函数调用添加日志
	newList := make([]ast.Stmt, 0, len(fd.Body.List)+1)
	newList = append(newList, i.MyStmt)
	newList = append(newList, fd.Body.List...)
	fd.Body.List = newList
	return nil
}

func (i *Injector) InjectFile(path string) error {
	// 解析目标文件
	fset := token.NewFileSet()
	fbytes, err := os.ReadFile(path)
	if err != nil {
		return errors.Wrap(err, "inject read file error")
	}
	index := strings.LastIndex(path, `/`)
	f, err := parser.ParseFile(fset, path[index+1:], fbytes, 0)
	if err != nil {
		return err
	}

	for _, decl := range f.Decls {
		err = i.InjectFunc(decl)
		if err != nil {
			return errors.Wrap(err, "inject func error")
		}
	}
	return nil
}

package patch

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"strings"
)

type Parser struct {
	filename  string
	tokenFSet *token.FileSet
	astFile   *ast.File
}

func NewParser(filename string) *Parser {
	return &Parser{
		filename:  filename,
		tokenFSet: token.NewFileSet(),
	}
}

func (entity *Parser) Parse() error {
	src, err := os.ReadFile(entity.filename)
	index := strings.LastIndex(entity.filename, `/`)
	f, err := parser.ParseFile(entity.tokenFSet, entity.filename[index+1:], src, 0)
	if err != nil {
		return err
	}
	entity.astFile = f
	return nil
}

func (entity *Parser) PkgName() string {
	return entity.astFile.Name.Name
}

func (entity *Parser) ForEachDecl(f func(ast.Decl) error) (err error) {
	for _, decl := range entity.astFile.Decls {
		err = f(decl)
		if err != nil {
			return err
		}
	}
	return nil
}

func (entity *Parser) GetAst() *ast.File {
	return entity.astFile
}

func (entity *Parser) GetTokenFSet() *token.FileSet {
	return entity.tokenFSet
}

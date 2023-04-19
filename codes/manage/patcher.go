package manager

import (
	"bytes"
	"github.com/LiveAlone/GoLibSourceAnalyse/codes/patch"
	"github.com/pkg/errors"
	"go/ast"
	"go/printer"
	"os"
)

// Patcher 补偿程序
type Patcher struct {
	// 补偿目录
	path string
	dirs []string
	deep int

	// 注入工具
	inject *patch.Injector
	backup patch.Backup
}

func NewPatcher(path string, deep int) *Patcher {
	return &Patcher{
		path: path,
		deep: deep,
		dirs: patch.TreeDir(path, deep),

		inject: patch.NewInjector(),
	}
}

// Backup 备份
func (entity *Patcher) Backup() (err error) {
	for _, dir := range entity.dirs {
		err = entity.backup.BackupPath(dir)
		if err != nil {
			return err
		}
	}
	return nil
}

// Inject 注入
func (entity *Patcher) Inject() (err error) {
	for _, dir := range entity.dirs {
		for _, filename := range patch.ListGoFile(dir, false) {

			parser := patch.NewParser(filename)
			if err := parser.Parse(); err != nil {
				return errors.Wrap(err, "parser parse error")
			}

			err = parser.ForEachDecl(func(decl ast.Decl) error {
				return entity.inject.InjectFunc(decl)
			})
			if err != nil {
				return errors.Wrap(err, "inject func error")
			}

			// write to new file
			var buf bytes.Buffer
			err = printer.Fprint(&buf, parser.GetTokenFSet(), parser.GetAst())
			if err != nil {
				return errors.Wrap(err, "printer fprint error")
			}
			//fmt.Println(buf.String())
			if err := os.WriteFile(filename, buf.Bytes(), os.ModeExclusive); err != nil {
				return errors.Wrap(err, "write file error")
			}

			// write helper
			dgHelper := patch.NewDogHelper(dir, parser.PkgName())
			err = dgHelper.WriteDogHelper()
			if err != nil {
				return errors.Wrap(err, "write dog helper error")
			}
		}
	}

	return nil
}

// Recover 恢复备份文件
func (entity *Patcher) Recover() (err error) {
	for _, dir := range entity.dirs {
		err = entity.backup.RestorePath(dir)
		if err != nil {
			return err
		}

		gofiles := patch.ListGoFile(dir, false)
		if len(gofiles) == 0 {
			continue
		}
		parser := patch.NewParser(gofiles[0])
		if err := parser.Parse(); err != nil {
			return err
		}
		dgHelper := patch.NewDogHelper(dir, parser.PkgName())
		if err := dgHelper.EraseDogHelper(); err != nil && !os.IsNotExist(err) {
			return errors.Wrap(err, "erase dog helper error")
		}
	}
	return nil
}

package patch

import (
	"bytes"
	"github.com/pkg/errors"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
	pkgPath "path"
	"runtime"
	"strings"
)

type DogHelper struct {
	path string
	pkg  string
}

func NewDogHelper(path, pkg string) *DogHelper {
	return &DogHelper{
		path: path,
		pkg:  pkg,
	}
}
func (d *DogHelper) WriteDogHelper() error {
	_, filename, _, _ := runtime.Caller(0)

	fset := token.NewFileSet()
	fbytes, err := os.ReadFile(pkgPath.Join(pkgPath.Dir(filename)) + "/dogHelper.go")
	if err != nil {
		return errors.Wrap(err, "read dogHelper.go error")
	}
	f, err := parser.ParseFile(fset, "dogHelper.go", fbytes, 0)
	if err != nil {
		return errors.Wrap(err, "parse dogHelper.go error")
	}

	f.Name.Name = d.pkg

	var buf bytes.Buffer
	err = printer.Fprint(&buf, fset, f)
	if err != nil {
		return errors.Wrap(err, "printer fprint error")
	}
	if err := os.WriteFile(genPath(d.path, d.pkg), buf.Bytes(), os.ModePerm); err != nil {
		return errors.Wrap(err, "write dogHelper.go error")
	}
	return nil
}

func (d *DogHelper) EraseDogHelper() error {
	return os.Remove(genPath(d.path, d.pkg))
}

func genPath(path, pkg string) string {
	suffix := ""
	if !strings.HasSuffix(path, "/") && len(path) != 0 {
		suffix = "/"
	}
	return path + suffix + "gen_" + pkg + HelperSuffix
}

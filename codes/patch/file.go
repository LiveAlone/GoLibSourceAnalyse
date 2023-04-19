package patch

import (
	"log"
	"os"
	"strings"
)

const BackupSuffix = ".ld"
const HelperSuffix = "_lzdgen.go"

// 文件过滤支持

func ListGoFile(path string, jumpBacked bool) []string {
	return listSuffixFile(path, []string{".go"}, jumpBacked, "_test.go", HelperSuffix)
}

func ListGoFileByPaths(paths []string, jumpBacked bool) []string {
	var ret []string
	for _, path := range paths {
		fs := ListGoFile(path, jumpBacked)
		ret = append(ret, fs...)
	}
	return ret
}

func listSuffixFile(path string, include []string, jumpBacked bool, exclude ...string) (fs []string) {
	fis, err := os.ReadDir(path)
	if err != nil {
		log.Fatalf("list file error: %v", err)
	}

	for _, fi := range fis {
		if fi.IsDir() {
			continue
		}
		for _, v := range exclude {
			if strings.HasSuffix(fi.Name(), v) {
				continue
			}
		}

		// backup file exists
		if jumpBacked {
			if _, err := os.Stat(path + "/" + fi.Name() + BackupSuffix); !os.IsNotExist(err) {
				continue
			}
		}

		for _, v := range include {
			if strings.HasSuffix(fi.Name(), v) {
				fs = append(fs, path+"/"+fi.Name())
			}
		}
	}
	return fs
}

// TreeDir 递归文件目录
func TreeDir(path string, deep int) []string {
	// deep
	if deep == 0 {
		return []string{}
	}
	if deep != -1 {
		deep = deep - 1
	}

	paths := []string{path}
	fis, err := os.ReadDir(path)
	if err != nil {
		log.Fatalf("list file error: %v", err)
	}
	for _, fi := range fis {
		if fi.IsDir() {
			paths = append(paths, TreeDir(path+"/"+fi.Name(), deep)...)
		}
	}
	return paths
}

// Backup 文件备份
type Backup struct {
}

func (j *Backup) BackupPath(path string) error {
	files := ListGoFile(path, true)
	for _, fn := range files {
		if err := copyFile(fn, backupFileName(fn)); err != nil {
			return err
		}
	}
	return nil
}

func (j *Backup) RestorePath(path string) error {
	files := listSuffixFile(path, []string{".go" + BackupSuffix}, false)
	for _, fn := range files {
		if err := copyFile(fn, restoreFileName(fn)); err != nil {
			return err
		}

		if err := os.Remove(fn); err != nil {
			return err
		}
	}
	return nil
}

func backupFileName(fileName string) string {
	sps := strings.Split(fileName, `/`)
	sps[len(sps)-1] = sps[len(sps)-1] + BackupSuffix
	return strings.Join(sps, "/")
}

func restoreFileName(fileName string) string {
	sps := strings.Split(fileName, `/`)
	if !strings.HasSuffix(fileName, BackupSuffix) {
		return ""
	}
	sps[len(sps)-1] = strings.Replace(sps[len(sps)-1], BackupSuffix, "", 1)
	return strings.Join(sps, "/")
}

// 文件操作

func copyFile(src string, dst string) error {
	data, err := os.ReadFile(src)
	if err != nil {
		return err
	}
	return os.WriteFile(dst, data, 0644)
}

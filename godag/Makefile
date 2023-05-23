pkg_name=.
pkg_path=/Users/zyb/workspace/zyb/trans
depth=1
dot_file=godag.dot

gen:
	go run main.go --pkg_name=$(pkg_name) --pkg_path=$(pkg_path) --depth=$(depth) --dot_file_path=$(dot_file)

png:
	dot -Tpng $(dot_file) > godag.png

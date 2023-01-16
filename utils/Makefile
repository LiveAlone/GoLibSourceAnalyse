add:
	git add . && git commit -m "none"

model:
	go run main.go model -d dest

build:
	go build

convert:
	go run main.go convert -f=dest/from.txt -t=dest/to.txt

clean:
	rm -rf utils
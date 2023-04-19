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

api-brick:
	go run main.go api -p brick -a 315713 -d dest

api-hxxmis:
	go run main.go api -p hxxmis -a 643396 -d dest
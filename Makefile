.PHONY: build clean test

build:
	go build -o xuuid

clean:
	rm -f xuuid

test: build
	./test.sh

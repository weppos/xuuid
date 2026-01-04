.PHONY: build clean

build:
	go build -o xuuid

clean:
	rm -f xuuid

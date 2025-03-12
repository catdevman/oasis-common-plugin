.PHONY: all clean build
all: clean build
clean:
	rm -rf common
build:
	go build -o common *.go

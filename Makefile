PREFIX=/usr/local/bin

all: build

build:
	go build -o dist/zet cmd/main.go

install:
	cp dist/zet ${PREFIX}/zet

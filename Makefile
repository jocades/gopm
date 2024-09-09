.PHONY: all server client

all: server client

server:
	go build -o bin/server src/server/main.go

client:
	go build -o bin/client src/client/main.go


.PHONY: default build exec

SRC := $(wildcard *.go)
BIN := desktop_calander

default: build exec

build: $(BIN)

$(BIN): $(SRC)
	go build -o $@ $<

exec:
	./$(BIN)

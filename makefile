.PHONY: default build dllgen exec

SRC := \
	main.go  \
	zsyscall_windows.go \

BIN := desktop_calander

default: build exec

build: dllgen $(BIN)

$(BIN): $(SRC)
	go build -o $@ $<

dllgen:
	cd util && go generate && mv zsyscall_windows.go ..

exec:
	./$(BIN)

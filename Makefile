BIN_SRC=$(wildcard *.go)
BIN=bin/proxypool

PLUSIN_SRC=$(wildcard crawler/*.go)
PLUSIN=$(patsubst %.go,bin/%.so,$(notdir $(PLUSIN_SRC)))

build: $(PLUSIN) $(BIN)

$(PLUSIN): bin/%.so:crawler/%.go
	go build -o $@ -buildmode=plugin $<

$(BIN): $(BIN_SRC)
	go build -o $@ $(BIN_SRC)

install:

clean:
	rm -rf bin/*

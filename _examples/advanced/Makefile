.PHONY: run build clean

PROG_NAME := $(shell basename $(CURDIR))

default: info run

info:
	@echo "PROG_NAME: $(PROG_NAME)"

clean:
	@rm -f $(PROG_NAME)

run:
	@go run *.go

build: clean
	@go build .
	@./$(PROG_NAME)
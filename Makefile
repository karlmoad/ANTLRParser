.PHONY: all build dist
.DEFAULT_GOAL := all

DIST_PATH := build

all: build

build: clean
	go build -o $(DIST_PATH)/PlSqlParser ./main.go

clean:
	rm -rf $(DIST_PATH)/*
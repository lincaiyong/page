.PHONY: all
all: frontend target

.PHONY: rebuild
rebuild: clean all

.PHONY: target
target: build/bin/CodeEdge.app/Contents/MacOS/codeedge

build/bin/CodeEdge.app/Contents/MacOS/codeedge: $(shell ls *.go)
	rm -rf build/bin/demo.app/Contents/MacOS/demo
	wails build

.PHONY: frontend
frontend: frontend/dist/index.html

frontend/dist/index.html: frontend/frontend.go
	cd frontend && go run .

.PHONY: clean
clean:
	rm -rf build/bin frontend/dist/index.html

.PHONY: open
open:
	open build/bin/demo.app

.PHONY: debug
debug: frontend target
	wails dev
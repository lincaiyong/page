.PHONY: all
all: frontend/dist/index.html build/bin/demo.app/Contents/MacOS/demo

.PHONY: rebuild
rebuild: clean all

build/bin/demo.app/Contents/MacOS/demo: main.go
	rm -rf build/bin/demo.app/Contents/MacOS/demo
	wails build

frontend/dist/index.html: frontend/frontend.go
	cd frontend && go run .

.PHONY: clean
clean:
	rm -rf build/bin frontend/dist/index.html

.PHONY: open
open:
	open build/bin/demo.app
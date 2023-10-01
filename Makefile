.PHONY: build
build:
	go build -o bin/chat cmd/chat/main.go

.PHONY: run
run:
	./bin/chat

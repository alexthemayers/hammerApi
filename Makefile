.phony: run build clean


run: build
	./hammerApi --rate=20000 --debug=true

build:
	make clean && make hammerApi

clean:
	rm -rf hammerApi

hammerApi:
	go mod tidy
	goimports -w .
	go vet ./...
	go build .

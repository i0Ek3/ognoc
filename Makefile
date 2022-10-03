.PHONY: build test clean

GO=go

build:
	@$(GO) build ./cmd/ognoc/main.go

test:
	@$(GO) test -v .

clean:
	@rm ./cmd/ognoc/ognoc

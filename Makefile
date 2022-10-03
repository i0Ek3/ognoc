.PHONY: build test clean

GO=go

build:
	@$(GO) build -o ognoc ./cmd/ognoc/main.go

test:
	@$(GO) test -v .

clean:
	@rm ognoc

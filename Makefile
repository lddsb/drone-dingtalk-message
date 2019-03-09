GO ?= go
PACKAGES ?= $(shell $(GO) list ./...)

vet:
	$(GO) vet $(PACKAGES)

test:
	@$(GO) test -v -cover -coverprofile coverage.txt $(PACKAGES) && echo "\n==>\033[32m Ok\033[m\n" || exit 1

coverage:
	sed -i '/main.go/d' coverage.txt
REPO=github.com/sean9999/go-stable-map
BRANCH := $$(git branch --show-current)
REF := $$(git describe --dirty --tags --always)

info:
	@printf "REPO:\t%s\nBRANCH:\t%s\nREF:\t%s\n" $(REPO) $(BRANCH) $(REF)


tidy:
	go mod tidy

clean:
	go clean
	go clean -modcache
	rm bin/*

pkgsite:
	if [ -z "$$(command -v pkgsite)" ]; then go install golang.org/x/pkgsite/cmd/pkgsite@latest; fi

docs: pkgsite
	pkgsite -open .

publish:
	GOPROXY=https://goproxy.io,direct go list -m ${REPO}@${REF}

test:
	git restore testdata
	go test ./...
	git restore testdata

.PHONY: test

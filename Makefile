bin-vendor: cmd internal
	go build -o ./bin-vendor ./cmd/bin-vendor/

bin: bin.yaml bin-vendor
	./bin-vendor

dist: bin
	./bin/goreleaser --snapshot --skip-publish --rm-dist

clean:
	rm -rf bin/
	rm -rf dist/
	rm -rf bin-vendor

lint: bin
	./bin/staticcheck ./...
	./bin/golangci-lint run

release: bin
	./bin/goreleaser

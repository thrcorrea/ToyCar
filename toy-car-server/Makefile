run:
	go run ./cmd/main.go

watch:
	gotestsum --format=pkgname --watch -- -v -race -short ./internal/...

mocks:
	cd internal && rm -rf mocks && mockery --all --keeptree --inpackage --with-expecter

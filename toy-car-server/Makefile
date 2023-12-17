run:
	go run ./cmd/main.go

unit-test:
	gotestsum -- ./internal/... -failfast -race -coverprofile ./coverage.out

watch:
	gotestsum --format=pkgname --watch -- -v -race -short ./internal/...

mocks:
	cd internal && rm -rf mocks && mockery --all --keeptree --inpackage --with-expecter

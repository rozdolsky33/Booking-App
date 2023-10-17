.PHONY: run
run:
	./run.sh

.PHONY: test
test:
	go test -timeout=3s -race -count=10 -failfast -short ./...
	go test -timeout=3s -race -count=1 -failfast ./...

# Code tidy
.PHONY: tidy
tidy:
	go mod tidy
	go fmt ./...

# Runs test coverage check
.PHONY: check-coverage
check-coverage:
	go test ./... -coverprofile=./cover.out -covermode=atomic -coverpkg=./...
	go run ./main.go --config=./.github/.testcoverage.yml
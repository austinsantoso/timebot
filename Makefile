build:
	go build main.go

test:
	go test ./...

testAndCoverage:
	go test -coverprofile coverage.out ./...
	go tool cover -func=coverage.out

testAndCoverageHtml:
	go test -coverprofile coverage.out ./...
	go tool cover -html coverage.out -o coverage.html

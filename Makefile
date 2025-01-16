.PHONY: test test-cover test-cover-report

test:
	@clear
	@go test -v ./ds/... -coverprofile=cover.out

test-cover: test
	@go tool cover -func=cover.out
	@rm cover.out

test-cover-report: test
	@go tool cover -html=cover.out
	@rm cover.out

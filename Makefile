.PHONY: test mocks

mocks:
	mockgen -source=metrics/metrics.go -destination=metrics/metrics_mock.go -package=metrics

test: mocks
	go test -v ./...

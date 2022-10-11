run:
	go run cmd/*.go -config=./config.yaml

debug:
	dlv debug cmd/*.go

test:
	go test -v -count=1 ./...

race:
	go test -v -race -count=1 ./...

cover:
	go test -short -count=1 -race -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out
	rm coverage.out

mockgen:
	mockgen -source=internal/repository/repository.go \
		-destination=internal/repository/mocks/repository.go

setup_test_db:
	sqlite3 sport.db < data.sql

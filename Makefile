build:
	go build -o bin/app cmd/main.go

run:
	./bin/app

# test the handlers and models packages
test:
	go test -v ./handlers/... ./models/...

all: test vet fmt lint build

test:
	go test ./...

vet:
	go vet ./...

fmt:
	go list -f '{{.Dir}}' ./... | grep -v /vendor/ | xargs -L1 gofmt -l
	test -z $$(go list -f '{{.Dir}}' ./... | grep -v /vendor/ | xargs -L1 gofmt -l)

lint:
	go list ./... | grep -v /vendor/ | xargs -L1 golint -set_exit_status

build:
	go build -o bin/apiserver ./exec/api
	go build -o bin/kvstore ./exec/cli
	tar -czvf bin/apiserver.tar.gz  bin/apiserver
	tar -czvf bin/kvstore.tar.gz  bin/kvstore

runserver:
	go run ./exec/api/main.go
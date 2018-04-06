dep:
	dep ensure

run:
	go run cmd/cmd.go

test:
	go test -v -race ./

lint-install:
	go get -u github.com/alecthomas/gometalinter
	gometalinter --install

lint:
	@gometalinter  --disable-all --enable=goimports --enable=vetshadow --enable=gofmt  --vendor ./...

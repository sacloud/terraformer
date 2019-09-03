TEST?=$$(go list ./... | grep -v vendor)
VETARGS?=-all
GOFMT_FILES?=$$(find . -name '*.go' | grep -v vendor)
export GO111MODULE=on

default: gen fmt goimports

clean:
	rm -f providers/sacloud/*_gen.go; rm ./terraformer

gen:
	go run providers/sakuracloud/generator/*.go

vet: golint
	go vet ./...

golint:
	test -z "$$(golint ./... | grep -v 'vendor/' | grep -v '_string.go' | tee /dev/stderr )"

goimports: fmt
	goimports -l -w $(GOFMT_FILES)

fmt:
	gofmt -s -l -w $(GOFMT_FILES)

install:
	go install -v

build: gen fmt goimports
	go build -o terraformer main.go


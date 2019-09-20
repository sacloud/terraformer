# Copyright 2019 Kazumichi Yamamoto.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

TEST?=$$(go list ./... | grep -v vendor)
VETARGS?=-all
GOFMT_FILES?=$$(find . -name '*.go' | grep -v vendor)
LDFLAGS = "-s -w"
export GO111MODULE=on

default: gen fmt goimports

clean:
	rm -f providers/sacloud/*_gen.go; rm -r bin/*

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

build: gen 
	CGO_ENABLED=0 go build -mod=vendor -ldflags $(LDFLAGS) -o bin/terraformer main.go

build-linux: gen
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -mod=vendor -ldflags $(LDFLAGS) -o bin/terraformer main.go
	(cd bin/; zip -rm terraformer-linux-amd64.zip terraformer)

build-darwin: gen 
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -mod=vendor -ldflags $(LDFLAGS) -o bin/terraformer main.go
	(cd bin/; zip -rm terraformer-darwin-amd64.zip terraformer)

build-windows: gen 
	GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -mod=vendor -ldflags $(LDFLAGS) -o bin/terraformer.exe main.go
	(cd bin/; zip -rm terraformer-windows-amd64.zip terraformer.exe)

build-all: build-linux build-windows build-darwin

build-on-docker:
	docker run -it --rm -v $(PWD):$(PWD) -w $(PWD) golang:1.12.1 sh -c "apt-get update; apt-get install -y zip; make build-all"

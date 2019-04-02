BINARY=consul-kv-search
VERSION=1.0.0
BUILD_TIME=`date +%FT%T%z`
GOX_OSARCH="darwin/amd64 linux/386 linux/amd64 windows/386 windows/amd64"

default: build

clean:
	rm -rf ./bin

build:
	GO111MODULE=on \
	CGO_ENABLED=0 \
	go build -a -o ./bin/${BINARY}-${VERSION} *.go

build-linux:
	GO111MODULE=on \
	CGO_ENABLED=0 \
	GOARCH=amd64 \
	GOOS=linux \
	go build -ldflags "-X main.Version=${VERSION}" -a -o ./bin/${BINARY}-${VERSION} *.go

build-gox:
	GO111MODULE=on \
	gox -ldflags "-X main.Version=${VERSION}" -osarch=${GOX_OSARCH} -output="bin/${VERSION}/{{.Dir}}_{{.OS}}_{{.Arch}}"

deps:
	go get;

release:
	ghr -u pteich -r consul-kv-search ${VERSION} bin/${VERSION}/;

test:
	go test

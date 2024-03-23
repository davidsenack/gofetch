build:
	go build --compiler gccgo -gccgoflags "-Ofast -s -w" -o build-output/current/gofetch main.go


run:
	go run main.go


release:
	echo "Compiling for Linux x86_64"
	mkdir bin/gofetch-current-linux-amd64
	GOOS=linux GOARCH=amd64 go build --compiler gccgo -gccgoflags "-Ofast -s -w" -o bin/gofetch-current-linux-amd64/gofetch main.go
	cp README.md bin/gofetch-current-linux-amd64
	cp LICENSE bin/gofetch-current-linux-amd64
	tar -czvf bin/gofetch-current-linux-amd64.tar.gz bin/gofetch-current-linux-amd64
	rm -rf bin/gofetch-current-linux-amd64


package:
	echo "Packaging the binary into a tar.gz file"
	tar -czvf bin/main-linux-amd64.tar.gz bin/main-linux-amd64


all: build compile
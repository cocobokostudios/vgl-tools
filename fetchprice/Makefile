build: clean
	go build

clean:
	rm -f ./fetchprice

install:
	go install

uninstall:
	go clean -i

run:
	go run main.go

test: build
	go test ./lib

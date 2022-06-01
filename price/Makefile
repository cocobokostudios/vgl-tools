build: clean
	go build

clean:
	rm -f ./price

install:
	go install

uninstall:
	go clean -i

run:
	go run main.go

test: build
	go test ./lib

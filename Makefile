SHELL := /bin/bash
.PHONY: clean docker deps install run test build swagger-spec shell all

clean:
	rm -f ./bin/loadis

docker:
	docker build -t loadis .

deps:
	go install

install:
	go install

build:
	mkdir -p bin
	go build -o ./bin/loadis .

shell:
	docker run --rm -it -P --name loadis \
		-v "`pwd`:/go/src/github.com/replicatedcom/loadis" \
		loadis

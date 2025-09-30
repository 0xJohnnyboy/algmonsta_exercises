test-all:
	go test -v ./*/*

pattern ?= .

test:
	go test -v ./$(PATTERN)/...

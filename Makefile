PATH_FSYS = github.com/trussle/fsys

.PHONY: all
all: install
	
.PHONY: install
install:
	go get github.com/Masterminds/glide
	glide install

FORCE:

.PHONY: unit-tests
unit-tests:
	go test -v ./ioext/... ./lock/... ./mmap/... .

.PHONY: integration-tests
integration-tests:
	go test -v -tags=integration ./ioext/... ./lock/... ./mmap/... .

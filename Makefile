BIN := httptaskrunner
VERSION := 1.0.59

MAKE_FILE := $(CURDIR)/$(lastword $(MAKEFILE_LIST))

all: build

build: autoinc-version
	@docker run                                                             \
	    --rm                                                                \
	    -v "$$(pwd)/.go:/go"                                                \
	    -v "$$(pwd)/bin:/go/bin"                                            \
	    -v "$$(pwd)/src:/go/src/$(BIN)"                                     \
	    -v "$$(pwd)/.go/cache:/.cache"                                      \
            -w /go/src/$(BIN)                                                   \
	    golang                                                              \
	    /bin/sh -c "                                                        \
	        go get -v &&                                                    \
                go generate &&                                                  \
	        go install -v -ldflags \"-X main.VERSION=${VERSION} -s -w\"     \
	    "                                                                   \
	&& cp ./src/httptaskrunner.yml ./src/httptaskrunner.service ./bin/
#	&& ./bulid_deb.sh

version:
	@echo $(VERSION)

autoinc-version:
	$(eval VERSION := $(shell echo $(VERSION) | sed -n 's/\([0-9]\+\.[0-9]\+\.\)\([0-9]\+\)/"\1$$$$\(\(\2+1\)\)"/p'))
	@sed -i'' -r 's/(.*VERSION\s*:=\s*)([0-9]+\.[0-9]+\.[0-9]+)(.*)/echo "\1${VERSION}\3"/ge' ${MAKE_FILE}

serve:
	./bin/$(BIN) --conf ./bin/$(BIN).yml

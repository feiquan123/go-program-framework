PROJECTNAME=$(shell basename "$(PWD)")
GODOCFILE = index.html
GOURL= /pkg/github.com/feiquan123/$(PROJECTNAME)/
GOURLRE = $(shell python ./scripts/convert.py -s $(GOURL))

## compile: Compile the binary of server and client
compile: build

## build-server: build server
build-server:
	go build -o bin/server $(GOPATH)/src/github.com/feiquan123/go-program-framework/src/cmd/server

## build-client: build client
build-client:
	go build -o bin/client $(GOPATH)/src/github.com/feiquan123/go-program-framework/src/cmd/client

## build: build client and server
build: build-server build-client

## clean-server: clean server
clean-server:
	rm -rf bin/server

## clean-client: clean client
clean-client:
	rm -rf bin/client

## clean: clean server and client
clean: clean-server clean-client

## start-server: start server
start-server: ./bin/server -c ./configs/app.yaml

## start-client: start client
start-client: ./bin/client -c ./configs/app.yaml

## godoc: start godoc server
godoc:
	# @ln -fs $(GOPATH)/src/github.com $(GOROOT)/src/
	@godoc -http=127.0.0.1:6060 -play -notes=NOTE -goroot=. -url=$(GOURL) > doc/$(GODOCFILE)
	# repleace godoc file
	@sed -i "" '/padding\-left/{n;s/href="/href="$(GOURLRE)/;}' doc/$(GODOCFILE)
	godoc -http=127.0.0.1:6060 -play -goroot=.

## wire: go wire
wire: 
	wire src/cmd/server/wire.go

## di-clean: docker image clean
di-clean:
	@docker rmi $$(docker images -q)

## d-clean: docker clean
d-clean: 
	@docker rm $$(docker ps -aq)

## d-build: docker build
d-build:
	@docker build -t $(PROJECTNAME) -f ./docker/DockerFile .

## md: rename readme
md: 
	mv readme__.md readme.md
	mv readme.zh__.md readme.zh.md

.PHONY: help
all: help
help: Makefile
	@echo
	@echo " Choose a command run in "$(PROJECTNAME)":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo

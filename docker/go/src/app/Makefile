.PHONY: swaggerui \
		models

##### goa ######

REPO:=app

init: depend bootstrap
gen: clean generate
rerun: clean generate run
install: submodule depend

submodule:
	git submodule init
	git submodule update

depend:
	go get -u github.com/goadesign/goa/goagen
	go get -u github.com/alecthomas/gometalinter/...
	gometalinter --install --update --force
	go get -u github.com/goadesign/goa
	go get -u github.com/goadesign/gorma
	go get -u github.com/jinzhu/gorm
	go get -u github.com/deadcheat/goacors

bootstrap:
	goagen bootstrap -d $(REPO)/design

main:
	goagen main -d $(REPO)/design

clean:
	rm -rf app
	rm -rf client
	rm -rf tool
	rm -rf swagger
	rm -rf schema
	rm -f build

generate:
	goagen app     -d $(REPO)/design
	goagen swagger -d $(REPO)/design
	goagen client -d $(REPO)/design
	goagen schema -d $(REPO)/design

models:
	ls models | grep -v '_defined.go' | xargs rm -f
	goagen --design=$(REPO)/design gen --pkg-path=github.com/goadesign/gorma

##### etc #####

ENV:=develop
PORT:=:8080
NOSECURE:=false

run:
	go run server.go

build:
	go build -o goapp .

swaggerui:
	open http://localhost:8080/swaggerui/index.html

lint:
	@if [ "`gometalinter ./... --deadline=1000s --config=lint_config.json | tee /dev/stderr`" ]; then \
    		echo "^ - lint err" && echo && exit 1; \
    fi

##### Change package name ######

BEFORE:=app
AFTER:=path/to

change-package:
	which gorep || go get -v github.com/novalagung/gorep
	gorep -path="./" \
          -from="$(BEFORE)/app" \
          -to="$(AFTER)/app"
	gorep -path="./" \
          -from="$(BEFORE)/design/client" \
          -to="$(AFTER)/design/client"
	gorep -path="./" \
          -from="$(BEFORE)/design/constant" \
          -to="$(AFTER)/design/constant"
	gorep -path="./" \
          -from="$(BEFORE)/design/media" \
          -to="$(AFTER)/design/media"
	gorep -path="./" \
          -from="$(BEFORE)/design/resource" \
          -to="$(AFTER)/design/resource"
	gorep -path="./" \
          -from="$(BEFORE)/app/test" \
          -to="$(AFTER)/app/test"
	gorep -path="./" \
          -from="$(BEFORE)/controller" \
          -to="$(AFTER)/controller"
	gorep -path="./" \
          -from="$(BEFORE)/mymiddleware" \
          -to="$(AFTER)/mymiddleware"
	gorep -path="./" \
          -from="$(BEFORE)/server" \
          -to="$(AFTER)/server"

##### Appengine ######

#init: depend bootstrap appengine
#gen: clean generate appengine
#rerun: clean generate appengine run

appengine:
	which gorep || go get -v github.com/novalagung/gorep
	gorep -path="./vendor/github.com/goadesign/goa" \
          -from="context" \
          -to="golang.org/x/net/context"
	gorep -path="./models" \
          -from="context" \
          -to="golang.org/x/net/context"
	gorep -path="./app" \
          -from="context" \
          -to="golang.org/x/net/context"
	gorep -path="./client" \
          -from="context" \
          -to="golang.org/x/net/context"
	gorep -path="./tool" \
          -from="context" \
          -to="golang.org/x/net/context"
	gorep -path="./" \
          -from="../app" \
          -to="$(REPO)/app"
	gorep -path="./" \
          -from="../client" \
          -to="$(REPO)/client"
	gorep -path="./" \
          -from="../tool/cli" \
          -to="$(REPO)/tool/cli"

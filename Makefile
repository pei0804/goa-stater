##### Convenience targets ######

REPO:=github.com/path/to

init: depend bootstrap
gen: clean generate
rerun: clean generate run

depend:
	@which goagen || go get -v github.com/goadesign/goa/goagen
	@go get -v ./...

bootstrap:
	@goagen bootstrap -d $(REPO)/design

main:
	@goagen main -d $(REPO)/design

clean:
	@rm -rf app
	@rm -rf client
	@rm -rf tool
	@rm -rf swagger
	@rm -rf schema
	@rm -rf js
	@rm -f build

generate:
	@goagen app     -d $(REPO)/design
	@goagen swagger -d $(REPO)/design
	@goagen client -d $(REPO)/design
	@goagen js -d $(REPO)/design
	@goagen schema -d $(REPO)/design
	@go build -o build

swaggerUI:
	@open http://localhost:8080/swagger/index.html

model:
	@rm -rf models
	@goagen --design=$(REPO)/design gen --pkg-path=github.com/goadesign/gorma

run:
	@go run main.go

build:
	@go build -o build

##### Appengine targets ######

#init: depend bootstrap appengine
#gen: clean generate appengine
#rerun: clean generate appengine run

appengine:
	@which gorep || go get -v github.com/novalagung/gorep
	@gorep -path="./vendor/github.com/goadesign/goa" \
          -from="context" \
          -to="golang.org/x/net/context"
	@gorep -path="./models" \
          -from="context" \
          -to="golang.org/x/net/context"
	@gorep -path="./app" \
          -from="context" \
          -to="golang.org/x/net/context"
	@gorep -path="./client" \
          -from="context" \
          -to="golang.org/x/net/context"
	@gorep -path="./tool" \
          -from="context" \
          -to="golang.org/x/net/context"
	@gorep -path="./" \
          -from="../app" \
          -to="$(REPO)/app"
	@gorep -path="./" \
          -from="../client" \
          -to="$(REPO)/client"
	@gorep -path="./" \
          -from="../tool/cli" \
          -to="$(REPO)/tool/cli"

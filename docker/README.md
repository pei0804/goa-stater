# docker

```console
$ go get github.com/pei0804/goa-stater
$ cd docker
$ make -C go/src/app gen
$ make docker_server
```

- [http://localhost:8080/](http://localhost:8080/)
- [Swagger UI](http://localhost:8080/swaggerui/index.html)

## change package name

1.Makefile update

```
BEFORE:=app
AFTER:=path/to <--- change
```

2.Run

```
$ make change-package
```
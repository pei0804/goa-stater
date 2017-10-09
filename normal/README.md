# normal

```console
$ go get github.com/pei0804/goa-stater
$ cd normal
$ make install
$ make gen
$ make run
```

- [http://localhost:8080/](http://localhost:8080/)
- [Swagger UI](http://localhost:8080/swaggerui/index.html)

## change package name

1.Makefile update

```
BEFORE:=github.com/pei0804/goa-stater/normal
AFTER:=github.com/path/to <---change
```

2.Run

```
$ make change-package
```
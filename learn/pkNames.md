# package names

## CompileDaemon

to watch file change and run go build when on save

- install for internal use

> go get github.com/githubnemo/CompileDaemon

- ### install globally

so that we could run it from command line

> go install github.com/githubnemo/CompileDaemon

## dotEnv

to manipulate .env for go lang

- install
  > go get github.com/joho/godotenv

## gin

REST api for golang (most popular one)

- install

  > go get -u github.com/gin-gonic/gin

## gorm

Object Relational Mapping (ORM) for golang.
to be able to use human readable method to talk with
database. one translator to all db.

- install gorm lib

> go get -u gorm.io/gorm

- install database driver (postgres,sql,sqlite)
  install with postgres driver

> go get -u gorm.io/driver/postgres

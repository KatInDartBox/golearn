# gorm

Object Relational Mapping (ORM) for golang.
to be able to use human readable method to talk with
database. one translator to all db.

- install gorm lib

> go get -u gorm.io/gorm

- install database driver (postgres,sql,sqlite)
  install with postgres driver

> go get -u gorm.io/driver/postgres

## model

```go
type User struct {
  gorm.Model
  Name string
}
// equals
type User struct {
  ID        uint           `gorm:"primaryKey"`
  CreatedAt time.Time
  UpdatedAt time.Time
  DeletedAt gorm.DeletedAt `gorm:"index"`
  Name string
}

```

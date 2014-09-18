package persistence

import (
  "github.com/coopernurse/gorp"
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
  "log"
)

type Mysql struct {
  dbMap *gorp.DbMap
}

func (s *Mysql) Initialize(settings map[string]string) {
  db, err := sql.Open("mysql", settings["user"] + ":" + settings["password"] + "@/" + settings["database"])
  log.Println("DB: ", db.Ping())

  if err != nil {
    panic("Cannot connect to Mysql: " + err.Error())
  }
  dialect := gorp.MySQLDialect{"InnoDB", "UTF8"}
  s.dbMap = &gorp.DbMap{Db: db, Dialect: dialect}
}

func (s Mysql) DbMap() *gorp.DbMap {
  return s.dbMap
}

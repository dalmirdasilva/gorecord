package persistence

import (
  "github.com/coopernurse/gorp"
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
)

type Mysql struct {
  dbMap *gorp.DbMap
}

var defaultSettings = map[string]string {
  "user": "root",
  "password": "password",
  "database": "database",
  "engine": "InnoDB",
  "encoding": "UTF8",
}

func (s *Mysql) Initialize(settings map[string]string) {
  mergeSettings(settings)
  db, err := sql.Open("mysql", settings["user"] + ":" + settings["password"] + "@/" + settings["database"])
  panicOnError(err, "Cannot open driver")
  err = db.Ping()
  panicOnError(err, "Cannot connect to the server")
  dialect := gorp.MySQLDialect{settings["engine"], settings["encoding"]}
  s.dbMap = &gorp.DbMap{Db: db, Dialect: dialect}
}

func (s Mysql) DbMap() *gorp.DbMap {
  return s.dbMap
}

func mergeSettings(settings map[string]string) {
  for k, v := range defaultSettings {
    if e, ok := settings[k]; !ok || e == "" {
      settings[k] = v
    }
  }
}

func panicOnError(err error, msg string) {
  if err != nil {
    panic(msg + ": " + err.Error())
  }
}

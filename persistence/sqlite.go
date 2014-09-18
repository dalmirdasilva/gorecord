package persistence

import (
  "github.com/coopernurse/gorp"
  "database/sql"
  _ "github.com/mattn/go-sqlite3"
)

type Sqlite struct {
  dbMap *gorp.DbMap
}

func (s *Sqlite) Initialize(settings map[string]string) {
  db, err := sql.Open("sqlite3", settings["file"])
  if err != nil {
    panic("Cannot open file " + settings["file"] + ".")
  }
  s.dbMap = &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}
}

func (s Sqlite) DbMap() *gorp.DbMap {
  return s.dbMap
}

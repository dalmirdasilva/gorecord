package persistence

import "github.com/coopernurse/gorp"

type Database interface {
  Initialize(settings map[string]string)
  DbMap() *gorp.DbMap
}

var db Database

func GetDatabase() Database {
  if db == nil {
    panic("Database not initialized.")
  }
  return db
}

func Initialize(databaseName string, settings map[string]string) {
  db = newDatabase(databaseName, settings)
}

func RegisterTables(tables map[string]interface{}) {
  for tableName, object := range tables {
    db.DbMap().AddTableWithName(object, tableName).SetKeys(true, "Id")
  }
  db.DbMap().CreateTablesIfNotExists()
}

func newDatabase(databaseName string, settings map[string]string) Database {
  var database Database
  if databaseName == "sqlite3" {
    database = new(Sqlite)
  }
  if databaseName == "mysql" {
    database = new(Mysql)
  }
  if database != nil {
    database.Initialize(settings)
  }
  return database
}

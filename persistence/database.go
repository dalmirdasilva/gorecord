package persistence

import (
  "github.com/coopernurse/gorp"
  "reflect"
  "log"
)

type Database interface {
  Initialize(settings map[string]string)
  DbMap() *gorp.DbMap
}

var db Database
var registeredTables map[string]string

func GetDatabase() Database {
  if db == nil {
    panic("Database not initialized.")
  }
  return db
}

func Initialize(databaseName string, settings map[string]string) {
  log.Println("Initializing database.")
  db = newDatabase(databaseName, settings)
  log.Println(db)
}

func RegisterTables(tables map[string]interface{}) {
  log.Println("Registering tables")
  registeredTables = make(map[string]string)
  for tableName, object := range tables {
    log.Println("Table registered: ", tableName)
    db.DbMap().AddTableWithName(object, tableName)
    registeredTables[reflect.TypeOf(object).Name()] = tableName
  }
  db.DbMap().CreateTablesIfNotExists()
  log.Println("Done, tables: ", registeredTables)
}

func GetTableFrom(entry interface{}) string {
  for name, instance := range registeredTables {
    if reflect.TypeOf(instance).Kind() ==  reflect.TypeOf(entry).Kind() {
      return name
    }
  }
  return "unknown"
}

func newDatabase(databaseName string, settings map[string]string) Database {
  var database Database
  if databaseName == "sqlite3" {
    log.Println("Instantiating Sqlite")
    database = new(Sqlite)
  }
  if databaseName == "mysql" {
    log.Println("Instantiating Mysql")
    database = new(Mysql)
  }
  if database != nil {
    database.Initialize(settings)
  }
  log.Println("Database is: ", database)
  return database
}

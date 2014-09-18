package model

import (
  "github.com/dalmirdasilva/way/persistence"
  "github.com/coopernurse/gorp"
  "errors"
  "log"
)

type fn func(* gorp.DbMap) (interface{}, error)

type Entry struct {
  Id int64
  Created int64
  Updated int64
}

func (e *Entry) Save() error  {
  _, err := run(func(db *gorp.DbMap) (interface{}, error) {
    var err error = nil
    var rows int64 = 0
    log.Println(e.Id)
    return nil, nil
    if e.Id == 0 {
      log.Println("Iserting.")
      err = db.Insert(e)
    } else {
      rows, err = db.Update()
    }
    return rows, err
  })
  return err
}

func (e *Entry) Delete() error  {
  _, err := run(func(db *gorp.DbMap) (interface{}, error) {
    var err error = nil
    var rows int64 = 0
    if e.Id != 0 {
      rows, err = db.Delete(e)
    }
    return rows, err
  })
  return err
}

func (e *Entry) Find(id int64) (*Entry, error)  {
  var result Entry
  var ok bool
  found, err := run(func(db *gorp.DbMap) (interface{}, error) {
    var entry Entry
    tableName := persistence.GetTableFrom(e)
    db.SelectOne(&entry, "select * from " + tableName)
    return entry, nil
  })
  if err != nil {
    result, ok = found.(Entry)
    if !ok {
      err = errors.New("Cannot cast result to Entry.")
    }
  }
  return &result, err
}

func run(f fn) (interface{}, error) {
  pdb := persistence.GetDatabase()
  db := pdb.DbMap()
  return f(db)
}


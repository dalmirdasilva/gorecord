package model

import (
  "github.com/coopernurse/gorp"
  "github.com/dalmirdasilva/gorecord/persistence"
  "log"
  "reflect"
)

type fn func(* gorp.DbMap) (interface{}, error)

type Entry struct {
  Id int64
  Created int64
  Updated int64
  Self interface{} `db:"-"`
}



func (e *Entry) Save() error  {
  _, err := run(func(db *gorp.DbMap) (interface{}, error) {
    var err error = nil
    var rows int64 = 0
    if e.Id == 0 {
      err = db.Insert(e.Self)
    } else {
      rows, err = db.Update(e.Self)
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

func (e Entry) Find(id int64) (*Entry, error)  {
  log.Println("find with: ", reflect.TypeOf(e))
  /*var result Entry
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
  */
  return nil, nil
}

func run(f fn) (interface{}, error) {
  pdb := persistence.GetDatabase()
  db := pdb.DbMap()
  return f(db)
}


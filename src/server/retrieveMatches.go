package main

import (
  "database/sql"
  "log"
)

type Keyobj struct {
  Key string `json:"key"`

}

func (k *Keyobj) getStrictMatches(db *sql.DB ) []string {
   var returnString []string
   rows, err := db.Query("SELECT match FROM strict_match WHERE key=$1",k.Key)
   if err != nil {
    log.Fatal(err)
    log.Println("inside error")
   }
   defer rows.Close()

   for rows.Next() {
     var name string

     if err := rows.Scan(&name); err != nil {
       log.Fatal(err)
     }
     log.Printf("name is %s\n", name)
     returnString = append(returnString,name)
   }
   return returnString
}

func (k *Keyobj) getPotentialTierOne( db *sql.DB) []string{
  var returnString []string
  rows, err := db.Query("SELECT match FROM image_tier_one WHERE key = $1", k.Key)
  if err != nil {
    log.Fatal(err)
  }
  defer rows.Close()

  for rows.Next() {
    var name string

    if err := rows.Scan(&name); err != nil {
      log.Fatal(err)
    }
    log.Printf("name is %s\n", name)
    returnString = append(returnString,name)
  }
  return returnString
}

func (k *Keyobj) getPotentialTierTwo( db *sql.DB ) []string{
  var returnString []string
  rows, err := db.Query("SELECT match FROM image_tier_two WHERE key = $1",k.Key)
  if err != nil {
    log.Fatal(err)
  }
  defer rows.Close()

  for rows.Next() {
    var name string

    if err := rows.Scan(&name); err != nil {
      log.Fatal(err)
    }
    log.Printf("name is %s\n", name)
    returnString = append(returnString,name)
  }
  return returnString
}

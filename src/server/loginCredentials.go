// model.go

package main

import (
    "database/sql"
    //"errors"
    "log"
)

type loginCred struct {
    ID    int     `json:"id"`
    UserName  string  `json:"username"`
    Password  string  `json:"password"`
}

func (l *loginCred) checkLoginCred(db *sql.DB) (error, bool) {
  i := 0
  count := &i
  err := db.QueryRow("SELECT COUNT(username) FROM logincredentials WHERE username=$1 AND password=$2",
  //err := db.QueryRow("SELECT COUNT(username) FROM test_one WHERE username=$1 AND password=$2",
          l.UserName, l.Password).Scan(count)
  if err != nil {
    log.Fatal(err)
    return err, false
  }

  log.Println(*count)
  log.Println(l.UserName, l.Password)
  return nil, *count == 1
}

/*
 * the login cred object will have new username and/or password but the id will stay the same
 */
func (l *loginCred) addLoginCred(db *sql.DB) error {
  _, err := db.Exec("INSERT INTO logincredentials(username,password) VALUES($1,$2)", l.UserName, l.Password)

  return err
}

func (l *loginCred) deleteLoginCred(db *sql.DB) error {
      _, err := db.Exec("DELETE FROM logincredentials WHERE id=$1", l.ID)

    return err
}

/*
func getAllTimes(db *sql.DB) error {

}

*/

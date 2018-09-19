// app.go

package main

import (
	"database/sql"
	"fmt"
	"log"

	"encoding/json"
	"github.com/gorilla/mux"
  "github.com/rs/cors"
  _ "github.com/lib/pq"
	"net/http"
  "io/ioutil"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (a *App) Initialize(user, password, dbname string) {
	connectionString :=
		//fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, password, dbname)
    //fmt.Sprintf("user=%s dbname=%s sslmode=disable", user, dbname)
    fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable", "localhost","5432","test","test_one")

	var err error
	a.DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

  errPing := a.DB.Ping()
  if errPing != nil{
    log.Println("ping error")
  }

  log.Println("successful")
	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

func (a *App) initializeRoutes() {

	a.Router.HandleFunc("/login/add", a.addCredentials)
	//a.Router.HandleFunc("/getTimes",a.getTimes)
	a.Router.HandleFunc("/login/check", a.checkCredentials)
}

/*
func (a *App) getTimes(w http.ResponseWriter, r *http.Request)  {
  var l loginCred
  respondWithJSON(w,http.StatusOK, loginCred.getAllTimes())
}
*/
func (a *App) Run(addr string) {
  //log.Fatal(http.ListenAndServe(":12345", handlers.CORS(handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD"}), handlers.AllowedOrigins([]string{"*"}))(a.Router)))

  // Solves Cross Origin Access Issue
  c := cors.New(cors.Options{
    AllowedOrigins: []string{"http://localhost:4200"},
    AllowedMethods: []string{"GET","POST","PUT","HEAD"},
  })
  handler := c.Handler(a.Router)

  srv := &http.Server{
    Handler: handler,
    Addr:    ":12345",
  }

  log.Fatal(srv.ListenAndServe())
}

func (a *App) checkCredentials(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	//_, errId := strconv.Atoi(vars["id"])
	//username := vars["username"]
	//password := vars["password"]

  contents, _ := ioutil.ReadAll(r.Body)
  var l loginCred
  json.Unmarshal(contents,&l)
  log.Println(l.UserName, l.Password)

  //log.Println("something", vars)
  //l := loginCred{-1, username, password}
  //err := l.addLoginCred(a.DB)


	//l := loginCred{-1, username, password}
	err, correctCred := l.checkLoginCred(a.DB)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "login not made yet")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	log.Println("correctCred",correctCred)
	if !correctCred {
		respondWithCorrectValue(w, http.StatusNotFound, false)
		log.Println("inside if correctCred")

	}else{
    respondWithCorrectValue(w, http.StatusOK, true)
  }


}


func (a *App) addCredentials(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	//username := vars["username"]
	//password := vars["password"]

  contents, _ := ioutil.ReadAll(r.Body)
  var l loginCred
  json.Unmarshal(contents,&l)
  //log.Println(res.UserName, res.Password)

	//log.Println("something", vars)
	//l := loginCred{-1, username, password}
  err := l.addLoginCred(a.DB)

  if err != nil {
    respondWithError(w,http.StatusInternalServerError, err.Error())
    log.Println(err)
  }else {
    respondWithCorrectValue(w, http.StatusOK, true)
  }
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}
func respondWithCorrectValue(w http.ResponseWriter, code int, value bool) {
  var correct string
  //log.Println("value:",value)
  if value {
    correct = "True"
  } else {
    correct = "False"
  }
  //log.Println("correct:",correct)
	respondWithJSON(w, code, map[string]string{"correctValue": correct})
}
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {


	response, _ := json.Marshal(payload)
  log.Println(string(response))
	log.Println("response",string(response))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)

}

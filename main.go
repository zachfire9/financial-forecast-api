package main

import (
    "fmt"
    "log"
    "net/http"
    "os"

    "github.com/gorilla/mux"
    "gopkg.in/mgo.v2"
)

var db *mgo.Database

func init() {
    databaseConnection := os.Getenv("DATABASE_CONNECTION")

    if len(databaseConnection) == 0 {
        databaseConnection = "localhost"
    }

    databaseName := os.Getenv("DATABASE_NAME")

    if len(databaseName) == 0 {
        databaseName = "go"
    }


    session, err := mgo.Dial(databaseConnection)
    if err != nil {
        panic(err)
    }
    db = session.DB(databaseName)
}

func main() {

    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/", Index)

    router.HandleFunc("/users", UsersGetAll).Methods("GET")
    router.HandleFunc("/users/{userId}", UsersGet).Methods("GET")
    router.HandleFunc("/users", UsersPost).Methods("POST")
    router.HandleFunc("/users/{userId}", UsersPut).Methods("PUT")
    router.HandleFunc("/users/{userId}", UsersDelete).Methods("DELETE")

    router.HandleFunc("/investments", InvestmentsGetAll).Methods("GET")
    router.HandleFunc("/investments/{investmentId}", InvestmentsGet).Methods("GET")
    router.HandleFunc("/investments", InvestmentsPost).Methods("POST")
    router.HandleFunc("/investments/{investmentId}", InvestmentsPut).Methods("PUT")
    router.HandleFunc("/investments/{investmentId}", InvestmentsDelete).Methods("DELETE")

    appPort := os.Getenv("APP_PORT")

    if len(appPort) == 0 {
        appPort = ":8080"
    }

    log.Fatal(http.ListenAndServe(appPort, router))
}

func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "You can send a RESTful request to the /users collection.")
    fmt.Fprintln(w, "You can send a RESTful request to the /investments collection.")
}
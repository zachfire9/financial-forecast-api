package main

import (
    "encoding/json"
    "net/http"
    "time"

    "github.com/gorilla/mux"
    "gopkg.in/mgo.v2/bson"
)

type User struct {
    ID                          bson.ObjectId `json:"id" bson:"_id,omitempty"`
    Name                        string `json:"name"`
    Email                       string `json:"email"`
    RetirementGoal              float64 `json:"retirementGoal,string"`
    RetirementDate              time.Time `json:"retirementDate"`
    RetirementLivingAmount      float64 `json:"retirementLivingAmount,string"`
    InflationAnnualExpected     float64 `json:"inflationAnnualExpected,string"`
    CreatedAt                   time.Time `json:"createdAt"`
    UpdatedAt                   time.Time `json:"updatedAt"`
}

type Users []User

func UsersGet(w http.ResponseWriter, r *http.Request) {

    params := mux.Vars(r)
    userId := params["userId"]

    var user User
    err := db.C("users").FindId(bson.ObjectIdHex(userId)).One(&user)

    if err != nil {
        panic(err)
    }

    respBody, err := json.MarshalIndent(user, "", "  ")

    if err != nil {
        panic(err)
    }

    w.Header().Set("Content-Type", "application/json; charset=utf-8")
    w.WriteHeader(http.StatusOK)
    w.Write(respBody)
}

func UsersGetAll(w http.ResponseWriter, r *http.Request) {

    var users []User
    err := db.C("users").Find(bson.M{}).All(&users)

    if err != nil {
        panic(err)
    }

    respBody, err := json.MarshalIndent(users, "", "  ")

    if err != nil {
        panic(err)
    }

    w.Header().Set("Content-Type", "application/json; charset=utf-8")
    w.WriteHeader(http.StatusOK)
    w.Write(respBody)
}

func UsersPost(w http.ResponseWriter, r *http.Request) {

    var user User
    decoder := json.NewDecoder(r.Body)
    err := decoder.Decode(&user)

    user.ID = bson.NewObjectId()
    user.CreatedAt = time.Now()
    user.UpdatedAt = time.Now()

    if err != nil {
        panic(err)
    }

    err = db.C("users").Insert(user)

    if err != nil {
        panic(err)
    }

    w.Header().Set("Content-Type", "application/json; charset=utf-8")
    w.Header().Set("Location", r.URL.Path+"/"+user.ID.Hex())
    w.WriteHeader(http.StatusCreated)
}

func UsersPut(w http.ResponseWriter, r *http.Request) {

    params := mux.Vars(r)
    userId := params["userId"]

    var user User
    decoder := json.NewDecoder(r.Body)
    err := decoder.Decode(&user)

    user.UpdatedAt = time.Now()

    if err != nil {
        panic(err)
    }

    err = db.C("users").Update(bson.M{"_id": bson.ObjectIdHex(userId)}, &user)

    if err != nil {
        panic(err)
    }

    w.WriteHeader(http.StatusNoContent)
}

func UsersDelete(w http.ResponseWriter, r *http.Request) {

    params := mux.Vars(r)
    userId := params["userId"]

    err := db.C("users").Remove(bson.M{"_id": bson.ObjectIdHex(userId)})

    if err != nil {
        panic(err)
    }

    w.WriteHeader(http.StatusNoContent)
}
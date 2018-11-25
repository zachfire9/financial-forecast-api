package main

import (
    "encoding/json"
    "net/http"
    "time"

    "github.com/gorilla/mux"
    "gopkg.in/mgo.v2/bson"
)

type Investment struct {
    ID                          bson.ObjectId `json:"id" bson:"_id,omitempty"`
    UserId                      bson.ObjectId `json:"userId" bson:"userId"`
    Name                        string `json:"name"`
    CurrentAmount               float64 `json:"currentAmount,string"`
    InterestAnnualExpected      float64 `json:"interestAnnualExpected,string"`
    PaymentMonthly              float64 `json:"paymentMonthly,string"`
    CreatedAt                   time.Time `json:"createdAt"`
    UpdatedAt                   time.Time `json:"updatedAt"`
}

type Investments []Investment

func InvestmentsGet(w http.ResponseWriter, r *http.Request) {

    params := mux.Vars(r)
    investmentId := params["investmentId"]

    var investment Investment
    err := db.C("investments").FindId(bson.ObjectIdHex(investmentId)).One(&investment)

    if err != nil {
        panic(err)
    }

    respBody, err := json.MarshalIndent(investment, "", "  ")

    if err != nil {
        panic(err)
    }

    w.Header().Set("Content-Type", "application/json; charset=utf-8")
    w.WriteHeader(http.StatusOK)
    w.Write(respBody)
}

func InvestmentsGetAll(w http.ResponseWriter, r *http.Request) {

    var investments []Investment
    err := db.C("investments").Find(bson.M{}).All(&investments)

    if err != nil {
        panic(err)
    }

    respBody, err := json.MarshalIndent(investments, "", "  ")

    if err != nil {
        panic(err)
    }

    w.Header().Set("Content-Type", "application/json; charset=utf-8")
    w.WriteHeader(http.StatusOK)
    w.Write(respBody)
}

func InvestmentsPost(w http.ResponseWriter, r *http.Request) {

    var investment Investment
    decoder := json.NewDecoder(r.Body)
    err := decoder.Decode(&investment)

    investment.ID = bson.NewObjectId()
    investment.CreatedAt = time.Now()
    investment.UpdatedAt = time.Now()

    if err != nil {
        panic(err)
    }

    err = db.C("investments").Insert(investment)

    if err != nil {
        panic(err)
    }

    w.Header().Set("Content-Type", "application/json; charset=utf-8")
    w.Header().Set("Location", r.URL.Path+"/"+investment.ID.Hex())
    w.WriteHeader(http.StatusCreated)
}

func InvestmentsPut(w http.ResponseWriter, r *http.Request) {

    params := mux.Vars(r)
    investmentId := params["investmentId"]

    var investment Investment
    decoder := json.NewDecoder(r.Body)
    err := decoder.Decode(&investment)

    investment.UpdatedAt = time.Now()

    if err != nil {
        panic(err)
    }

    err = db.C("investments").Update(bson.M{"_id": bson.ObjectIdHex(investmentId)}, &investment)

    if err != nil {
        panic(err)
    }

    w.WriteHeader(http.StatusNoContent)
}

func InvestmentsDelete(w http.ResponseWriter, r *http.Request) {

    params := mux.Vars(r)
    investmentId := params["investmentId"]

    err := db.C("investments").Remove(bson.M{"_id": bson.ObjectIdHex(investmentId)})

    if err != nil {
        panic(err)
    }

    w.WriteHeader(http.StatusNoContent)
}
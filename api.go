package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Item -
type Item struct {
	Seq    int
	Result map[string]int
}

// Message -
type Message struct {
	Dept    string
	Subject string
	Time    int64
	Detail  []Item
}

func getJSON() ([]byte, error) {
	var pass = make(map[string]int)
	pass["x"] = 50
	pass["c"] = 60
	var item1 = Item{100, pass}

	var reject = make(map[string]int)
	reject["l"] = 11
	reject["d"] = 20
	var item2 = Item{200, reject}

	var detail = []Item{item1, item2}
	var m = Message{"it", "kpi", time.Now().Unix(), detail}
	return json.MarshalIndent(m, "", "")
}

func apiPerson(w http.ResponseWriter, r *http.Request) {
	var res, err = getJSON()
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, string(res))
}

func apiAddMoney(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "{all_money:2450}")
}

func apiUserLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "{login:succes, user_id:3213819283ANF2C}")
}

func main() {
	APIDocAddHandleFun(apiPerson, APIs{Info{title: "All person", url: "/api/person"}, nil})
	APIDocAddHandleFun(apiAddMoney, APIs{Info{title: "add money", url: "/api/add/money"}, []APITag{
		{Name: "user_id", Type: "string", Tip: "find user"},
		{Name: "add_money", Type: "int", Tip: "add the money for all_money"},
	}})
	APIDocAddHandleFun(apiUserLogin, APIs{Info{title: "user login", url: "/api/user/login"}, []APITag{
		{Name: "phone", Type: "string", Tip: "phoneNumber"},
		{Name: "password", Type: "string", Tip: ""},
	}})
	http.ListenAndServe(":5000", nil)
	fmt.Printf("listen localhost:5000")
}

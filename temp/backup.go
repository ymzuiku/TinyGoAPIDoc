package main

import (
	"encoding/json"
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

func main() {}

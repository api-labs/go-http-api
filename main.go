package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

type Payload struct {
	ID       string `json:"id"`
	TimeSent int64  `json:"time_sent"`
	Message  string `json:"message"`
	TimeRecv int64  `json:"time_recv"`
}

func echo(w http.ResponseWriter, r *http.Request) {
	timeRecv := time.Now().UnixNano() / 1000
	query := r.URL.Query()
	if sleepIntv, err := strconv.Atoi(query.Get("sleep_intv")); err == nil {
		time.Sleep(time.Millisecond * time.Duration(sleepIntv))
	}
	timeSent, _ := strconv.ParseInt(query.Get("time_sent"), 10, 64)
	payload := Payload{query.Get("id"), timeSent, query.Get("payload"), timeRecv}
	out, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}

func main() {
	http.HandleFunc("/", helloWorld)
	http.HandleFunc("/echo", echo)
	http.ListenAndServe("0.0.0.0:8000", nil)
}

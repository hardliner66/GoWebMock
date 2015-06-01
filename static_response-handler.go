package main

import (
	"io/ioutil"
	"net/http"
)

func (g *StaticResponse) GetPath() string {
	return g.Path
}

func (g *StaticResponse) WebPostHandler(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()

	w.Header().Add("Access-Control-Allow-Origin", "*")

	if r.Method == "OPTIONS" {
		w.Header().Add("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, apiKey")
		w.Header().Add("Allow", "GET")
		return
	}

	w.Header().Add("content-type", g.ContentType)
	// Return encoded entry.
	w.WriteHeader(http.StatusOK)

	var response string

	response = g.Response

	if len(g.File) != 0 {
		responseString, err := ioutil.ReadFile(g.File)
		if err == nil {
			response = string(responseString)
		}
	}

	w.Write([]byte(response))
}

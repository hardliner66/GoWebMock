package main

import (
	"log"
	//	"fmt"
	"encoding/json"

	"github.com/robertkrimen/otto"
	"io/ioutil"
	"net/http"
	//	"reflect"
)

func (g *DynamicResponse) GetPath() string {
	return g.Path
}

func (g *DynamicResponse) WebOptionsHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Access-Control-Allow-Origin", "*")
}

func Clone(a, b interface{}) {
	val, err := json.Marshal(a)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(val, b)
	if err != nil {
		log.Fatal(err)
	}
}

func (g *DynamicResponse) WebPostHandler(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()

	w.Header().Add("Access-Control-Allow-Origin", "*")

	if r.Method == "OPTIONS" {
		w.Header().Add("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, apiKey")
		w.Header().Add("Allow", "GET")
		return
	}

	var response string

	dat, err := ioutil.ReadFile(g.Handler)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	} else {
		g.vm.Set("externalIp", func(call otto.FunctionCall) otto.Value {
			ip, err := externalIP()
			if err != nil {
				val, _ := g.vm.ToValue("123.4.56.789")
				return val
			}
			val, err := g.vm.ToValue(ip)
			if err != nil {
				val, _ := g.vm.ToValue("123.4.56.789")
				return val
			}
			return val
		})

		w.Header().Add("content_type", g.ContentType)

		header := make(map[string]string)

		config := make(map[string]interface{})

		Clone(g.GlobalConfig, &config)

		g.vm.Set("request", r)
		g.vm.Set("config", config)
		g.vm.Set("storage", g.Storage)
		g.vm.Set("header", header)

		val, err := g.vm.Run(string(dat))

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		} else {
			val, err := val.ToString()
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
				return
			} else {
				response = val
				for s := range header {
					if len(w.Header().Get(s)) > 0 {
						w.Header().Set(s, header[s])
					} else {
						w.Header().Add(s, header[s])
					}
				}
			}
		}

	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}

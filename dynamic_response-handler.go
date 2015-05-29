package main

import (
	"github.com/robertkrimen/otto"
	"io/ioutil"
	"net/http"
)

func (g *DynamicResponse) GetPath() string {
	return g.Path
}

func (g *DynamicResponse) WebOptionsHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Access-Control-Allow-Origin", "*")
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

		g.vm.Set("request", r)
		g.vm.Set("config", g.GlobalConfig)
		g.vm.Set("storage", g.Storage)
		g.vm.Set("content_type", g.ContentType)

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
				g.Storage, err = g.vm.Get("storage")
				ct, _ := g.vm.Get("content_type")
				ctStr, _ := ct.ToString()
				w.Header().Add("content-type", ctStr)
			}
		}

	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}

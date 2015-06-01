package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	_ "net/http/pprof"
)

func main() {

	var config_path = flag.String("cfg", "./autoexec.json", "the config file")
	flag.Parse()

	var config Config

	config_data, err := ioutil.ReadFile(*config_path)

	if err == nil {
		err = json.Unmarshal(config_data, &config)
	} else {
		fmt.Println(err)
		fmt.Println("Using default config...")
		err = nil
		config = Config{}
	}

	if err != nil {
		fmt.Println(err)
	} else {
		if config.Port == 0 {
			config.Port = 80
		}

		configJson, err := json.MarshalIndent(config, "", "    ")
		if err == nil {
			fmt.Println("Config:")
			fmt.Println(string(configJson))
		}

		fmt.Println("Serving on port: " + strconv.Itoa(config.Port))

		http.Handle("/", http.FileServer(http.Dir("static")))

		ip, err := externalIP()

		if err != nil {
			ip = ""
		}

		http.HandleFunc("/ipaddress", func(w http.ResponseWriter, r *http.Request) {
			defer r.Body.Close()

			w.Header().Add("Access-Control-Allow-Origin", "*")

			if r.Method == "OPTIONS" {
				w.Header().Add("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, apiKey")
				w.Header().Add("Allow", "GET")
				return
			}

			w.Header().Add("content-type", "application/json; charset=utf-8")
			// Return encoded entry.
			w.WriteHeader(http.StatusOK)

			if !strings.HasPrefix(r.RemoteAddr, "127.0.0.1") {
				if strings.Contains(r.RemoteAddr, ":") {
					ip = strings.Split(r.RemoteAddr, ":")[0]
				} else {
					ip = r.RemoteAddr
				}
			}
			w.Write([]byte(ip))
		})

		// add static responses
		for sresp := range config.StaticResponses {
			sr := NewStaticResponse(config.StaticResponses[sresp])
			http.HandleFunc(sr.GetPath(), sr.WebPostHandler)
		}

		// add dynamic responses
		for dresp := range config.DynamicResponses {
			if len(config.DynamicResponses[dresp].Handler) > 0 {
				_, err := ioutil.ReadFile(config.DynamicResponses[dresp].Handler)
				if err == nil {
					dr := NewDynamicResponse(config.DynamicResponses[dresp])
					http.HandleFunc(dr.GetPath(), dr.WebPostHandler)
				} else {
					fmt.Println("Handler not registered")
					fmt.Println(err)
					fmt.Println(config.DynamicResponses[dresp])
				}
			} else {
				fmt.Println("Handler not registered, no filename given")
				fmt.Println(config.DynamicResponses[dresp])
			}
		}

		log.Fatal(http.ListenAndServe(":"+strconv.Itoa(config.Port), nil))
	}
}

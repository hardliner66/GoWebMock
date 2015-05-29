package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"strconv"
)

func main() {

	var cfg = flag.String("cfg", "./autoexec.json", "the config file")
	flag.Parse()

	var config Config

	dat, err := ioutil.ReadFile(*cfg)

	if err == nil {
		err = json.Unmarshal(dat, &config)
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
			w.Write([]byte(ip))
		})

		for sresp := range config.StaticResponses {
			sr := NewStaticResponse(config.StaticResponses[sresp])
			http.HandleFunc(sr.GetPath(), sr.WebPostHandler)
		}

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

func externalIP() (string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}
		addrs, err := iface.Addrs()
		if err != nil {
			return "", err
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip == nil || ip.IsLoopback() {
				continue
			}
			ip = ip.To4()
			if ip == nil {
				continue // not an ipv4 address
			}
			return ip.String(), nil
		}
	}
	return "", errors.New("are you connected to the network?")
}

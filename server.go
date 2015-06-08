package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"sync"

	_ "net/http/pprof"
)

var GlobalStorage map[string]interface{}
var StorageMutex *sync.Mutex
var Closing bool

func cleanup() {
	if !Closing {
		Closing = true

		b, err := json.Marshal(GlobalStorage)
		if err == nil {
			os.Mkdir("."+string(filepath.Separator)+"__data", 0644)
			ioutil.WriteFile("__data/global.json", b, 0644)
		}
	}
}

func main() {
	Closing = false

	StorageMutex = &sync.Mutex{}

	var wg sync.WaitGroup

	var config_path = flag.String("cfg", "./autoexec.json", "the config file")
	var numCPU = flag.Int("cpus", runtime.NumCPU(), "the number of cpus to use")

	flag.Parse()

	runtime.GOMAXPROCS(*numCPU)

	fmt.Println("Running on " + strconv.Itoa(*numCPU) + " CPUs!")

	b, err := ioutil.ReadFile("__data/global.json")
	if err == nil {
		err = json.Unmarshal(b, &GlobalStorage)
		if err != nil {
			fmt.Println(err)
			GlobalStorage = make(map[string]interface{})
		}
	} else {
		fmt.Println(err)
		GlobalStorage = make(map[string]interface{})
	}

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

		if config.SSLPort == 0 {
			config.SSLPort = 443
		}

		if _, err := os.Stat(config.PrivateKeyPath); os.IsNotExist(err) {
			fmt.Println("Private key not found at: " + config.PrivateKeyPath)
			config.SSLPort = -1
		}

		if _, err := os.Stat(config.PublicKeyPath); os.IsNotExist(err) {
			fmt.Println("Public key not found at: " + config.PrivateKeyPath)
			config.SSLPort = -1
		}

		configJson, err := json.MarshalIndent(config, "", "    ")
		if err == nil {
			fmt.Println("Config:")
			fmt.Println(string(configJson))
		}

		if config.Port < 0 && config.SSLPort < 0 {
			config.Port = 80
		}

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

		// Cleanup on Ctrl-C and on Close
		defer cleanup()
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		go func() {
			for sig := range c {
				if os.Interrupt == sig || os.Kill == sig {
					cleanup()
					os.Exit(0)
				}
			}
		}()

		if config.SSLPort > 0 {
			fmt.Println("Serving HTTPS on port: " + strconv.Itoa(config.SSLPort))
			wg.Add(1)
			go func() {
				log.Fatal(http.ListenAndServeTLS(":"+strconv.Itoa(config.SSLPort), config.PublicKeyPath, config.PrivateKeyPath, nil))
				wg.Done()
			}()
		}

		if config.Port > 0 {
			fmt.Println("Serving HTTP on port: " + strconv.Itoa(config.Port))
			wg.Add(1)
			go func() {
				log.Fatal(http.ListenAndServe(":"+strconv.Itoa(config.Port), nil))
				wg.Done()
			}()
		}

		wg.Wait()
	}
}

package main

import (
	"github.com/robertkrimen/otto"
	"net/http"
)

type StaticResponse struct {
	Path        string `json:"path"`
	Response    string `json:"response"`
	File        string `json:"file"`
	ContentType string `json:"contentType"`
}

type DynamicResponse struct {
	Path         string      `json:"path"`
	Handler      string      `json:"handler"`
	ContentType  string      `json:"contentType"`
	GlobalConfig interface{} `json:"config"`
	Storage      interface{} `json:"storage"`
	vm           *otto.Otto
}

type Config struct {
	Port             int               `json:"port"`
	Beautify         bool              `json:"beautify"`
	StaticResponses  []StaticResponse  `json:"staticResponses"`
	DynamicResponses []DynamicResponse `json:"dynamicResponses"`
}

type Context struct {
	Request     *http.Request `json:"request"`
	Cfg         interface{}   `json:"config"`
	Storage     interface{}   `json:"storage"`
	ContentType string        `json:"contentType"`
}

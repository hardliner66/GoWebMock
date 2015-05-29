package main

import (
	"github.com/robertkrimen/otto"
)

func NewDynamicResponse(dr DynamicResponse) *DynamicResponse {
	o := DynamicResponse{}
	o.GlobalConfig = dr.GlobalConfig
	o.Handler = dr.Handler
	o.Path = dr.Path
	o.Storage = dr.Storage
	o.vm = otto.New()
	return &o
}

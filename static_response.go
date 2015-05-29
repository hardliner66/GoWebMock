package main

func NewStaticResponse(config StaticResponse) *StaticResponse {
	o := StaticResponse{}
	o.Path = config.Path
	o.Response = config.Response
	o.File = config.File
	return &o
}

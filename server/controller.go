package server

import (
	"encoding/json"
	"fmt"
	"github.com/demo-api/service"
	"github.com/valyala/fasthttp"
	"log"
)


func (s *Server) GetSoftwareDetails(ctx *fasthttp.RequestCtx) {

	fmt.Println(string(ctx.RequestURI()))
	//var soft service.Software
	fmt.Println("I am in controller")
	var res []byte
	var err error
	/*
	err = json.Unmarshal(ctx.QueryArgs().Peek("softwareName"),&soft)
	if err!=nil{
		log.Println(err)
		return
	}
	log.Println(soft)

	 */
	res, err = json.Marshal(service.GetSoftwareDetails(string(ctx.QueryArgs().Peek("softwareName")), s.mongoService))
	if err != nil {
		log.Printf("Unable to get details from database")
		ctx.Response.Header.Set("Access-Control-Allow-Origin", "*")
		ctx.Response.Header.Set("Content-Type", "application/json; charset=UTF-8")
		ctx.Response.SetStatusCode(400)
		return
	}

	ctx.Response.Header.Set("Access-Control-Allow-Origin", "*")
	ctx.Response.Header.Set("Content-Type", "application/json; charset=UTF-8")
	ctx.Response.SetStatusCode(200)
	if err == nil {
		ctx.Response.SetBody(res)
	}

	return
}



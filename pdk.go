package pdk

import (
	"github.com/karuppiah7890/go-pdk/client"
	"github.com/karuppiah7890/go-pdk/ip"
	"github.com/karuppiah7890/go-pdk/log"
	"github.com/karuppiah7890/go-pdk/nginx"
	"github.com/karuppiah7890/go-pdk/node"
	"github.com/karuppiah7890/go-pdk/request"
	"github.com/karuppiah7890/go-pdk/response"
	"github.com/karuppiah7890/go-pdk/router"
	"github.com/karuppiah7890/go-pdk/service"
	service_request "github.com/karuppiah7890/go-pdk/service/request"
	service_response "github.com/karuppiah7890/go-pdk/service/response"
)

// PDK go pdk module
type PDK struct {
	Client          client.Client
	Log             log.Log
	Nginx           nginx.Nginx
	Request         request.Request
	Response        response.Response
	Router          router.Router
	IP              ip.Ip
	Node            node.Node
	Service         service.Service
	ServiceRequest  service_request.Request
	ServiceResponse service_response.Response
}

// Init initialize go pdk
func Init(ch chan interface{}) *PDK {
	return &PDK{
		Client:          client.New(ch),
		Log:             log.New(ch),
		Nginx:           nginx.New(ch),
		Request:         request.New(ch),
		Response:        response.New(ch),
		Router:          router.New(ch),
		IP:              ip.New(ch),
		Node:            node.New(ch),
		Service:         service.New(ch),
		ServiceRequest:  service_request.New(ch),
		ServiceResponse: service_response.New(ch),
	}
}

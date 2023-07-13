package main

import (
	"log"
	"net"
	"project3/kitex_gen/kitex/demo/studentservice"
	// demo "project3/kitex_gen/kitex/demo/studentservice"
	"github.com/cloudwego/kitex/server"
)

func main() {
	// svr := demo.NewServer(new(StudentServiceImpl))
	var svr server.Server
	
	
	addr, _ := net.ResolveTCPAddr("tcp", ":9999")
	svr = studentservice.NewServer(new(StudentServiceImpl), server.WithServiceAddr(addr))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}

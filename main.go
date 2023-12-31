package main

import (
	"log"
	"net"
	"kitex-project/kitex_gen/kitex/demo/studentservice"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	// "github.com/cloudwego/kitex/server/genericserver"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func main() {
	r, err := etcd.NewEtcdRegistry([]string{"127.0.0.1:2379"})
	if err != nil {
		log.Fatal(err)
	}

	// svr := demo.NewServer(new(StudentServiceImpl))
	var svr server.Server
	
	// p, err := generic.NewThriftFileProvider("./idl/student.thrift")
    if err != nil {
        panic(err)
    }
	
    // g, err := generic.JSONThriftGeneric(p)
    if err != nil {
        panic(err)
    }

	addr, _ := net.ResolveTCPAddr("tcp", ":9999")
	//svr = studentservice.NewServer(new(StudentServiceImpl), server.WithServiceAddr(addr))
	svr = studentservice.NewServer(new(StudentServiceImpl),server.WithRegistry(r), server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
		ServiceName: "student",
	}), server.WithServiceAddr(addr) )

	// svr = genericserver.NewServer(new(StudentServiceImpl),g)
	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}

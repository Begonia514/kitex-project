package main

import (
	"context"
	"log"
	"time"

	// "time"

	demo "kitex-project/kitex_gen/kitex/demo"
	"kitex-project/kitex_gen/kitex/demo/studentservice"

	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
)


func main(){


	r, err := etcd.NewEtcdResolver([]string{"127.0.0.1:2379"})

	if err != nil{
		log.Fatal(err)
	}

	//c,err:= studentservice.NewClient("example",client.WithHostPorts("127.0.0.1:9999"))
	c := studentservice.MustNewClient("student",client.WithResolver(r))

	if err!=nil{
		log.Fatal(err)
	}



	RegiReq := &demo.Student{
		Id: 10,
		Name: "begonia",
		College: &demo.College{
			Name: "nju",
			Address: "no",
		},
		Email: nil,
	}

	for{
		ctx, cancel := context.WithTimeout(context.Background(),time.Second*3)
		resp, err := c.Register(ctx,RegiReq)
		cancel()
		if err != nil{
			log.Fatal(err)
		}
		log.Println(resp)
		time.Sleep(time.Second)
	}
	/*
	RegiResp,err := c.Register(context.Background(),RegiReq)

	if err != nil{
		log.Fatal(err)
	}
	log.Println(RegiResp)

	*/


	// QueryResp :=&demo.Student{
	// 	Id: 10,
	// 	Name: "begonia",
	// 	College: &demo.College{
	// 		Name: "nju",
	// 		Address: "no",
	// 	},
	// 	Email: nil,
	// }



}

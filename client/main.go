package main

import (
	"log"
	"context"
	// "time"

	"project3/kitex_gen/kitex/demo/studentservice"
	demo "project3/kitex_gen/kitex/demo"
	"github.com/cloudwego/kitex/client"
)


func main(){
	c,err:= studentservice.NewClient("example",client.WithHostPorts("127.0.0.1:9999"))
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

	RegiResp,err := c.Register(context.Background(),RegiReq)

	if err != nil{
		log.Fatal(err)
	}
	log.Println(RegiResp)


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

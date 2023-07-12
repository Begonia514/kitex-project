package main

import (
	"context"
	demo "project3/kitex_gen/kitex/demo"

	"github.com/cloudwego/kitex/pkg/klog"
)

// StudentServiceImpl implements the last service interface defined in the IDL.
type StudentServiceImpl struct{}

// Register implements the StudentServiceImpl interface.
func (s *StudentServiceImpl) Register(ctx context.Context, student *demo.Student) (resp *demo.RegisterResp, err error) {
	// TODO: Your code here...


	//stub
	klog.Infof("Received: ID: %d, Name: %s",student.Id,student.Name )
	resp = &demo.RegisterResp{
		Message: "OK!",
	}
	
	return
}

// Query implements the StudentServiceImpl interface.
func (s *StudentServiceImpl) Query(ctx context.Context, req *demo.QueryReq) (resp *demo.Student, err error) {
	// TODO: Your code here...


	//stub
	klog.Infof("Your query: %d",req.Id)

	resp = &demo.Student{
		Id: 10,
		Name: "begonia",
		College: &demo.College{
			Name: "NJU",
			Address: "no",
		},
		Email: nil,
	}
	return 
}

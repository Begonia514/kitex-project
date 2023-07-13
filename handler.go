package main

import (
	"context"
	demo "project3/kitex_gen/kitex/demo"

	"github.com/cloudwego/kitex/pkg/klog"
)

// StudentServiceImpl implements the last service interface defined in the IDL.
type StudentServiceImpl struct{}

var studentmap map[int]demo.Student = make(map[int]demo.Student)

// Register implements the StudentServiceImpl interface.
func (s *StudentServiceImpl) Register(ctx context.Context, student *demo.Student) (resp *demo.RegisterResp, err error) {
	// TODO: Your code here...

	_, ok := studentmap[int(student.Id)]

	if ok {
		resp = &demo.RegisterResp{
			Message: "id exist!",
			Success: false,
		}
		return
	}

	studentmap[int(student.Id)] = *student

	//stub
	klog.Infof("Received: ID: %d, Name: %s", student.Id, student.Name)
	resp = &demo.RegisterResp{
		Message: "OK!",
		Success: true,
	}

	return
}

// Query implements the StudentServiceImpl interface.
func (s *StudentServiceImpl) Query(ctx context.Context, req *demo.QueryReq) (resp *demo.Student, err error) {
	// TODO: Your code here...

	value, ok := studentmap[int(req.Id)]
	if !ok {
		resp = nil
		return
	}

	resp = &value

	//stub
	klog.Infof("query: %d", req.Id)
	/*
		resp = &demo.Student{
			Id: 10,
			Name: "begonia",
			College: &demo.College{
				Name: "NJU",
				Address: "no",
			},
			Email: nil,
		}
	*/
	return
}

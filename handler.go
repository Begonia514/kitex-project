package main

import (
	"context"
	demo "project3/kitex_gen/kitex/demo"
)

// StudentServiceImpl implements the last service interface defined in the IDL.
type StudentServiceImpl struct{}

// Register implements the StudentServiceImpl interface.
func (s *StudentServiceImpl) Register(ctx context.Context, student *demo.Student) (resp *demo.RegisterResp, err error) {
	// TODO: Your code here...
	var retErr error
	ret := new(demo.RegisterResp)

	//stub
	ret.Message="OK"
	ret.Success=true
	return ret,retErr
}

// Query implements the StudentServiceImpl interface.
func (s *StudentServiceImpl) Query(ctx context.Context, req *demo.QueryReq) (resp *demo.Student, err error) {
	// TODO: Your code here...

	var retErr error
	ret :=new(demo.Student)
	
	//stub
	ret.Id=10
	ret.Name="begonia"
	ret.College=new(demo.College)
	ret.Email=nil
	return ret,retErr
}

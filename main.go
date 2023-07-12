package main

import (
	"log"
	demo "project3/kitex_gen/demo/studentservice"
)

func main() {
	svr := demo.NewServer(new(StudentServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}

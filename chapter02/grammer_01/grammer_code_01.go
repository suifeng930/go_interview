package main

type Param map[string]interface{}

type Show struct {
	Param
}

func main1()  {

	s :=new(Show)             // new 无法初始化Show结构体中的param
	s.Param["RMB"]=10000     // panic: assignment to entry in nil map
}
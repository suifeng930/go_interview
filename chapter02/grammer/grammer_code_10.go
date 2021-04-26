package main

type query func(string) string

func exec(name string,vs ...query) string  {
	ch :=make(chan string)

	fn := func(i int) {
		ch <- vs[i](name)
	}
	for i, _ := range vs {
		go fn(i)
	}
	return <-ch
}

//func main()  {
//
//	querys :=[]query{}
//	querys1 := func(n string) string {
//		return n+"func1"
//	}
//	querys2:= func(n string) string {
//		return n+"func2"
//	}
//	querys3:= func(n string) string {
//		return n+"func3"
//	}
//	querys4:= func(n string) string {
//		return n+"func4"
//	}
//	querys=append(querys,querys1,querys2,querys3,querys4)
//	ret :=exec("111", querys...)
//	fmt.Println(ret)
//
//}
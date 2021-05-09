package main

//func main() {

	//out := make(chan int)
	//wg := sync.WaitGroup{}
	//wg.Add(2)
	////chann := channelGO{
	////	WaitGroup: wg,
	////	out:       out,
	////}
	////chann.makeNumber()
	////chann.printNumber()
	//
	//go func() {
	//	defer wg.Done()
	//	for i := 0; i < 5; i++ {
	//		out <- rand.Intn(5)
	//	}
	//	close(out)
	//}()
	//
	//go func() {
	//	defer wg.Done()
	//	for i := range out {
	//		fmt.Println(i)
	//	}
	//}()
	//wg.Wait()

	//

//	var foo=new(sync.RWMutex)
//	m := make(map[string]*entry)
//	rwMap :=&Map{
//		c:   m,
//		rmx: foo,
//	}
//
//	rwMap.Out("hehe","hello")
//	rwMap.Out("haha"," haha deng dao le ")
//	rwMap.Rd("haha",time.Second*10)
//	//select {
//	//
//	//}
//}

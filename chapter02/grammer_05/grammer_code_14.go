package main

import (
	"encoding/hex"
	"log"
)

type ServiceData struct {
	Id        string      `json:"id"`
	ProductID int32       `json:"productId"`
	DeviceID  int32       `json:"deviceId"`
	Params    interface{} `json:"params"`
}

func main() {

	//a := "aaaa"
	//ssh := *(*reflect.StringHeader)(unsafe.Pointer(&a))
	//b := *(*[]byte)(unsafe.Pointer(&ssh))
	//fmt.Printf("%v \n", b)
	//ss := ServiceData{}
	//fmt.Printf("%s \n", ss)

	sync := false
	//bytw :=[]byte{123, 10, 32, 32, 32, 32, 34, 107, 101, 121, 34, 58, 32, 34, 95, 51, 50, 48, 48, 95, 48, 95, 53, 55, 53, 48, 34, 44, 10, 32, 32, 32, 32, 34, 118, 97, 108, 117, 101, 34, 58, 32, 34, 70, 70, 100, 50, 50, 48, 50, 49, 45, 48, 53, 45, 49, 56, 32, 48, 57, 58, 52, 53, 58, 48, 55, 46, 51, 57, 51, 49, 55, 49, 32, 43, 48, 56, 48, 48, 32, 67, 83, 84, 32, 109, 61, 43, 48, 46, 48, 48, 48, 52, 48, 48, 49, 53, 53, 34, 44, 10, 32, 32, 32, 32, 34, 109, 105, 100, 34, 58, 32, 34, 68, 50, 34, 10, 125}
	//bytw :=[]byte{232, 22 ,118, 55, 70, 70 ,101, 52, 50 ,48 ,50 ,49 ,45 ,48, 53 ,45, 49, 56, 32, 49, 48, 58, 49, 49, 58, 51, 53, 46 ,50, 51, 53, 56, 53, 56, 32, 43, 48, 56, 48, 48, 32, 67, 83, 84, 32, 109, 61, 43, 48, 46 ,48, 48, 48, 52, 53, 56, 55, 55, 54}
	bytw := []byte{232, 22, 118, 55, 70, 70, 101, 52}

	for {
		sync=false
		if !sync {

			log.Println(string(bytw))
		}
		//log.Printf("lll")
		sync = true

	}

	toString := hex.EncodeToString(bytw)
	log.Println(toString)

	//
	//for i := 0; i < 15; i++ {
	//
	//	r := rand.Intn(25)
	//	log.Println(r)
	//}

}

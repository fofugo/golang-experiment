package main

import (
	"fmt"
	"reflect"
)

type Data struct {
	Name string
	Age  int
}
type Data2 struct {
	Name string
	Age  int
}

func main() {
	var datatype Data
	data := getData()
	//datatype=data <- will occur error
	assertedData := data.(Data)
	fmt.Println(reflect.TypeOf(assertedData))
	switch v := data.(type) {
	case Data:
		fmt.Println("this is Data")
		datatype = v
		fmt.Println(datatype)
		break
	case Data2:
		fmt.Println("this is Data2")
	default:
		fmt.Println("nothing")
	}
	var data2 Data2
	data2 = Data2(data.(Data))
	fmt.Println(data2)
}

func getData() interface{} {
	data := Data{Name: "alice", Age: 55}
	return data
}

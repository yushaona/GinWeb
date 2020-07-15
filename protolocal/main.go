package main

//https://www.php.cn/be/go/441604.html

import (
	"encoding/json"
	"fmt"
	"GinWeb/protolocal/test"
	"io/ioutil"
	"os"
	"sync"

	"github.com/golang/protobuf/proto"
)

func write() {

	p1 := &test.Person{
		Id:   1,
		Name: "小张",
		Phones: []*test.Phone{
			{Type: test.PhoneType_HOME, Number: "111111111"},
			{Type: test.PhoneType_WORK, Number: "222222222"},
		},
	}

	p2 := &test.Person{
		Id:   2,
		Name: "小王",
		Phones: []*test.Phone{
			{Type: test.PhoneType_HOME, Number: "333333333"},
			{Type: test.PhoneType_WORK, Number: "444444444"},
		},
	}

	//创建地址簿
	book := &test.ContactBook{}
	book.Persons = append(book.Persons, p1)
	book.Persons = append(book.Persons, p2)

	//编码数据
	j, _ := json.Marshal(book)
	fmt.Println(string(j))
	data, _ := proto.Marshal(book)

	//把数据写入文件
	ioutil.WriteFile("./test.txt", data, os.ModePerm)
}

func read() {

	//读取文件数据

	data, _ := ioutil.ReadFile("./test.txt")

	book := &test.ContactBook{}

	//解码数据

	proto.Unmarshal(data, book)

	for _, v := range book.Persons {

		fmt.Println(v.Id, v.Name)

		for _, vv := range v.Phones {

			fmt.Println(vv.Type, vv.Number)

		}

	}

}

type PName struct {
	a string
}

type MapWithLock struct {
	sync.RWMutex
	m map[string]int
}

// f returns 1

func f() (result int) {

	defer func() {

		result++

	}()

	return 0

}

func main() {
	fmt.Println(f())
	return
	mapVal := make(map[string]int, 2)
	mapVal["sss"] = 2
	mapVal["eee"] = 3

	ml := MapWithLock{
		m: mapVal,
	}

	ml.m["nash"] = 1
	ml.Lock()
	for _, v := range ml.m {
		fmt.Println(v)
	}
	ml.Unlock()

	return

	slic := make([]int, 1, 1)

	s := slic

	slic[0] = 2

	slic = append(slic, 3)

	fmt.Println(s)
	fmt.Println(slic)

	m := make(map[string]int, 1)

	m1 := m

	m["aa"] = 1

	m1["bb"] = 2
	fmt.Println(m)
	fmt.Println(m1)

	return

	// var s []*PName = []*PName{
	// 	&PName{a: "1"},
	// 	&PName{a: "2"},
	// }

	// for _, v := range s {
	// 	//fmt.Println(k)
	// 	fmt.Println(v.a)
	// }

	// return
	write()

	read()

}

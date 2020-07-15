package main

import (
	"GinWeb/stars/module"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gogo/protobuf/proto"
)

func main() {

	slic := make([]string, 0, 2)
	slic = append(slic, "3")
	slic = append(slic, "2")

	fmt.Println(slic[0])

	slic = append(slic, "1")
	fmt.Println(slic[2])
	return
	resp, err := http.Get("http://localhost:8080/protobuf")
	if err != nil {
		fmt.Println(err)
	} else {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		} else {
			user := &module.User{}
			proto.UnmarshalMerge(body, user)
			fmt.Println(*user)
		}

	}

	// resp, err := http.Get("http://localhost:8080/protobuf_map")
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	defer resp.Body.Close()
	// 	body, err := ioutil.ReadAll(resp.Body)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	} else {
	// 		m := make(map[string]interface{})
	// 		proto.UnmarshalMerge(body, m)
	// 		fmt.Println(m)
	// 	}

	// }
}

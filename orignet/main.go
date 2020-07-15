package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

var js map[string]interface{}

// func init() {
// 	js = make(map[string]interface{})
// }

func main() {

	// var str string = "你好吗"
	// var buffer bytes.Buffer
	// for _, r := range str {
	// 	cvt := string(r)
	// 	fmt.Println(cvt)
	// 	fmt.Printf("%c - %d - %U - %x \n", r, r, r, r)
	// 	if r >= 128 {
	// 		cvt = fmt.Sprintf("\\u%04x", int64(r))
	// 		fmt.Println(cvt)
	// 	}
	// 	buffer.WriteString(cvt)
	// }

	// fmt.Println(buffer.String())

	http.HandleFunc("/", Index)

	log.Fatal(http.ListenAndServe(":8080", nil))

	// writeContentType(w, jsonContentType)
	// jsonBytes, err := json.Marshal(obj)
	// if err != nil {
	// 	return err
	// }
	// _, err = w.Write(jsonBytes)
	// return err

	return
}

type jsonOut struct {
	ID   int    `json:"id"`
	Name string `json:"nAme"`
	Sex  int    `json:"SEX"`
}

func Index(w http.ResponseWriter, r *http.Request) {

	header := w.Header()
	if val := header["Content-Type"]; len(val) == 0 {
		header["Content-Type"] = []string{"application/json; charset=utf-8"}
	}

	var records []interface{} = make([]interface{}, 0)

	for i := 0; i < 2; i++ {
		records = append(records, map[string]interface{}{
			"Userid": strconv.Itoa(i),
			"Name":   "小名",
		})
	}

	js = map[string]interface{}{
		"code":    1,
		"state":   "ok",
		"records": records,
	}

	//还可以再嵌套一层 数组

	var res []interface{} = make([]interface{}, 0)

	res = append(res, js)
	//utf-8编码的字节流
	
	jsonBytes, err := json.Marshal(jsonOut{1, "2222", 3})
	if err != nil {
		return
	}
	_, err = w.Write(jsonBytes)

	// fmt.Fprint(w,"Blog:www.flysnow.org\nwechat:flysnow_org")
}

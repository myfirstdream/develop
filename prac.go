package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type info struct {
	Name string
	Age  int
}

type infomation struct {
	Infolice []info
}

func main() {

	var newinfo infomation
	strinfo := `{"Infolice":[{"Name":"Mantouone","Age":2},{"Name":"Mantoutwo","Age":3}]}` //json格式的字符串
	err := json.Unmarshal([]byte(strinfo), &newinfo)                                      //转化为json用struct接收
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(newinfo)

	newinfo.Infolice = append(newinfo.Infolice, info{Name: "mantouthree", Age: 4}) //append意思为向第一个参数后面加上第二个参数
	b, err := json.Marshal(newinfo)                                                //转化问字符串，接收
	if err != nil {
		fmt.Println(err)
	}
	if b != nil {

	}
	//fmt.Println(string(b))

	//对json文件的操作，比上面多了一个文件操作
	content, err := ioutil.ReadFile("F:/MyGo/src/hello/hello.json") //读文件内容，用content接收，括号里是文件地址
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("File contents: %s", content)

	data := `{"Infolice":[{"Name":"Mantouone","Age":2}]}` //将json格式的字符串转化为二进制流再写入
	var d1 = []byte(data)
	err = ioutil.WriteFile("F:/MyGo/src/hello/hello.json", d1, 0666) //括号里是文件地址，第二个是要写入的内容，第三个我也不知道就这么写
	if err != nil {
		log.Fatal(err)
	}

}

package main

import (
	"net/http"
	"fmt"
	"strings"
	"log"
)

//处理函数
func handler (w http.ResponseWriter,r *http.Request){
	r.ParseForm()  //解析参数，默认是不会解析的
	fmt.Println(r.Form)  //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello astaxie!") //这个写入到w的是输出到客户端的
}

func main(){
	http.HandleFunc("/",handler)
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

//浏览器输入地址 localhost:1200进行测试
//func TestHttp(t *testing.T){
//	t.Log("hello baby")
//	srv := http.Server{
//		Addr: ":1200",
//		Handler: &mhttp{},
//	}
//	go func() {
//		if err := srv.ListenAndServe(); err != nil {
//			log.Fatal(err)
//		}
//	}()
//	for ; ;  {
//		time.Sleep(time.Second)
//	}
//
//}
//
//func (mhttp)ServeHTTP(w http.ResponseWriter,r *http.Request) {
//	fmt.Println("hello baby ")
//	fmt.Fprintf(w, "Hello baby!") //这个写入到w的是输出到客户端的
//}

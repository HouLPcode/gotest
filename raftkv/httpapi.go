package main

import (
	"net/http"
	"strconv"
	"log"
	"io/ioutil"
	"fmt"
)

type httpKVApi struct {

}

func (h *httpKVApi)ServeHTTP(w http.ResponseWriter,r *http.Request){
	key := r.RequestURI
	switch {
	case r.Method == "PUT"://curl -XPUT 127.0.0.1:8080/key -d val
		v,_ := ioutil.ReadAll(r.Body)
		w.Write([]byte(fmt.Sprintln("PUT ","key :",key,"val :",v)))
	case r.Method == "GET"://curl -X GET 127.0.0.1:8080/key
		w.Write([]byte(fmt.Sprintln("GET ","key :",key)))
	}
}

//服务启动接口
func serveHttpKVAPI(port int,errorC <-chan error) {
	srv := http.Server{
		Addr: ":" + strconv.Itoa(port),
		Handler: &httpKVApi{},
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	// exit when raft goes down
	if err, ok := <-errorC; ok {
		log.Fatal(err)
	}
}




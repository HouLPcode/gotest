package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
	"testing"
)

type httpAPI struct{}

func (h httpAPI) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprint(w, err)
	}
	fmt.Println(string(buf))
	w.Write([]byte(strconv.Itoa(int(time.Now().Unix()))))
}

func TestHttp(t *testing.T) {
	server := http.Server{
		Addr:    ":1234",
		Handler: httpAPI{},
	}

	errC := make(chan struct{})
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			fmt.Println(err)
			errC <- struct{}{}
		}
	}()

	<-errC
}

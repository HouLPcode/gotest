package main

import "fmt"

func main(){
	fmt.Println("hello baby")
	err := make(chan error)
	serveHttpKVAPI(8080,err)
}

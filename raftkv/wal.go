package main

import (
	"github.com/etcd-io/etcd/wal"
	"log"
	"os"
	//"go.uber.org/zap"
	"fmt"
	"github.com/coreos/etcd/wal/walpb"
)

var mwal *wal.WAL

func openWal(dirpath string) *wal.WAL {
	if !wal.Exist(dirpath) {
		//创建文件夹
		if err := os.Mkdir(dirpath, 0750); err != nil {
			log.Fatalf("raftexample: cannot create dir for wal (%v)", err)
		}
		metadata := []byte("somedata")
		w, err := wal.Create(dirpath, metadata) //appending模式?????????????????????????
		if err != nil {
			fmt.Println("create wal err", err)
		}
		//关闭wal
		w.Close()
	}
	w, err := wal.Open(dirpath, walpb.Snapshot{})
	if err != nil {
		fmt.Println("wal open err", err)
	}
	mdata, state, ents, err := w.ReadAll()
	if err != nil {
		fmt.Println("readAll err", err)
	}
	fmt.Println("readAll data")
	fmt.Println("metadata = ", string(mdata))
	fmt.Println("state = ", state)
	fmt.Println("ents:")
	for _, v := range ents {
		fmt.Print(v)
	}
	return w
}

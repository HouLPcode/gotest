package main

import (
	"fmt"
	"github.com/etcd-io/etcd/raft"
	"time"
	"context"
	"github.com/coreos/etcd/raft/raftpb"
)

var (
	raftNode   raft.Node
	proposeC   chan string
	confChangeC chan raftpb.ConfChange
)

func init(){
	proposeC = make(chan string)
	confChangeC = make(chan raftpb.ConfChange)
}

func startRaft() {
	config := &raft.Config{
		ID:            uint64(1),
		ElectionTick:  10,
		HeartbeatTick: 1,
		MaxSizePerMsg:   1024 * 1024,
		MaxInflightMsgs: 256,
	}
	peers := []raft.Peer{{ID: uint64(1)}}

	raftNode = raft.StartNode(config, peers)
	//fmt.Println(node)
}

func serveChannel() {
	//逻辑时钟，100ms每次
	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()

	go func() {
		for{
			select {
			case  v:= <- proposeC:
				fmt.Println("send kv to raft")
				if err := raftNode.Propose(context.TODO(),[]byte(v));err != nil{
					fmt.Println(err)
				}
			case v:= <- confChangeC:
				fmt.Println("send config to raft")
				if err := raftNode.ProposeConfChange(context.TODO(),v);err != nil{
					fmt.Println(err)
				}
			}
		}
	}()

	go func() {
		for {
			select {
			case <-ticker.C:
				fmt.Println("ticker")
				raftNode.Tick()
			case rd := <-raftNode.Ready():
				//处理rd信息
				mwal.Save(rd.HardState,rd.Entries)
				fmt.Println("receive kv data from raft",rd)
				raftNode.Advance()
			}
		}
	}()
}

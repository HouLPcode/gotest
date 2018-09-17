package main

import (
	"bytes"
	"encoding/gob"
	"log"
)

type kv struct {
	Key string
	Val string
}

func encode(key, val string) string {
	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(kv{Key: key, Val: val}); err != nil {
		log.Fatal(err)
	}
	return buf.String()
}

func decode(data string) kv {
	var datakv kv
	if err := gob.NewDecoder(bytes.NewBufferString(data)).Decode(&datakv); err != nil {
		log.Fatal(err)
	}
	return datakv
}

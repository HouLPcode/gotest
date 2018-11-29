package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	type ColorGroup struct {
		ID     int
		Name   string
		Colors []string
	}
	group := ColorGroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}

	// 编码
	b, err := json.Marshal(group)
	if err != nil {
		fmt.Println("error:", err)
	}
	os.Stdout.Write(b)

	// 解码
	var group1 ColorGroup
	err = json.Unmarshal(b, &group1)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("\n%+v\n", group1)

	dec := json.NewDecoder(strings.NewReader(b))
	for {
		var m ColorGroup
		if err := dec.Decode(&m); err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
		fmt.Printf("%s: %s\n", m.ID, m.Name, m.Colors)
	}
}

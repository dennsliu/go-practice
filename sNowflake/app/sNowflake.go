package main

import (
	"fmt"
	"os"

	"github.com/bwmarrin/sNowflake"
)

func main() {
	n, err := sNowflake.NewNode(1)
	if err != nil {
		println(err)
		os.Exit(1)
	}
	for i := 0; i < 3; i++ {
		id := n.Generate()
		fmt.Println("id", id)
		fmt.Println(
			"node: ", id.Node(), "step: ", id.Step(), "time: ", id.Time(), "\n")
	}
}

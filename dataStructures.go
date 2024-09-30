package main

import "github.com/tyspice/data-structures/linkedList"

func main() {
	lst := linkedList.New("Hello")
	lst.AddNode("snake")
	lst.AddNode("poop")
	lst.AddNode("chunder")
	lst.AddNode("snook")
	lst.PrintNodeData()
}

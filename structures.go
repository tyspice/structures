package main

import "github.com/tyspice/structures/linkedlist"

func main() {
	lst := linkedlist.New("Hello")
	lst.AddNode("snake")
	lst.AddNode("poop")
	lst.AddNode("chunder")
	lst.AddNode("snook")
	lst.PrintNodeData()
}

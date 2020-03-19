package main

import "fmt"

type SLLNode struct {
	next  *SLLNode
	value int
}

func (sNode *SLLNode) SetValue(v int) {
	sNode.value = v
}

func (sNode *SLLNode) GetValue() int {
	return sNode.value
}

func NewSLLNode() *SLLNode {
	return new(SLLNode)
}

func main() {
	node := NewSLLNode()
	node.SetValue(4)
	fmt.Println("Node is of value: ", node.GetValue())
}

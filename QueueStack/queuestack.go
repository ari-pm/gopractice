package queuestack

import (
)

type Node int

type Queue []*Node

type Stack []*Node


func (s *Queue) Push(node *Node) {
	*s = append(*s, node)
}
func (s *Stack) Push(node *Node) {
	*s = append(*s, node)
}
func (s *Queue) Pop() *Node{
	length := len(*s)
	if length <= 0 {
		return nil
	}
	output := (*s)[0]
	*s = (*s)[1:]
	return output
}
func (s *Stack) Pop() *Node{
	length := len(*s)
	if length <= 0 {
		return nil
	}
	output := (*s)[length-1]
	*s = (*s)[:length-1]
	return output
}
func (s *Stack) Len() int {
	return len(*s)
}
func (s *Queue) Len() int {
	return len(*s)
}
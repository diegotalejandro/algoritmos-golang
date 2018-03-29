package main
import (
  "fmt"
)
type node struct {
         value int
         balance int
         left_child *node
         right_child *node
}

func balance(root *node) int {
	if root != nil {
		return root.balance
	}
	return -1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func insert(root *node, value int) *node {
	if root == nil {
		root = &node{value, 0, nil, nil}
		root.balance = max(balance(root.left_child), balance(root.right_child)) + 1
		return root
	}

	if value < root.value {
		root.left_child = insert(root.left_child, value)
	}

	if value > root.value {
		root.right_child = insert(root.right_child, value)
	}

	root.balance = max(balance(root.left_child), balance(root.right_child)) + 1
	return root
}
func AVLcheck(v *node) bool{
 if v==nil{
  return false
  }
  v.balanceAVL()
  fmt.Println(balance(v.left_child)-balance(v.right_child))
  if balance(v.left_child)-balance(v.right_child) < 1 {
  return true
  }else{
  return false
  }
  return false
}
func (n *node) balanceAVL(){
  if n == nil{
    return
  }
  n.left_child.balanceAVL()
  n.right_child.balanceAVL()
  n.balance=max(balance(n.left_child), balance(n.right_child)) + 1
}
func main(){
  raiz:=&node{4,0,nil,nil}
  insert(raiz,2)
  insert(raiz,6)
  insert(raiz,1)
  insert(raiz,3)
  insert(raiz,5)
  insert(raiz,7)
  fmt.Println(AVLcheck(raiz.right_child))

}

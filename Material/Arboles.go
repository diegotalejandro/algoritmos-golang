package tree
import()
type tree struct {
root *node
}
type node struct {
left *node
right *node
value int
}
func PrintInOrder(n* node,numero int) {
  if n == nil{
    return
  }
  n.left.PrintInOrder()
  fmt.Println(n.value)
  n.right.PrintInOrder()
}

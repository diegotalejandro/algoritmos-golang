package main
import ("fmt"
)
type node struct {
Value int
Left *node
Right *node
}
/*func PrintInOrder(n* node) {
  if n == nil{
    return
  }
  n.Left.PrintInOrder()
  fmt.Println(n.Value)
  n.Right.PrintInOrder()
}*/
func Insert(t *node, v int) *node {
	if t == nil {
		return &node{v,nil,nil}
	}
	if v < t.Value {
		t.Left = Insert(t.Left, v)
		return t
	}
	t.Right = Insert(t.Right, v)
	return t
}
var contador []int
var boolian bool

func (t *node)Distance(v int) {
  if t == nil{
    return
  }
  if boolian == false {
    contador=append(contador,t.Value)
  }else{ return}
  t.Left.Distance(v)
  if t.Value == v {
    boolian=true
    return
  }
    contador=append(contador[:len(contador)-1])
  t.Right.Distance(v)
  /*contador=append(contador[:len(contador)-1])
  if t.Value == v {
contador=append(contador,t.Value)
    boolian=true
    return
  }*/
}

func main(){
raiz:=&node{0,nil,nil}
var aux int
var nvariables int
fmt.Scan(&nvariables)
for i:=0 ; i < nvariables ; i++ {
fmt.Scan(&aux)
raiz= Insert(raiz,aux)
 }
var a int
var b int
var num int
 for i:=0 ; i < nvariables ; i++ {
 fmt.Scan(&a)
 fmt.Scan(&b)
 raiz.Distance(a)
 arr:=contador
 contador=append(contador[:0])
 boolian=false
 raiz.Distance(b)
 brr:=contador
 contador=append(contador[:0])
 boolian=false
 for j:=0; j< nvariables ;j++{
  if arr[j] == brr[j]{
     num++
     fmt.Println(arr,brr)
     break
   }else{
     break
   }
 }
 fmt.Println(len(arr)+len(brr)-(num*2))
 num=0
  }

}

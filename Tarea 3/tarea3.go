package main
import (
  "fmt"
  "math/rand"
  "time"
)

type KEsimo struct {
    data []int // datos desordenados
    value int // valor del nodo (sólo si data == nil)
    lessThan int // cuantos valores hay a la izquierda
    left *KEsimo // hijo izquierdo
    right *KEsimo // hijo derecho
}

func New(data []int) *KEsimo {
    return &KEsimo{data: data}
}

func (ke *KEsimo) Get(k int) int{
 if ke == nil{return 0}
  if ke.data == nil && ke.lessThan == k {
    return ke.value
  }
  if ke.data == nil && ke.lessThan < k {
    return ke.right.Get(k-ke.lessThan-1)
  }
  if ke.data == nil && ke.lessThan > k {
    return ke.left.Get(k)
  }
  if ke.data != nil {
    ke.value = ke.data[rand.Intn(len(ke.data))]
    var auxR []int
    var auxL []int
    for _, item := range ke.data {
      switch {
			case item < ke.value:
        auxL = append(auxL , item)
        ke.left=New(auxL)
			case item > ke.value:
        auxR = append(auxR , item)
        ke.right=New(auxR)
			}
		}
    ke.data=nil
    ke.lessThan=len(auxL)
    ke.Get(k)
    ke.left.Get(k)
    ke.right.Get(k)
  }
  return 0
}
func main(){
var arr []int
for i:=1; i<= 1000; i++{
  arr=append(arr, i)
}
raiz:=New(arr)
a:=time.Now()
raiz.Get(1)
b:=time.Now()
fmt.Println(b.Sub(a))
for i:=0;i < 1000;i++{
  arr[i]=rand.Intn(1000000)
}
raiz=New(arr)
a=time.Now()
raiz.Get(1)
b=time.Now()
fmt.Println(b.Sub(a))
for i:=0;i < 100;i++{
  arr[i]=i
}
for i:=100;i < 1000;i++{
  arr[i]=rand.Intn(1000000)
}
raiz=New(arr)
a=time.Now()
raiz.Get(1)
b=time.Now()
fmt.Println(b.Sub(a))
for i:=0;i < 100;i++{
  arr[i]=i
}
for i:=100;i < 200;i++{
  arr[i]=rand.Intn(1000000)
}
for i:=200;i < 400;i++{
  arr[i]=i
}
for i:=400;i < 600;i++{
  arr[i]=rand.Intn(1000000)
}
for i:=600;i < 1000;i++{
  arr[i]=i
}
raiz=New(arr)
a=time.Now()
raiz.Get(1)
b=time.Now()
fmt.Println(b.Sub(a))
}

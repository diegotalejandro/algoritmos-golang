package main
import ("fmt"
"math/rand"
)
func Slice(n []int) []int{
if len(n)<=1{
  return nil
}
var numeros []int
  pivot:=rand.Intn(len(n)-1)
  left:=n[:pivot]
  right:=n[pivot+1:]
  for i:=0;i< len(left);i++{
    if n[pivot]==left[i]{
      L:=left[:i]
      R:=left[i+1:]
      L=append(L,R...)
      left=L
      numeros=append(numeros,n[pivot])
      numeros=append(numeros,Slice(L)...)
      break
    }
  }
    for i:=0;i< len(right);i++{
      if n[pivot]==right[i]{
        L:=right[:i]
        R:=right[i+1:]
        L=append(L,R...)
        right=L
        numeros=append(numeros,n[pivot])
        numeros=append(numeros,Slice(L)...)
        break
      }
  }
  left=append(left,right...)
numeros=append(numeros,left...)
  return numeros
 }


func main(){
  var slice []int
  slice= append(slice,1,2,3,2,1)
  fmt.Println(Slice(slice))
}

package main

import "fmt"

func main() {
var t,n,k,arr int
var a []int
var aux []int
var temp []int
    fmt.Scan(&t)
    for i:=0 ; i < t ; i++{
      fmt.Scan(&n)
      fmt.Scan(&k)
      for j:=0 ; j < n ; j++{
        fmt.Scan(&arr)
      if i==0{
      a=append(a,arr)
      }else{
      	a[j]=arr
      }
    }
    for y:=0 ; k < 3;y++ {
      for x:=0; x < k ; x++{
        if y==0 {
          aux=append(aux,a[x])
        }else{
          aux[x]=a[x]
        }
      }

      for z:=1 ; z < k; z++{
        if aux[z-1]>aux[z]{
          aux[z-1], aux[z] = aux[z], aux[z-1]
        }//funciona
      }
      fmt.Print("----------------",aux[len(aux)-1],"-----------------")
      fmt.Print(aux)
      for w:=1; w < len(a); w++{
      	if y==0 {
      		temp=append(temp,a[w])
      	}else{
      		temp[w-1]=a[w]
      	}
      }
      a=temp
      fmt.Print(a)
    }
  }
}

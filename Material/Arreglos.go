package main
import ("fmt")
 func  main(){
 var arr []int
 for i:=0; i < 10; i++{
   arr=append(arr,i+1)
 }
 arr=append(arr[:1])
 fmt.Println(arr)
 }

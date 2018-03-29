package main
import ("fmt"
)
/*func Asd(X [10]int , y int)  (int,int,bool) {
for i, j := 0, len(X) - 1; i < len(X) && j >= 0; {

if X[i] + X[j] < y {
i++
}else if X[i] + X[j] > y {
j--
}else{
return i, j, true
  }

 }
return -1, -1, false
}
func Asd2(X [10]int)  (int,int,int,bool) {
for k:=len(X)-1;k >=0;k--{
for i, j := 0, len(X) - 1; i < len(X) && j >= 0; {
if X[i] + X[j] < X[k] {
i++
}else if X[i] + X[j] > X[k] {
j--
}else{
return i, j, k , true
   }
  }
 }
return -1, -1, -1, false
}*/
func Asd3(X [10]int)  (int,int,int,bool) {
for k:=len(X)-1;k >=0;k--{
for i, j := 0, len(X) - 1; i < len(X) && j >= 0; {
if X[i] + X[j] + X[k] < 0 {
i++
}else if X[i] + X[j] + X[k] > 0 {
j--
}else{
return i, j, k , true
   }
  }
 }
return -1, -1, -1, false
}
func main(){
  arr:=[10]int{-1,2,3,4,5,6,7,8,9,10}
  i,j,k,booliano:=Asd3(arr)
  fmt.Println(i,j,k,booliano)

}

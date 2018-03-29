package main
import ("fmt"
)
func BusquedaBinaria(arr [10]int, valor int,pos int) int {

	inf := 0
	sup := len(arr)-1

	for inf <= sup {

		medio := (inf + sup) / 2

		if arr[medio] < valor {
			inf = medio + 1
      fmt.Println(medio)

		} else {
			sup = medio - 1
		}
	}

	if inf == len(arr) || arr[inf] != valor {
		return -1
	} else {
		return inf
	}

}

func BusquedaExponencial(arr [10]int, tamano int, valor int) int {
    if (tamano == 0) {
        fmt.Println("no existe")
        return -1
    }

    pos := 1
    for  pos < tamano && arr[pos] < valor {
        pos *= 2
    }
fmt.Println(arr, valor, Min(pos/2, tamano))
    i:=BusquedaBinaria(arr, valor, Min(pos/2,tamano))
    return i
}
func Min(x int, y int) int{
  if x>=y{
    return y
  }else {return x}
}
func main(){
  arr:=[10]int{1,2,3,4,5,6,7,8,9,10}
  numero:=10
  posicion:=BusquedaExponencial(arr,len(arr),numero)
  fmt.Println(posicion)

}

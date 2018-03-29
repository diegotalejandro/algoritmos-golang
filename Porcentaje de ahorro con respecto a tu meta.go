package main
import "fmt"
func main(){
var dinero int
var total int
fmt.Println("Introducir dinero ahorrado: ")
fmt.Scan(&dinero)
fmt.Println("Introducir dinero total a ahorrar: ")
fmt.Scan(&total)
fmt.Println("Llevas el ",((dinero*100)/total),"% del total")
}

package main
import (
	"fmt"
	"math/rand"
	"io/ioutil"
	"log"
	"strings"
	"strconv"
	"testing"
)
func Swap(arrayzor []int, i, j int) {
	tmp := arrayzor[j]
	arrayzor[j] = arrayzor[i]
	arrayzor[i] = tmp
}
func BubbleSort(data []int) []int {
	unsorted := true
	for unsorted {
		unsorted = false
		for i := len(data) - 1; i > 0; i-- {
			if data[i] < data[i-1] {
				data[i], data[i-1] = data[i-1], data[i]
				unsorted = true
			}
		}
	}
	return data
}
func SelectionSort(data []int) []int {
	for i := 0; i < len(data)-1; i++ {
		min := i
		for j := i + 1; j < len(data)-1; j++ {
			if data[j] < data[min] {
				min = j
			}
		}
		Swap(data, i, min)
	}
	return data
}
func InsertionSort(data []int) []int {
  for i := 1; i < len(data); i++ {
  		for j := i; j > 0 && data[j] < data[j - 1]; j-- {
  			Swap(data, j, j - 1)
  		}
  	}
    return data
}
func MergeSort(data []int) []int {
  if len(data) <= 1 {
      return data
    }

    left := make([]int, 0)
    right := make([]int, 0)
    m := len(data) / 2

    for i, x := range data {
      switch {
      case i < m:
        left = append(left, x)
      case i >= m:
        right = append(right, x)
      }
 }
 left = MergeSort(left)
   right = MergeSort(right)

   return Merge(left, right)
 }




 func Merge(left, right []int) []int {

   results := make([]int, 0)

   for len(left) > 0 || len(right) > 0 {
     if len(left) > 0 && len(right) > 0 {
       if left[0] <= right[0] {
         results = append(results, left[0])
         left = left[1:len(left)]
       } else {
         results = append(results, right[0])
         right = right[1:len(right)]
       }
     } else if len(left) > 0 {
       results = append(results, left[0])
       left = left[1:len(left)]
     } else if len(right) > 0 {
       results = append(results, right[0])
       right = right[1:len(right)]
     }
   }

   return results
 }

func QuickSort(data []int) []int {
	length := len(data)

		if length <= 1 {
			dataCopy := make([]int, length)
			copy(dataCopy, data)
			return dataCopy
		}

		m := data[rand.Intn(length)]

		less := make([]int, 0, length)
		middle := make([]int, 0, length)
		more := make([]int, 0, length)

		for _, item := range data {
			switch {
			case item < m:
				less = append(less, item)
			case item == m:
				middle = append(middle, item)
			case item > m:
				more = append(more, item)
			}
		}

		less, more = QuickSort(less), QuickSort(more)

		less = append(less, middle...)
		less = append(less, more...)

		return less
}
func Convert(numerosstring []string) []int {
  var numerosint []int
		for i:=0;i < len(numerosstring) ; i++ {
			//conv,err := strconv.Atoi(numerosstring[i])
			//fmt.Println(numerosstring[i])
			if s, err := strconv.Atoi(numerosstring[i]); err == nil {
				//fmt.Print(s)
      numerosint = append(numerosint,s)
		}
	}
	return numerosint
}
func benchmarkBubbleSort(array []int, bk *testing.B) {
	for n := 0; n < bk.N; n++ {
		BubbleSort(array)
	}
}
func main() {
 var directorio string = "./tarea2"
 files, err := ioutil.ReadDir(directorio)
 if err != nil {
	 log.Fatal(err)
 }
 var numerosstring []string
 var numerosint []int
 //recorrido de los archivos comprobando cada funcion
 for _, file := range files {
fmt.Println(file.Name()+": ")
	 x, _ := ioutil.ReadFile(directorio + "/" + file.Name())//lee el texto
	 numerosstring = strings.Fields(string(x))
	 numerosint = Convert(numerosstring)
BubbleS:=BubbleSort(numerosint)
benchmarkBubbleSort(numerosint)
SelectionS:=SelectionSort(numerosint)
InsertionS:=InsertionSort(numerosint)
MergeS:=MergeSort(numerosint)
QuickS:=QuickSort(numerosint)
fmt.Println("listo",BubbleS[0],SelectionS[0],InsertionS[0],MergeS[0],QuickS[0])
break
 }
}

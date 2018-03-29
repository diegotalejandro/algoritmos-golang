package sort

import ("testing"
        "io/ioutil"
        "strings"
        "log"
        //"fmt"
        "strconv"
)

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
func All(nombre string) []int {
  files, err := ioutil.ReadDir("./tarea2")
  if err != nil {
   log.Fatal(err)}
  var numerosstring []string
  var numerosint []int
  for _,file:=range files{
if nombre == file.Name() {
       //fmt.Println(file.Name()+": ")
    	 x, _ := ioutil.ReadFile("./tarea2/" + file.Name())//lee el texto
    	 numerosstring = strings.Fields(string(x))
    	 numerosint = Convert(numerosstring)
       break
     }
 }
 return numerosint
}

func benchmarkBubbleSort(array []int, bk *testing.B) {
	for n := 0; n < bk.N; n++ {
		BubbleSort(array)
	}
}

func benchmarkSelectionSort(array []int, bk *testing.B) {
	for n := 0; n < bk.N; n++ {
		SelectionSort(array)
	}
}

func benchmarkInsertionSort(array []int, bk *testing.B) {
	for n := 0; n < bk.N; n++ {
		InsertionSort(array)
	}
}

func benchmarkMergeSort(array []int, bk *testing.B) {
	for n := 0; n < bk.N; n++ {
		MergeSort(array)
	}
}

func benchmarkQuickSort(array []int, bk *testing.B) {
	for n := 0; n < bk.N ; n++ {
		QuickSort(array)
	}
}
/*
func BenchmarkBubbleSortRandom1k(bk *testing.B) {
var numerosint []int = All("random1k.txt")
	benchmarkBubbleSort(numerosint, bk)
}
func BenchmarkBubbleSortRandom10k(bk *testing.B) {
var numerosint []int = All("random10k.txt")
	benchmarkBubbleSort(numerosint, bk)
}
func BenchmarkBubbleSortRandom100k(bk *testing.B) {
var numerosint []int = All("random100k.txt")
	benchmarkBubbleSort(numerosint, bk)
}
func BenchmarkBubbleSortRandom1m(bk *testing.B) {
var numerosint []int = All("random1m.txt")
	benchmarkBubbleSort(numerosint, bk)
}
func BenchmarkSelectionSortRandom1k(bk *testing.B) {
	var numerosint []int = All("random1k.txt")
	benchmarkSelectionSort(numerosint, bk)
}
func BenchmarkSelectionSortRandom10k(bk *testing.B) {
	var numerosint []int = All("random10k.txt")
	benchmarkSelectionSort(numerosint, bk)
}
func BenchmarkSelectionSortRandom100k(bk *testing.B) {
	var numerosint []int = All("random100k.txt")
	benchmarkSelectionSort(numerosint, bk)
}
func BenchmarkSelectionSortRandom1m(bk *testing.B) {
	var numerosint []int = All("random1m.txt")
	benchmarkSelectionSort(numerosint, bk)
}

func BenchmarkInsertSortRandom1k(bk *testing.B) {
	var numerosint []int = All("random1k.txt")
	benchmarkInsertionSort(numerosint, bk)
}
func BenchmarkInsertSortRandom10k(bk *testing.B) {
	var numerosint []int = All("random10k.txt")
	benchmarkInsertionSort(numerosint, bk)
}
func BenchmarkInsertSortRandom100k(bk *testing.B) {
	var numerosint []int = All("random100k.txt")
	benchmarkInsertionSort(numerosint, bk)
}
func BenchmarkInsertSortRandom1m(bk *testing.B) {
	var numerosint []int = All("random1m.txt")
	benchmarkInsertionSort(numerosint, bk)
}

func BenchmarkMergeSortRandom1k(bk *testing.B) {
	var numerosint []int = All("random1k.txt")
	benchmarkMergeSort(numerosint, bk)
}
func BenchmarkMergeSortRandom10k(bk *testing.B) {
	var numerosint []int = All("random10k.txt")
	benchmarkMergeSort(numerosint, bk)
}
func BenchmarkMergeSortRandom100k(bk *testing.B) {
	var numerosint []int = All("random100k.txt")
	benchmarkMergeSort(numerosint, bk)
}
func BenchmarkMergeSortRandom1m(bk *testing.B) {
	var numerosint []int = All("random1m.txt")
	benchmarkMergeSort(numerosint, bk)
}
*/
func BenchmarkQuickSortRandom1k(bk *testing.B) {
	var numerosint []int = All("random1k.txt")
	benchmarkQuickSort(numerosint, bk)
} /*
func BenchmarkQuickSortRandom10k(bk *testing.B) {
	var numerosint []int = All("random10k.txt")
	benchmarkQuickSort(numerosint, bk)
}
func BenchmarkQuickSortRandom100k(bk *testing.B) {
	var numerosint []int = All("random100k.txt")
	benchmarkQuickSort(numerosint, bk)
}
func BenchmarkQuickSortRandom1m(bk *testing.B) {
	var numerosint []int = All("random1m.txt")
	benchmarkQuickSort(numerosint, bk)
}
*/

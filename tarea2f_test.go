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

func BenchmarkBubbleSortDuplicates1k(bk *testing.B) {
var numerosint []int = All("duplicates1k.txt")
	benchmarkBubbleSort(numerosint, bk)
}
func BenchmarkBubbleSortDuplicates10k(bk *testing.B) {
var numerosint []int = All("duplicates10k.txt")
	benchmarkBubbleSort(numerosint, bk)
}
func BenchmarkBubbleSortDuplicates100k(bk *testing.B) {
var numerosint []int = All("duplicates100k.txt")
	benchmarkBubbleSort(numerosint, bk)
}
func BenchmarkBubbleSortDuplicates1m(bk *testing.B) {
var numerosint []int = All("duplicates1m.txt")
	benchmarkBubbleSort(numerosint, bk)
}


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


func BenchmarkBubbleSortReversed1k(bk *testing.B) {
var numerosint []int = All("reversed1k.txt")
	benchmarkBubbleSort(numerosint, bk)
}
func BenchmarkBubbleSortReversed10k(bk *testing.B) {
var numerosint []int = All("reversed10k.txt")
	benchmarkBubbleSort(numerosint, bk)
}
func BenchmarkBubbleSortReversed100k(bk *testing.B) {
var numerosint []int = All("reversed100k.txt")
	benchmarkBubbleSort(numerosint, bk)
}
func BenchmarkBubbleSortReversed1m(bk *testing.B) {
var numerosint []int = All("reversed1m.txt")
	benchmarkBubbleSort(numerosint, bk)
}


func BenchmarkBubbleSortSorted1k(bk *testing.B) {
var numerosint []int = All("sorted1k.txt")
	benchmarkBubbleSort(numerosint, bk)
}
func BenchmarkBubbleSortSorted10k(bk *testing.B) {
var numerosint []int = All("sorted10k.txt")
	benchmarkBubbleSort(numerosint, bk)
}
func BenchmarkBubbleSortSorted100k(bk *testing.B) {
var numerosint []int = All("sorted100k.txt")
	benchmarkBubbleSort(numerosint, bk)
}
func BenchmarkBubbleSortSorted1m(bk *testing.B) {
var numerosint []int = All("sorted1m.txt")
	benchmarkBubbleSort(numerosint, bk)
}



func BenchmarkSelectionSortDuplicates1k(bk *testing.B) {
	var numerosint []int = All("duplicates1k.txt")
	benchmarkSelectionSort(numerosint, bk)
}
func BenchmarkSelectionSortDuplicates10k(bk *testing.B) {
	var numerosint []int = All("duplicates10k.txt")
	benchmarkSelectionSort(numerosint, bk)
}
func BenchmarkSelectionSortDuplicates100k(bk *testing.B) {
	var numerosint []int = All("duplicates100k.txt")
	benchmarkSelectionSort(numerosint, bk)
}
func BenchmarkSelectionSortDuplicates1m(bk *testing.B) {
	var numerosint []int = All("duplicates1m.txt")
	benchmarkSelectionSort(numerosint, bk)
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


func BenchmarkSelectionSortReversed1k(bk *testing.B) {
	var numerosint []int = All("reversed1k.txt")
	benchmarkSelectionSort(numerosint, bk)
}
func BenchmarkSelectionSortReversed10k(bk *testing.B) {
	var numerosint []int = All("reversed10k.txt")
	benchmarkSelectionSort(numerosint, bk)
}
func BenchmarkSelectionSortReversed100k(bk *testing.B) {
	var numerosint []int = All("reversed100k.txt")
	benchmarkSelectionSort(numerosint, bk)
}
 func BenchmarkSelectionSortReversed1m(bk *testing.B) {
	var numerosint []int = All("reversed1m.txt")
	benchmarkSelectionSort(numerosint, bk)
}


func BenchmarkSelectionSortSorted1k(bk *testing.B) {
	var numerosint []int = All("sorted1k.txt")
	benchmarkSelectionSort(numerosint, bk)
}
func BenchmarkSelectionSortSorted10k(bk *testing.B) {
	var numerosint []int = All("sorted10k.txt")
	benchmarkSelectionSort(numerosint, bk)
}
func BenchmarkSelectionSortSorted100k(bk *testing.B) {
	var numerosint []int = All("sorted100k.txt")
	benchmarkSelectionSort(numerosint, bk)
}
 func BenchmarkSelectionSortSorted1m(bk *testing.B) {
	var numerosint []int = All("sorted1m.txt")
	benchmarkSelectionSort(numerosint, bk)
}



func BenchmarkInsertSortDuplicates1k(bk *testing.B) {
	var numerosint []int = All("duplicates1k.txt")
	benchmarkInsertionSort(numerosint, bk)
}
func BenchmarkInsertSortDuplicates10k(bk *testing.B) {
	var numerosint []int = All("duplicates10k.txt")
	benchmarkInsertionSort(numerosint, bk)
}
func BenchmarkInsertSortDuplicates100k(bk *testing.B) {
	var numerosint []int = All("duplicates100k.txt")
	benchmarkInsertionSort(numerosint, bk)
}
 func BenchmarkInsertSortDuplicates1m(bk *testing.B) {
	var numerosint []int = All("duplicates1m.txt")
	benchmarkInsertionSort(numerosint, bk)
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

func BenchmarkInsertSortReversed1k(bk *testing.B) {
	var numerosint []int = All("reversed1k.txt")
	benchmarkInsertionSort(numerosint, bk)
}
func BenchmarkInsertSortReversed10k(bk *testing.B) {
	var numerosint []int = All("reversed10k.txt")
	benchmarkInsertionSort(numerosint, bk)
}
func BenchmarkInsertSortReversed100k(bk *testing.B) {
	var numerosint []int = All("reversed100k.txt")
	benchmarkInsertionSort(numerosint, bk)
}
 func BenchmarkInsertSortReversed1m(bk *testing.B) {
	var numerosint []int = All("reversed1m.txt")
	benchmarkInsertionSort(numerosint, bk)
}


func BenchmarkInsertSortSorted1k(bk *testing.B) {
	var numerosint []int = All("sorted1k.txt")
	benchmarkInsertionSort(numerosint, bk)
}
func BenchmarkInsertSortSorted10k(bk *testing.B) {
	var numerosint []int = All("sorted10k.txt")
	benchmarkInsertionSort(numerosint, bk)
}
func BenchmarkInsertSortSorted100k(bk *testing.B) {
	var numerosint []int = All("sorted100k.txt")
	benchmarkInsertionSort(numerosint, bk)
}
 func BenchmarkInsertSortSorted1m(bk *testing.B) {
	var numerosint []int = All("sorted1m.txt")
	benchmarkInsertionSort(numerosint, bk)
}



func BenchmarkMergeSortDuplicates1k(bk *testing.B) {
	var numerosint []int = All("duplicates1k.txt")
	benchmarkMergeSort(numerosint, bk)
}
func BenchmarkMergeSortDuplicates10k(bk *testing.B) {
	var numerosint []int = All("duplicates10k.txt")
	benchmarkMergeSort(numerosint, bk)
}
func BenchmarkMergeSortDuplicates100k(bk *testing.B) {
	var numerosint []int = All("duplicates100k.txt")
	benchmarkMergeSort(numerosint, bk)
}
 func BenchmarkMergeSortDuplicates1m(bk *testing.B) {
	var numerosint []int = All("duplicates1m.txt")
	benchmarkMergeSort(numerosint, bk)
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


func BenchmarkMergeSortReversed1k(bk *testing.B) {
	var numerosint []int = All("reversed1k.txt")
	benchmarkMergeSort(numerosint, bk)
}
func BenchmarkMergeSortReversed10k(bk *testing.B) {
	var numerosint []int = All("reversed10k.txt")
	benchmarkMergeSort(numerosint, bk)
}
func BenchmarkMergeSortReversed100k(bk *testing.B) {
	var numerosint []int = All("reversed100k.txt")
	benchmarkMergeSort(numerosint, bk)
}
 func BenchmarkMergeSortReversed1m(bk *testing.B) {
	var numerosint []int = All("reversed1m.txt")
	benchmarkMergeSort(numerosint, bk)
}


func BenchmarkMergeSortSorted1k(bk *testing.B) {
	var numerosint []int = All("sorted1k.txt")
	benchmarkMergeSort(numerosint, bk)
}
func BenchmarkMergeSortSorted10k(bk *testing.B) {
	var numerosint []int = All("sorted10k.txt")
	benchmarkMergeSort(numerosint, bk)
}
func BenchmarkMergeSortSorted100k(bk *testing.B) {
	var numerosint []int = All("sorted100k.txt")
	benchmarkMergeSort(numerosint, bk)
}
 func BenchmarkMergeSortSorted1m(bk *testing.B) {
	var numerosint []int = All("sorted1m.txt")
	benchmarkMergeSort(numerosint, bk)
}



func BenchmarkQuickSortDuplicates1k(bk *testing.B) {
	var numerosint []int = All("duplicates1k.txt")
	benchmarkQuickSort(numerosint, bk)
}
func BenchmarkQuickSortDuplicates10k(bk *testing.B) {
	var numerosint []int = All("duplicates10k.txt")
	benchmarkQuickSort(numerosint, bk)
}
func BenchmarkQuickSortDuplicates100k(bk *testing.B) {
	var numerosint []int = All("duplicates100k.txt")
	benchmarkQuickSort(numerosint, bk)
}
 func BenchmarkQuickSortDuplicates1m(bk *testing.B) {
	var numerosint []int = All("duplicates1m.txt")
	benchmarkQuickSort(numerosint, bk)
}


func BenchmarkQuickSortRandom1k(bk *testing.B) {
	var numerosint []int = All("random1k.txt")
	benchmarkQuickSort(numerosint, bk)
}
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


func BenchmarkQuickSortReversed1k(bk *testing.B) {
	var numerosint []int = All("reversed1k.txt")
	benchmarkQuickSort(numerosint, bk)
}
func BenchmarkQuickSortReversed10k(bk *testing.B) {
	var numerosint []int = All("reversed10k.txt")
	benchmarkQuickSort(numerosint, bk)
}
func BenchmarkQuickSortReversed100k(bk *testing.B) {
	var numerosint []int = All("reversed100k.txt")
	benchmarkQuickSort(numerosint, bk)
}
 func BenchmarkQuickSortReversed1m(bk *testing.B) {
	var numerosint []int = All("reversed1m.txt")
	benchmarkQuickSort(numerosint, bk)
}


func BenchmarkQuickSortSorted1k(bk *testing.B) {
	var numerosint []int = All("sorted1k.txt")
	benchmarkQuickSort(numerosint, bk)
}
func BenchmarkQuickSortSorted10k(bk *testing.B) {
	var numerosint []int = All("sorted10k.txt")
	benchmarkQuickSort(numerosint, bk)
}
func BenchmarkQuickSortSorted100k(bk *testing.B) {
	var numerosint []int = All("sorted100k.txt")
	benchmarkQuickSort(numerosint, bk)
}
 func BenchmarkQuickSortSorted1m(bk *testing.B) {
	var numerosint []int = All("sorted1m.txt")
	benchmarkQuickSort(numerosint, bk)
}

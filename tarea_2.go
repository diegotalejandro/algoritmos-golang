package main

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
		algoutils.Swap(data, i, min)
	}
	return data
}
func InsertionSort(data []int) []int {
  for i := 1; i < len(data); i++ {
  		for j := i; j > 0 && data[j] < data[j - 1]; j-- {
  			algoutils.Swap(data, j, j - 1)
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

    for i, x := range a {
      switch {
      case i < m:
        left = append(left, x)
      case i >= m:
        right = append(right, x)
      }
}
func QuickSort(data []int) []int {

}
func main() {

}

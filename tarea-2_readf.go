package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func ReadF(Directory string) string {
	files, err := ioutil.ReadDir(Directory)
	if err != nil {
		log.Fatal(err)
	}
	var temp string
	for _, file := range files {
		x, _ := ioutil.ReadFile(Directory + "/" + file.Name()) //lee el texto
		temp = string(x)
	}
	return temp
}

/*func dfs( adjunta [][]int, nodes map[int][]int, fn func (int)) {
    dfs_recur(node, map[int]bool{}, fn)
}

func dfs_recur(node int, v map[int]bool, fn func (int)) {
    v[node] = true
    fn(node)
    for _, n := range nodes[node] {
        if _, ok := v[n]; !ok {
            dfs_recur(n, v, fn)
        }
    }
}*/
func main() {
	r := ReadF("/Users/macbookretinadeeugenio/trabajos/texto_sampled")
	arr := strings.Fields(r)
	fmt.Println(arr)
	/*var arrint []int
	for i := 0; i < len(arr); i++ {
		temp, _ := strconv.Atoi(arr[i])
		arrint = append(arrint, temp)
	}*/
	//n := arrint[0]
	/*var adjunta [][]int
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			adjunta[i][j] = append(adjunta, 0)
		}
	}*/

	//fmt.Println(adjunta)
}

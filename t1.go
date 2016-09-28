package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type Blog struct {
	numero string
	genero string
	edad   string
	txt    string
}

// ----------- Lee el nombre y texto de cada archivo de txt y lo mete en un obejto Blog-------------------
func ReadF(Directory string) []Blog {
	files, err := ioutil.ReadDir(Directory)
	if err != nil {
		log.Fatal(err)
	}
	var blogs []Blog
	for _, file := range files {
		split := strings.Split(file.Name(), ".")
		number := split[0]
		gender := split[1]
		age := split[2]
		x, _ := ioutil.ReadFile(Directory + "/" + file.Name())
		temp := Blog{numero: number, genero: gender, edad: age, txt: string(x)}
		blogs = append(blogs, temp)
		//fmt.Println("\n","number:",split[0])
		//fmt.Println(file.Name())
	}
	return blogs
}

type word struct {
	word  string
	count int
}

//Parte 1
func PalabraMasRepetida(blogs []Blog) []word {

	for i := 0; i < len(blogs); i++ {
		aux := blogs[i]
		temptxt := aux.txt
		temptxt = strings.Join(string.Split(temptxt, "."), ", ")
		temptxt = strings.Join(string.Split(temptxt, ","), ", ")
		temptxt = strings.Join(string.Split(temptxt, ":"), ", ")
		for temptxt == "" {
			arrpalabra := string.Split(temptxt, " ")
			palabra := arrpalabra[0]
			contador := strings.Count(temptxt, palabra)

		}
	}
}
func main() {
	var directorio string = "C:/Users/Diego/Desktop/Codigos Go -Atom/sampled"
	var blogs []Blog = ReadF(directorio)
	temp := blogs[0]
	fmt.Println(temp.edad)

	// ----------- Read an entire file -------------------
	//x, _ := ioutil.ReadFile("C:/Users/Diego/Desktop/Codigos Go -Atom/sampled/63420.male.40.txt")
	//lee el txt y lo mete en un string
	//for _, f := range x {
	//fmt.Println(string(f))
	//fmt.Println(string(x))
	//}
	//	---------------------------------------------------
	// file, _ := os.Open("./ej/tweets.txt")
	// scanner := bufio.NewScanner(file)
	// var a []string
	// for scanner.Scan() {
	// 	a = append(a, scanner.Text())
	// }
	// fmt.Println(a[0], "---", a[1])
	//-----------------------------------------------------
}

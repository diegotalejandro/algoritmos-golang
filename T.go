package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strings"
)

//separa por rango de edad, rangos: 13-17,23-27,33-47
func separar(slice []string) ([]string, []string, []string) {
	var rango1 []string
	var rango2 []string
	var rango3 []string
	for i := 0; i < len(slice); i++ {
		comparacion := strings.Split(slice[i], ".")
		if comparacion[2] >= "13" && comparacion[2] <= "17" {
			rango1 = append(rango1, slice[i])
		}
		if comparacion[2] >= "23" && comparacion[2] <= "27" {
			rango2 = append(rango2, slice[i])
		}
		if comparacion[2] >= "33" && comparacion[2] <= "47" {
			rango3 = append(rango3, slice[i])
		}
	}
	return rango1, rango2, rango3
}

// ----------- Lee el nombre y texto de cada archivo de txt y lo mete en un obejto Blog-------------------
func ReadF(Directory string) []string {
	files, err := ioutil.ReadDir(Directory)
	if err != nil {
		log.Fatal(err)
	}
	var blogs []string
	for _, file := range files {
		//x, _ := ioutil.ReadFile(Directory + "/" + file.Name())//lee el texto
		blogs = append(blogs, file.Name())
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
func PalabraMasRepetida(blogs []string, Directory string) []word {
	var PMR []word
	var palabrastxt string
	for i := 0; i < len(blogs); i++ {
		temptxt := blogs[i]
		x, _ := ioutil.ReadFile(Directory + "/" + temptxt)
		temptxt = string(x)
		temptxt = strings.Replace(temptxt, ".", " ", -1)
		temptxt = strings.Replace(temptxt, ",", " ", -1)
		temptxt = strings.Replace(temptxt, ":", " ", -1)
		temptxt = strings.Replace(temptxt, ";", " ", -1)
		temptxt = strings.Replace(temptxt, "!", " ", -1)
		temptxt = strings.Replace(temptxt, "¡", " ", -1)
		temptxt = strings.Replace(temptxt, "?", " ", -1)
		temptxt = strings.Replace(temptxt, "¿", " ", -1)
		temptxt = strings.Replace(temptxt, "(", " ", -1)
		temptxt = strings.Replace(temptxt, ")", " ", -1)
		temptxt = strings.Replace(temptxt, "*", " ", -1)
		temptxt = strings.Replace(temptxt, "-", " ", -1)
		temptxt = strings.Replace(temptxt, "'", "", -1)
		temptxt = strings.Replace(temptxt, "\"", " ", -1)
		//temptxt = strings.Replace(temptxt, "", " ", -1)
		temptxt = strings.ToLower(temptxt)
		a := strings.Fields(temptxt)
		a[0] = " " + a[0]
		a[len(a)-1] = a[len(a)-1] + " "
		temptxt = strings.Join(a, " ")

		//fmt.Println(temptxt)
		for strings.Join(strings.Fields(temptxt), "") != "" {
			arrpalabra := strings.Fields(temptxt) //lo convierte en el arreglo de palabras
			palabra := arrpalabra[0]
			//fmt.Print(len(arrpalabra), "(", i, ")")     //indica la primera palabra
			contador := strings.Count(temptxt, palabra) //cuenta la cantidad de veces q se repite la palabra en el string
			//fmt.Println(palabra, contador)
			tempo := word{word: palabra, count: contador}
			if i > 0 {
				if strings.Index(palabrastxt, " "+palabra+" ") != -1 {
					for j := 0; j < len(PMR); j++ {
						if PMR[j].word == palabra {
							PMR[j].count = PMR[j].count + contador
						}
					}
				} else {
					palabrastxt = palabrastxt + " " + palabra + " "
					PMR = append(PMR, tempo)
				}

			} else {
				palabrastxt = palabrastxt + palabra + " "
				PMR = append(PMR, tempo)
			}
			temptxt = strings.Join(strings.Split(temptxt, " "+palabra+" "), " ")
		}
		fmt.Println("texto numero ", i+1, " listo---------------------------------------------------------------------------------")
		//todoPMR = append(todoPMR, PMR...)
		//PMR = PMR[:0]
	}
	return PMR
}

type ByCount []word

func (a ByCount) Len() int           { return len(a) }
func (a ByCount) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByCount) Less(i, j int) bool { return a[i].count > a[j].count }

func main() {
	var directorio string = "C:/Users/Diego/Desktop/Codigos Go -Atom/sampled"
	var blogs []string = ReadF(directorio)
	var xd1, xd2, xd3 []word
	R1, R2, R3 := separar(blogs)
	xd1 = PalabraMasRepetida(R1, directorio)
	sort.Sort(ByCount(xd1))
	println("Edad entre 13 y 17:")
	for i := 0; i < 20; i++ {
		println("Palabra:    ", xd1[i].word, "           Cantidad:    ", xd1[i].count)
	}
	xd2 = PalabraMasRepetida(R2, directorio)
	sort.Sort(ByCount(xd2))
	println("Edad entre 23 y 27:")
	for i := 0; i < 20; i++ {
		println("Palabra:    ", xd2[i].word, "           Cantidad:    ", xd2[i].count)
	}
	xd3 = PalabraMasRepetida(R3, directorio)
	sort.Sort(ByCount(xd3))
	println("Edad entre 33 y 47:")
	for i := 0; i < 20; i++ {
		println("Palabra:    ", xd3[i].word, "           Cantidad:    ", xd3[i].count)
	}
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

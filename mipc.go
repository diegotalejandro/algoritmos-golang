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

//separa por rango de edad, rangos: 13-17,23-27,33-47
func separar(slice []Blog) ([]Blog, []Blog, []Blog) {
	var rango1 []Blog
	var rango2 []Blog
	var rango3 []Blog
	for i := 0; i < len(slice); i++ {
		if slice[i].edad >= "13" && slice[i].edad <= "17" {
			rango1 = append(rango1, slice[i])
		}
		if slice[i].edad >= "23" && slice[i].edad <= "27" {
			rango2 = append(rango2, slice[i])
		}
		if slice[i].edad >= "33" && slice[i].edad <= "47" {
			rango3 = append(rango3, slice[i])
		}
	}
	return rango1, rango2, rango3
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
	var PMR []word
	var todoPMR []word
	for i := 0; i < len(blogs); i++ {
		aux := blogs[i]
		temptxt := aux.txt
		temptxt = strings.Join(strings.Split(temptxt, "."), " ")
		temptxt = strings.Join(strings.Split(temptxt, ","), " ")
		temptxt = strings.Join(strings.Split(temptxt, ":"), " ")
		temptxt = strings.Join(strings.Split(temptxt, ";"), " ")
		temptxt = strings.Join(strings.Split(temptxt, "!"), " ")
		temptxt = strings.Join(strings.Split(temptxt, "¡"), " ")
		temptxt = strings.Join(strings.Split(temptxt, "?"), " ")
		temptxt = strings.Join(strings.Split(temptxt, "¿"), " ")
		temptxt = strings.Join(strings.Split(temptxt, "("), " ")
		temptxt = strings.Join(strings.Split(temptxt, ")"), " ")
		temptxt = strings.Join(strings.Split(temptxt, "*"), " ")
		temptxt = strings.Join(strings.Split(temptxt, "-"), " ")
		temptxt = strings.Join(strings.Split(temptxt, "'"), "")
		temptxt = strings.Join(strings.Split(temptxt, "\""), "")
		temptxt = strings.ToLower(temptxt)
		//fmt.Println(temptxt)
		for strings.Join(strings.Fields(temptxt), "") != "" {
			arrpalabra := strings.Fields(temptxt) //lo convierte en el arreglo de palabras
			palabra := arrpalabra[0]
			fmt.Println(len(arrpalabra))                //indica la primera palabra
			contador := strings.Count(temptxt, palabra) //cuenta la cantidad de veces q se repite la palabra en el string
			temptxt = strings.Join(strings.Split(temptxt, " "+palabra+" "), " ")
			tempo := word{word: palabra, count: contador}
			if i==0
			{PMR = append(PMR, tempo)}
			else
			{
			for j:=0; j < len(PMR); j++{
				//aun me falta xd
			}
		}
		}
		fmt.Println("texto numero ", i+1, " listo---------------------------------------------------------------------------------")
		//todoPMR = append(todoPMR, PMR...)
		//PMR = PMR[:0]
	}

	return todoPMR
}
func main() {
	var directorio string = "C:/Users/Diego/Desktop/Codigos Go -Atom/sampled"
	var blogs []Blog = ReadF(directorio)
	R1, R2, R3 := separar(blogs)
	fmt.Println(R1[0].edad, R2[0].edad, R3[0].edad)
	xd := PalabraMasRepetida(R1)
	fmt.Println(xd[0].word)

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

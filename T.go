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
	var palabrastxt string
	for i := 0; i < len(blogs); i++ {
		i = 87
		aux := blogs[i]
		temptxt := aux.txt
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
		temptxt = strings.Replace(temptxt, "'", " ", -1)
		temptxt = strings.Replace(temptxt, "\"", " ", -1)
		//temptxt = strings.Replace(temptxt, "", " ", -1)
		temptxt = strings.ToLower(temptxt)
		//fmt.Println(temptxt)
		for strings.Join(strings.Fields(temptxt), "") != "" {
			arrpalabra := strings.Fields(temptxt) //lo convierte en el arreglo de palabras
			palabra := arrpalabra[0]
			fmt.Print(len(arrpalabra), "(", i, ")")     //indica la primera palabra
			contador := strings.Count(temptxt, palabra) //cuenta la cantidad de veces q se repite la palabra en el string
			fmt.Println(palabra, contador)
			tempo := word{word: palabra, count: contador}
			if i > 0 {
				if strings.Index(palabrastxt, " "+palabra+" ") != -1 {
					for j := 0; j < len(PMR); j++ {
						if PMR[j].word == palabra {
							PMR[j].count = PMR[j].count + contador
						}
					}
				} else {
					palabrastxt = palabrastxt + palabra + " "
					PMR = append(PMR, tempo)
				}

			} else {
				palabrastxt = palabrastxt + palabra + " "
				PMR = append(PMR, tempo)
			}
			//if len(arrpalabra) >= 2000 {
			//fmt.Println(temptxt)
			//}
			temptxt = strings.Join(strings.Split(temptxt, " "+palabra+" "), " ")

			if strings.Index(temptxt, " "+palabra+" ") != -1 {
				println("alguno entro")
				temptxt = strings.Join(strings.Split(temptxt, " "+palabra), " ")
				if strings.Index(temptxt, " "+palabra+" ") != -1 {
					temptxt = strings.Join(strings.Split(temptxt, palabra+" "), " ")
					if strings.Index(temptxt, " "+palabra+" ") != -1 {
						temptxt = strings.Join(strings.Split(temptxt, palabra), " ")
					}
				}
			}
			//	if len(arrpalabra) == 1455 {
			//	fmt.Println(temptxt)
			//break
			//	}
		}
		fmt.Println("texto numero ", i+1, " listo---------------------------------------------------------------------------------")
		break
		//todoPMR = append(todoPMR, PMR...)
		//PMR = PMR[:0]
	}
	return PMR
}
func main() {
	var directorio string = "C:/Users/Diego/Desktop/Codigos Go -Atom/sampled"
	var blogs []Blog = ReadF(directorio)
	R1, R2, R3 := separar(blogs)
	fmt.Println(len(R1), len(R2), len(R3))
	xd := PalabraMasRepetida(R1)
	for i := 0; i < 20; i++ {
		fmt.Println(xd[i].count)
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

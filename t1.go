package main
import ("fmt"
"io/ioutil"
"log"
"strings"
)
type Blog struct {
	numero string
  genero string
	edad string
	txt string
}
// ----------- Read file names on a Directory ------------------
func ReadFiles(Directory string,number int ) []string{
	files, err := ioutil.ReadDir(Directory)
	if err != nil {
	 log.Fatal(err)
	}
	var Blogs[] string Blog
	var number string
	var gender string
	var age string
	var txt string
	for _, file := range files {
	split :=strings.Split(file.Name(), ".")
	number= append(data,split[0])
	gender= append(data,split[1])
	age= append(data,split[2])
	x, _ := ioutil.ReadFile(Directory+"/"+file.Name())
	temp :=Blog{numero:number ,genero:gender ,edad:age ,txt:string(x)}
	//fmt.Println("\n","number:",split[0])
	//fmt.Println(file.Name())
 }

 return data
}
func ()  {

}
func main(){
//var number[] string = ReadFiles("C:/Users/Diego/Desktop/Codigos Go -Atom/sampled",0)
//var gender[] string = ReadFiles("C:/Users/Diego/Desktop/Codigos Go -Atom/sampled",1)
//var age[] string = ReadFiles("C:/Users/Diego/Desktop/Codigos Go -Atom/sampled",2)
//var DynamicDirectory[2] string ={"C:/Users/Diego/Desktop/Codigos Go -Atom/sampled/"}
//files, err := ioutil.ReadDir("C:/Users/Diego/Desktop/Codigos Go -Atom/sampled")
//if err != nil {
// log.Fatal(err)
//}
//var i int
//crear lista para objetos de Blog
//for _,file:=range files {
//DynamicDirectory[1]=file.Name()
//  x, y := ioutil.ReadFile(join(DynamicDirectory,""))
//	temp :=Blog{numero:number[i] ,genero:gender[i] ,edad:age[i] ,txt:string(x)}
//	temp=temp <- siguiente
//	i++
//}
//fmt.Println(temp.numero)



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

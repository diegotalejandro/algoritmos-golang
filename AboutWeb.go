package main
import (
  "fmt"
  "bytes"
  "encoding/xml"
  "net/http"
  "log"
  "io/ioutil"
)
type html struct {
	resp resp `xml:"resp"`
}
type resp struct {
	Content string `xml:",innerxml"`
}
func main(){
resp, err := http.Get("http://www.yapo.cl/region_metropolitana/television_camaras?ca=15_s&l=0&q=sony+a6000&w=1&cmn=")
if err != nil {
  log.Fatal(err)
}
defer resp.Body.Close()
//Info de la pagina
body,err:= ioutil.ReadAll(resp.Body)
if err != nil {
  log.Fatal(err)
}
//Pagina en Slice de Bytes
txt:=string(body)//Slice de Bytes a texto
b:=[]byte(txt)//texto a Slice Bytes,pero sin necesidad de conversion
h := html{}
er := xml.NewDecoder(bytes.NewBuffer(b)).Decode(&h)//De Slice de Bytes a Body en html
if er != nil {
  fmt.Println("error", er)
  return
}

fmt.Println(h)
}

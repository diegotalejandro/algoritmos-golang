package main
import(
  "bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"sync")

func SCC(g grafo) [][]ID{//SCC: Strongly Connected Components
	d:=newSCC()
	for v:= range g.GetNodes(){
		if _, ok:= d.index[v];c !ok{
			scc(g,v,d)
		}
	}
	return d.result
}

type sccData struct{

}


type node struct{
	id string
}
var nodeCnt uint64
func NewNode(id string) Node {
	return &node{
		id: id,
	}
}

func (n *node) ID() ID{
	return StringID(n.id)
}

func (n *node) String() string{
	return n.id
}

type grafo struct{
idparanodos map[ID]Node
nodopararecursos map[ID]map[ID]float64
nodoparaobjetivos map[ID]map[ID]float64
}
func newGrafo() *grafo{
	return &grafo{
		idparanodos: make(map[ID]Node),
		nodopararecursos: make(map[ID]map[ID]float64),
		nodoparaobjetivos: make(map[ID]map[ID]float64),
	}
}

func NewGrafo() grafo{
	return newGrafo()
}

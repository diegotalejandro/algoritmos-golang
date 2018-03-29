type List interface {
    // Add agrega val en la posición i de la lista
    func Add(i, val int){
    aux= &node
    aux = head
    for n:=0; n < i; n++ {
    if aux.next == nil{
    return -1
    }
    aux=aux.next
  }
    temp=&node{value: val}
    temp.next = aux.next
    aux.next =temp
  }
    // PushFront agrega val al inicio de la lista
    func PushFront(val int)

    // PushBack agrega val al final de la lista
    func PushBack(val int)

    // Get obtiene la posición i
    func Get(i int) int

    // First obtiene el primer elemento
    func First() int

    // Last obtiene el último elemento
    func Last() int

    // Remove elimina la posición i de la lista
    func Remove(i int)

    // RemoveLast elimina el último elemento de la lista
    func RemoveLast()

    // RemoveFirst elimina el primer elemento de la lista
    func RemoveFirst()

    // Reverse invierte la lista
    func Reverse()
}

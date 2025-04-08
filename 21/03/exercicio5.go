package main
import "fmt"
// Estrutura do nó (tem ponteiro para o próximo e o anterior)
type Node struct {
data int
next *Node
prev *Node
}
// Estrutura da lista duplamente encadeada
type DoublyLinkedList struct {
head *Node
tail *Node
}
// Adiciona um elemento no final da lista
func (l *DoublyLinkedList) Append(data int) {
newNode := &Node{data: data}
if l.tail == nil {
l.head = newNode
l.tail = newNode
return
}
l.tail.next = newNode
newNode.prev = l.tail
l.tail = newNode
}
// Remove o último elemento da lista
func (l *DoublyLinkedList) RemoveLast() {
if l.tail == nil {
return
}
if l.tail == l.head { // Apenas um elemento na lista
l.head = nil
l.tail = nil
return
}
l.tail = l.tail.prev
l.tail.next = nil
}
// Exibe os elementos da lista do início ao fim
func (l *DoublyLinkedList) PrintForward() {
current := l.head
for current != nil {
fmt.Print(current.data, " <-> ")
current = current.next
}
fmt.Println("nil")
}
// Exibe os elementos da lista do fim para o início
func (l *DoublyLinkedList) PrintBackward() {
current := l.tail
for current != nil {
fmt.Print(current.data, " <-> ")
current = current.prev
}
fmt.Println("nil")
}
func main() {
list := DoublyLinkedList{}
list.Append(1)
list.Append(2)
list.Append(3)
fmt.Println("Lista do início para o fim:")
list.PrintForward()
fmt.Println("\nLista do fim para o início:")
list.PrintBackward()
list.RemoveLast()
fmt.Println("\nApós remover o último elemento:")
list.PrintForward()
}
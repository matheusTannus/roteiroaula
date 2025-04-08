package main
import "fmt"
// Estrutura do nó (cada tarefa)
type Task struct {
	name string
	next *Task
	
/continue a implementação/
}


// Estrutura da lista encadeada (lista de tarefas)
type TaskList struct {
	head *Task
	
/continue a implementação/
}


// Método para adicionar uma nova tarefa à lista (Append)
func (t *TaskList) AddTask(name string) {
	newTask := &Task{name: name}
	if t.head == nil{
		t.head = newTask
		return
	}
	current := t.head
	for current.next != nil{
		current = current.next
	}
	current.next = newTask
		
/continue a implementação/
}


// Método para remover a primeira tarefa (RemoveFirst)
func (t *TaskList) CompleteTask() {
	if t.head != nil{
		t.head = t.head.next
	}
	
/continue a implementação/
}


// Método para exibir todas as tarefas pendentes
func (t *TaskList) ShowTasks() {
	current:= t.head
	for current != nil{
		fmt.Println(current.name)
		current = current.next
	}
	fmt.Println("-----------------")
/continue a implementação/
}


func main() {
lista := TaskList{}
// Adicionando tarefas
lista.AddTask("Estudar Go")
lista.AddTask("Fazer compras")
lista.AddTask("Academia")
lista.ShowTasks() // Exibe as tarefas pendentes
lista.CompleteTask() // Remove a primeira tarefa concluída
lista.ShowTasks() // Exibe as tarefas restantes
lista.CompleteTask()
lista.ShowTasks()
}
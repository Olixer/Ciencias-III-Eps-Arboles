package main

import (
	"fmt"
)



type Node struct {
  	PrimerNombre string
	SegundoNombre string
	PrimerApellido string
	SegundoApellido string
	TipoIdentificacion string
	NumeroIdentificacion string
	Sintomas string
	HoraLlegada string
	eps string
}

func (n *Node) String() string {
	return fmt.Sprint("Paciente de nombre ",n.PrimerNombre, " ", n.SegundoNombre, " ", n.PrimerApellido, " ", n.SegundoApellido, " con identificacion ", n.TipoIdentificacion, " ", n.NumeroIdentificacion, " ,posee sintomas como ", n.Sintomas, ",su hora de llegada fue ", n.HoraLlegada, " y viene de parte de la EPS ", n.eps)
}

// NewQueue returns a new queue with the given initial size.

func NewQueue(size int) *Queue {
	return &Queue{
		nodes: make([]*Node, size),
		size:  size,
	}
}

// Queue is a basic FIFO queue based on a circular list that resizes as needed.


type Queue struct {
	nodes []*Node
	size  int
	head  int
	tail  int
	count int
}

// Push adds a node to the queue.


func (q *Queue) Push(n *Node) {
	if q.head == q.tail && q.count > 0 {
		nodes := make([]*Node, len(q.nodes)+q.size)
		copy(nodes, q.nodes[q.head:])
		copy(nodes[len(q.nodes)-q.head:], q.nodes[:q.head])
		q.head = 0
		q.tail = len(q.nodes)
		q.nodes = nodes
	}
	q.nodes[q.tail] = n
	q.tail = (q.tail + 1) % len(q.nodes)
	q.count++
}

// Pop removes and returns a node from the queue in first to last order.
func (q *Queue) Pop() *Node {
	if q.count == 0 {
		return nil
	}
	node := q.nodes[q.head]
	q.head = (q.head + 1) % len(q.nodes)
	q.count--
	return node
}

func main() {
	
	var pacientes int;
	var primerNombre string
	var segundoNombre string
	var primerApellido string
	var segundoApellido string
	var tipoIdentificacion string
	var numeroIdentificacion string
	var sintomas string
	var horaLlegada string
	var eps string
	var epsSanitas int
	var epsFamisanar int


	fmt.Println("Sistema Medico\n\n")
	fmt.Println("Listado EPS\n")
	fmt.Println("1. Famisanar\n")
	epsFamisanar=0;
	fmt.Println("2. Sanitas EPS\n")
	epsSanitas=0;

	q := NewQueue(1)

	fmt.Println("Cantidad de pacientes a registrar\n\n")
	fmt.Scanf("%v/n", &pacientes)
	fmt.Scanf("%v/n", &pacientes)


	fmt.Println("Registro de pacientes\n\n")

	for i:=0; i<pacientes; i++{
		fmt.Println("\nPaciente #: ",i+1,"\n")
		fmt.Println("\nPrimer Nombre: ")
		fmt.Scanf("%v\n",&primerNombre)
		fmt.Println("\nSegundo Nombre: ")
		fmt.Scanf("%v\n",&segundoNombre)
		fmt.Println("\nPrimer Apellido: ")
		fmt.Scanf("%v\n",&primerApellido)
		fmt.Println("\nSegundo Apellido: ")
		fmt.Scanf("%v\n",&segundoApellido)
		fmt.Println("\nTipo Identificacion: ")
		fmt.Scanf("%v\n",&tipoIdentificacion)
		fmt.Println("\nNumero Identificacion: ")
		fmt.Scanf("%v\n",&numeroIdentificacion)
		fmt.Println("\nSintomas: ")
		fmt.Scanf("%v\n",&sintomas)
		fmt.Println("\nHora de Llegada: ")
		fmt.Scanf("%v\n",&horaLlegada)
		fmt.Println("\neps: ")
		fmt.Scanf("%v\n",&eps)
		if(eps=="famisanar"){
			epsFamisanar++;
		}else if(eps=="sanitas"){
			epsSanitas++;
		}else{
			fmt.Println("\nEps sin registrar ")
			eps=eps+" que no se encuentra afiliada"
		}

	q.Push(&Node{primerNombre,segundoNombre,primerApellido,segundoApellido,tipoIdentificacion,numeroIdentificacion,sintomas,horaLlegada,eps})

	}
	fmt.Println("\nHistoria clinica\n")
	for i:=0; i<pacientes; i++{
	
		fmt.Println(q.Pop())
	}

	fmt.Println("\nPacientes Atendidos de Famisanar: ",epsFamisanar,"\n")
	fmt.Println("\nPacientes Atendidos de Sanitas: ",epsSanitas,"\n")
}

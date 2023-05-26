package main

import "fmt"

func main() {
	search(NewGraph())
}
func search(graph Graph) bool {
	var q Queue
	q.EnqueFromSlice(graph["you"])
	searched := make(map[string]bool)
	for !q.IsEmpty() {
		person, _ := q.Dequeue()

		if person_is_seller(person) {
			fmt.Println("Person %v is a seller!", person)
			return true
		} else {
			q.EnqueFromSlice(graph[person])
			searched[person] = true
		}
	}
	return false
}

type Graph map[string][]string

func NewGraph() Graph {
	g := make(Graph)
	g["you"] = []string{"alice", "bob", "claire"}
	g["bob"] = []string{"anuj", "peggy"}
	g["alice"] = []string{"peggy"}
	g["claire"] = []string{"thom", "jonny"}
	g["anuj"] = []string{}
	g["peggy"] = []string{}
	g["thom"] = []string{}
	g["thom"] = []string{}
	return g
}

func person_is_seller(person string) bool {
	if person[len(person)-1:] == "m" {
		return true
	}
	return false
}

// Que.
type Queue []string

// Adds an element to the end of the queue
func (q *Queue) Enqueue(i string) {
	*q = append(*q, i)
}

// Removes an element from the start of the queue
func (q *Queue) Dequeue() (string, bool) {
	if q.IsEmpty() {
		return "", false
	}
	element := (*q)[0]
	*q = (*q)[1:]
	return element, true
}

// Checks if the queue is empty
func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
func (q *Queue) EnqueFromSlice(sl []string) {
	for _, el := range sl {
		q.Enqueue(el)
	}
}

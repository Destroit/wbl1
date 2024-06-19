package main

import "fmt"

// Parent
type Human struct {
	Name   string
	status string
}

// Creating methods for parent
func (h *Human) Run() {
	h.status = "running"
}

func (h *Human) Sit() {
	h.status = "sitting"
}

func (h *Human) Sleep() {
	h.status = "sleeping"
}

func (h *Human) GetStatus() string {
	return h.status
}

// Child
type Action struct {
	// Embedding parent struct
	Human
}

func main() {
	var a Action
	// We can access Human's methods directly
	a.Run()
	fmt.Println(a.GetStatus())

	// Or by including embedded type's name
	a.Human.Sit()
	fmt.Println(a.Human.GetStatus())

	a.Human.Sleep()
	fmt.Println(a.GetStatus())
}

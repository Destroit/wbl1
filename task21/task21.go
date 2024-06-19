package main

import "fmt"

// Adaptee structs with different methods

// Notebook doesn't need adapter
type Notebook struct{}

func (n *Notebook) Write() {
	fmt.Println("Writing with a pen")
}

type Pc struct{}

func (pc *Pc) HitKeys(isTyping bool) {
	if isTyping {
		fmt.Println("Typing code")
	}
}

type Cave struct{}

func (cave *Cave) DrawOnWall() {
	fmt.Println("Drawing mammoth")
}

// Target interface
type WriteAdapter interface {
	Write()
}

// Concrete adapters for adaptee structs
type PcAdapter struct {
	*Pc
}

type CaveAdapter struct {
	*Cave
}

// Adapters methods realisation
func (adapter *PcAdapter) Write() {
	adapter.HitKeys(true)
}

func (adapter *CaveAdapter) Write() {
	adapter.DrawOnWall()
}

// Adapters constructors
func NewPcAdapter(pc *Pc) WriteAdapter {
	return &PcAdapter{pc}
}

func NewCaveAdapter(cave *Cave) WriteAdapter {
	return &CaveAdapter{cave}
}

func main() {
	s := [3]WriteAdapter{NewCaveAdapter(&Cave{}), NewPcAdapter(&Pc{}), &Notebook{}}
	for _, v := range s {
		v.Write()
	}
}

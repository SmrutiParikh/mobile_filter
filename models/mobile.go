package models

import "fmt"

// Mobile is ...
type Mobile struct {
	Model  string
	Brand  Brand
	Memory int
}

//Equals ...
func (current Mobile) Equals(that Mobile) bool {
	return current.Model == that.Model && current.Brand == that.Brand && current.Memory == that.Memory
}

//Print ...
func (current Mobile) Print() {
	fmt.Printf("\t %s\t| by %s\t| %d GB\n", current.Model, current.Brand, current.Memory)
}

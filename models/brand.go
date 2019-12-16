package models

// Brand is ...
type Brand int

func (b Brand) String() string {
	return brands[b]
}

// Parse is ...
func Parse(value string) Brand {
	var b Brand
	for index, element := range brands {
		if value == element {
			b = Brand(index)
		}
	}
	return b
}

//  ...
const (
	VIVO Brand = iota
	APPLE
	REALME
)

var brands = [...]string{
	"VIVO",
	"APPLE",
	"REALME",
}

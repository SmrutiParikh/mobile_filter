package models

//FilterType is ...
type FilterType int

//...
const (
	MODEL FilterType = iota
	BRAND
	MEMORY
)

var filters = [...]string{
	"MODEL",
	"BRAND",
	"MEMORY",
}

func (f FilterType) String() string {
	return filters[f]
}

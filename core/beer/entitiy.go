package beer

type Beer struct {
	ID    int64     `json: "id"`
	Name  string    `json: "name"`
	Type  BeerType  `json:"type"`
	Style BeerStyle `json:"style"`
}

type BeerType int

const (
	TypeAle   = 1
	TypeLager = 2
	TypemMalt = 3
	TypeStout = 4
)

func (t BeerType) String() string {
	switch t {
	case TypeAle:
		return "Ale"
	case TypeLager:
		return "Lager"
	case TypeStout:
		return "Stout"
	case TypemMalt:
		return "Malt"
	}
	return "Unknown"
}

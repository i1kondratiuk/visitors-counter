package value

type VisitType int

const (
	URI VisitType = iota
)

// Visit represents a visit
type Visit struct {
	Type    VisitType
	Value   string
}

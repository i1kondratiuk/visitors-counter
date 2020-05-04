package value

type VisitType string

const (
	ResourcePath VisitType = "Resource Path"
)

// Visit represents a visit
type Visit struct {
	Type  VisitType
	Value string
}

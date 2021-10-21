package value

// ValueObject
// Interface that all value objects should implement
type ValueObject interface {
	Value() interface{}
	String() string
}

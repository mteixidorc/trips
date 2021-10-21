package value

import "strconv"

// ValueObjectID
// A wrapper for an int64 type, it's esencially an indentification 
type ValueObjectID struct {
	value int64
}

func NewValueObjectID(value int64) ValueObjectID {
	return ValueObjectID{value}
}

func (vo ValueObjectID) Value() interface{} {
	return vo.value
}

func (vo ValueObjectID) String() string {
	return strconv.FormatInt(vo.value, 10)
}

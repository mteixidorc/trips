package value

import "github.com/google/uuid"

// ValueObjectUniqueID
// A wrapper for an uuid type, it's esencially a unique identification generator
type ValueObjectUniqueID struct {
	id uuid.UUID
}

func NewValueObjectUniqueID() ValueObjectUniqueID {
	return ValueObjectUniqueID{
		uuid.New(),
	}
}

func NewValueObjectUniqueIDFromString(id string) (*ValueObjectUniqueID, error) {
	uuidFromString, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	return &ValueObjectUniqueID{
		uuidFromString,
	}, nil
}

func NewValueObjectUniqueIDFromStringMust(id string) ValueObjectUniqueID {
	uuidFromString, _ := NewValueObjectUniqueIDFromString(id)
	return *uuidFromString
}

func (valueObject ValueObjectUniqueID) Value() interface{} {
	return valueObject.id
}

func (valueObject ValueObjectUniqueID) String() string {
	return valueObject.id.String()
}

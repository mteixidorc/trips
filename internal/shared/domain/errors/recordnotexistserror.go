package errors

import (
	"fmt"

	"github.com/mteixidorc/trips/internal/shared/domain/value"
)

// recordNotExistsError
// shared error type usefull for all bounded contexts
type recordNotExistsError struct {
	recordId value.ValueObject
}

func NewRecordNotExistsError(id value.ValueObject) recordNotExistsError {
	return recordNotExistsError{id}
}

func (err recordNotExistsError) Error() string {
	return fmt.Sprintf("record with id %s not exists", err.recordId.String())
}

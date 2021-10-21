package value_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/mteixidorc/trips/internal/shared/domain/value"
)

func TestValueObjectUniqueIDMustFail(t *testing.T) {
	_, err := value.NewValueObjectUniqueIDFromString("not a valid uuid")
	if err == nil {
		t.Fail()
	}
}

func TestValueObjectUniqueIDMustWork(t *testing.T) {
	id, err := value.NewValueObjectUniqueIDFromString("f4724b93-fcd2-4a3b-922a-71ed432f7edf")
	if err != nil {
		t.FailNow()
	}

	_, ok := id.Value().(uuid.UUID)
	if !ok {
		t.Fail()
	}
}

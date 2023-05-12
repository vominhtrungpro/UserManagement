package uid

import (
	"github.com/google/uuid"
	pkgerrors "github.com/pkg/errors"
)

// Generate generates a new UUID.
func Generate() (string, error) {
	uid, err := uuid.NewUUID()
	if err != nil {
		return "", pkgerrors.WithStack(err)
	}

	return uid.String(), nil
}

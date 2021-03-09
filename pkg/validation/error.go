package validation

import (
	"errors"

	"github.com/xh3b4sd/tracer"
)

var invalidMetadataError = &tracer.Error{
	Kind: "invalidMetadataError",
}

func IsInvalidMetadata(err error) bool {
	return errors.Is(err, invalidMetadataError)
}

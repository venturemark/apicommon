package validation

import (
	"github.com/xh3b4sd/tracer"

	"github.com/venturemark/apicommon/pkg/metadata"
)

func Audience(m map[string]string) error {
	{
		_, ok := m[metadata.AudienceID]
		if !ok {
			return tracer.Maskf(invalidMetadataError, "%s must not be empty", metadata.AudienceID)
		}
	}

	return nil
}

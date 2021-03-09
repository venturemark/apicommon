package validation

import (
	"github.com/xh3b4sd/tracer"

	"github.com/venturemark/apicommon/pkg/metadata"
)

func Timeline(m map[string]string) error {
	{
		_, ok := m[metadata.TimelineID]
		if !ok {
			return tracer.Maskf(invalidMetadataError, "%s must not be empty", metadata.TimelineID)
		}
	}

	{
		_, ok := m[metadata.VentureID]
		if !ok {
			return tracer.Maskf(invalidMetadataError, "%s must not be empty", metadata.VentureID)
		}
	}

	return nil
}

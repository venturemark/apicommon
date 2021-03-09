package validation

import (
	"github.com/xh3b4sd/tracer"

	"github.com/venturemark/apicommon/pkg/metadata"
)

func Message(m map[string]string) error {
	{
		_, ok := m[metadata.MessageID]
		if !ok {
			return tracer.Maskf(invalidMetadataError, "%s must not be empty", metadata.MessageID)
		}
	}

	{
		_, ok := m[metadata.TimelineID]
		if !ok {
			return tracer.Maskf(invalidMetadataError, "%s must not be empty", metadata.TimelineID)
		}
	}

	{
		_, ok := m[metadata.UpdateID]
		if !ok {
			return tracer.Maskf(invalidMetadataError, "%s must not be empty", metadata.UpdateID)
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

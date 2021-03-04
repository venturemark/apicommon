package resource

import (
	"crypto/sha256"
	"fmt"

	"github.com/venturemark/apicommon/pkg/key"
	"github.com/venturemark/apicommon/pkg/metadata"
)

func Hash(m map[string]string) string {
	var h string

	if isVenture(m) {
		h = hashVenture(m)
	}

	return h
}

func hash(s string) string {
	h := sha256.New()

	_, err := h.Write([]byte(s))
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("%x", h.Sum(nil))
}

func hashVenture(m map[string]string) string {
	var vid string
	{
		vid = m[metadata.VentureID]
	}

	return hash(fmt.Sprintf(key.Venture, vid))
}

func isVenture(m map[string]string) bool {
	// If there is no venture ID provided we do not expect to create a resource
	// hash for a venture.
	{
		_, ok := m[metadata.VentureID]
		if !ok {
			return false
		}
	}

	// If there is a timeline ID provided we do not expect to create a resource
	// hash for a venture, because in this case we would expect to create a
	// resource hash for a timeline.
	{
		_, ok := m[metadata.TimelineID]
		if ok {
			return false
		}
	}

	return true
}

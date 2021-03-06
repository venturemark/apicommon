package hash

import (
	"crypto/sha256"
	"fmt"

	"github.com/venturemark/apicommon/pkg/key"
	"github.com/venturemark/apicommon/pkg/metadata"
)

func Audience(m map[string]string) string {
	var aui string
	{
		aui = m[metadata.AudienceID]
	}

	return hash(fmt.Sprintf(key.Audience, aui))
}

func Message(m map[string]string) string {
	var tii string
	{
		tii = m[metadata.TimelineID]
	}

	var upi string
	{
		upi = m[metadata.UpdateID]
	}

	var vei string
	{
		vei = m[metadata.VentureID]
	}

	return hash(fmt.Sprintf(key.Message, vei, tii, upi))
}

func Timeline(m map[string]string) string {
	var vei string
	{
		vei = m[metadata.VentureID]
	}

	return hash(fmt.Sprintf(key.Timeline, vei))
}

func Update(m map[string]string) string {
	var tii string
	{
		tii = m[metadata.TimelineID]
	}

	var vei string
	{
		vei = m[metadata.VentureID]
	}

	return hash(fmt.Sprintf(key.Update, vei, tii))
}

func Venture(m map[string]string) string {
	var vei string
	{
		vei = m[metadata.VentureID]
	}

	return hash(fmt.Sprintf(key.Venture, vei))
}

func hash(s string) string {
	h := sha256.New()

	_, err := h.Write([]byte(s))
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("%x", h.Sum(nil))
}

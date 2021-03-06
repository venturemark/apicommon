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

	return hash(key.Audience, aui)
}

func Message(m map[string]string) string {
	var mei string
	{
		mei = m[metadata.MessageID]
	}

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

	return hash(key.Message, vei, tii, upi, mei)
}

func Timeline(m map[string]string) string {
	var tii string
	{
		tii = m[metadata.TimelineID]
	}

	var vei string
	{
		vei = m[metadata.VentureID]
	}

	return hash(key.Timeline, vei, tii)
}

func Update(m map[string]string) string {
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

	return hash(key.Update, vei, tii, upi)
}

func Venture(m map[string]string) string {
	var vei string
	{
		vei = m[metadata.VentureID]
	}

	return hash(key.Venture, vei)
}

func hash(f string, v ...interface{}) string {
	h := sha256.New()

	_, err := h.Write([]byte(fmt.Sprintf(f+":%s", v...)))
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("%x", h.Sum(nil))
}

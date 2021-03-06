package hash

import (
	"crypto/sha256"
	"fmt"

	"github.com/venturemark/apicommon/pkg/key"
	"github.com/venturemark/apicommon/pkg/metadata"
)

func Message(m map[string]string) string {
	var tid string
	{
		tid = m[metadata.TimelineID]
	}

	var uid string
	{
		uid = m[metadata.UpdateID]
	}

	var vid string
	{
		vid = m[metadata.VentureID]
	}

	return hash(fmt.Sprintf(key.Message, vid, tid, uid))
}

func Timeline(m map[string]string) string {
	var vid string
	{
		vid = m[metadata.VentureID]
	}

	return hash(fmt.Sprintf(key.Timeline, vid))
}

func Update(m map[string]string) string {
	var tid string
	{
		tid = m[metadata.TimelineID]
	}

	var vid string
	{
		vid = m[metadata.VentureID]
	}

	return hash(fmt.Sprintf(key.Update, vid, tid))
}

func Venture(m map[string]string) string {
	var vid string
	{
		vid = m[metadata.VentureID]
	}

	return hash(fmt.Sprintf(key.Venture, vid))
}

func hash(s string) string {
	h := sha256.New()

	_, err := h.Write([]byte(s))
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("%x", h.Sum(nil))
}

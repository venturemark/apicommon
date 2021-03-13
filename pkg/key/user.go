package key

import (
	"fmt"
	"strconv"

	"github.com/venturemark/apicommon/pkg/metadata"
)

// User is the storage key for mapping user IDs. A user can have multiple
// associations with multiple resources.
func User(m map[string]string) *Key {
	var rei string
	{
		switch m[metadata.ResourceKind] {
		case "audience":
			rei = Audience(m).Elem()
		case "message":
			rei = Message(m).Elem()
		case "timeline":
			rei = Timeline(m).Elem()
		case "update":
			rei = Update(m).Elem()
		case "venture":
			rei = Venture(m).Elem()
		}
	}

	var usi string
	{
		usi = m[metadata.UserID]
	}

	var k *Key
	{
		var id *ID
		if usi != "" {
			f, err := strconv.ParseFloat(usi, 64)
			if err != nil {
				panic(err)
			}
			s := usi

			id = &ID{
				f: f,
				s: s,
			}
		}

		var ele string
		{
			ele = fmt.Sprintf("res:%s:use", hash("res:%s:use:%s", rei, usi))
		}

		var lis string
		{
			lis = fmt.Sprintf("res:%s:use", hash("res:%s:use", rei))
		}

		k = &Key{
			id:  id,
			ele: ele,
			lis: lis,
		}
	}

	return k
}

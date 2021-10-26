package schema

type Update struct {
	Obj UpdateObj `json:"obj"`
}

type UpdateObj struct {
	Metadata map[string]string `json:"metadata"`
	Property UpdateObjProperty `json:"property"`
}

type UpdateObjProperty struct {
	Attachments []UpdateObjPropertyAttachment `json:"attachments"`
	Head        string                        `json:"head"`
	Text        string                        `json:"text"`
}

type UpdateObjPropertyAttachment struct {
	Addr string `json:"addr"`
	Type string `json:"type"`
}

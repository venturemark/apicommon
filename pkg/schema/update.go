package schema

type Update struct {
	Obj UpdateObj `json:"obj"`
}

type UpdateObj struct {
	Metadata map[string]string `json:"metadata"`
	Property UpdateObjProperty `json:"property"`
}

type UpdateObjProperty struct {
	Head string                  `json:"head"`
	Link []UpdateObjPropertyLink `json:"link"`
	Text string                  `json:"text"`
}

type UpdateObjPropertyLink struct {
	Addr string `json:"addr"`
	Text string `json:"text"`
}

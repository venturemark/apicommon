package schema

type Venture struct {
	Obj VentureObj `json:"obj"`
}

type VentureObj struct {
	Metadata map[string]string  `json:"metadata"`
	Property VentureObjProperty `json:"property"`
}

type VentureObjProperty struct {
	Desc string                   `json:"desc"`
	Link []VentureObjPropertyLink `json:"link"`
	Name string                   `json:"name"`
}

type VentureObjPropertyLink struct {
	Addr string `json:"addr"`
	Text string `json:"text"`
}

package schema

type Invite struct {
	Obj InviteObj `json:"obj"`
}

type InviteObj struct {
	Metadata map[string]string `json:"metadata"`
	Property InviteObjProperty `json:"property"`
}

type InviteObjProperty struct {
	Mail string `json:"mail"`
	Stat string `json:"stat"`
}

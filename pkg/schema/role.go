package schema

type Role struct {
	Obj RoleObj `json:"obj"`
}

type RoleObj struct {
	Metadata map[string]string `json:"metadata"`
}

package schema

type Role struct {
	Obj RoleObj `json:"obj"`
}

type RoleObj struct {
	Metadata map[string]string `json:"metadata"`
	Property RoleObjProperty   `json:"property"`
}

type RoleObjProperty struct {
	Kin string `json:"kin"`
	Res string `json:"res"`
	Sub string `json:"sub"`
}

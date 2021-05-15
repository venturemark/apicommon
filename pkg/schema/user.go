package schema

type User struct {
	Obj UserObj `json:"obj"`
}

type UserObj struct {
	Metadata map[string]string `json:"metadata"`
	Property UserObjProperty   `json:"property"`
}

type UserObjProperty struct {
	Desc string                `json:"desc"`
	Mail string                `json:"mail"`
	Name string                `json:"name"`
	Prof []UserObjPropertyProf `json:"prof"`
}

type UserObjPropertyProf struct {
	Desc string `json:"desc"`
	Vent string `json:"vent"`
}

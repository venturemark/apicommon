package key

const (
	// Audience is the storage key for mapping audience IDs. An organization
	// owns an audience, while a user can be part of multiple organizations as
	// well as multiple audiences.
	Audience = "org:%s:aud"
)

const (
	// Timeline is the storage key for mapping timeline IDs. An organization can
	// own multiple timelines.
	Timeline = "org:%s:tml"
)

const (
	// Update is the storage key for mapping update IDs. An audience can have
	// multiple updates.
	Update = "org:%s:tml:%s:upd"
)

const (
	// Message is the storage key for mapping message IDs. An update can have
	// multiple messages.
	Message = "org:%s:tml:%s:upd:%s:mes"
)

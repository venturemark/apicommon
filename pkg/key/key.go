package key

const (
	// Audience is the storage key for mapping audience IDs. An organization
	// owns an audience, while a user can be part of multiple organizations as
	// well as multiple audiences.
	Audience = "org:%s:aud"
)

const (
	// Message is the storage key for mapping message IDs. An update can have
	// multiple messages.
	Message = "org:%s:tml:%s:upd:%s:mes"
)

const (
	// Role is the storage key for mapping role IDs. An user can have multiple
	// roles. The inserted resource key is the filled, hashed version of any of
	// the other resource keys defined here, e.g. sha265("org:al9qy:tml") for
	// records of timeline roles.
	Role = "res:%s:rol"
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

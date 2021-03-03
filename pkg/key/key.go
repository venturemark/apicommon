package key

const (
	// Audience is the storage key for mapping audience IDs. A venture can have
	// multiple audiences, while a user can be part of multiple ventures as well
	// as multiple audiences.
	Audience = "ven:%s:aud"
)

const (
	// Message is the storage key for mapping message IDs. A message can have
	// multiple replies, which are just messages as well.
	Message = "ven:%s:tml:%s:upd:%s:mes"
)

const (
	// Role is the storage key for mapping role IDs. A user can have multiple
	// roles associated to resources. The inserted resource key is the filled,
	// hashed version of any of the other resource keys defined here, e.g.
	// sha265("ven:al9qy:tml") for records of timeline roles.
	Role = "res:%s:rol"
)

const (
	// Timeline is the storage key for mapping timeline IDs. A timeline can have
	// multiple updates.
	Timeline = "ven:%s:tml"
)

const (
	// Update is the storage key for mapping update IDs. An update can have
	// multiple messages.
	Update = "ven:%s:tml:%s:upd"
)

const (
	// Venture is the storage key for mapping venture IDs. A venture can have
	// multiple timelines.
	Venture = "ven:%s"
)

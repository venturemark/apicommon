package key

const (
	AudienceOwner = "org:%s:aud:own"
	// AudienceResource is the storage key for mapping audience IDs. An
	// organization owns an audience, while a user can be part of multiple
	// organizations as well as multiple audiences.
	AudienceResource = "org:%s:aud:res"
)

const (
	TimelineOwner = "org:%s:tml:own"
	// TimelineResource is the storage key for mapping timeline IDs. An
	// organization can own multiple timelines.
	TimelineResource = "org:%s:tml:res"
)

const (
	UpdateOwner = "org:%s:tml:%s:upd:own"
	// UpdateResource is the storage key for mapping update IDs. An audience can
	// have multiple updates.
	UpdateResource = "org:%s:tml:%s:upd:res"
)

const (
	MessageOwner = "org:%s:tml:%s:upd:%s:mes:own"
	// MessageResource is the storage key for mapping message IDs. An update can
	// have multiple messages.
	MessageResource = "org:%s:tml:%s:upd:%s:mes:res"
)

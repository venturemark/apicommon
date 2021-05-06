package metadata

const (
	// PermissionModel indicates which role kinds to prefer for resource role
	// creation. E.g. defining the visibility of a resource as "private" can be
	// persisted using this metadata key, so that roles can be created according
	// to the permission model.
	PermissionModel = "permission.venturemark.co/model"
)

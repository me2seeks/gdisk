package jobtype

// RemoveDeletedObject remove deleted object
type RemoveDeletedObject struct {
	Sn string
}

// RemoveDeletedObjectPayload pay success notify user
type RemoveDeletedObjectPayload struct {
	Identity string
}

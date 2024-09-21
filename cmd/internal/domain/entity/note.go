package domain

type Note struct {
	ID int
	// Profile does not refer to any other tables, it's not a user
	// nor an id, instead, it is just a simple, generic name
	// users must provide to separate or label the note as theirs.
	//
	// Currently, the front-end is always setting this value to "root",
	// but this behavior may change in the future if new users want to
	// have access to this feature.
	Profile      string
	Name         string
	Content      string
	LastModified int64
	// Used for soft-deletion.
	Active bool
}

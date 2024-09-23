package domain

type Note struct {
	ID int
	// Profile does not refer to any other tables, it's not a user
	// nor an id, instead, it is just a simple, generic name
	// users must provide to separate or label the note as theirs.
	//
	// This value should never be returned in any API endpoints
	// since it is hashed on insertion.
	Profile      string
	Name         string
	Content      string
	CreatedAt    int64
	LastModified int64
	// Used for soft-deletion.
	Active bool
}

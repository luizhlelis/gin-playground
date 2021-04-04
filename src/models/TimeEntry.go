package models

// TimeEntry entry that register employee worked hors
type TimeEntry struct {
	User     string `form:"user" json:"user" xml:"user"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

// NewTimeEntry returns a pointer to a new TimeEntry struct
func NewTimeEntry(user string, password string) (*TimeEntry, error) {

	timeEntry := &TimeEntry{
		User:     user,
		Password:  password,
	}

	return timeEntry, nil
}

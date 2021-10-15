package sendinglist

// SendingList contains slice of users to share data with
type SendingList struct {
	Users []string
}

// New creates new SendingList and returns it
func New() *SendingList {
	return &SendingList{}
}

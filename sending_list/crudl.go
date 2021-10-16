package sendinglist

// Create appends new username to s.Users
func (s *SendingList) Create(username string) {
	s.Users = append(s.Users, username)
}

// Read returns the s.Users
func (s *SendingList) Read() []string {
	return s.Users
}

// Update updates the s.Users with usernames slice
func (s *SendingList) Update(usernames []string) {
	s.Users = usernames
}

// Delete removes the username from s.Users
func (s *SendingList) Delete(username string) {
	for i, name := range s.Users {
		if name == username {
			newUsers := append(s.Users[:i], s.Users[i+1:]...)
			s.Users = newUsers
		}
	}
}

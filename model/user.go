package model

type User struct {
    ID    string `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
    Age   int32  `json:"age"`
}

// In-memory storage for demonstration
type UserStore struct {
    users map[string]*User
}

func NewUserStore() *UserStore {
    return &UserStore{
        users: make(map[string]*User),
    }
}

func (s *UserStore) Get(id string) (*User, bool) {
    user, exists := s.users[id]
    return user, exists
}

func (s *UserStore) List() []*User {
    users := make([]*User, 0, len(s.users))
    for _, user := range s.users {
        users = append(users, user)
    }
    return users
}

func (s *UserStore) Create(user *User) {
    s.users[user.ID] = user
}

func (s *UserStore) Update(user *User) bool {
    if _, exists := s.users[user.ID]; !exists {
        return false
    }
    s.users[user.ID] = user
    return true
}

func (s *UserStore) Delete(id string) bool {
    if _, exists := s.users[id]; !exists {
        return false
    }
    delete(s.users, id)
    return true
}

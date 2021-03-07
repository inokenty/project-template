package user

type User struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type GetUserByIDReply struct {
	User *User `json:"user"`
}

type ListUsersReply struct {
	Items []*User `json:"items"`
}

type CreateUserArgs struct {
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type CreateUserReply struct {
	ID uint `json:"id"`
}

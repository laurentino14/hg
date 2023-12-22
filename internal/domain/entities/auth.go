package entities

type Auth struct {
	User
}

func NewAuth(user User) *Auth {
	return &Auth{
		User: user,
	}
}

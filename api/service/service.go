package service

import (
	"github.com/jadengore/Ricetta/api/service/query"
)

type Svc struct {
	Query *query.Query
}

func NewService(uri string) *Svc {
	s := &Svc{
		query.NewQuery(uri),
	}
	return s
}

func (s Svc) CreateNewUser(handle, email, passwordHash string) bool {
	return s.Query.CreateUser(handle, email, passwordHash)
}

func (s Svc) HandleIsUnique(handle string) bool {
	return s.Query.HandleUnique(handle)
}

func (s Svc) EmailIsUnique(email string) bool {
	return s.Query.EmailUnique(email)
}

func (s Svc) GetHashedPassword(handle string) (hashedPassword []byte, ok bool) {
	return s.Query.GetHashedPassword(handle)
}

func (s Svc) SetGetNewAuthToken(handle string) (token string, ok bool) {
	return s.Query.SetGetNewAuthTokenForUser(handle)
}

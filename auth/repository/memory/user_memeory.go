package memory

import (
	"learn_gqlgen/auth/entity"
	errors "learn_gqlgen/auth/errors_auth"
)

type MemoryImpl struct {
	stockUser []*entity.User
}

func NewAuthRepository(stockUser []*entity.User) *MemoryImpl {
	return &MemoryImpl{
		stockUser: stockUser,
	}
}

func (mi *MemoryImpl) CreateAccount(user *entity.User) (*entity.User, error) {
	mi.stockUser = append(mi.stockUser, user)
	return user, nil
}

func (mi *MemoryImpl) FindUserByEmail(email string) (*entity.User, error) {
	for _, user := range mi.stockUser {
		if user.Email == email {
			return user, nil 
		}
	}
	return nil, errors.ErrEmailAlreadyExists
}
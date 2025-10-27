package user

import "ecommerce/domain"

type service struct {
	userRepo UserRepo
}

func NewService(userRepo UserRepo) Service {
	return &service{
		userRepo: userRepo,
	}
}

func (svc *service) Create(user domain.User) (*domain.User, error){
	createdUser, err := svc.userRepo.Create(user)
	if err != nil {
		return nil, err
	}
	if createdUser == nil {
		return nil, nil
	}
	return createdUser, nil
}

func (svc *service) Find(email string, pass string) (*domain.User, error){
	user, err := svc.userRepo.Find(email, pass)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, nil
	}
	return user, nil
}
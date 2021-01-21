package user

import (
	"Jinder/jinder-api/pkg/domain/registration"
	"Jinder/jinder-api/pkg/repository"
	"Jinder/jinder-api/pkg/service/dto"
	"errors"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo repository.User
}

func (s UserService) Register(user dto.User) (uuid.UUID, error) {
	if user.Password != user.PasswordConfirmation {
		return uuid.UUID{}, errors.New("passwords are not the same")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
	if err != nil {
		return uuid.UUID{}, err
	}

	id, err := s.repo.Register(Mapping(user, string(hash)))
	if err != nil {
		return uuid.UUID{}, err
	}

	return id, nil
}

func (s UserService) Update(user dto.User) (uuid.UUID, error) {
	if user.Password != user.PasswordConfirmation {
		return uuid.UUID{}, errors.New("passwords are not the same")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
	if err != nil {
		return uuid.UUID{}, err
	}

	id, err := s.repo.Update(Mapping(user, string(hash)))
	if err != nil {
		return uuid.UUID{}, err
	}

	return id, nil
}

func (s UserService) Get(userId uuid.UUID) (dto.User, error) {
	user, err := s.repo.Get(userId)
	if err != nil {
		return dto.User{}, err
	}

	return MappingToDto(user), nil
}

func NewService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}

func Mapping(user dto.User, hash string) registration.User {
	return registration.User{
		Id: user.Id,
		Person: registration.Person{
			Name:     user.Name,
			Surname:  user.Surname,
			Birthday: user.Birthday,
		},
		Credentials: registration.Credentials{
			Email:        user.Email,
			PasswordHash: hash,
		},
	}
}

func MappingToDto(user registration.User) dto.User {
	return dto.User{
		Id: user.Id,
		Person: dto.Person{
			Name:     user.Name,
			Surname:  user.Surname,
			Birthday: user.Birthday,
		},
		Credentials: dto.Credentials{
			Email: user.Email,
		},
	}
}

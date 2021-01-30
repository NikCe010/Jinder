package service

import (
	"Jinder/jinder-api/jobs/pkg/repository"
	"Jinder/jinder-api/jobs/pkg/service/dto/user"
	"errors"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo repository.User
}

func (s UserService) Register(userDto user.User) (uuid.UUID, error) {
	if userDto.Password != userDto.PasswordConfirmation {
		return uuid.UUID{}, errors.New("passwords are not the same")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(userDto.Password), bcrypt.MinCost)
	if err != nil {
		log.Error(err.Error())
		return uuid.UUID{}, err
	}

	id, err := s.repo.Register(user.ToDomain(userDto, string(hash)))
	if err != nil {
		log.Error(err.Error())
		return uuid.UUID{}, err
	}

	return id, nil
}

func (s UserService) UpdateUser(userDto user.User) (uuid.UUID, error) {
	if userDto.Password != userDto.PasswordConfirmation {
		error := errors.New("passwords are not the same")
		log.Error(error.Error())
		return uuid.UUID{}, error
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(userDto.Password), bcrypt.MinCost)
	if err != nil {
		log.Error(err.Error())
		return uuid.UUID{}, err
	}

	id, err := s.repo.Update(user.ToDomain(userDto, string(hash)))
	if err != nil {
		log.Error(err.Error())
		return uuid.UUID{}, err
	}

	return id, nil
}

func (s UserService) GetUser(userId uuid.UUID) (user.User, error) {
	u, err := s.repo.Get(userId)
	if err != nil {
		log.Error(err.Error())
		return user.User{}, err
	}

	return user.ToDto(u), nil
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}

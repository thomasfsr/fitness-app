package usecase

import (
    "github.com/yourusername/fitness-app/internal/domain/user"
)

type UserUseCase struct {
    repo user.Repository
}

func NewUserUseCase(r user.Repository) *UserUseCase {
    return &UserUseCase{repo: r}
}

func (u *UserUseCase) CreateUser(input *user.User) error {
    return u.repo.Create(input)
}

func (u *UserUseCase) GetUser(id uint32) (*user.User, error) {
    return u.repo.GetByID(id)
}

func (u *UserUseCase) UpdateUser(input *user.User) error {
    return u.repo.Update(input)
}

func (u *UserUseCase) DeleteUser(id uint32) error {
    return u.repo.Delete(id)
}

func (u *UserUseCase) ListUsers() ([]user.User, error) {
    return u.repo.List()
}

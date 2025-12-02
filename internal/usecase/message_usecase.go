package usecase

import (
    "github.com/yourusername/fitness-app/internal/domain/message"
    u "github.com/yourusername/fitness-app/internal/domain/user"
)

type MessageUseCase struct {
    repo message.Repository
    userRepo u.Repository
}

func NewMessageUseCase(r message.Repository, ur u.Repository) *MessageUseCase {
    return &MessageUseCase{repo: r, userRepo: ur}
}

func (muc *MessageUseCase) CreateMessage(m *message.Message) error {
    if _, err := muc.userRepo.GetByID(m.UserId); err != nil {
        return err
    }
    return muc.repo.Create(m)
}

func (muc *MessageUseCase) ListByUser(userId uint32) ([]message.Message, error) {
    return muc.repo.ListByUser(userId)
}

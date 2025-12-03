package repository

import (
	"time"
	"gorm.io/gorm"
	"github.com/thomasfsr/fitness-app/internal/domain/message"
)

type MessageGormRepository struct {
    db *gorm.DB
}

type MessageModel struct {
    ID        uint64    `gorm:"primaryKey;autoIncrement"`
    UserId    uint32
    Role      string    `gorm:"type:ENUM('user','assistant')"`
    Message   string    `gorm:"type:char(100)"`
    CreatedAt time.Time
}

func NewMessageGormRepository(db *gorm.DB) message.Repository {
    return &MessageGormRepository{db: db}
}

func (r *MessageGormRepository) Create(m *message.Message) error {
    mm := MessageModel{
        UserId:  m.UserId,
        Role:    m.Role,
        Message: m.Message,
    }
    return r.db.Create(&mm).Error
}

func (r *MessageGormRepository) ListByUser(userId uint32) ([]message.Message, error) {
    var ms []MessageModel
    if err := r.db.Where("user_id = ?", userId).Find(&ms).Error; err != nil {
        return nil, err
    }
    out := make([]message.Message, 0, len(ms))
    for _, m := range ms {
        out = append(out, message.Message{
            ID:        m.ID,
            UserId:    m.UserId,
            Role:      m.Role,
            Message:   m.Message,
            CreatedAt: m.CreatedAt,
        })
    }
    return out, nil
}

package repository

import (
    "time"

    "gorm.io/gorm"

    "github.com/yourusername/fitness-app/internal/domain/user"
)

type UserGormRepository struct {
    db *gorm.DB
}

type UserModel struct {
    ID        uint32    `gorm:"primaryKey;autoIncrement"`
    FirstName string    `gorm:"type:char(20)"`
    LastName  string    `gorm:"type:char(50)"`
    Whatsapp  uint64
    Active    bool
    CreatedAt time.Time
    UpdatedAt time.Time
}

func NewUserGormRepository(db *gorm.DB) user.Repository {
    return &UserGormRepository{db: db}
}

func (r *UserGormRepository) toDomain(m *UserModel) *user.User {
    return &user.User{
        ID:        m.ID,
        FirstName: m.FirstName,
        LastName:  m.LastName,
        Whatsapp:  m.Whatsapp,
        Active:    m.Active,
        CreatedAt: m.CreatedAt,
        UpdatedAt: m.UpdatedAt,
    }
}

func (r *UserGormRepository) Create(u *user.User) error {
    m := UserModel{
        FirstName: u.FirstName,
        LastName:  u.LastName,
        Whatsapp:  u.Whatsapp,
        Active:    u.Active,
    }
    return r.db.Create(&m).Error
}

func (r *UserGormRepository) GetByID(id uint32) (*user.User, error) {
    var m UserModel
    if err := r.db.First(&m, id).Error; err != nil {
        return nil, err
    }
    return r.toDomain(&m), nil
}

func (r *UserGormRepository) Update(u *user.User) error {
    var m UserModel
    if err := r.db.First(&m, u.ID).Error; err != nil {
        return err
    }
    m.FirstName = u.FirstName
    m.LastName = u.LastName
    m.Whatsapp = u.Whatsapp
    m.Active = u.Active
    m.UpdatedAt = time.Now()
    return r.db.Save(&m).Error
}

func (r *UserGormRepository) Delete(id uint32) error {
    return r.db.Delete(&UserModel{}, id).Error
}

func (r *UserGormRepository) List() ([]user.User, error) {
    var ms []UserModel
    if err := r.db.Find(&ms).Error; err != nil {
        return nil, err
    }
    out := make([]user.User, 0, len(ms))
    for _, m := range ms {
        out = append(out, *r.toDomain(&m))
    }
    return out, nil
}

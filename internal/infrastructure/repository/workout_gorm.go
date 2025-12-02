package repository

import (
    "time"

    "gorm.io/gorm"

    "github.com/thomasfsr/fitness-app/internal/domain/workout"
)

type WorkoutGormRepository struct {
    db *gorm.DB
}

type WorkoutModel struct {
    ID        uint64    `gorm:"primaryKey;autoIncrement"`
    UserId    uint32
    Exercise  string    `gorm:"type:char(100)"`
    Weight    uint16
    Reps      uint8
    CreatedAt time.Time
    UpdatedAt time.Time
}

func NewWorkoutGormRepository(db *gorm.DB) workout.Repository {
    return &WorkoutGormRepository{db: db}
}

func (r *WorkoutGormRepository) toDomain(m *WorkoutModel) *workout.WorkoutSet {
    return &workout.WorkoutSet{
        ID:        m.ID,
        UserId:    m.UserId,
        Exercise:  m.Exercise,
        Weight:    m.Weight,
        Reps:      m.Reps,
        CreatedAt: m.CreatedAt,
        UpdatedAt: m.UpdatedAt,
    }
}

func (r *WorkoutGormRepository) Create(w *workout.WorkoutSet) error {
    m := WorkoutModel{
        UserId:   w.UserId,
        Exercise: w.Exercise,
        Weight:   w.Weight,
        Reps:     w.Reps,
    }
    return r.db.Create(&m).Error
}

func (r *WorkoutGormRepository) GetByID(id uint64) (*workout.WorkoutSet, error) {
    var m WorkoutModel
    if err := r.db.First(&m, id).Error; err != nil {
        return nil, err
    }
    return r.toDomain(&m), nil
}

func (r *WorkoutGormRepository) Update(w *workout.WorkoutSet) error {
    var m WorkoutModel
    if err := r.db.First(&m, w.ID).Error; err != nil {
        return err
    }
    m.Exercise = w.Exercise
    m.Weight = w.Weight
    m.Reps = w.Reps
    m.UpdatedAt = time.Now()
    return r.db.Save(&m).Error
}

func (r *WorkoutGormRepository) Delete(id uint64) error {
    return r.db.Delete(&WorkoutModel{}, id).Error
}

func (r *WorkoutGormRepository) ListByUser(userId uint32) ([]workout.WorkoutSet, error) {
    var ms []WorkoutModel
    if err := r.db.Where("user_id = ?", userId).Find(&ms).Error; err != nil {
        return nil, err
    }
    out := make([]workout.WorkoutSet, 0, len(ms))
    for _, m := range ms {
        out = append(out, *r.toDomain(&m))
    }
    return out, nil
}

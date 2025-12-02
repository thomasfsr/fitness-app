package workout

type Repository interface {
    Create(w *WorkoutSet) error
    GetByID(id uint64) (*WorkoutSet, error)
    Update(w *WorkoutSet) error
    Delete(id uint64) error
    ListByUser(userId uint32) ([]WorkoutSet, error)
}

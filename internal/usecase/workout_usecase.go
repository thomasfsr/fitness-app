package usecase

import (
    "github.com/thomasfsr/fitness-app/internal/domain/workout"
    u "github.com/thomasfsr/fitness-app/internal/domain/user"
)

type WorkoutUseCase struct {
    repo workout.Repository
    userRepo u.Repository
}

func NewWorkoutUseCase(r workout.Repository, ur u.Repository) *WorkoutUseCase {
    return &WorkoutUseCase{repo: r, userRepo: ur}
}

func (w *WorkoutUseCase) CreateWorkout(input *workout.WorkoutSet) error {
    // validate user exists
    if _, err := w.userRepo.GetByID(input.UserId); err != nil {
        return err
    }
    return w.repo.Create(input)
}

func (w *WorkoutUseCase) GetWorkout(id uint64) (*workout.WorkoutSet, error) {
    return w.repo.GetByID(id)
}

func (w *WorkoutUseCase) UpdateWorkout(input *workout.WorkoutSet) error {
    return w.repo.Update(input)
}

func (w *WorkoutUseCase) DeleteWorkout(id uint64) error {
    return w.repo.Delete(id)
}

func (w *WorkoutUseCase) ListByUser(userId uint32) ([]workout.WorkoutSet, error) {
    return w.repo.ListByUser(userId)
}

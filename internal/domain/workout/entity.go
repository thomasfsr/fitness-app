package workout

import "time"

type WorkoutSet struct {
    ID        uint64
    UserId    uint32
    Exercise  string
    Weight    uint16
    Reps      uint8
    CreatedAt time.Time
    UpdatedAt time.Time
}

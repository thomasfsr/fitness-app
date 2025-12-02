package user

import "time"

type User struct {
    ID        uint32
    FirstName string
    LastName  string
    Whatsapp  uint64
    Active    bool
    CreatedAt time.Time
    UpdatedAt time.Time
}

package message

import "time"

type Message struct {
    ID        uint64
    UserId    uint32
    Role      string
    Message   string
    CreatedAt time.Time
}

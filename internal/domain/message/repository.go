package message

type Repository interface {
    Create(m *Message) error
    ListByUser(userId uint32) ([]Message, error)
}

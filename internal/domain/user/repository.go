package user

type Repository interface {
    Create(u *User) error
    GetByID(id uint32) (*User, error)
    Update(u *User) error
    Delete(id uint32) error
    List() ([]User, error)
}

package types

type User struct {
	ID        int    `json:"id"`
	Password  string `json:"-"`
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	CreatedAt string `json:"createdAt"`
}

type Record struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Category    string `json:"category"`
	Amount      int    `json:"amount"`
	CreatedAt   string `json:"createAt"`
}

type UserStore interface {
	GetUserByEmail(email string) (User, error)
	GetUserByID(id int) (*User, error)
	CreateUser(User) (userId int, err error)
}

type RegisterUserPayload struct {
	Email     string `json:"email" validate:"required"`
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
	Password  string `json:"password" validate:"required"`
}

type LoginUserPayload struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type RecordStore interface {
	GetUserRecords(userId string) ([]Record, error)
	GetRecordById(id string) (Record, error)
	GetUserRecordsByCategory(userId string, category string) ([]Record, error)
	CreateUserRecord(userId string, record Record) error
	UpdateRecord(recordId string, updates Record) error
	DeleteRecord(recordId string) error
}

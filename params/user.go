package params

import "time"

type CreateUser struct {
	ID        int        `json:"int"`
	Age       int        `json:"age"`
	Email     string     `json:"email"`
	Password  string     `json:"password,omitempty"`
	Username  string     `json:"username"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

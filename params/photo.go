package params

import "time"

type CreatePhoto struct {
	ID       int    `json:"id"`
	UserID   int    `json:"user_id"`
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url"`
}

type GetAllPhotos struct {
	ID        int        `json:"id"`
	Title     string     `json:"title"`
	Caption   string     `json:"caption"`
	PhotoUrl  string     `json:"photo_url"`
	UserID    int        `json:"user_id"`
	CreatedAt time.Time  `json:"created_at,omitempty"`
	UpdateAt  time.Time  `json:"updated_at,omitempty"`
	User      *UserPhoto `json:"user,omitempty"`
}

type UserPhoto struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}

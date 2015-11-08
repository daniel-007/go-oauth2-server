package oauth

import "time"

// Client ...
type Client struct {
	ID          int
	ClientID    string `sql:"type:varchar(254);unique;not null"`
	Secret      string `sql:"type:varchar(60);not null"`
	RedirectURI string `sql:"type:varchar(200)"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// Scope ...
type Scope struct {
	ID          int
	Scope       string `sql:"type:varchar(200);unique;not null"`
	Description string
	IsDefault   bool `sql:"default:false"`
}

// User ...
type User struct {
	ID        int
	Username  string `sql:"type:varchar(254);unique;not null"`
	Password  string `sql:"type:varchar(60);not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// RefreshToken ...
type RefreshToken struct {
	ID           int
	RefreshToken string    `sql:"type:varchar(40);unique;not null"`
	ExpiresAt    time.Time `sql:"not null"`
}

// AccessToken ...
type AccessToken struct {
	ID             int
	AccessToken    string    `sql:"type:varchar(40);unique;not null"`
	ExpiresAt      time.Time `sql:"not null"`
	Scope          string    `sql:"type:varchar(200);not null"`
	Client         Client
	ClientID       int `sql:"index;not null"`
	User           User
	UserID         int `sql:"index"`
	RefreshToken   RefreshToken
	RefreshTokenID int `sql:"index"`
}

// AuthCode ...
type AuthCode struct {
	ID          int
	Code        string    `sql:"type:varchar(40);unique;not null"`
	RedirectURI string    `sql:"type:varchar(200)"`
	ExpiresAt   time.Time `sql:"not null"`
	Scope       string    `sql:"type:varchar(200);not null"`
	Client      Client
	ClientID    int `sql:"index;not null"`
	User        User
	UserID      int `sql:"index"`
}
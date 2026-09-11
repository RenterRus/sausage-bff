package entity

import "time"

// region: task
type Task struct {
	TransactionUUID *string
	Title           string
	Comment         *string
	Priority        int32
	StartDate       time.Time
	Complete        bool
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type UpsertTaskInput struct {
	TransactionUUID *string
	UserID          string
	Title           string
	Comment         *string
	Priority        int32
	StartDate       time.Time
	Complete        bool
}

type User struct {
	UUID  string
	Login string
}

// region: auth

type AcceptRequest struct {
	Login   string
	OtpCode string
}

// RevokeSession uses pointers to represent the `oneof` field.
type RevokeSessionRequest struct {
	Current *RevokeSessionCurrent
	All     *RevokeSessionAll
}
type RevokeSessionCurrent struct {
	Hash string
}
type RevokeSessionAll struct {
	Login string
}

type Tokens struct {
	Access  string
	Refresh string
}

type LoginOTPRequest struct {
	Login     string
	UserAgent string
	Code      string
}

type RefreshRequest struct {
	Login        string
	UserAgent    string
	RefreshToken string
}

package entity

import "time"

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

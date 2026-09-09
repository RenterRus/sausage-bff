package entity

import "errors"

var (
	ErrNoRequiredParams = errors.New("required params not found")
)

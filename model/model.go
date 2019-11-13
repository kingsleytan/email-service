package model

import (
	"time"
)

// Model :
type Model struct {
	CreatedDateTime time.Time `json:"createdAt"`
	UpdatedDateTime time.Time `json:"updatedAt"`
}

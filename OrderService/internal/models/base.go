package models

import "time"

type BaseModel struct {
	ID        string    `bson:"_id,omitempty"`
	CreatedAt time.Time `bson:"createdAt"`
	UpdatedAt time.Time `bson:"updatedAt"`
}

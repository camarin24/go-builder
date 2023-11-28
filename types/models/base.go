package models

import "time"

type BaseModel struct {
	ID               string
	Created, Updated time.Time
}

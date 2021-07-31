package models

import (
	"errors"
	"time"
)

var ErrNoRecord = errors.New("models: no maching record found")

type Pegion struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

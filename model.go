package main

import (
	"time"
)

type Product struct {
	ID        uint
	CreatedAt time.Time
	Code      string
	Price     int
}

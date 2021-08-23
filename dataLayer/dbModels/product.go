package dbModels

import "time"

type Product struct {
	Id               int64
	BrandId          int64
	CategoryId       int64
	Name             string
	LongDescription  string
	ShortDescription string
	Cost             int64
	CardImage        string
	CreatedAt        *time.Time
}

package models

import "time"

// gorm.Model definition
type File struct {
	Name      string // file name
	Bucket    string // minio bucket name
	CreatedAt time.Time
	UpdatedAt time.Time
}

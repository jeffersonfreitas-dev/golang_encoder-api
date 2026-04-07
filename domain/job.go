package domain

import "time"

type StatusJob string

const (
	StatusSuccess StatusJob = "success"
	StatusFailed  StatusJob = "failed"
)

type Job struct {
	ID               string
	OutputBucketPath string
	Status           StatusJob
	Video            *Video
	Error            string
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

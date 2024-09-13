package models

import "time"

type BidStatus string

const (
	BidCreated   BidStatus = "CREATED"
	BidPublished BidStatus = "PUBLISHED"
	BidCanceled  BidStatus = "CANCELED"
	BidApproved  BidStatus = "APPROVED"
	BidRejected  BidStatus = "REJECTED"
)

type Bid struct {
	ID             uint      `gorm:"primaryKey" json:"id"`
	Name           string    `gorm:"not null" json:"name"`
	Description    string    `json:"description"`
	Status         BidStatus `gorm:"default:CREATED" json:"status"`
	Version        int       `gorm:"default:1" json:"version"`
	TenderID       uint      `gorm:"not null" json:"tenderId"`
	AuthorID       uint      `gorm:"not null" json:"authorId"`
	OrganizationID uint      `gorm:"not null" json:"organizationId"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}

type BidVersion struct {
	ID          uint      `gorm:"primaryKey"`
	BidID       uint      `gorm:"not null"`
	Version     int       `gorm:"not null"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
}

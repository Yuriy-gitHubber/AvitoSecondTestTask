package models

import "time"

type TenderStatus string

const (
	TenderCreated   TenderStatus = "CREATED"
	TenderPublished TenderStatus = "PUBLISHED"
	TenderClosed    TenderStatus = "CLOSED"
)

type Tender struct {
	ID             uint         `gorm:"primaryKey" json:"id"`
	Name           string       `gorm:"not null" json:"name"`
	Description    string       `json:"description"`
	Status         TenderStatus `gorm:"default:CREATED" json:"status"`
	Version        int          `gorm:"default:1" json:"version"`
	OrganizationID uint         `gorm:"not null" json:"organizationId"`
	CreatedAt      time.Time    `json:"createdAt"`
	UpdatedAt      time.Time    `json:"updatedAt"`
}

type TenderVersion struct {
	ID          uint      `gorm:"primaryKey"`
	TenderID    uint      `gorm:"not null"`
	Version     int       `gorm:"not null"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
}

package models

import "time"

type OrganizationType string

const (
	IE  OrganizationType = "IE"
	LLC OrganizationType = "LLC"
	JSC OrganizationType = "JSC"
)

type Organization struct {
	ID          uint             `gorm:"primaryKey" json:"id"`
	Name        string           `gorm:"not null" json:"name"`
	Description string           `json:"description"`
	Type        OrganizationType `gorm:"type:organization_type" json:"type"`
	CreatedAt   time.Time        `json:"createdAt"`
	UpdatedAt   time.Time        `json:"updatedAt"`
}

type OrganizationResponsible struct {
	ID             uint `gorm:"primaryKey" json:"id"`
	OrganizationID uint `gorm:"not null;constraint:OnDelete:CASCADE;" json:"organizationId"`
	UserID         uint `gorm:"not null;constraint:OnDelete:CASCADE;" json:"userId"`
}

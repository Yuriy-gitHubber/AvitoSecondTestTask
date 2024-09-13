package models

import (
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) error {
	err := db.AutoMigrate(&Employee{}, &Organization{}, &OrganizationResponsible{}, &Tender{}, &TenderVersion{}, &Bid{}, &BidVersion{})
	if err != nil {
		return err
	}

	// Create custom types if necessary (e.g., organization_type)
	err = db.Exec(`DO $$ BEGIN
		IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'organization_type') THEN
			CREATE TYPE organization_type AS ENUM ('IE', 'LLC', 'JSC');
		END IF;
	END$$;`).Error
	if err != nil {
		return err
	}

	return nil
}

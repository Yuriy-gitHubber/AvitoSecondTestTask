package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"ZADANIE-6105/models"

	"ZADANIE-6105/utils"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type TenderController struct {
	DB *gorm.DB
}

func (tc *TenderController) GetTenders(w http.ResponseWriter, r *http.Request) {
	var tenders []models.Tender
	tc.DB.Find(&tenders)
	utils.RespondJSON(w, http.StatusOK, tenders)
}

func (tc *TenderController) CreateTender(w http.ResponseWriter, r *http.Request) {
	var tender models.Tender
	if err := json.NewDecoder(r.Body).Decode(&tender); err != nil {
		utils.RespondError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	tender.Status = models.TenderCreated
	if err := tc.DB.Create(&tender).Error; err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondJSON(w, http.StatusCreated, tender)
}

func (tc *TenderController) EditTender(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	tenderID, _ := strconv.Atoi(params["tenderId"])

	var tender models.Tender
	if err := tc.DB.First(&tender, tenderID).Error; err != nil {
		utils.RespondError(w, http.StatusNotFound, "Tender not found")
		return
	}

	var updatedTender models.Tender
	if err := json.NewDecoder(r.Body).Decode(&updatedTender); err != nil {
		utils.RespondError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// Save current version
	tenderVersion := models.TenderVersion{
		TenderID:    tender.ID,
		Version:     tender.Version,
		Name:        tender.Name,
		Description: tender.Description,
		CreatedAt:   tender.UpdatedAt,
	}
	tc.DB.Create(&tenderVersion)

	// Update tender
	tender.Name = updatedTender.Name
	tender.Description = updatedTender.Description
	tender.Version += 1
	tc.DB.Save(&tender)

	utils.RespondJSON(w, http.StatusOK, tender)
}

func (tc *TenderController) RollbackTender(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	tenderID, _ := strconv.Atoi(params["tenderId"])
	version, _ := strconv.Atoi(params["version"])

	var tender models.Tender
	if err := tc.DB.First(&tender, tenderID).Error; err != nil {
		utils.RespondError(w, http.StatusNotFound, "Tender not found")
		return
	}

	var tenderVersion models.TenderVersion
	if err := tc.DB.Where("tender_id = ? AND version = ?", tenderID, version).First(&tenderVersion).Error; err != nil {
		utils.RespondError(w, http.StatusNotFound, "Version not found")
		return
	}

	// Update tender to the selected version
	tender.Name = tenderVersion.Name
	tender.Description = tenderVersion.Description
	tender.Version = tenderVersion.Version + 1 // Increment version after rollback
	tc.DB.Save(&tender)

	utils.RespondJSON(w, http.StatusOK, tender)
}

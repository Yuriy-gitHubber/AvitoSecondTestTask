package controllers

import (
	"ZADANIE-6105/models"
	"ZADANIE-6105/utils"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type BidController struct {
	DB *gorm.DB
}

func (bc *BidController) CreateBid(w http.ResponseWriter, r *http.Request) {
	var bid models.Bid
	if err := json.NewDecoder(r.Body).Decode(&bid); err != nil {
		utils.RespondError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	bid.Status = models.BidCreated
	if err := bc.DB.Create(&bid).Error; err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondJSON(w, http.StatusCreated, bid)
}

func (bc *BidController) EditBid(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bidID, _ := strconv.Atoi(params["bidId"])

	var bid models.Bid
	if err := bc.DB.First(&bid, bidID).Error; err != nil {
		utils.RespondError(w, http.StatusNotFound, "Bid not found")
		return
	}

	var updatedBid models.Bid
	if err := json.NewDecoder(r.Body).Decode(&updatedBid); err != nil {
		utils.RespondError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// Save current version
	bidVersion := models.BidVersion{
		BidID:       bid.ID,
		Version:     bid.Version,
		Name:        bid.Name,
		Description: bid.Description,
		CreatedAt:   bid.UpdatedAt,
	}
	bc.DB.Create(&bidVersion)

	// Update bid
	bid.Name = updatedBid.Name
	bid.Description = updatedBid.Description
	bid.Version += 1
	bc.DB.Save(&bid)

	utils.RespondJSON(w, http.StatusOK, bid)
}

func (bc *BidController) RollbackBid(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bidID, _ := strconv.Atoi(params["bidId"])
	version, _ := strconv.Atoi(params["version"])

	var bid models.Bid
	if err := bc.DB.First(&bid, bidID).Error; err != nil {
		utils.RespondError(w, http.StatusNotFound, "Bid not found")
		return
	}

	var bidVersion models.BidVersion
	if err := bc.DB.Where("bid_id = ? AND version = ?", bidID, version).First(&bidVersion).Error; err != nil {
		utils.RespondError(w, http.StatusNotFound, "Version not found")
		return
	}

	// Update bid to the selected version
	bid.Name = bidVersion.Name
	bid.Description = bidVersion.Description
	bid.Version = bidVersion.Version + 1 // Increment version after rollback
	bc.DB.Save(&bid)

	utils.RespondJSON(w, http.StatusOK, bid)
}

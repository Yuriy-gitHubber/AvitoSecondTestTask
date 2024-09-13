package controllers

import "gorm.io/gorm"

var TenderCtrl TenderController
var BidCtrl BidController

func InitControllers(db *gorm.DB) {
	TenderCtrl = TenderController{DB: db}
	BidCtrl = BidController{DB: db}
}

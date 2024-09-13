package routes

import (
	"ZADANIE-6105/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	// Health check
	router.HandleFunc("/api/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	}).Methods("GET")

	// Tenders
	router.HandleFunc("/api/tenders", controllers.TenderCtrl.GetTenders).Methods("GET")
	router.HandleFunc("/api/tenders/new", controllers.TenderCtrl.CreateTender).Methods("POST")
	router.HandleFunc("/api/tenders/{tenderId}/edit", controllers.TenderCtrl.EditTender).Methods("PATCH")
	router.HandleFunc("/api/tenders/{tenderId}/rollback/{version}", controllers.TenderCtrl.RollbackTender).Methods("PUT")

	// Bids
	router.HandleFunc("/api/bids/new", controllers.BidCtrl.CreateBid).Methods("POST")
	router.HandleFunc("/api/bids/{bidId}/edit", controllers.BidCtrl.EditBid).Methods("PATCH")
	router.HandleFunc("/api/bids/{bidId}/rollback/{version}", controllers.BidCtrl.RollbackBid).Methods("PUT")

	return router
}

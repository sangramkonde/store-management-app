package routes

import (
	"github.com/gorilla/mux"
	"github.com/sangramkonde/store-management-app/pkg/controllers"
)

var RegisterStoreRoutes = func(router *mux.Router){
     router.HandleFunc("/store/", controllers.CreateStore).Methods("POST")
	 router.HandleFunc("/store/", controllers.GetStore).Methods("GET")
	 router.HandleFunc("/store/{storeId}", controllers.GetStoreById).Methods("GET")
	 router.HandleFunc("/store/{storeId}", controllers.UpdateStore).Methods("PUT")
	 router.HandleFunc("/store/{storeId}", controllers.DeleteStore).Methods("DELETE")
}


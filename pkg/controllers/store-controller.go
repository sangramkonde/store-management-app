package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sangramkonde/store-management-app/pkg/models"
	"github.com/sangramkonde/store-management-app/pkg/utils"
)

var NewStore models.Store

func GetStore(w http.ResponseWriter, r *http.Request){
	newStores := models.GetAllStores()
	res, _ := json.Marshal(newStores)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetStoreById(w http.ResponseWriter, r *http.Request){
    vars := mux.Vars(r)
	storeId := vars["storeId"]
	ID, err := strconv.ParseInt(storeId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
    storeDetails, _ := models.GetStoreById(ID)
	res, _ := json.Marshal(storeDetails)
    w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateStore(w http.ResponseWriter, r *http.Request){
	CreateStore := &models.Store{}
	utils.ParseBody(r, CreateStore)
	s := CreateStore.CreateStore()
    res, _ := json.Marshal(s)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteStore(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	storeId := vars["storeId"]
	ID, err := strconv.ParseInt(storeId, 0, 0)
	if err != nil{
		fmt.Println("error while parsing")
	}
	store := models.DeleteStore(ID)
	res, _ := json.Marshal(store)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateStore(w http.ResponseWriter, r *http.Request){
	updateStore := &models.Store{}
	utils.ParseBody(r, updateStore)
	vars := mux.Vars(r)
    storeId := vars["storeId"]
	ID, err := strconv.ParseInt(storeId, 0, 0)
	if err != nil{
		fmt.Println("error while parsing")
	}
	storeDetails, db := models.GetStoreById(ID)
	if updateStore.Name != "" {
	   storeDetails.Name = updateStore.Name	
	}
	if updateStore.Owner != "" {
		storeDetails.Owner = updateStore.Owner	
	}
	if updateStore.EstablishedYear != "" {
		storeDetails.EstablishedYear = updateStore.EstablishedYear	
	}
	db.Save(&storeDetails)
	res, _ := json.Marshal(storeDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
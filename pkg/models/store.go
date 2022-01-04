package models

import (
	"github.com/jinzhu/gorm"
	"github.com/sangramkonde/store-management-app/pkg/config"
)

var db *gorm.DB

type Store struct{
	gorm.Model
	Name string `gorm:""json:"name"`
	Owner string `json:"owner"`
	EstablishedYear string `json:"establishedYear"`
}

func init(){
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Store{})
}

func (s *Store) CreateStore() *Store{
   db.NewRecord(s)
   db.Create(&s)
   return s
}

func GetAllStores() []Store{
	var stores []Store
	db.Find(&stores)
	return stores
}

func GetStoreById(Id int64) (*Store, *gorm.DB){
	var getStore Store
	db := db.Where("ID=?", Id).Find(&getStore)
	return &getStore, db
}

func DeleteStore(Id int64) Store{
	var store Store
	db.Where("ID=?", Id).Delete(store)
	return store
}
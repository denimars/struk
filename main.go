package main

import (
    "log"

    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

type Business struct {
    ID     uint
    Name   string `gorm:"not null"`
    Tables Tables `gorm:"ForeignKey:BusinessID"`
}

type Table struct {
    ID         uint
    Ref        string `gorm:"not null"`
    Business   Business
    BusinessID uint
}

type Tables []Table
type Businesses []Business

func main() {
    var err error
    var db *gorm.DB

    db, err = gorm.Open("mysql", "root@/ptk?charset=utf8&parseTime=True&loc=Local")
    if err != nil {
        log.Fatal(err)
    }else{
		log.Println("Berhasil")
	}

    defer db.Close()

    db.LogMode(true)
    db.AutoMigrate(&Business{})
    db.AutoMigrate(&Table{})

    bus := Business{
        Name: "Test",
        Tables: Tables{
            Table{Ref: "A1"},
        },
    }
    db.Create(&bus)
    var businesses Businesses
    db.Preload("Tables").Find(&businesses)
	log.Println(businesses)

	var kk Businesses
	db.Preload("Tables").Find(&kk, 1)
	log.Println(kk[0].Name)
	log.Println(kk[0].Tables[0].Ref)
	/**log.Println(kk.Tables.ID)
	log.Println(kk.Tables.Ref)
	log.Println(kk.Tables.BusinessID)*/
	
	
	
}
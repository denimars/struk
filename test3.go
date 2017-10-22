package main

import (
    "log"

    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

type Business struct {
    ID     uint
    Name   string `gorm:"not null"`
    Kucing Tablet `gorm:"ForeignKey:BusinessID"`
}

type Tablet struct {
    ID         uint
    Ref        string `gorm:"not null"`
    BusinessID uint
}

type Hasil struct{
	Name string
	Ref string
}

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
	db.SingularTable(true)

    //db.LogMode(true)
    db.AutoMigrate(&Business{})
	db.AutoMigrate(&Tablet{})
	
   
	var businesses []Business
	db.Preload("Kucing").Find(&businesses)
	log.Println(businesses)

	for _, i := range businesses{
		log.Println(i.ID)
		log.Println(i.Name)
		log.Println(i.Kucing.Ref)
		log.Println("________________________")
	}

	var hasil []Hasil
	db.Raw("select business.name, tablet.ref from business, tablet where business.id = tablet.business_id").Scan(&hasil)
	log.Println(hasil)
	

	/**var kk Business
	db.Preload("Tables").Find(&kk, 1)
	log.Println(kk[0].Name)ß
	log.Println(kk[0].Tables[0].Ref)
	/**log.Println(kk.Tables.ID)
	log.Println(kk.Tables.Ref)
	log.Println(kk.Tables.BusinessID)*/
	
	
	
}
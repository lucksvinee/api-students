package db

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)
  
	type Student struct {
		gorm.Model
		Name string
		CPF int
		Email string
		Age int
		Active bool
	}
  func Init() *gorm.DB {
	dbPath := "/home/hitsugaya/Área de trabalho/api-students/student.db"
	if _, err := os.Stat("/home/hitsugaya/Área de trabalho/api-students"); os.IsNotExist(err) {
        os.MkdirAll("/home/hitsugaya/Área de trabalho/api-students", os.ModePerm)
    }

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&Student{})

	return db
  }

  func AddStudent(){
	db := Init()

	student := Student{
		Name: "Lucas Inacio",
		CPF: 35674513324,
		Email: "dsdadjuw@example.com",
		Age: 29,
		Active: true,
	}
	
	if result := db.Create(&student); result.Error != nil {
		fmt.Println("Error creating student:")
	}  else {
        fmt.Println("Student created:", student)
    }
	
  }

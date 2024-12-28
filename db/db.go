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
		Name string `json:"name"`
		CPF int `json:"cpf"`
		Email string `json:"email"`
		Age int `json:"age"`
		Active bool `json:"registration"`
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

  func AddStudent(student Student) error {
	db := Init()

	if result := db.Create(&student); result.Error != nil {
		return result.Error
	}  
        fmt.Println("Student created:")
		return nil
  }

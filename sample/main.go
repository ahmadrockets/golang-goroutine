package main

import (
	"ahmadrockets/golang-goroutine/sample/config"
	"ahmadrockets/golang-goroutine/sample/model"
	"ahmadrockets/golang-goroutine/sample/repository"
	"ahmadrockets/golang-goroutine/sample/service"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Init Config
	conf, err := config.InitConfig(".")
	if err != nil {
		log.Fatalf("Error load config with err : %v", err)
	}

	// Init sqlLite
	sqlLite := config.InitSqlLite(conf)

	// init repository
	studentRepo := repository.NewStudentRepository(sqlLite.Client)

	// init service
	studentService := service.NewStudentService(studentRepo)

	// Init sample data students
	students := model.SampleStudent()

	// Insert Data Students With No Goroutine
	timeNow := time.Now()
	err = studentService.InsertStudentsWithWg(conf.MaxProcs, students)
	if err != nil {
		log.Printf("Error inserting data student %v \n", err.Error())
	}
	finishTime := time.Since(timeNow)
	log.Printf("[NO GOROUTINE] Finish inserting %d data student, take %s time", len(students), finishTime.String())

	// Insert Data Students with Goroutine
	timeNowWg := time.Now()
	err = studentService.InsertStudents(students)
	if err != nil {
		log.Printf("Error inserting data student %v \n", err.Error())
	}
	finishTimeWg := time.Since(timeNowWg)
	log.Printf("[GOROUTINE] Finish inserting %d data student with goroutine (Waitgroup), take %s time", len(students), finishTimeWg.String())

	// Get All Data Student
	resStudents, err := studentService.GetAllStudents()
	if err != nil {
		log.Printf("Error get all data student %v \n", err.Error())
	}
	log.Printf("Total Data Inserted in database (student table) = %d", len(resStudents))

}

package repository

import (
	"ahmadrockets/golang-goroutine/sample/model"
	"database/sql"
	"log"
	"sync"
)

type StudentRepository interface {
	InsertStudent(student model.Student) error
	InsertStudentWithWg(wg *sync.WaitGroup, student model.Student)
	GetAllStudent() (students []model.Student, err error)
}

type StudentRepoImpl struct {
	SqlLiteClient *sql.DB
}

func NewStudentRepository(client *sql.DB) StudentRepository {
	return &StudentRepoImpl{
		SqlLiteClient: client,
	}
}

func (r *StudentRepoImpl) InsertStudent(student model.Student) (err error) {
	insertStudentSQL := `INSERT INTO student(code, name, program) VALUES (?, ?, ?)`
	statement, err := r.SqlLiteClient.Prepare(insertStudentSQL)

	if err != nil {
		log.Printf("Error prepare query insert student : %s", err.Error())
		return
	}
	_, err = statement.Exec(student.Code, student.Name, student.Program)
	if err != nil {
		log.Printf("Error execute insert student : %s", err.Error())
		return
	}

	return
}

func (r *StudentRepoImpl) InsertStudentWithWg(wg *sync.WaitGroup, student model.Student) {
	defer wg.Done()

	insertStudentSQL := `INSERT INTO student(code, name, program) VALUES (?, ?, ?)`
	statement, err := r.SqlLiteClient.Prepare(insertStudentSQL)

	if err != nil {
		log.Printf("Error prepare query insert student : %s", err.Error())
	}
	_, err = statement.Exec(student.Code, student.Name+" [GO]", student.Program)
	if err != nil {
		log.Printf("Error execute insert student : %s", err.Error())
	}
}

func (r *StudentRepoImpl) GetAllStudent() (students []model.Student, err error) {
	row, err := r.SqlLiteClient.Query("SELECT * FROM student ORDER BY name")
	if err != nil {
		log.Printf("Error get all data student : %v", err.Error())
		return
	}
	defer row.Close()
	for row.Next() {
		var id int
		var code string
		var name string
		var program string
		row.Scan(&id, &code, &name, &program)

		students = append(students, model.Student{
			Code:    code,
			Name:    name,
			Program: program,
		})
	}

	return
}

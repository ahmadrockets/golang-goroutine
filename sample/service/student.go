package service

import (
	"ahmadrockets/golang-goroutine/sample/model"
	"ahmadrockets/golang-goroutine/sample/repository"
	"runtime"
	"sync"
)

type StudentService interface {
	InsertStudent(student model.Student) (err error)
	InsertStudents(students []model.Student) (err error)
	InsertStudentsWithWg(maxProcs int, students []model.Student) (err error)
	GetAllStudents() (students []model.Student, err error)
}

type StudentServiceImpl struct {
	studentRepo repository.StudentRepository
}

func NewStudentService(studentRepo repository.StudentRepository) StudentService {
	return &StudentServiceImpl{
		studentRepo: studentRepo,
	}
}

func (s *StudentServiceImpl) InsertStudent(student model.Student) (err error) {
	err = s.studentRepo.InsertStudent(student)
	if err != nil {
		return
	}
	return nil
}

func (s *StudentServiceImpl) InsertStudents(students []model.Student) (err error) {
	// Process Insert data
	var totalSuccess, totalFailed int
	for _, student := range students {
		err = s.studentRepo.InsertStudent(student)
		if err != nil {
			totalFailed++
			continue
		}
		totalSuccess++
	}
	return
}

func (s *StudentServiceImpl) InsertStudentsWithWg(maxProcs int, students []model.Student) (err error) {
	runtime.GOMAXPROCS(maxProcs)
	var wg sync.WaitGroup
	wg.Add(len(students))
	for _, student := range students {
		go s.studentRepo.InsertStudentWithWg(&wg, student)
	}
	wg.Wait()
	return
}

func (s *StudentServiceImpl) GetAllStudents() (students []model.Student, err error) {
	students, err = s.studentRepo.GetAllStudent()
	return
}

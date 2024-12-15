package model

type Student struct {
	Code    string
	Name    string
	Program string
}

func SampleStudent() (students []Student) {
	students = appendStudent(students, "0001", "Liana Kim", "Bachelor")
	students = appendStudent(students, "0002", "Glen Rangel", "Bachelor")
	students = appendStudent(students, "0003", "Martin Martins", "Master")
	students = appendStudent(students, "0004", "Alayna Armitage", "PHD")
	students = appendStudent(students, "0005", "Marni Benson", "Bachelor")
	students = appendStudent(students, "0006", "Derrick Griffiths", "Master")
	students = appendStudent(students, "0007", "Leigh Daly", "Bachelor")
	students = appendStudent(students, "0008", "Marni Benson", "PHD")
	students = appendStudent(students, "0009", "Klay Correa", "Bachelor")

	return
}

func appendStudent(students []Student, code, name, program string) (resStudent []Student) {
	resStudent = append(students, Student{
		Code:    code,
		Name:    name,
		Program: program,
	})
	return
}

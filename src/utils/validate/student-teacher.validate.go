package validate

import (
	"errors"

	"github.com/Leonardo-Antonio/api.driving-school/src/entity"
)

var (
	errIdStudent = errors.New("the student's id is required")
	errIdTeacher = errors.New("the teacher id is required")
)

func StudentTeacher(studentTeacher entity.AssignStudentTeacher) error {
	if studentTeacher.IdClient.IsZero() {
		return errIdStudent
	}

	if studentTeacher.IdClient.IsZero() {
		return errIdTeacher
	}

	return nil
}

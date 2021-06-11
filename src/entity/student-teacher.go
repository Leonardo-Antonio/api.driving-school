package entity

type (
	StudentTeacher struct {
		Client  []User `bson:"client,omitempty" json:"client,omitempty" xml:"client,omitempty"`
		Teacher []User `bson:"teacher,omitempty" json:"teacher,omitempty" xml:"teacher,omitempty"`
	}

	TeacherStudents struct {
		Teacher  User   `bson:"teacher,omitempty" json:"teacher,omitempty" xml:"teacher,omitempty"`
		Students []User `bson:"students,omitempty" json:"students,omitempty" xml:"students,omitempty"`
	}
)

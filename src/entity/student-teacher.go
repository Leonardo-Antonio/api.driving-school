package entity

type StudentTeacher struct {
	Client  []User `bson:"client,omitempty" json:"client,omitempty" xml:"client,omitempty"`
	Teacher []User `bson:"teacher,omitempty" json:"teacher,omitempty" xml:"teacher,omitempty"`
}

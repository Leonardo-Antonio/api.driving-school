package entity

type Dni struct {
	Dni             string `bson:"dni,omitempty" json:"dni,omitempty"`
	Names           string `bson:"nombres,omitempty" json:"nombres,omitempty"`
	FathersLastName string `bson:"apellidoPaterno,omitempty" json:"apellidoPaterno,omitempty"`
	MothersLastName string `bson:"apellidoMaterno,omitempty" json:"apellidoMaterno,omitempty"`
	Code            string `bson:"codVerifica,omitempty" json:"codVerifica,omitempty"`
}

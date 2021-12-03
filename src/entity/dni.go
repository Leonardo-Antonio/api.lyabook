package entity

type Dni struct {
	NombreCompleto  string `json:"nombre,omitempty" xml:"nombre,omitempty" bson:"nombre,omitempty"`
	Numero          string `json:"numeroDocumento,omitempty" xml:"numeroDocumento,omitempty" bson:"numeroDocumento,omitempty"`
	TipoDocumento   string `json:"tipoDocumento,omitempty" xml:"tipoDocumento,omitempty" bson:"tipoDocumento,omitempty"`
	ApellidoPaterno string `json:"apellidoPaterno,omitempty" xml:"apellidoPaterno,omitempty" bson:"apellidoPaterno,omitempty"`
	ApellidoMaterno string `json:"apellidoMaterno,omitempty" xml:"apellidoMaterno,omitempty" bson:"apellidoMaterno,omitempty"`
	Nombres         string `json:"nombres,omitempty" xml:"nombres,omitempty" bson:"nombres,omitempty"`
}

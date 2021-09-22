package entity

type ResponseClaim struct {
	Name    string `json:"name,omitempty" xml:"name,omitempty" bson:"name,omitempty" validmor:"required"`
	Type    string `json:"type,omitempty" xml:"type,omitempty" bson:"type,omitempty" validmor:"required"`
	Message string `json:"message,omitempty" xml:"message,omitempty" bson:"message,omitempty" validmor:"required"`
	Email   string `json:"email,omitempty" xml:"email,omitempty" bson:"email,omitempty" validmor:"mail,required"`
}

package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DataPerson struct {
	Name           string `bson:"name,omitempty" json:"name,omitempty" xml:"name,omitempty" validmor:"required"`
	LastName       string `bson:"last_name,omitempty" json:"last_name,omitempty" xml:"last_name,omitempty" validmor:"required"`
	TypeDocument   string `bson:"type_document,omitempty" json:"type_document,omitempty" xml:"type_document,omitempty" validmor:"required"`
	NumberDocument string `bson:"number_document,omitempty" json:"number_document,omitempty" xml:"number_document,omitempty" validmor:"required"`
}

type DataSendingReply struct {
	Phone               string `bson:"phone,omitempty" json:"phone,omitempty" xml:"phone,omitempty" validmor:"required"`
	AdditionalTelephone string `bson:"additional_telephone,omitempty" json:"additional_telephone,omitempty" xml:"additional_telephone,omitempty"`
	Email               string `bson:"email,omitempty" json:"email,omitempty" xml:"email,omitempty" validmor:"mail,required"`
	CommunicationMethod string `bson:"communication_method,omitempty" json:"communication_method,omitempty" xml:"communication_method,omitempty"`
	Address             string `bson:"address,omitempty" json:"address,omitempty" xml:"address,omitempty"`
	Department          string `bson:"department,omitempty" json:"department,omitempty" xml:"department,omitempty" validmor:"required"`
	Province            string `bson:"province,omitempty" json:"province,omitempty" xml:"province,omitempty" validmor:"required"`
	District            string `bson:"district,omitempty" json:"district,omitempty" xml:"district,omitempty" validmor:"required"`
}

type ClaimData struct {
	IdentificationContractedAsset string `bson:"identification_contracted_asset,omitempty" json:"identification_contracted_asset,omitempty" xml:"identification_contracted_asset,omitempty" validmor:"required"`
	Type                          string `bson:"type,omitempty" json:"type,omitempty" xml:"type,omitempty" validmor:"required"`
	Details                       string `bson:"details,omitempty" json:"details,omitempty" xml:"details,omitempty" validmor:"required"`
	Order                         string `bson:"order,omitempty" json:"order,omitempty" xml:"order,omitempty" validmor:"required"`
}

type ComplaintsBook struct {
	Id               primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty" xml:"_id,omitempty"`
	DataPerson       DataPerson         `bson:"data_person,omitempty" json:"data_person,omitempty" xml:"data_person,omitempty" validmor:"required"`
	DataSendingReply DataSendingReply   `bson:"data_sending_reply,omitempty" json:"data_sending_reply,omitempty" xml:"data_sending_reply,omitempty" validmor:"required"`
	ClaimData        ClaimData          `bson:"claim_data,omitempty" json:"claim_data,omitempty" xml:"claim_data,omitempty" validmor:"required"`
	CreatedAt        time.Time          `bson:"created_at,omitempty" json:"created_at,omitempty" xml:"created_at,omitempty"`
	UpdatedAt        time.Time          `bson:"updated_at,omitempty" json:"updated_at,omitempty" xml:"updated_at,omitempty"`
	DeletedAt        time.Time          `bson:"deleted_at,omitempty" json:"deleted_at,omitempty" xml:"deleted_at,omitempty"`
	Active           bool               `bson:"active,omitempty" json:"active,omitempty" xml:"active,omitempty"`
}

type ComplaintsBooks []ComplaintsBook

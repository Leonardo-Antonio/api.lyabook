package entity

type MessageEmail struct {
	From    string `json:"from,omitempty" xml:"from,omitempty"`
	Name    string `json:"name,omitempty" xml:"name,omitempty"`
	Subject string `json:"subject,omitempty" xml:"subject,omitempty"`
	Message string `json:"message,omitempty" xml:"message,omitempty"`
}

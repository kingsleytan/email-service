package model

// Mail : Parent Key -> Root
type Mail struct {
	ID           string `json:"id"`
	To           string `json:"to"`
	From         string `json:"from"`
	Domain       string `json:"domain"`
	Subject      string `json:"subject"`
	TemplateData string `json:"templateData"`
	Template     string `json:"template"`
	ReferenceID  string `json:"referenceID"`
	Status       string `json:"status"`
	Events       string `json:"events"`
	Model
}

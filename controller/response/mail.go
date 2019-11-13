package response

// Mail : Parent Key -> Root
type Mail struct {
	To           string         `json:"to"`
	From         string         `json:"from"`
	Domain       string         `json:"domain"`
	TemplateData string         `json:"templateData"`
	Template     string         `json:"template"`
	ReferenceID  string         `json:"referenceID"`
}

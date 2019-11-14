package response

// Mail : Parent Key -> Root
type Mail struct {
	To           string `json:"to"`
	From         string `json:"from"`
	Domain       string `json:"domain"`
	TemplateData struct {
		Title string `json:"title"`
		Body  string `json:"body"`
	} `json:"templateData" validate:"required"`
	Template        string `json:"template"`
	ReferenceID     string `json:"referenceID"`
	CreatedDateTime string `json:"created"`
	UpdatedDateTime string `json:"updated"`
}

package v1

import (
	"context"
	"email-service/config"
	"email-service/controller/response"
	"email-service/model"
	"email-service/package/validator"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/mailgun/mailgun-go/v3"
)

var i struct {
	ID           string `json:"id" validate:"required"`
	To           string `json:"to" validate:"required"`
	From         string `json:"from"`
	Domain       string `json:"domain"`
	Subject      string `json:"subject"`
	TemplateData struct {
		Title string `json:"title"`
		Body  string `json:"body"`
	} `json:"templateData" validate:"required"`
	Template    string `json:"template" validate:"required"`
	ReferenceID string `json:"referenceID"`
	Status      string `json:"status" validate:"required"`
	Events      string `json:"events"`
}

// ListTemplates :
func ListTemplates(c echo.Context) error {
	mg := mailgun.NewMailgun(config.MailgunDomain, config.MailgunKey)

	it := mg.ListTemplates(nil)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	var page, result []mailgun.Template
	for it.Next(ctx, &page) {
		result = append(result, page...)
	}

	if it.Err() != nil {
		return it.Err()
	}

	return c.JSON(http.StatusOK, map[string]string{
		"result": fmt.Sprintf("result: %s", result),
	})
}

// ListTemplateVersions :
func ListTemplateVersions(c echo.Context) error {
	templateName := c.Param("name")
	mg := mailgun.NewMailgun(config.MailgunDomain, config.MailgunKey)

	it := mg.ListTemplateVersions(templateName, nil)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	var page, result []mailgun.TemplateVersion
	for it.Next(ctx, &page) {
		result = append(result, page...)
	}

	if it.Err() != nil {
		return it.Err()
	}

	return c.JSON(http.StatusOK, map[string]string{
		"result": fmt.Sprintf("result: %s", result),
	})
}

var j struct {
	Template string `json:"template" validate:"required"`
	Version  string `json:"version" validate:"required"`
	Active   bool   `json:"active" validate:"required"`
	Comment  string `json:"comment"`
}

// UpdateTemplateVersion :
func UpdateTemplateVersion(c echo.Context) error {

	// bind input
	if err := c.Bind(&j); err != nil {
		return err
	}

	// validation checking
	if _, err := validator.Validate(&j); err != nil {
		return err
	}

	mg := mailgun.NewMailgun(config.MailgunDomain, config.MailgunKey)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	err := mg.UpdateTemplateVersion(ctx, j.Template, &mailgun.TemplateVersion{
		Comment: j.Comment,
		Tag:     j.Version,
		Active:  j.Active,
	})
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"result": fmt.Sprintf("result: Successfully updated template (%s), version (%s), active (%s)", j.Template, j.Version, j.Active),
	})
}

// SendWithTemplate :
func SendWithTemplate(c echo.Context) error {
	// bind input
	if err := c.Bind(&i); err != nil {
		return err
	}

	// validation checking
	if _, err := validator.Validate(&i); err != nil {
		return err
	}

	m := model.Mail{}
	m.ID = i.ID
	m.To = i.To
	m.From = i.From
	m.Subject = i.Subject
	m.Domain = i.Domain
	m.TemplateData.Title = i.TemplateData.Title
	m.TemplateData.Body = i.TemplateData.Body
	m.Template = i.Template
	m.ReferenceID = i.ReferenceID
	m.Status = i.Status
	m.Events = i.Events

	templateName := m.Template
	mg := mailgun.NewMailgun(config.MailgunDomain, config.MailgunKey)

	var err error

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	// Create a new message with template
	msg := mg.NewMessage(m.From, m.Subject, "")
	msg.SetTemplate(templateName)

	// Add recipients
	msg.AddRecipient(m.To)

	// Add the variables to be used by the template
	msg.AddVariable("title", m.TemplateData.Title)
	msg.AddVariable("body", m.TemplateData.Body)

	_, id, err := mg.Send(ctx, msg)
	result := fmt.Sprintf("Queued: %s", id)
	if err != nil {
		return c.JSON(400, err.Error())
	}

	r := response.Mail{}
	r.To = m.To
	r.From = m.From
	r.Domain = m.Domain
	r.TemplateData.Title = m.TemplateData.Title
	r.TemplateData.Body = m.TemplateData.Body
	r.Template = m.Template
	r.ReferenceID = m.ReferenceID
	r.CreatedDateTime = time.Now().Format(time.RFC3339)
	r.UpdatedDateTime = time.Now().Format(time.RFC3339)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"item":   r,
		"result": fmt.Sprintf("result: %s", result),
	})
}

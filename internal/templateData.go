package internal

import validation "github.com/go-ozzo/ozzo-validation/v4"

type Fields struct {
	Name string
	Email string
}

type TemplateData struct {
	Fields Fields
	Errors validation.Errors
	Members *Members
}

func NewTemplateData() TemplateData {
	members := NewMembers()
	return TemplateData{
		Fields: Fields{
			Name: "",
			Email: "",
		},
		Errors: make(validation.Errors),
		Members: &members,
	}
}

func (t TemplateData) HasError(field string) bool {
	if _, ok := t.Errors[field]; ok {
		return true
	}
	return false
}

func (t TemplateData) GetError(field string) string {
	if err, ok := t.Errors[field]; ok {
		return err.Error()
	}
	return ""
}

func (t TemplateData) HasMembers() bool  {
	return len(*(t.Members)) > 0
}
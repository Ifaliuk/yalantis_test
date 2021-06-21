package internal

import (
	"errors"
	"fmt"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/go-ozzo/ozzo-validation/v4"
	"time"
)

type Member struct {
	Id       int
	Name     string
	Email    string
	Register time.Time
}

func NewMember(name string, email string) Member {
	return Member{Name: name, Email: email}
}

func (m Member) FormatRegisterTime() string {
	return m.Register.Format("01-02-2006 15:04:05")
}

func (m Member) String() string {
	return fmt.Sprintf("{Name: '%s', Email: '%s', Register: '%s'}",
		m.Name, m.Email, m.FormatRegisterTime(),
	)
}

type Members []Member

func NewMembers() Members {
	return Members{}
}

func (m Members) GetList() Members {
	return m
}

func (m Members) validateMember(member Member) validation.Errors {
	err := validation.ValidateStruct(&member,
		validation.Field(&member.Name, validation.Required),
		validation.Field(&member.Email, validation.Required, is.Email),
	)
	if err != nil {
		res, _ := err.(validation.Errors)
		return res
	}
	return nil
}

func (m *Members) AddMember(member *Member) validation.Errors {
	err := m.validateMember(*member)
	if err != nil {
		return err
	}
	if hasEmail := (*m).hasEmail((*member).Email); hasEmail {
		err := make(validation.Errors)
		err["Email"] = errors.New("user with this email already exists")
		return err
	}
	member.Id = len(*m) + 1
	member.Register = time.Now()
	*m = append(*m, *member)
	return nil
}

func (m Members) hasEmail(email string) bool {
	for _, member := range m {
		if member.Email == email {
			return true
		}
	}
	return false
}

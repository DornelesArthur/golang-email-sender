package campaign

import (
	"errors"
	"net/mail"
	"time"

	"github.com/rs/xid"
)

type Contact struct {
	Email string
}

type Campaign struct {
	ID       string
	Name     string
	Created  time.Time
	Content  string
	Contacts []Contact
}

func NewCampaign(name string, content string, emails []string) (*Campaign, error) {
	if name == "" {
		return nil, errors.New("name is required")
	}
	if content == "" {
		return nil, errors.New("content is required")
	}
	if len(emails) == 0 {
		return nil, errors.New("email list is required")
	}

	contacts := make([]Contact, len(emails))
	for index, email := range emails {
		_, err := mail.ParseAddress(email)
		if err != nil {
			return nil, errors.New("email list contains invalid email")
		}
		contacts[index].Email = email
	}

	return &Campaign{
		ID:       xid.New().String(),
		Name:     name,
		Content:  content,
		Created:  time.Now(),
		Contacts: contacts,
	}, nil
}

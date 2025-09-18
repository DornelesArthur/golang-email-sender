package campaign

import (
	"errors"
	"fmt"
	internalerrors "golang-email-sender/internal/internalErrors"
	"net/mail"
	"time"

	"github.com/rs/xid"
)

type Contact struct {
	Email string `validate:"email"`
}

type Campaign struct {
	ID       string    `validate:"required"`
	Name     string    `validate:"min=5,max=24"`
	Created  time.Time `validate:"required"`
	Content  string    `validate:"min=5,max=1024"`
	Contacts []Contact `validate:"min=1,dive"`
}

func NewCampaign(name string, content string, emails []string) (*Campaign, error) {
	contacts := make([]Contact, len(emails))
	for index, email := range emails {
		_, err := mail.ParseAddress(email)
		if err != nil {
			return nil, errors.New("email list contains invalid email")
		}
		contacts[index].Email = email
	}

	campaign := &Campaign{
		ID:       xid.New().String(),
		Name:     name,
		Content:  content,
		Created:  time.Now(),
		Contacts: contacts,
	}
	fmt.Print(name)
	err := internalerrors.ValidateStruct(campaign)

	if err == nil {
		return campaign, nil
	}
	return nil, err
}

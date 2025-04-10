package campaign_test

import (
	"golang-email-sender/internal/domain/campaign"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	name            = "Campaing X"
	content         = "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Duis pharetra dolor non felis placerat luctus. Curabitur ac accumsan urna. Sed."
	contacts        = []string{"email1@gmail.com", "anotheremail3@yahoo.com", "thirdemail@hotmail.com"}
	invalidContacts = []string{"email1@gmail.com", "anotheremail3@yahoo@.com", "thirdemailhotmail.com"}
)

func Test_NewCampaign_CreateNewCampaign(t *testing.T) {
	assert := assert.New(t)

	campaign, _ := campaign.NewCampaign(name, content, contacts)

	assert.Equal(campaign.Name, name, "Name should be "+name)
	assert.Equal(campaign.Content, content, "Content should be "+name)
	assert.Equal(len(campaign.Contacts), len(contacts), "Number of Contacts dont match")
}

func Test_NewCampaign_IDIsNotNil(t *testing.T) {
	assert := assert.New(t)

	campaign, _ := campaign.NewCampaign(name, content, contacts)

	assert.NotNil(campaign.ID, "ID shouldn'tbe nil")
}

func Test_NewCampaign_CreatedMustBeNow(t *testing.T) {
	assert := assert.New(t)
	beforeCreated := time.Now()

	campaign, _ := campaign.NewCampaign(name, content, contacts)

	assert.WithinRange(campaign.Created, beforeCreated, time.Now().Add(time.Minute))
}

func Test_NewCampaign_ValidateNameIsEmpty(t *testing.T) {
	assert := assert.New(t)

	_, err := campaign.NewCampaign("", content, contacts)

	assert.Equal("name is required", err.Error())
}

func Test_NewCampaign_ValidateContentIsEmpty(t *testing.T) {
	assert := assert.New(t)

	_, err := campaign.NewCampaign(name, "", contacts)

	assert.Equal("content is required", err.Error())
}

func Test_NewCampaign_ValidateContactsIsEmpty(t *testing.T) {
	assert := assert.New(t)

	_, err := campaign.NewCampaign(name, content, []string{})

	assert.Equal("email list is required", err.Error())
}

func Test_NewCampaign_ValidateContactsContainsInvalidEmail(t *testing.T) {
	assert := assert.New(t)

	_, err := campaign.NewCampaign(name, content, invalidContacts)

	assert.Equal("email list contains invalid email", err.Error())
}

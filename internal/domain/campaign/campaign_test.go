package campaign_test

import (
	"golang-email-sender/internal/domain/campaign"
	"testing"
	"time"

	"github.com/jaswdr/faker/v2"
	"github.com/stretchr/testify/assert"
)

var (
	name            = "Campaign X"
	content         = "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Duis pharetra dolor non felis placerat luctus. Curabitur ac accumsan urna. Sed."
	contacts        = []string{"email1@gmail.com", "anotheremail3@yahoo.com", "thirdemail@hotmail.com"}
	invalidContacts = []string{"email1@gmail.com", "anotheremail3@yahoo@.com", "thirdemailhotmail.com"}
	fake            = faker.New()
)

func Test_NewCampaign_CreateNewCampaign(t *testing.T) {
	assert := assert.New(t)

	campaign, _ := campaign.NewCampaign(name, content, contacts)

	assert.Equal(campaign.Name, name, "Name should be "+name)
	assert.Equal(campaign.Content, content, "Content should be "+content)
	assert.Equal(len(campaign.Contacts), len(contacts), "Number of Contacts dont match")
}

func Test_NewCampaign_IDIsNotNil(t *testing.T) {
	assert := assert.New(t)

	campaign, _ := campaign.NewCampaign(name, content, contacts)

	assert.NotNil(campaign.ID, "ID shouldn't be nil")
}

func Test_NewCampaign_CreatedMustBeNow(t *testing.T) {
	assert := assert.New(t)
	beforeCreated := time.Now()

	campaign, _ := campaign.NewCampaign(name, content, contacts)

	assert.WithinRange(campaign.Created, beforeCreated, time.Now().Add(time.Minute))
}

func Test_NewCampaign_ValidateNameMin(t *testing.T) {
	assert := assert.New(t)

	_, err := campaign.NewCampaign("1234", content, contacts)

	assert.Equal("name is required with min 5", err.Error())
}

func Test_NewCampaign_ValidateNameMax(t *testing.T) {
	assert := assert.New(t)

	_, err := campaign.NewCampaign(fake.Lorem().Text(30), content, contacts)

	assert.Equal("name is required with max 24", err.Error())
}

func Test_NewCampaign_ValidateContentMin(t *testing.T) {
	assert := assert.New(t)

	_, err := campaign.NewCampaign(name, "1234", contacts)

	assert.Equal("content is required with min 5", err.Error())
}

func Test_NewCampaign_ValidateContentMax(t *testing.T) {
	assert := assert.New(t)

	_, err := campaign.NewCampaign(name, fake.Lorem().Text(1040), contacts)

	assert.Equal("content is required with max 1024", err.Error())
}

func Test_NewCampaign_ValidateContactsMin(t *testing.T) {
	assert := assert.New(t)

	_, err := campaign.NewCampaign(name, content, []string{})

	assert.Equal("contacts is required with min 1", err.Error())
}

func Test_NewCampaign_ValidateContactsContainsInvalidEmail(t *testing.T) {
	assert := assert.New(t)

	_, err := campaign.NewCampaign(name, content, invalidContacts)

	assert.Equal("email list contains invalid email", err.Error())
}

package campaign_test

import (
	"golang-email-sender/internal/domain/campaign"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	name     = "Campaing X"
	content  = "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Duis pharetra dolor non felis placerat luctus. Curabitur ac accumsan urna. Sed."
	contacts = []string{"email1@gmail.com", "anotheremail3@yahoo.com", "thirdemail@hotmail.com"}
)

func Test_NewCampaign_CreateNewCampaign(t *testing.T) {
	assert := assert.New(t)

	campaign := campaign.NewCampaign(name, content, contacts)

	assert.Equal(campaign.Name, name, "Name should be "+name)
	assert.Equal(campaign.Content, content, "Content should be "+name)
	assert.Equal(len(campaign.Contacts), len(contacts), "Number of Contacts dont match")
}

func Test_NewCampaign_IDIsNotNil(t *testing.T) {
	assert := assert.New(t)

	campaign := campaign.NewCampaign(name, content, contacts)

	assert.NotNil(campaign.ID, "ID shouldn'tbe nil")
}

func Test_NewCampaign_CreatedMustBeNow(t *testing.T) {
	assert := assert.New(t)
	beforeCreated := time.Now()

	campaign := campaign.NewCampaign(name, content, contacts)

	assert.WithinRange(campaign.Created, beforeCreated, time.Now().Add(time.Minute))
}

package campaign_test

import (
	"golang-email-sender/internal/domain/campaign"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCampaign(t *testing.T) {
	assert := assert.New(t)
	name := "Campaing X"
	content := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Duis pharetra dolor non felis placerat luctus. Curabitur ac accumsan urna. Sed."
	contacts := []string{"email1@gmail.com", "anotheremail3@yahoo.com", "thirdemail@hotmail.com"}

	campaign := campaign.NewCampaign(name, content, contacts)

	assert.Equal(campaign.ID, "1", "ID should 1")
	assert.Equal(campaign.Name, name, "Name should be "+name)
	assert.Equal(campaign.Content, content, "Content should be "+name)
	assert.Equal(len(campaign.Contacts), len(contacts), "Number of Contacts dont match")
}

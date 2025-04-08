package campaign_test

import (
	"golang-email-sender/internal/domain/campaign"
	"testing"
)

func TestNewCampaign(t *testing.T) {
	name := "Campaing X"
	content := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Duis pharetra dolor non felis placerat luctus. Curabitur ac accumsan urna. Sed."
	contacts := []string{"email1@gmail.com", "anotheremail3@yahoo.com", "thirdemail@hotmail.com"}

	campaign := campaign.NewCampaign(name, content, contacts)

	if campaign.ID != "1" {
		t.Error("expected 1", campaign.ID)
	}
	if campaign.Name != name {
		t.Error("expected the correct name")
	}
	if campaign.Content != content {
		t.Error("expected the correct content")
	}
	if len(campaign.Contacts) != len(contacts) {
		t.Error("expected the amount of contacts")
	}
}

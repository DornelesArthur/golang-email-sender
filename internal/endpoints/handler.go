package endpoints

import "golang-email-sender/internal/domain/campaign"

type Handler struct {
	CampaignService campaign.Service
}

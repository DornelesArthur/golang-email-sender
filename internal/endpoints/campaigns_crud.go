package endpoints

import (
	"fmt"
	"golang-email-sender/internal/contract"
	"net/http"

	"github.com/go-chi/render"
)

func (h *Handler) CampaignPost(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	fmt.Println(h)
	var campaign contract.NewCampaign
	render.DecodeJSON(r.Body, &campaign)
	id, err := h.CampaignService.Create(campaign)
	return map[string]string{"id": id}, 201, err
}

func (h *Handler) CampaignGet(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	campaigns, err := h.CampaignService.Repository.Get()
	return campaigns, 200, err
}

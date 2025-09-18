package endpoints

import (
	"errors"
	"fmt"
	"golang-email-sender/internal/contract"
	internalerrors "golang-email-sender/internal/internalErrors"
	"net/http"

	"github.com/go-chi/render"
)

func (h *Handler) CampaignPost(w http.ResponseWriter, r *http.Request) {
	fmt.Println(h)
	var campaign contract.NewCampaign
	render.DecodeJSON(r.Body, &campaign)
	id, err := h.CampaignService.Create(campaign)
	if err != nil {
		if errors.Is(err, internalerrors.ErrInternal) {
			render.Status(r, 500)
		} else {
			render.Status(r, 400)
		}
		render.JSON(w, r, map[string]string{"error": err.Error()})
		return
	}
	render.Status(r, 201)
	render.JSON(w, r, map[string]string{"id": id})
}

func (h *Handler) CampaignGet(w http.ResponseWriter, r *http.Request) {
	render.Status(r, 200)
	render.JSON(w, r, h.CampaignService.Repository.Get())
}

package response

import "github.com/fazriachyar/cloudBread/models"

type JsonResponse struct {
	Type string `json:"type"`
	Data []models.Bread `json:"data"`
	Message string `json:"message"`
}
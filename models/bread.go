package models

type JsonResponse struct {
	Type string `json:"type"`
	Data []Bread `json:"data"`
	Message string `json:"message"`
}

type Bread struct {
	BreadID string `json:"breadid"`
	BreadName string `json:"breadname"`
	BreadPrice string `json:"breadprice"`
	ImgURL string `json:"imgurl"`
}

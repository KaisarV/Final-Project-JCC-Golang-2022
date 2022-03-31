package model

type ProductReview struct {
	ID        int    `json:"id"`
	UserId    int    `json:"userid"`
	ProductId int    `json:"productid"`
	Review    string `json:"review"`
	Rating    int    `json:"rating"`
	Date      string `json:"Date"`
}

type ProductReviewsResponse struct {
	Status  int             `json:"status"`
	Message string          `json:"message"`
	Data    []ProductReview `json:"data,omitempty"`
}

type ProductReviewResponse struct {
	Status  int           `json:"status"`
	Message string        `json:"message"`
	Data    ProductReview `json:"data,omitempty"`
}

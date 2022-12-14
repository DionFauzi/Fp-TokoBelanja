package response

import "time"

type ProductResponse struct {
	ID         int       `json:"id"`
	Title      string    `json:"title"`
	Price      int       `json:"price"`
	Stock      int       `json:"stock"`
	CategoryID int       `json:"category_Id"`
	CreatedAt  time.Time `json:"created_at"`
}

type ProductPutResponseBody struct {
	ID         int       `json:"id"`
	Title      string    `json:"title"`
	Price      int       `json:"price"`
	Stock      int       `json:"stock"`
	CategoryID int       `json:"CategoryId"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

type ProductPutResponse struct {
	Product ProductPutResponseBody
}

type ProductDeleteResponse struct {
	Message string `json:"message"`
}

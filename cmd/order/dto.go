package order

type CreateRequest struct {
	Customer_name string `json:"customer_name" binding:"required"`
	Product_id    int    `json:"product_id" binding:"required"`
	Quantity      int    `json:"quantity" binding:"required"`
	Status        string `json:"status" binding:"required"`
}

type EditRequest struct {
	Id            int    `json:"id" binding:"required"`
	Customer_name string `json:"customer_name" binding:"required"`
	Product_id    int    `json:"product_id" binding:"required"`
	Quantity      int    `json:"quantity" binding:"required"`
	Status        string `json:"status" binding:"required"`
}

type HapusRequest struct {
	Id int `json:"id" binding:"required"`
}

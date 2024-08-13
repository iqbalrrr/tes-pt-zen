package product

type CreateRequest struct {
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description" binding:"required"`
	Price       float32 `json:"price" binding:"required"`
	Stock       int     `json:"stock" binding:"required"`
}

type EditRequest struct {
	Id          int     `json:"id" binding:"required"`
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description" binding:"required"`
	Price       float32 `json:"price" binding:"required"`
	Stock       int     `json:"stock" binding:"required"`
}

type HapusRequest struct {
	Id int `json:"id" binding:"required"`
}

package model

type ItemResponse struct {
	ID            int    `json:"item_id"`
	Name          string `json:"item_name"`
	CurrentAmount int    `json:"current_amount"`
	ImgUrl        string `json:"img_url"`
	BorrowCount   int    `json:"borrow_count"`
}

type NewItemRequest struct {
	Name          string `json:"item_name"`
	CurrentAmount int    `json:"current_amount"`
	ImgUrl        string `json:"img_url"`
}

type UpdateItemRequest struct {
	ID            *int    `json:"item_id"`
	Name          string `json:"item_name"`
	CurrentAmount *int    `json:"current_amount"`
	ImgUrl        string `json:"img_url"`
	BorrowCount   *int    `json:"borrow_count"`
}

type ItemService interface {
	GetItems() ([]ItemResponse, error)
	GetFrequentlyBorrowed() ([]ItemResponse, error)
	GetItem(int) (*ItemResponse, error)
	CreateItem(NewItemRequest) (*ItemResponse, error)
	UpdateItem(UpdateItemRequest) (*ItemResponse, error)
	DeleteItem(int) (error)
}

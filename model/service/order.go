package model

type ItemInOrderResponse struct {
	ID     int    `json:"item_id"`
	Name   string `json:"item_name"`
	Amount int    `json:"amount"`
	ImgUrl string `json:"img_url"`
}

type ItemInOrderRequest struct {
	ID     int    `json:"item_id"`
	Name   string `json:"item_name"`
	Amount int    `json:"amount"`
}

type OrderResponse struct {
	ID              int                   `json:"order_id"`
	UserId          string                `json:"user_id"`
	UserOrg         string                `json:"user_org"`
	BorrowDatetime  string                `json:"borrow_datetime"`
	ReturnDatetime  string                `json:"return_datetime"`
	PickupDatetime  string                `json:"pickup_datetime"`
	DropoffDatetime string                `json:"dropoff_datetime"`
	Items           []ItemInOrderResponse `json:"items"`
}

type NewOrderRequest struct {
	UserId         string               `json:"user_id"`
	UserOrg        string               `json:"user_org"`
	BorrowDatetime string               `json:"borrow_datetime"`
	ReturnDatetime string               `json:"return_datetime"`
	Items          []ItemInOrderRequest `json:"items"`
}

type UpdateOrderRequest struct {
	ID             int                  `json:"order_id"`
	UserId         string               `json:"user_id"`
	UserOrg        string               `json:"user_org"`
	BorrowDatetime string               `json:"borrow_datetime"`
	ReturnDatetime string               `json:"return_datetime"`
	Items          []ItemInOrderRequest `json:"items"`
}

type CheckOrderRequest struct {
	ID    int                  `json:"order_id"`
	Items []ItemInOrderRequest `json:"items"`
}

type OrderService interface {
	GetOrders() ([]OrderResponse, error)
	GetOrder(int) (*OrderResponse, error)

	CreateOrder(NewOrderRequest) (*OrderResponse, error)
	UpdateOrder(UpdateOrderRequest) (*OrderResponse, error)
	DeleteOrder(int) error

	PickupOrder(CheckOrderRequest) (*OrderResponse, error)
	DropoffOrder(CheckOrderRequest) (*OrderResponse, error)
}

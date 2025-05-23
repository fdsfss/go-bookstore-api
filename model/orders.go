package model

type Orders struct {
	Id         int    `json:"id"`
	CostumerId int    `json:"costumer_id"`
	OrderDate  string `json:"order_date"`
}

type OrderItems struct {
	Id       int `json:"id"`
	OrderId  int `json:"order_id"`
	Quantity int `json:"quantity"`
	BookId   int `json:"book_id"`
}

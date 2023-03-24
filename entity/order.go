package entity

type Order struct {
	Id              int64  `json:"id"`
	Quantity        int    `json:"quantity"`
	CustomerAddress string `json:"customer_address"`
}

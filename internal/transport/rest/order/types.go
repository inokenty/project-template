package order

import "time"

type Order struct {
	ID        uint       `json:"id"`
	CreatedAt *time.Time `json:"created_at"`
	Sum       int64      `json:"sum"`
	UserID    uint       `json:"user_id"`
}

type GetOrderByIDReply struct {
	Order *Order `json:"order"`
}

type ListOrdersReply struct {
	Items []*Order `json:"items"`
}

type CreateOrderArgs struct {
	UserID uint  `json:"user_id"`
	Sum    int64 `json:"sum"`
}

type CreateOrderReply struct {
	ID uint `json:"id"`
}

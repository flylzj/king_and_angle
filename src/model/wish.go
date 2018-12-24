package model

type WishModel struct {
	Wish    string   	`json:"wish" bind:"required"`
}

type WishStatusModel struct {
	WishStatus 	uint	`json:"wish_status" bind:"required"`
}

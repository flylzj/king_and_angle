package model

type WishModel struct {
	Wish    string   	`json:"wish" bind:"required"`
}

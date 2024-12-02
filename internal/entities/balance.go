package entities

type Balance struct {
	Current   float32 `json:"balance"`
	Withdrawn float32 `json:"withdrawn"`
}

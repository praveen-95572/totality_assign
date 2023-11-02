package domain

type User struct {
	Id      int64   `json:"id"`
	Fname   string  `json:"fname"`
	City    string  `json:"city"`
	Phone   int64   `json:"phone"`
	Height  float64 `json:"height"`
	Married bool    `json:"married"`
}

package request

type ReviewPost struct {
	Rating      float64 `json:"rating"`
	Description string  `json:"description"`
}

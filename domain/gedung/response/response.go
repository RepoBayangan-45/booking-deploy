package response

type ResponsePost struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Location    string `json:"location"`
	Price       string `json:"price"`
	Latitude    string `json:"latitude"`
	Longitude   string `json:"longitude"`
	Description string `json:"description"`
}

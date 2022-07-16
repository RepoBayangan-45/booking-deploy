package request

type BookingRequest struct {
	Status      string `json:"status"`
	BookingCode string `json:"bookingcode"`
	OrderDate   string `json:"orderdate"`
	CheckIn     string `json:"checkin"`
	CheckOut    string `json:"checkout"`
}

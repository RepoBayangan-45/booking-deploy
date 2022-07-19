package response

type BookingResponse struct {
	ID          int      `json:"id"`
	Status      string   `json:"status"`
	BookingCode string   `json:"bookingcode"`
	OrderDate   string   `json:"orderdate"`
	CheckIn     string   `json:"checkin"`
	CheckOut    string   `json:"checkout"`
	User        []User   `gorm:"Foreignkey:IDBooking;" json:"user"`
	Gedung      []Gedung `gorm:"Foreignkey:IDBooking;" json:"gedung"`
	Jenis       []Jenis  `gorm:"Foreignkey:IDBooking;" json:"jenis"`
}

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Fullname string `json:"fullname"`
	Alamat   string `json:"alamat"`
	Phone    string `json:"phone"`
}

type Gedung struct {
	ID          int    `json:"id" gorm:"PrimaryKey"`
	Name        string `json:"name"`
	Price       string `json:"price"`
	Location    string `json:"location"`
	Latitude    string `json:"latitude"`
	Longitude   string `json:"longitude"`
	Description string `json:"description"`
}
type Jenis struct {
	ID    int    `json:"id"`
	Jenis string `json:"jenis"`
}
